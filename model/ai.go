// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
	"io"
	"net/http"
)

type Ai struct {
}

func (ai *Ai) ToJson() string {
	b, err := json.Marshal(ai)
	if err != nil {
		return ""
	}

	return string(b)
}

func AiFromJson(data io.Reader) (*Ai, *AppError) {
	decoder := json.NewDecoder(data)
	var ai Ai
	err := decoder.Decode(&ai)
	if err == nil {
		return &ai, nil
	}

	return nil, NewAppError("AiFromJson", "JSON decoding error", http.StatusBadRequest)
}
