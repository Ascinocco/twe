package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// @TODO: Create messaging folder, place this there, along with other ws specific stuff other than http endpoint
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			// @TODO: better error handling
			log.Println("err reading message")
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			// @TODO: Better error handling
			log.Println("err writing message", err)
			return
		}
	}
}

// @TODO: WS endpoint needs to be public bc, you cant send http headers, so no bearer token. Message handler will need to verify
// authorization of user.
func EstablishWsConn(w http.ResponseWriter, r *http.Request) {
	// @TODO: restrict origin
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		// @TODO: proper error handling
		fmt.Println("err establishing socket conn")
	}

	reader(ws)
}
