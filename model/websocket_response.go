// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
	"io"
)

type WebSocketResponse struct {
	Status        string                 `json:"status"`
	ReplySequence int64                  `json:"reply_sequence,omitempty"`
	Data          map[string]interface{} `json:"data,omitempty"`
	Error         *AppError              `json:"error,omitempty"`
}

func NewWebSocketResponse(status string, replySequence int64, data map[string]interface{}) *WebSocketResponse {
	return &WebSocketResponse{
		Status:        status,
		ReplySequence: replySequence,
		Data:          data,
	}
}

func (w *WebSocketResponse) Add(key string, value interface{}) {
	w.Data[key] = value
}

func (w *WebSocketResponse) IsValid() bool {
	return w.Status != ""
}

func (w *WebSocketResponse) ToJson() string {
	json, err := json.Marshal(w)
	if err != nil {
		return ""
	}

	return string(json)
}

func WebSocketResponseFromJson(data io.Reader) *WebSocketResponse {
	decoder := json.NewDecoder(data)
	var w WebSocketResponse
	err := decoder.Decode(&w)

	if err != nil {
		return nil
	}

	return &w
}
