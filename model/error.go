// Copyright (c) 2016 David Lu
// See License.txt

package model

import (
	"encoding/json"
	"io"
	"strconv"
)

type Error struct {
	Id         string                 `json:"id"`
	Message    string                 `json:"message"`
	RequestId  string                 `json:"request_id"`
	StatusCode int                    `json:"status_code"`
	Where      string                 `json:"-"`
	params     map[string]interface{} `json:"-"`
}

func (er *Error) ToString() string {
	return er.Where + ": " + er.Message + ", " + strconv.Itoa(er.StatusCode)
}

func (er *Error) ToJson() string {
	b, err := json.Marshal(er)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ErrorFromJson(data io.Reader) *Error {
	decoder := json.NewDecoder(data)
	var er Error
	err := decoder.Decode(&er)
	if err == nil {
		return &er
	} else {
		return NewLocError("ErrorFromJson", "JSON decoding error", nil, err.Error())
	}
}

func NewLocError(where string, id string, params map[string]interface{}, details string) *Error {
	er := &Error{}
	er.Id = id
	er.params = params
	er.Message = id
	er.Where = where
	er.StatusCode = 500
	return er
}
