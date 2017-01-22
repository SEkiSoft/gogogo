// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
	"io"
)

type Move struct {
	PlayerID string `json:"player_id"`
	GameID   string `json:"game_id"`
	ID       string `json:"id"`
	X        uint   `json:"move_x"`
	Y        uint   `json:"move_y"`
	CreateAt int64  `json:"create_at"`
}

func (m *Move) ToJson() string {
	json, err := json.Marshal(m)
	if err != nil {
		return ""
	}

	return string(json)
}

func (m *Move) PreSave() {
	m.CreateAt = GetMillis()

	if m.ID == "" {
		m.ID = NewID()
	}
}

func (m *Move) IsValid(game *Game) *Error {
	currentPiece, err := game.GetBoardPiece(m.X, m.Y)

	if err != nil {
		return err
	} else if currentPiece != 0 {
		return NewLocError("Move.IsValid", "Spot is occupied", nil, "")
	}

	return nil
}

func MoveFromJson(data io.Reader) *Move {
	decoder := json.NewDecoder(data)
	var m Move
	err := decoder.Decode(&m)
	if err == nil {
		return &m
	}
	return nil
}

func MovesToJson(m []*Move) string {
	json, err := json.Marshal(m)
	if err == nil {
		return string(json)
	}

	return "[]"
}

func MovesFromJson(data io.Reader) []*Move {
	decoder := json.NewDecoder(data)
	var o []*Move
	err := decoder.Decode(&o)
	if err == nil {
		return o
	}

	return nil
}
