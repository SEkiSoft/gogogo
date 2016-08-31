// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	WEBSOCKET_URL = "/api/websocket"
)

type WebSocketClient struct {
	Url             string
	Connection      *websocket.Conn
	AuthId          string
	Sequence        int64
	EventChannel    chan *WebSocketEvent
	ResponseChannel chan *WebSocketResponse
}

func NewWebSocketClient(url, authId string) (*WebSocketClient, *Error) {
	header := http.Header{}
	header.Set(HEADER_AUTH, HEADER_BEAR+authId)

	conn, _, err := websocket.DefaultDialer.Dial(url+WEBSOCKET_URL, header)

	if err != nil {
		return nil, NewLocError("NewWebSocketClient", "WebSocket client connection fail", nil, err.Error())
	}

	return &WebSocketClient{
		Url:             url,
		Connection:      conn,
		AuthId:          authId,
		Sequence:        1,
		EventChannel:    make(chan *WebSocketEvent, 100),
		ResponseChannel: make(chan *WebSocketResponse, 100),
	}, nil
}

func (w *WebSocketClient) Connect() *Error {
	header := http.Header{}
	header.Set(HEADER_AUTH, HEADER_BEAR+w.AuthId)

	var err error
	w.Connection, _, err = websocket.DefaultDialer.Dial(w.Url+WEBSOCKET_URL, header)

	if err != nil {
		return NewLocError("WebSocketClientConnect", "WebSocket client connection fail", nil, err.Error())
	}

	return nil
}

func (w *WebSocketClient) Close() {
	w.Connection.Close()
}

func (w *WebSocketClient) Listen() {
	go func() {
		for {
			var rawMessage json.RawMessage
			var err error
			if _, rawMessage, err = w.Connection.ReadMessage(); err != nil {
				return
			}

			var event WebSocketEvent
			if err := json.Unmarshal(rawMessage, &event); err == nil && event.IsValid() {
				w.EventChannel <- &event
				continue
			}

			var response WebSocketResponse
			if err := json.Unmarshal(rawMessage, &response); err == nil && response.IsValid() {
				w.ResponseChannel <- &response
				continue
			}
		}
	}()
}

func (w *WebSocketClient) SendMessage(action string, data map[string]interface{}) {
	request := &WebSocketRequest{
		Sequence: w.Sequence,
		Action:   action,
		Data:     data,
	}

	w.Sequence++

	w.Connection.WriteJSON(request)
}
