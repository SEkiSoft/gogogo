// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

type WebSocketMessage interface {
	ToJson() string
	IsValid() bool
}

func NewWebSocketError(seqReply int64, err *Error) *WebSocketResponse {
	return &WebSocketResponse{Status: STATUS_FAIL, ReplySequence: seqReply, Error: err}
}
