// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"strings"
	"testing"
)

func TestMoveToJson(t *testing.T) {
	move := Move{
		ID:       NewID(),
		GameID:   NewID(),
		PlayerID: NewID(),
		CreateAt: GetMillis(),
		X:        0,
		Y:        0,
	}

	json := move.ToJson()
	rMove := MoveFromJson(strings.NewReader(json))

	if rMove.ID != move.ID {
		t.Fatal("IDs do not match")
	}
}

func TestMoveFromJson(t *testing.T) {
	move := Move{
		ID:       NewID(),
		GameID:   NewID(),
		PlayerID: NewID(),
		CreateAt: GetMillis(),
		X:        0,
		Y:        0,
	}

	json := move.ToJson()
	rMove := MoveFromJson(strings.NewReader(json))
	rJson := rMove.ToJson()

	if json != rJson {
		t.Fatal("JSONs do not match")
	}
}

func TestMovesToJson(t *testing.T) {
	move0 := Move{
		ID:       NewID(),
		GameID:   NewID(),
		PlayerID: NewID(),
		CreateAt: GetMillis(),
		X:        0,
		Y:        0,
	}

	move1 := Move{
		ID:       NewID(),
		GameID:   NewID(),
		PlayerID: NewID(),
		CreateAt: GetMillis(),
		X:        1,
		Y:        1,
	}

	moves := []*Move{&move0, &move1}

	json := MovesToJson(moves)
	rMoves := MovesFromJson(strings.NewReader(json))

	if rMoves[0].ID != move0.ID || rMoves[1].ID != move1.ID {
		t.Fatal("IDs do not match")
	}
}

func TestMovesFromJson(t *testing.T) {
	move0 := Move{
		ID:       NewID(),
		GameID:   NewID(),
		PlayerID: NewID(),
		CreateAt: GetMillis(),
		X:        0,
		Y:        0,
	}

	move1 := Move{
		ID:       NewID(),
		GameID:   NewID(),
		PlayerID: NewID(),
		CreateAt: GetMillis(),
		X:        1,
		Y:        1,
	}

	moves := []*Move{&move0, &move1}

	json := MovesToJson(moves)
	rMoves := MovesFromJson(strings.NewReader(json))
	rJson := MovesToJson(rMoves)

	if json != rJson {
		t.Fatal("JSONs do not match")
	}
}

func TestMovePreSave(t *testing.T) {
	move := Move{
		PlayerID: NewID(),
		GameID:   NewID(),
		ID:       "",
		CreateAt: 0,
	}

	move.PreSave()

	if len(move.ID) == 0 {
		t.Fatal("ID should not be empty")
	}

	if move.CreateAt == 0 {
		t.Fatal("CreateAt should not be 0")
	}
}

func TestMoveIsValid(t *testing.T) {
	// TODO
	// Awaiting other issue
}
