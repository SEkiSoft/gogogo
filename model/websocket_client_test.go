// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import "testing"

func TestNewWebSocketClient(t *testing.T) {
	wsc, err := NewWebSocketClient("localhost", NewId())

	if err != nil {
		t.Fatal(err.Message)
	}

	if wsc.Connection == nil {
		t.Fatal("nil connection")
	}

	wsc.Close()
}

func TestWebSocketClientConnect(t *testing.T) {
	wsc, err := NewWebSocketClient("localhost", NewId())

	err = wsc.Connect()

	if err != nil {
		t.Fatal(err.Message)
	}

	wsc.Close()
}
