// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"strings"
	"testing"
)

func TestGameToJson(t *testing.T) {
	game := Game{
		ID:       NewID(),
		IDBlack:  NewID(),
		IDWhite:  NewID(),
		Board:    "",
		NumLines: MAX_NUMLINES,
		Turn:     0,
		CreateAt: 0,
		UpdateAt: 0,
		DeleteAt: 0,
		Finished: false,
	}

	json := game.ToJson()
	rgame := GameFromJson(strings.NewReader(json))

	if rgame.ID != game.ID {
		t.Fatal("ids do not match")
	}
}

func TestGameFromJson(t *testing.T) {
	game := Game{
		ID:       NewID(),
		IDBlack:  NewID(),
		IDWhite:  NewID(),
		Board:    "",
		NumLines: MAX_NUMLINES,
		Turn:     0,
		CreateAt: 0,
		UpdateAt: 0,
		DeleteAt: 0,
		Finished: false,
	}

	json := game.ToJson()
	rgame := GameFromJson(strings.NewReader(json))
	rjson := rgame.ToJson()

	if rjson != json {
		t.Fatal("jsons do not match")
	}
}

func TestGamesToJson(t *testing.T) {
	game1 := Game{
		ID:       NewID(),
		IDBlack:  NewID(),
		IDWhite:  NewID(),
		Board:    "",
		NumLines: MAX_NUMLINES,
		Turn:     0,
		CreateAt: 0,
		UpdateAt: 0,
		DeleteAt: 0,
		Finished: false,
	}

	game2 := Game{
		ID:       NewID(),
		IDBlack:  NewID(),
		IDWhite:  NewID(),
		Board:    "",
		NumLines: MAX_NUMLINES,
		Turn:     0,
		CreateAt: 0,
		UpdateAt: 0,
		DeleteAt: 0,
		Finished: false,
	}

	games := []*Game{&game1, &game2}

	json := GamesToJson(games)

	if json == "[]" {
		t.Fatal("games to json failed")
	}
}

func TestGameIsValid(t *testing.T) {
	// TODO by ian952
}

func TestGamePreSave(t *testing.T) {
	game := Game{
		ID:       NewID(),
		IDBlack:  NewID(),
		IDWhite:  NewID(),
		Board:    "",
		NumLines: MAX_NUMLINES,
		Turn:     0,
		CreateAt: 0,
		UpdateAt: 0,
		DeleteAt: 0,
		Finished: false,
	}

	game.PreSave()

	if game.CreateAt == 0 {
		t.Fatal("create at did not update")
	}
}

func TestGamePreUpdate(t *testing.T) {
	game := Game{
		ID:       NewID(),
		IDBlack:  NewID(),
		IDWhite:  NewID(),
		Board:    "",
		NumLines: MAX_NUMLINES,
		Turn:     0,
		CreateAt: 0,
		UpdateAt: 0,
		DeleteAt: 0,
		Finished: false,
	}

	game.PreUpdate()

	if game.UpdateAt == 0 {
		t.Fatal("update at did not update")
	}
}

func TestGameGetStats(t *testing.T) {
	// TODO
}

func TestGameHasPlayer(t *testing.T) {
	game := Game{
		ID:       NewID(),
		IDBlack:  NewID(),
		IDWhite:  NewID(),
		Board:    "",
		NumLines: MAX_NUMLINES,
		Turn:     0,
		CreateAt: 0,
		UpdateAt: 0,
		DeleteAt: 0,
		Finished: false,
	}

	if game.HasPlayer(game.IDBlack) == false {
		t.Fatal("should be true")
	}

	if game.HasPlayer(game.IDWhite) == false {
		t.Fatal("should be true")
	}

	if game.HasPlayer(NewID()) == true {
		t.Fatal("should be false")
	}
}

func TestGameGetBoardPiece(t *testing.T) {
	// TODO by ian952
}
