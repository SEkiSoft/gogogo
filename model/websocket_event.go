// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
	"io"
)

const (
	WEBSOCKET_EVENT_MATCHMAKING_START  = "matchmaking_start"
	WEBSOCKET_EVENT_MATCHMAKING_CANCEL = "matchmaking_cancel"
	WEBSOCKET_EVENT_MATCHMAKING_FOUND  = "matchmaking_found"
	WEBSOCKET_EVENT_GAME_MOVE          = "game_move"
	WEBSOCKET_EVENT_HELLO_WORLD        = "hello_world"
)

type WebSocketEvent struct {
	PlayerId string                 `json:"player_id"`
	GameId   string                 `json:"game_id"`
	Event    string                 `json:"event"`
	Data     map[string]interface{} `json:"data"`
}

func NewWebSocketEvent(playerId, gameId, event string) *WebSocketEvent {
	return &WebSocketEvent{
		PlayerId: playerId,
		GameId:   gameId,
		Event:    event,
		Data:     make(map[string]interface{}),
	}
}

func (w *WebSocketEvent) Add(key string, value interface{}) {
	w.Data[key] = value
}

func (w *WebSocketEvent) IsValid() bool {
	return w.Event != ""
}

func (w *WebSocketEvent) ToJson() string {
	json, err := json.Marshal(w)
	if err != nil {
		return ""
	}

	return string(json)
}

func WebSocketEventFromJson(data io.Reader) *WebSocketEvent {
	decoder := json.NewDecoder(data)
	var w WebSocketEvent
	err := decoder.Decode(&w)

	if err != nil {
		return nil
	}

	return &w
}
