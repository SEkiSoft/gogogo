// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
	"io"
)

const (
	TOKEN_DURATION = 1000 * 60 * 60 * 24 * 30 // 30 days
)

type Token struct {
	Id             string `json:"id"`
	CreateAt       int64  `json:"create_at"`
	ExpiresAt      int64  `json:"expires_at"`
	LastActivityAt int64  `json:"last_activity_at"`
	PlayerId       string `json:"player_id"`
	DeviceId       string `json:"device_id"`
	Roles          string `json:"roles"`
	IsOAuth        bool   `json:"is_oauth"`
}

func (t *Token) ToJson() string {
	json, err := json.Marshal(t)
	if err != nil {
		return ""
	}

	return string(json)
}

func TokenFromJson(data io.Reader) *Token {
	decoder := json.NewDecoder(data)
	var t Token
	err := decoder.Decode(&t)

	if err == nil {
		return &t
	}

	return nil
}

func (t *Token) PreSave() {
	t.Id = NewId()
	t.CreateAt = GetMillis()
	t.LastActivityAt = t.CreateAt
	t.ExpiresAt = t.CreateAt + TOKEN_DURATION
}

func (t *Token) PreUpdate() {
	t.LastActivityAt = GetMillis()
}

func (t *Token) IsExpired() bool {
	if t.ExpiresAt <= 0 {
		return false
	}

	return GetMillis() > t.ExpiresAt
}

func TokensToJson(tokens []*Token) string {
	if json, err := json.Marshal(tokens); err == nil {
		return string(json)
	}

	return "[]"
}
