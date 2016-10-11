// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"strings"
	"testing"
)

func TestNewWebSocketEvent(t *testing.T) {
	wse := NewWebSocketEvent(
		"dummy id1",
		"dummy id2",
		WEBSOCKET_EVENT_MATCHMAKING_START,
	)

	if wse == nil {
		t.Fatal("WebSocketEvent not created")
	}

	if wse.PlayerId != "dummy id1" {
		t.Fatal("WebSocketEvent creation failed")
	}

	if wse.GameId != "dummy id2" {
		t.Fatal("WebSocketEvent creation failed")
	}
}

func TestWebSocketEventAdd(t *testing.T) {
	wse := NewWebSocketEvent(
		NewId(),
		NewId(),
		WEBSOCKET_EVENT_MATCHMAKING_START,
	)

	wse.Add("test", "testing")

	if wse.Data["test"] != "testing" {
		t.Fatal("Add key to data failed")
	}
}

func TestWebSocketEventIsValid(t *testing.T) {
	wse := NewWebSocketEvent(
		NewId(),
		NewId(),
		WEBSOCKET_EVENT_MATCHMAKING_START,
	)

	if !wse.IsValid() {
		t.Fatal("should be true")
	}

	wse.Event = ""

	if wse.IsValid() {
		t.Fatal("should be false")
	}
}

func TestWebSocketEventToJson(t *testing.T) {
	wse := NewWebSocketEvent(
		NewId(),
		NewId(),
		WEBSOCKET_EVENT_MATCHMAKING_START,
	)

	json := wse.ToJson()
	rwse := WebSocketEventFromJson(strings.NewReader(json))

	if rwse.PlayerId != wse.PlayerId {
		t.Fatal("Player Ids do not match")
	}
}

func TestWebSocketEventFromJson(t *testing.T) {
	wse := NewWebSocketEvent(
		NewId(),
		NewId(),
		WEBSOCKET_EVENT_MATCHMAKING_START,
	)

	json := wse.ToJson()
	rwse := WebSocketEventFromJson(strings.NewReader(json))
	rjson := rwse.ToJson()

	if json != rjson {
		t.Fatal("JSONs do not match")
	}
}
