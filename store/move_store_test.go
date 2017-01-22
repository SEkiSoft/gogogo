// Copyright (c) 2016 SEkiSoft
// See License.txt

package store

import (
	"testing"

	"github.com/sekisoft/gogogo/model"
)

func TestMoveStoreSave(t *testing.T) {
	Setup()

	move := model.Move{
		ID:       model.NewID(),
		PlayerID: model.NewID(),
		GameID:   model.NewID(),
		X:        1,
		Y:        1,
	}

	if result := <-store.Move().Save(&move); result.Err != nil {
		t.Fatal("save move failed")
	}

	if result := <-store.Move().Delete(move.ID); result.Err != nil {
		t.Fatal("save move failed, delete")
	}
}

func TestMoveStoreDelete(t *testing.T) {
	Setup()

	move := model.Move{
		ID:       model.NewID(),
		PlayerID: model.NewID(),
		GameID:   model.NewID(),
		X:        1,
		Y:        1,
	}

	if result := <-store.Move().Save(&move); result.Err != nil {
		t.Fatal("delete move failed, save")
	}

	if result := <-store.Move().Delete(move.ID); result.Err != nil {
		t.Fatal("delete move failed")
	}
}

func TestMoveStoreGet(t *testing.T) {
	Setup()

	move := model.Move{
		ID:       model.NewID(),
		PlayerID: model.NewID(),
		GameID:   model.NewID(),
		X:        1,
		Y:        1,
	}

	if result := <-store.Move().Save(&move); result.Err != nil {
		t.Fatal("get move failed, save")
	}

	if result := <-store.Move().Get(move.ID); result.Err != nil {
		t.Fatal("get move failed, store")
	} else if result.Data.(*model.Move).GameID != move.GameID {
		t.Fatal("get move failed, not equal")
	}

	if result := <-store.Move().Delete(move.ID); result.Err != nil {
		t.Fatal("get move failed, delete")
	}
}

func TestMoveStoreGetByGame(t *testing.T) {
	Setup()

	move := model.Move{
		ID:       model.NewID(),
		PlayerID: model.NewID(),
		GameID:   model.NewID(),
		X:        1,
		Y:        1,
	}

	if result := <-store.Move().Save(&move); result.Err != nil {
		t.Fatal("get move failed, save")
	}

	if result := <-store.Move().GetByGame(move.GameID); result.Err != nil {
		t.Fatal("get move failed, store")
	} else if result.Data.(*model.Move).GameID != move.GameID {
		t.Fatal("get move failed, not equal")
	}

	if result := <-store.Move().Delete(move.ID); result.Err != nil {
		t.Fatal("get move failed, delete")
	}
}

func TestMoveStoreGetByPlayer(t *testing.T) {
	Setup()

	move := model.Move{
		ID:       model.NewID(),
		PlayerID: model.NewID(),
		GameID:   model.NewID(),
		X:        1,
		Y:        1,
	}

	if result := <-store.Move().Save(&move); result.Err != nil {
		t.Fatal("get move failed, save")
	}

	if result := <-store.Move().GetByPlayer(move.PlayerID); result.Err != nil {
		t.Fatal("get move failed, store")
	} else if result.Data.(*model.Move).GameID != move.GameID {
		t.Fatal("get move failed, not equal")
	}

	if result := <-store.Move().Delete(move.ID); result.Err != nil {
		t.Fatal("get move failed, delete")
	}
}

func TestMoveStoreGetAll(t *testing.T) {
	Setup()

	move1 := model.Move{
		ID:       model.NewID(),
		PlayerID: model.NewID(),
		GameID:   model.NewID(),
		X:        1,
		Y:        1,
	}

	move2 := model.Move{
		ID:       model.NewID(),
		PlayerID: model.NewID(),
		GameID:   model.NewID(),
		X:        1,
		Y:        1,
	}

	if result := <-store.Move().Save(&move1); result.Err != nil {
		t.Fatal("get all moves failed, save")
	}

	if result := <-store.Move().Save(&move2); result.Err != nil {
		t.Fatal("get all moves failed, save")
	}

	if result := <-store.Move().GetAll(); result.Err != nil {
		t.Fatal("get all moves failed, store")
	} else if len(result.Data.([]*model.Move)) != 2 {
		t.Fatal("get all moves failed, wrong size")
	}

	if result := <-store.Move().Delete(move1.ID); result.Err != nil {
		t.Fatal("get all moves failed, delete")
	}

	if result := <-store.Move().Delete(move2.ID); result.Err != nil {
		t.Fatal("get all moves failed, delete")
	}
}

func TestMoveStoreGetTotalMovesCount(t *testing.T) {
	Setup()

	move1 := model.Move{
		ID:       model.NewID(),
		PlayerID: model.NewID(),
		GameID:   model.NewID(),
		X:        1,
		Y:        1,
	}

	move2 := model.Move{
		ID:       model.NewID(),
		PlayerID: model.NewID(),
		GameID:   model.NewID(),
		X:        1,
		Y:        1,
	}

	if result := <-store.Move().Save(&move1); result.Err != nil {
		t.Fatal("get all moves count failed, save")
	}

	if result := <-store.Move().Save(&move2); result.Err != nil {
		t.Fatal("get all moves count failed, save")
	}

	if result := <-store.Move().GetTotalMovesCount(); result.Err != nil {
		t.Fatal("get all moves count failed, store")
	} else if result.Data.(int64) != 2 {
		t.Fatal("get all moves count failed, wrong size")
	}

	if result := <-store.Move().Delete(move1.ID); result.Err != nil {
		t.Fatal("get all moves count failed, delete")
	}

	if result := <-store.Move().Delete(move2.ID); result.Err != nil {
		t.Fatal("get all moves count failed, delete")
	}
}
