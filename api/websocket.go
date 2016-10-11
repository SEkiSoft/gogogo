// Copyright (c) 2016 sekisoft
// See License.txt

package api

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sekisoft/gogogo/model"

	l4g "github.com/alecthomas/log4go"
)

func InitWebSocket() {
	l4g.Info("Initializing Websockets")
	BaseRoutes.Players.Handle("/websocket", ApiPlayerRequired(connect)).Methods("GET")
	webHub.Start()
}

func connect(s *Session, w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		l4g.Error("Websocket connection error: %s", err.Error())
		s.Err = model.NewLocError("connect", "Upgrade to websocket connection failed", nil, "")
		return
	}

	wc := NewWebConn(s, ws)
	webHub.Register(wc)
	go wc.writePump()
	wc.readPump()
}
