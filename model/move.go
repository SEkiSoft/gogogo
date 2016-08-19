// Copyright (c) 2016 SEkiSoft
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
	p = Coordinate{m.X, m.Y}
	currentPiece, err := game.GetColor(p)

	if err != nil {
		return err
	}

	if currentPiece != EMPTY_COLOR {
		return NewLocError("Move.IsValid", "Spot is occupied", nil, "")
	}

	if p == game.KoPoint {
		return NewLocError(Move.IsValid, "Spot is Ko Point", nil, "")
	}
}

func (p *Coordinate) getNeighbors(game *Game) (neighbors *[]Coordinate) {
	var neighbors []int

	if p.X > 0 {
		neighbors.append(Coordinate{X: p.X - 1, Y: p.Y})
	}

	if p.X < game.NumLines-1 {
		neighbors.append(Coordinate{X: p.X + 1, Y: p.Y})
	}

	if p.Y > 0 {
		neighbors.append(Coordinate{X: p.X, Y: p.Y - 1})
	}

	if p.Y < game.NumLines-1 {
		neighbors.append(Coordinate{X: p.X, Y: p.Y + 1})
	}

	return
}

func (p *Coordinate) getLiberties(game *Game) (liberties *[]Coordinate) {
	var liberties []int

	myColor = game.getColor(p)
	fillColor = GetOppositeColor(myColor)
	game.SetPieceColor(fillColor, p.X, p.Y)

	neighbors = p.getNeighbors(game)

	for neighbor := range neighbors {
		if game.getColor(neighbor) == myColor {
			liberties.append(neighbor.getLiberties(game)...)
		} else if game.getColor(neighbor) == EMPTY_COLOR {
			liberties.append(neighbor)
		}
	}

	return liberties
}

func (p *Coordinate) isAtari(game *Game) bool {
	//this should make a copy
	game_copy := *game

	return len(p.getLiberties(game_copy)) == 1
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
