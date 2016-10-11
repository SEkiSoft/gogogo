// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	l4g "github.com/alecthomas/log4go"

	"github.com/sekisoft/gogogo/model"
)

func ApiWebSocketHandler(wh func(*model.WebSocketRequest) (map[string]interface{}, *model.Error)) *webSocketHandler {
	return &webSocketHandler{wh}
}

type webSocketHandler struct {
	handlerFunc func(*model.WebSocketRequest) (map[string]interface{}, *model.Error)
}

func (wh *webSocketHandler) ServeWebSocket(conn *WebConn, r *model.WebSocketRequest) {
	l4g.Debug("/api/websocket:%s", r.Action)

	r.Token = GetToken(conn.TokenId)

	var data map[string]interface{}
	var err *model.Error

	if data, err = wh.handlerFunc(r); err != nil {
		l4g.Error("Websocket handler error: action: %s seq: %s playerId: %s", r.Action, r.Sequence, r.Token.PlayerId)
		conn.Send <- model.NewWebSocketError(r.Sequence, err)
		return
	}

	conn.Send <- model.NewWebSocketResponse(model.STATUS_OK, r.Sequence, data)
}

func NewInvalidWebSocketParamError(action string, name string) *model.Error {
	return model.NewLocError("/api/websocket:"+action, "Invalid parameters", map[string]interface{}{"Name": name}, "")
}
