package controllers

import (
	"TheWarEconomy/api/messaging"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func EstablishWsConn(w http.ResponseWriter, r *http.Request) {
	// @TODO: restrict origin
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		// @TODO: proper error handling
		fmt.Println("err establishing socket conn")
	}

	messaging.WsReader(ws)
}
