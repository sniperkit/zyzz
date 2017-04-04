package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/thekb/zyzz/db"
	"fmt"
	"time"
)

const (
	STATUS_CREATED = iota
	STATUS_STREAMING
	STATUS_STOPPED
	STATUS_ERROR

	CREATE_STREAM_SERVER = `INSERT INTO stream_server (short_id, name, hostname, internal_ip, external_ip)
				VALUES (:short_id, :name, :hostname, :internal_ip, :external_ip) returning id;`

	GET_STREAM_SERVER_SHORT_ID = `SELECT A.* FROM stream_server A
				WHERE A.short_id=$1;`

	GET_STREAM_SERVER_ID = `SELECT A.* FROM stream_server A
				WHERE A.id=$1;`

	GET_DEFAULT_STREAM_SERVER = `SELECT A.* FROM stream_server A
				ORDER BY A.id ASC
				LIMIT 1;`

	CREATE_STREAM = `INSERT INTO stream (
				short_id, name, description,
				publish_url, subscribe_url, transport_url,
	 			stream_server_id, creator_id, event_id, language)
				VALUES (
				:short_id, :name, :description,
				:publish_url, :subscribe_url, :transport_url,
				:stream_server_id, :creator_id, :event_id, :language)
				returning id;`

	GET_STREAM_SHORT_ID = `SELECT A.* FROM stream A
				WHERE A.short_id=$1;`

	GET_STREAM_ID = `SELECT A.* FROM stream A
				WHERE A.id=$1;`

	GET_STREAMS = `SELECT A.* FROM stream A
				WHERE A.event_id=$1 and status=1
				ORDER By A.id ASC;`

	UPDATE_STREAM_STATUS = `UPDATE stream SET status=$1
				WHERE short_id=$2;`

	INCREMENT_STREAM_SUBSCRIBER_COUNT = `UPDATE stream
				SET subscriber_count = subscriber_count + 1
				WHERE short_id = $1;`

	INCREMENT_ACTIVE_LISTENERS_COUNT = `UPDATE stream
				SET active_listeners = active_listeners + 1
				WHERE short_id = $1;`

	DECREMENT_ACTIVE_LISTENERS_COUNT = `UPDATE stream
				SET active_listeners = active_listeners - 1
				WHERE short_id = $1;`

	STOP_STREAM = `UPDATE stream A
				SET A.status=$1, A.ended_at=CURRENT_TIMESTAMP
				WHERE A.short_id=$2;`
	GET_STREAM_FOR_USER_SHORT_ID = `select S.* from stream S inner join users U on S.creator_id = U.id
				and U.short_id = $1 and S.status=1;`

	UPDATE_STREAM = `UPDATE stream SET name=:name, description=:description, language=:language,
				match_source=:match_source WHERE stream.short_id=:short_id`
)

type Stream struct {
	Id              int `db:"id" json:"-"`
	ShortId         string `db:"short_id" json:"id"`
	Name            string `db:"name" json:"name"`
	Description     string `db:"description" json:"description"`
	StartedAt       time.Time `db:"started_at" json:"started_at"`
	EndedAt         time.Time `db:"ended_at" json:"ended_at"`
	Status          int `db:"status" json:"status"`
	PublishUrl      string `db:"publish_url" json:"publish_url"`
	SubscribeUrl    string `db:"subscribe_url" json:"subscribe_url"`
	SubscriberCount int `db:"subscriber_count" json:"subscriber_count"`
	ActiveListeners int `db:"active_listeners" json:"active_listeners"`
	CreatorId       int `db:"creator_id" json:"creator_id"`
	StreamServerId  int `db:"stream_server_id" json:"stream_server_id"`
	TransportUrl    string `db:"transport_url" json:"transport_url"`
	EventId		string `db:"event_id" json:"event_id"`
	User 		User `db:"_" json:"user"`
	Language        string `db:"language" json:"language"`
	MatchSource string `db:"match_source" json:"match_source"`
}

type StreamServer struct {
	Id         int `db:"id" json:"-"`
	ShortId    string `db:"short_id" json:"id"`
	Name       string `db:"name" json:"name"`
	HostName   string `db:"hostname" json:"host_name"`
	InternalIP string `db:"internal_ip" json:"internal_ip"`
	ExternalIP string `db:"external_ip" json:"external_ip"`
}

