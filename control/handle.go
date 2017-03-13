package control

import (
	"gopkg.in/kataras/iris.v6"
	ws "github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Control struct {
	DB *sqlx.DB
}



var upgrader = ws.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// returns logged in user id
func (c *Control) GetUserId(ctx *iris.Context) (int, error) {
	var userId int
	var err error

	// validate user session
	/*
	userId, err = strconv.ParseInt(ctx.RequestHeader("X-User-Id"), 10, 0)
	if err != nil {
		fmt.Println("unable to get user id from header:", err)
		return
	}
	*/


	userId, err = ctx.Session().GetInt("id")
	return userId, err
}


func (c *Control) Serve(ctx *iris.Context) {
	// verify if user is authenticated etc etc
	var err error
	var wsc *ws.Conn

	// upgrade current control socket get request to websocket
	wsc, err = upgrader.Upgrade(ctx.ResponseWriter, ctx.Request, ctx.ResponseWriter.Header())
	if err != nil {
		ctx.JSON(iris.StatusBadRequest, map[string]string{"Error":err.Error()})
		return
	}
	// upgrade to websocket end
	var userId int
	userId, err = c.GetUserId(ctx)
	if err != nil {
		fmt.Println("user id not found, publish not allowed:", err)
	}

	// once the upgrade is successful create control context
	controlContext := &ControlContext{}
	controlContext.Init(wsc, userId)
	defer controlContext.Close()
	var in []byte
	// read from websocket and push to stream
	for {
		fmt.Println("reading message")
		_, in, err = wsc.ReadMessage()
		if err != nil  {
			// if websocket connection is closed break out of the read loop
			if ws.IsUnexpectedCloseError(err) || ws.IsCloseError(err) {
				fmt.Println("websocket closed:", err)
				break
			}
			// else continue
			fmt.Println("unable to read from websocket:", err)
			continue
		}

		controlContext.HandleStreamMessage(c.DB, in)
	}


}




