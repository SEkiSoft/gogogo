// Copyright (c) 2016 sekisoft
// See License.txt

package model

import (
	"encoding/json"
	"io"
)

type GameStats struct {
}

func (gs *GameStats) ToJson() string {
	s, err := json.Marshal(gs)
	if err != nil {
		return ""
	}

	return string(s)
}

func GameStatsFromJson(data io.Reader) *GameStats {
	decoder := json.NewDecoder(data)
	var gs GameStats
	err := decoder.Decode(&gs)
	if err == nil {
		return &gs
	}

	return nil
}

func GameStatssToJson(gs []*GameStats) string {
	json, err := json.Marshal(gs)
	if err == nil {
		return string(json)
	}

	return "[]"
}

func GameStatsToJson(gs *Game) string {
	b, err := json.Marshal(gs)
	if err != nil {
		return ""
	}

	return string(b)
}

func (gs *GameStats) IsValid() *Error {
	return nil
}
