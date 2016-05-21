package model

import (
	"encoding/json"
	goi18n "github.com/nicksnyder/go-i18n/i18n"
	"io"
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
	return er.Where + ": " + er.Message + ", " + er.StatusCode
}

func (er *Error) Translate(T goi18n.TranslateFunc) {
	if er.params == nil {
		er.Message = T(er.Id)
	} else {
		er.Message = T(er.Id, er.params)
	}
}

func (er *Error) ToJson() string {
	b, err := json.Marshal(er)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ErrorFromJson(data io.reader) *Error {
	decoder := json.NewDecoder(data)
	var er Error
	err := decoder.Decode(&er)
	if err == nil {
		return &er
	} else {
		return NewLocError("ErrorFromJson", "model.error.decode_json.error", nil, err.Error())
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
