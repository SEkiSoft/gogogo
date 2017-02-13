// Copyright (c) 2016 SEkiSoft
// See License.txt

package store

import (
	"testing"

	"github.com/sekisoft/gogogo/model"
)

func TestGameStoreSave(t *testing.T) {
	Setup()

	game := model.Game{
		ID:       model.NewID(),
		IDBlack:  model.NewID(),
		IDWhite:  model.NewID(),
		Board:    "",
		NumLines: 9,
		Turn:     0,
		Finished: false,
	}

	if result := <-store.Game().Save(&game); result.Err != nil {
		t.Fatal("save game failed")
	}

	if result := <-store.Game().Delete(game.ID); result.Err != nil {
		t.Fatal("save game failed, delete")
	}
}

func TestGameStoreDelete(t *testing.T) {
	Setup()

	game := model.Game{
		ID:       model.NewID(),
		IDBlack:  model.NewID(),
		IDWhite:  model.NewID(),
		Board:    "",
		NumLines: 9,
		Turn:     0,
		Finished: false,
	}

	if result := <-store.Game().Save(&game); result.Err != nil {
		t.Fatal("save game failed")
	}

	if result := <-store.Game().Delete(game.ID); result.Err != nil {
		t.Fatal("save game failed, delete")
	}
}

func TestGameStoreUpdate(t *testing.T) {
	Setup()

	game := model.Game{
		ID:       model.NewID(),
		IDBlack:  model.NewID(),
		IDWhite:  model.NewID(),
		Board:    "",
		NumLines: 9,
		Turn:     0,
		Finished: false,
	}

	if result := <-store.Game().Save(&game); result.Err != nil {
		t.Fatal("get game failed, save")
	}

	game.Board = "01"

	if result := <-store.Game().Update(&game); result.Err != nil {
		t.Fatal("update game failed, store")
	} else if result.Data.(*model.Game).UpdateAt == result.Data.(*model.Game).CreateAt {
		t.Fatal("update game failed, update")
	}

	if result := <-store.Game().Delete(game.ID); result.Err != nil {
		t.Fatal("update game failed, delete")
	}
}

func TestGameStoreGet(t *testing.T) {
	Setup()

	game := model.Game{
		ID:       model.NewID(),
		IDBlack:  model.NewID(),
		IDWhite:  model.NewID(),
		Board:    "",
		NumLines: 9,
		Turn:     0,
		Finished: false,
	}

	if result := <-store.Game().Save(&game); result.Err != nil {
		t.Fatal("get game failed, save")
	}

	if result := <-store.Game().Get(game.ID); result.Err != nil {
		t.Fatal("get game failed, store")
	} else if result.Data.(*model.Game).IDBlack != game.IDBlack {
		t.Fatal("get game failed, not equal")
	}

	if result := <-store.Game().Delete(game.ID); result.Err != nil {
		t.Fatal("update game failed, delete")
	}
}

func TestGameStoreGetAll(t *testing.T) {
	Setup()

	game1 := model.Game{
		ID:       model.NewID(),
		IDBlack:  model.NewID(),
		IDWhite:  model.NewID(),
		Board:    "",
		NumLines: 9,
		Turn:     0,
		Finished: false,
	}

	game2 := model.Game{
		ID:       model.NewID(),
		IDBlack:  model.NewID(),
		IDWhite:  model.NewID(),
		Board:    "",
		NumLines: 9,
		Turn:     0,
		Finished: false,
	}

	if result := <-store.Game().Save(&game1); result.Err != nil {
		t.Fatal("get all games failed, save")
	}

	if result := <-store.Game().Save(&game2); result.Err != nil {
		t.Fatal("get all games failed, save")
	}

	if result := <-store.Game().GetAll(); result.Err != nil {
		t.Fatal("get all games failed, store")
	} else if len(result.Data.([]*model.Game)) != 2 {
		t.Fatal("get all games failed, wrong size")
	}

	if result := <-store.Game().Delete(game1.ID); result.Err != nil {
		t.Fatal("get all games failed, delete")
	}

	if result := <-store.Game().Delete(game2.ID); result.Err != nil {
		t.Fatal("get all games failed, delete")
	}
}
