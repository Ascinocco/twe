package messaging

import (
	"github.com/gorilla/websocket"
)

type Msg struct {
	conn    *websocket.Conn
	msgType string
	data    string
	userId  string
}

type HelloResponse struct {
	Message string `json:"message"`
}

func HandleMsg(msg Msg) {
	switch msg.msgType {
	case "hello":
		{
			msg.conn.WriteJSON(&HelloResponse{
				Message: "hi",
			})
		}
	}
}
