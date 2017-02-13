// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"

	"github.com/sekisoft/gogogo/model"
)

type WebSocketRouter struct {
	handlers map[string]*webSocketHandler
}

func NewWebSocketRouter() *WebSocketRouter {
	router := &WebSocketRouter{}
	router.handlers = make(map[string]*webSocketHandler)
	return router
}

func (wr *WebSocketRouter) Handle(action string, handler *webSocketHandler) {
	wr.handlers[action] = handler
}

func (wr *WebSocketRouter) ServeWebSocket(conn *WebConn, r *model.WebSocketRequest) {
	if r.Action == "" {
		err := model.NewAppError("ServeWebSocket", "No websocket action", http.StatusBadRequest)
		wr.ReturnWebSocketError(conn, r, err)
		return
	}

	if r.Sequence <= 0 {
		err := model.NewAppError("ServeWebSocket", "No websocket sequence", http.StatusBadRequest)
		wr.ReturnWebSocketError(conn, r, err)
		return
	}

	h, ok := wr.handlers[r.Action]
	if !ok {
		err := model.NewAppError("ServeWebSocket", "Websocket not ok", http.StatusBadRequest)
		wr.ReturnWebSocketError(conn, r, err)
		return
	}

	h.ServeWebSocket(conn, r)
}

func (wr *WebSocketRouter) ReturnWebSocketError(conn *WebConn, r *model.WebSocketRequest, err *model.AppError) {
	l4g.Error("Websocket server error: %s", err.Message)

	errorResp := model.NewWebSocketError(r.Sequence, err)

	conn.Send <- errorResp
}
