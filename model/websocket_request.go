// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
	"io"
)

type WebSocketRequest struct {
	Sequence int64                  `json:"sequence"`
	Action   string                 `json:"action"`
	Data     map[string]interface{} `json:"data"`
<<<<<<< bf935483bdee38a155347a0daeeb5db6eaef9887
	Token    *Token                 `json:"-"`
=======
	Token    Token                  `json:"-"`
>>>>>>> added websockets model
}

func (w *WebSocketRequest) ToJson() string {
	json, err := json.Marshal(w)
	if err != nil {
		return ""
	}

	return string(json)
}

func WebSocketRequestFromJson(data io.Reader) *WebSocketRequest {
	decoder := json.NewDecoder(data)
	var w WebSocketRequest
	err := decoder.Decode(&w)

	if err != nil {
		return nil
	}

	return &w
}
