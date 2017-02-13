// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"

	"github.com/sekisoft/gogogo/model"
)

func ApiWebSocketHandler(wh func(*model.WebSocketRequest) (map[string]interface{}, *model.AppError)) *webSocketHandler {
	return &webSocketHandler{wh}
}

type webSocketHandler struct {
	handlerFunc func(*model.WebSocketRequest) (map[string]interface{}, *model.AppError)
}

func (wh *webSocketHandler) ServeWebSocket(conn *WebConn, r *model.WebSocketRequest) {
	l4g.Debug("/api/websocket:%s", r.Action)

	r.Token = GetToken(conn.TokenID)

	var data map[string]interface{}
	var err *model.AppError

	if data, err = wh.handlerFunc(r); err != nil {
		l4g.Error("Websocket handler error: action: %s seq: %s playerID: %s", r.Action, r.Sequence, r.Token.PlayerID)
		conn.Send <- model.NewWebSocketError(r.Sequence, err)
		return
	}

	conn.Send <- model.NewWebSocketResponse(model.STATUS_OK, r.Sequence, data)
}

func NewInvalidWebSocketParamError(action string, name string) *model.AppError {
	return model.NewAppError("/api/websocket:"+action, "Invalid parameters", http.StatusBadRequest)
}
