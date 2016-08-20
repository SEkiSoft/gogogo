// Copyright (c) 2016 SEkiSoft
// See License.txt
//TODO: everything should be game object
//ProcessMoves
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
	json, err := json.Marshal(m)
	if err != nil {
		return ""
	}

	return string(json)
}

func (m *Move) PreSave() {
	m.CreateAt = GetMillis()

	if m.Id == "" {
		m.Id = NewId()
	}
}

func (m *Move) IsValid(game *Game) *Error {
	p := Coordinate{m.X, m.Y}
	pColor := game.GetColor(&p)

	if pColor != EMPTY_COLOR {
		return NewLocError("Move.IsValid", "Spot is occupied", nil, "")
	}

	if p == game.KoPoint {
		return NewLocError("Move.IsValid", "Spot is Ko Point", nil, "")
	}

	suicide := true
	neighbors := *game.getNeighbors(&p)
	for _, neighbor := range neighbors {
		if neighborC := game.GetColor(&neighbor); neighborC == EMPTY_COLOR {
			suicide = false
		} else if neighborC == pColor {
			if !game.inAtari(&neighbor) {
				suicide = false
			}
		} else if neighborC == GetOppositeColor(pColor) {
			if game.inAtari(&neighbor) {
				suicide = false
			}
		}
	}

	if suicide {
		return NewLocError("Move.IsValid", "Suicide Move", nil, "")
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
