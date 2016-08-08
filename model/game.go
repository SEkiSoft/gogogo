// Copyright (c) 2016 David Lu
// See License.txt

package model

import (
	"encoding/json"
	"io"
)

const (
	MAX_NUMLINES = 19
	MIN_NUMLINES = 13
)

type Game struct {
	Id       string `json:"id"`
	IdBlack  string `json:"id_black"`
	IdWhite  string `json:"id_white"`
	Board    string `json:"board"`
	NumLines uint   `json:"numlines"`
	Turn     uint   `json:"turn"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
	DeleteAt int64  `json:"delete_at"`
	Finished bool   `json:"finished"`
}

func (g *Game) ToJson() string {
	s, err := json.Marshal(g)
	if err != nil {
		return ""
	} else {
		return string(s)
	}
}

func GameFromJson(data io.Reader) *Game {
	decoder := json.NewDecoder(data)
	var g Game
	err := decoder.Decode(&g)
	if err == nil {
		return &g
	} else {
		return nil
	}
}

func (g *Game) IsValid() *Error {
	if g.NumLines < MIN_NUMLINES || g.NumLines > MAX_NUMLINES {
		return NewLocError("Game.IsValid", "Too many/few lines", nil, "")
	} else {
		return nil
	}
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

	return gs
}
