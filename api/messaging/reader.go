package messaging

import (
	"TheWarEconomy/api/middleware"
	"TheWarEconomy/api/utils"
	"encoding/json"
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
)

type MsgRcv struct {
	Token string `json:"token"`
	Data  string `json:"data"`
	Type  string `json:"type"`
}

type MsgSnd struct {
	Data  string `json:"data"`
	Error string `json:"error"`
	Type  string `json:"type"`
}

func authorize(tk string) (string, error) {
	if len(tk) <= 0 {
		return "", errors.New("Unauthorized")
	}

	token := &middleware.Token{}
	pt, err := jwt.ParseWithClaims(tk, token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(utils.EnvTokenSecret)), nil
	})

	if err != nil || !pt.Valid {
		return "", errors.New("Unauthorized")
	}

	return token.UserId, nil
}

func WsReader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()

		if err != nil {
			conn.WriteJSON(MsgSnd{
				Error: "Unable to parse message",
				Type:  "error",
			})
			return
		}

		message := MsgRcv{}
		err = json.Unmarshal(p, &message)

		if err != nil {
			conn.WriteJSON(MsgSnd{
				Error: "Unable to parse message",
				Type:  "error",
			})
			return
		}

		userId, err := authorize(message.Token)

		if err != nil {
			conn.WriteJSON(MsgSnd{
				Error: "Unauthorized",
				Type:  "error",
			})
			return
		}

		HandleMsg(Msg{
			conn:    conn,
			msgType: message.Type,
			data:    message.Data,
			userId:  userId,
		})
	}
}
