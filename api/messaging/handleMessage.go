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

// @TODO: figure out how to structure handling messages
// @TODO: Handle getting current user
// @TODO: Handle getting active pmc
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
