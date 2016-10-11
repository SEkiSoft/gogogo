// Copyright (c) 2016 sekisoft
// See License.txt

package model

import (
	"strings"
	"testing"
)

func TestMoveToJson(t *testing.T) {
	move := Move{
		Id:       NewId(),
		GameId:   NewId(),
		PlayerId: NewId(),
		CreateAt: GetMillis(),
		X:        0,
		Y:        0,
	}

	json := move.ToJson()
	rMove := MoveFromJson(strings.NewReader(json))

	if rMove.Id != move.Id {
		t.Fatal("Ids do not match")
	}
}

func TestMoveFromJson(t *testing.T) {
	move := Move{
		Id:       NewId(),
		GameId:   NewId(),
		PlayerId: NewId(),
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
		Id:       NewId(),
		GameId:   NewId(),
		PlayerId: NewId(),
		CreateAt: GetMillis(),
		X:        0,
		Y:        0,
	}

	move1 := Move{
		Id:       NewId(),
		GameId:   NewId(),
		PlayerId: NewId(),
		CreateAt: GetMillis(),
		X:        1,
		Y:        1,
	}

	moves := []*Move{&move0, &move1}

	json := MovesToJson(moves)
	rMoves := MovesFromJson(strings.NewReader(json))

	if rMoves[0].Id != move0.Id || rMoves[1].Id != move1.Id {
		t.Fatal("Ids do not match")
	}
}

func TestMovesFromJson(t *testing.T) {
	move0 := Move{
		Id:       NewId(),
		GameId:   NewId(),
		PlayerId: NewId(),
		CreateAt: GetMillis(),
		X:        0,
		Y:        0,
	}

	move1 := Move{
		Id:       NewId(),
		GameId:   NewId(),
		PlayerId: NewId(),
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
		PlayerId: NewId(),
		GameId:   NewId(),
		Id:       "",
		CreateAt: 0,
	}

	move.PreSave()

	if len(move.Id) == 0 {
		t.Fatal("Id should not be empty")
	}

	if move.CreateAt == 0 {
		t.Fatal("CreateAt should not be 0")
	}
}

func TestMoveIsValid(t *testing.T) {
	// TODO
	// Awaiting other issue
}
