// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"strings"
	"testing"
)

func TestNewWebSocketResponse(t *testing.T) {
	wsr := NewWebSocketResponse(
		"status1",
		1,
		make(map[string]interface{}),
	)

	if wsr == nil {
		t.Fatal("WebSocketResponse not created")
	}

	if wsr.Status != "status1" {
		t.Fatal("WebSocketResponse creation failed")
	}

	if wsr.ReplySequence != 1 {
		t.Fatal("WebSocketResponse creation failed")
	}
}

func TestWebSocketResponseAdd(t *testing.T) {
	wsr := NewWebSocketResponse(
		"status1",
		1,
		make(map[string]interface{}),
	)

	wsr.Add("test", "testing")

	if wsr.Data["test"] != "testing" {
		t.Fatal("Add key to data failed")
	}
}

func TestWebSocketResponseIsValid(t *testing.T) {
	wsr := NewWebSocketResponse(
		"status1",
		1,
		make(map[string]interface{}),
	)

	if !wsr.IsValid() {
		t.Fatal("should be true")
	}

	wsr.Status = ""

	if wsr.IsValid() {
		t.Fatal("should be false")
	}
}

func TestWebSocketResponseToJson(t *testing.T) {
	wsr := NewWebSocketResponse(
		"status1",
		1,
		make(map[string]interface{}),
	)

	json := wsr.ToJson()
	rwsr := WebSocketResponseFromJson(strings.NewReader(json))

	if rwsr.Status != wsr.Status {
		t.Fatal("Player Ids do not match")
	}
}

func TestWebSocketResponseFromJson(t *testing.T) {
	wsr := NewWebSocketResponse(
		"status1",
		1,
		make(map[string]interface{}),
	)

	json := wsr.ToJson()
	rwsr := WebSocketResponseFromJson(strings.NewReader(json))
	rjson := rwsr.ToJson()

	if json != rjson {
		t.Fatal("JSONs do not match")
	}
}
