package api

import (
	"net/http"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/ventu-io/go-shortid"
	"gopkg.in/redis.v5"
	"gopkg.in/kataras/iris.v6"
)

const (
	HEADER_CONTENT_TYPE = "Content-Type"
	JSON_CONTENT_TYPE = "application/json"
	SHORT_ID = "id"
)

var shortIdGenerator *shortid.Shortid

func init () {
	initializeShortId()
}

type Common struct {
	DB *sqlx.DB
	R *redis.Client
}

type Response struct {
	Error string `json:"error"`
	Data interface{} `json:"data"`
}

func (c *Common) decodeRequestJSON(r *http.Request, destination interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(destination)
	return err
}

func (c *Common) SendErrorJSON(rw http.ResponseWriter, error string, status int) {
	rw.Header().Set(HEADER_CONTENT_TYPE, JSON_CONTENT_TYPE)
	rw.WriteHeader(status)
	response := Response{Error: error}
	b, _ := json.Marshal(&response)
	rw.Write(b)
	return
}

func (c *Common) SendJSON(rw http.ResponseWriter, data interface{}) {
	rw.Header().Set(HEADER_CONTENT_TYPE, JSON_CONTENT_TYPE)
	rw.WriteHeader(http.StatusOK)
	response := Response{Data: data}
	b, _ := json.Marshal(&response)
	rw.Write(b)
}

func initializeShortId() {
	shortIdGenerator, _ = shortid.New(1, shortid.DefaultABC, 1729)
}

func getNewShortId() string {
	shortId, _ := shortIdGenerator.Generate()
	return shortId
}

func getUserId(ctx *iris.Context, r *redis.Client) (int, error) {
	var userId int
	var err error
	userId, err = ctx.Session().GetInt("id")
	if err != nil {
		sessionHeader := ctx.RequestHeader("X-Session-Token")
		redisValue := r.Get(sessionHeader)
		userId, err := redisValue.Int64()
		return int(userId), err
	}
	return userId, err
}