func CreateStreamServer(d *sqlx.DB, streamServer *StreamServer) (int64, error) {
	id, err := db.InsertStruct(d, CREATE_STREAM_SERVER, streamServer)
	if err != nil {
		fmt.Println("unable to create stream server:", err)
	}
	return id, err
}

func GetStreamServerForShortId(d *sqlx.DB, short_id string) (StreamServer, error) {
	var streamServer StreamServer
	err := db.Get(d, GET_STREAM_SERVER_SHORT_ID, &streamServer, short_id)
	if err != nil {
		fmt.Println("unable to get stream server:", err)
	}
	return streamServer, err
}

func GetStreamServerForId(d *sqlx.DB, id int64) (StreamServer, error) {
	var streamServer StreamServer
	err := db.Get(d, GET_STREAM_SERVER_ID, &streamServer, id)
	if err != nil {
		fmt.Println("unable to get stream server:", err)
	}
	return streamServer, err
}

func GetDefaultStreamServer(d *sqlx.DB) StreamServer {
	var streamServer StreamServer
	err := db.Get(d, GET_DEFAULT_STREAM_SERVER, &streamServer)
	if err != nil {
		fmt.Println("unable to get default stream server:", err)
	}
	return streamServer
}

func CreateStream(d *sqlx.DB, stream *Stream) (int64, error) {
	id, err := db.InsertStruct(d, CREATE_STREAM, stream)
	if err != nil {
		fmt.Println("unable to create stream:", err)
	}
	return id, err
}

func GetStreamForShortId(d *sqlx.DB, short_id string) (Stream, error) {
	var stream Stream
	err := db.Get(d, GET_STREAM_SHORT_ID, &stream, short_id)
	if err != nil {
		fmt.Println("unable to get stream:", err)
	}
	return stream, err
}

func GetStreamForId(d *sqlx.DB, id int64) (Stream, error) {
	var stream Stream
	err := db.Get(d, GET_STREAM_ID, &stream, id)
	if err != nil {
		fmt.Println("unable to get stream:", err)
	}
	return stream, err
}

func GetStreams(d *sqlx.DB, event_id string) ([]Stream, error) {
	var streams []Stream
	err := db.Select(d, GET_STREAMS, &streams, event_id)
	if err != nil {
		fmt.Println("unable to get streams:", err)
	}
	return streams, err
}

func SetStreamStatus(d *sqlx.DB, short_id string, status int) error {
	err := db.Update(d, UPDATE_STREAM_STATUS, status, short_id)
	if err != nil {
		fmt.Println("unable to update stream status:", err)
	}
	return err
}

func IncrementStreamSubscriberCount(d *sqlx.DB, short_id string) error {
	err := db.Update(d, INCREMENT_STREAM_SUBSCRIBER_COUNT, short_id)
	if err != nil {
		fmt.Println("unable to update subscriber count:", err)
	}
	return err
}

func StopStream(d *sqlx.DB, short_id string) error {
	err := db.Update(d, STOP_STREAM, STATUS_STOPPED, short_id)
	if err != nil {
		fmt.Println("unable to stop stream:", err)
	}
	return err
}

func IncrementActiveListenersCount(d *sqlx.DB, short_id string) error {
	err := db.Update(d, INCREMENT_ACTIVE_LISTENERS_COUNT, short_id)
	if err != nil {
		fmt.Println("unable to update active listeners count:", err)
	}
	return err
}

func DecrementActiveListenersCount(d *sqlx.DB, short_id string) error {
	err := db.Update(d, DECREMENT_ACTIVE_LISTENERS_COUNT, short_id)
	if err != nil {
		fmt.Println("unable to update active listeners count:", err)
	}
	return err
}

func GetUserActiveStream(d *sqlx.DB, short_id string) (Stream, error) {
	var stream Stream
	err := db.Get(d, GET_STREAM_FOR_USER_SHORT_ID, &stream, short_id)
	if err != nil {
		fmt.Println("unable to get stream:", err)
	}
	return stream, err
}

func UpdateStream(d *sqlx.DB, stream *Stream) (error) {
	err := db.UpdateObj(d, UPDATE_STREAM, stream)
	if err != nil {
		fmt.Println("unable to update stream:", err)
	}
	return err
}
