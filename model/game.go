package model

import (
	"encoding/json"
	"io"
	"unicode/utf8"
	"bytes"
	"encoding/base32"
)

const (
	ENCODING = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")
	ID_LENGTH = 26
	MAX_NUMLINES = 19
	MIN_NUMLINES = 13
)

type Game struct {
	Board 		[][]uint `json:"board"`
	NumLines 	uint	 `json:"numlines"`
	Turn 		bool	 `json:"turn"`
	Id 			string	 `json:"id"`
	IdBlack		string   `json:"id_black"`
	IdWhite		string   `json:"id_white"`
	CreateAt	int64	 `json:"create_at"`
	UpdateAt    int64  	 `json:"update_at"`
	DeleteAt    int64  	 `json:"delete_at"`
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

func (g *Game) IsValid() bool {
	if NumLines < MIN_NUMLINES || NumLines > MAX_NUMLINES {
		return false;
	} else {
		return true;
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
}

func (g *Game) PreUpdate() {
	g.UpdateAt = GetMillis()
}

func NewID() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(ID_LENGTH)
	return b.String()
}