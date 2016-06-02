// Copyright (c) 2016 David Lu
// See License.txt

package model

import (
	"encoding/json"
	"io"
)

type Move struct {
	PlayerId string `json:"player_id"`
	GameId   string `json:"game_id"`
	Id       string `json:"id"`
	X        uint   `json:"move_x"`
	Y        uint   `json:"move_y"`
	CreateAt int64  `json:"create_at"`
}

func (m *Move) ToJson() string {
	s, err := json.Marshal(m)
	if err != nil {
		return ""
	} else {
		return string(s)
	}
}

func (m *Move) PreSave() {
	m.CreateAt = GetMillis()

	if m.Id == "" {
		m.Id = NewId()
	}
}

func MoveFromJson(data io.Reader) *Move {
	decoder := json.NewDecoder(data)
	var m Move
	err := decoder.Decode(&m)
	if err == nil {
		return &m
	} else {
		return nil
	}
}
