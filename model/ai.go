// Copyright (c) 2016 David Lu
// See License.txt

package model

import (
	"encoding/json"
	"io"
	"strconv"
)

type Ai struct {
}

func (ai *Ai) ToJson() string {
	b, err := json.Marshal(ai)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func AiFromJson(data io.Reader) *Ai {
	decoder := json.NewDecoder(data)
	var ai Ai
	err := decoder.Decode(&ai)
	if err == nil {
		return &ai
	} else {
		return NewLocError("AiFromJson", "JSON decoding error", nil, err.Error())
	}
}
