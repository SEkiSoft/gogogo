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
	MIN_NUMLINES = 5
	WHITE        = 1
	BLACK        = 2
)

type Game struct {
	ID       string `json:"id"`
	IDBlack  string `json:"id_black"`
	IDWhite  string `json:"id_white"`
	Board    string `json:"board"`
	NumLines uint   `json:"numlines"`
	Turn     uint   `json:"turn"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
	DeleteAt int64  `json:"delete_at"`
	Finished bool   `json:"finished"`
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

func GamesToJson(g []*Game) string {
	json, err := json.Marshal(g)
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
	if g.ID == "" {
		g.ID = NewID()
		g.IDBlack = NewID()
		g.IDWhite = NewID()
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

func (g *Game) HasPlayer(playerID string) bool {
	return g.IDBlack == playerID || g.IDWhite == playerID
}

func (g *Game) GetBoardPiece(x, y uint) (int, *Error) {
	if x < g.NumLines && y < g.NumLines {
		piece, _ := strconv.ParseInt(string(g.Board[y*g.NumLines+x]), 10, 0)
		return int(piece), nil
	}
	return -1, NewLocError("Game.GetBoardPiece", "row/col out of range", nil, "")
}
