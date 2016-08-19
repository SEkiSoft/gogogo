// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
	"io"
	"strconv"
)

const (
	MAX_NUMLINES = 19
	MIN_NUMLINES = 13
	EMPTY_COLOR  = 0
)

type Game struct {
	Id       string     `json:"id"`
	IdBlack  string     `json:"id_black"`
	IdWhite  string     `json:"id_white"`
	Board    string     `json:"board"`
	KoPoint  Coordinate //TODO: Set Ko Point upon initialization
	NumLines uint       `json:"numlines"`
	Turn     uint       `json:"turn"`
	CreateAt int64      `json:"create_at"`
	UpdateAt int64      `json:"update_at"`
	DeleteAt int64      `json:"delete_at"`
	Finished bool       `json:"finished"`
}

type Coordinate struct {
	X uint
	Y uint
}

func (g *Game) ToJson() string {
	json, err := json.Marshal(g)
	if err != nil {
		return ""
	}

	return string(json)
}

func GameFromJson(data io.Reader) *Game {
	decoder := json.NewDecoder(data)
	var g Game
	err := decoder.Decode(&g)
	if err == nil {
		return &g
	}

	return nil
}

func GamesToJson(m []*Game) string {
	json, err := json.Marshal(m)
	if err == nil {
		return string(json)
	}

	return "[]"
}

func (g *Game) IsValid() *Error {
	if g.NumLines < MIN_NUMLINES || g.NumLines > MAX_NUMLINES {
		return NewLocError("Game.IsValid", "Too many/few lines", nil, "")
	} else if len(g.Board) != int(g.NumLines*g.NumLines) {
		return NewLocError("Game.IsValid", "Board does not match line number", nil, "")
	}
	return nil
}

func (g *Game) PreSave() {
	if g.Id == "" {
		g.Id = NewId()
		g.IdBlack = NewId()
		g.IdWhite = NewId()
	}
	g.CreateAt = GetMillis()
	g.UpdateAt = g.CreateAt
	g.Finished = false
}

func (g *Game) PreUpdate() {
	g.UpdateAt = GetMillis()
}

func (g *Game) GetStats() *GameStats {
	var gs GameStats

	// TODO GetGameStats

	return &gs
}

func (g *Game) GetPieceColor(x, y uint) (int, *Error) {
	if x < g.NumLines && y < g.NumLines {
		piece, _ := strconv.ParseInt(string(g.Board[y*g.NumLines+x]), 10, 0)
		return int(piece), nil
	}
	return -1, NewLocError("Game.GetPieceColor", "row/col out of range", nil, "")
}

func (g *Game) GetColor(p *Coordinate) int {
	color, _ := g.GetPieceColor(p.X, p.Y)
	return color
}

func GetOppositeColor(color int) int {
	return (3 - color)
}

func (g *Game) SetPieceColor(color int, x, y uint) *Error {
	if x < g.NumLines && y < g.NumLines {
		runeBoard := []rune(g.Board)
		runeBoard[y*g.NumLines+x] = rune(color)
		g.Board = string(runeBoard)
		return nil
	}
	return NewLocError("Game.SetPieceColor", "row/col out of range", nil, "")
}
