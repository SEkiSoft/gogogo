// Copyright (c) 2016 sekisoft
// See License.txt

package model

import (
	"strings"
	"testing"
)

func TestWebSocketRequestToJson(t *testing.T) {
	wsr := &WebSocketRequest{
		Sequence: 1,
		Action:   WEBSOCKET_EVENT_GAME_MOVE,
		Data:     make(map[string]interface{}),
		Token:    nil,
	}

	json := wsr.ToJson()
	rwsr := WebSocketRequestFromJson(strings.NewReader(json))

	if rwsr.Action != wsr.Action {
		t.Fatal("Actions do not match")
	}
}

func TestWebSocketRequestFromJson(t *testing.T) {
	wsr := &WebSocketRequest{
		Sequence: 1,
		Action:   WEBSOCKET_EVENT_GAME_MOVE,
		Data:     make(map[string]interface{}),
		Token:    nil,
	}

	json := wsr.ToJson()
	rwsr := WebSocketRequestFromJson(strings.NewReader(json))
	rjson := rwsr.ToJson()

	if json != rjson {
		t.Fatal("JSONs do not match")
	}
}
