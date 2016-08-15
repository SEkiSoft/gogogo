// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
)

type GameStats struct {
}

func (gs *GameStats) ToJson() string {
	s, err := json.Marshal(gs)
	if err != nil {
		return ""
	} else {
		return string(s)
	}
}

func (gs *GameStats) IsValid() *Error {
	return nil
}
