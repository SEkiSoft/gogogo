// Copyright (c) 2016 sekisoft
// See License.txt

package store

import (
	"testing"

	"github.com/sekisoft/gogogo/model"
)

func TestGameStoreSave(t *testing.T) {
	Setup()

	game := model.Game{
		Id:       model.NewId(),
		IdBlack:  model.NewId(),
		IdWhite:  model.NewId(),
		Board:    "",
		NumLines: 9,
		Turn:     0,
		Finished: false,
	}

	if result := <-store.Game().Save(&game); result.Err != nil {
		t.Fatal("save game failed")
	}

	if result := <-store.Game().Delete(game.Id); result.Err != nil {
		t.Fatal("save game failed, delete")
	}
}

func TestGameStoreDelete(t *testing.T) {
	Setup()

	game := model.Game{
		Id:       model.NewId(),
		IdBlack:  model.NewId(),
		IdWhite:  model.NewId(),
		Board:    "",
		NumLines: 9,
		Turn:     0,
		Finished: false,
	}

	if result := <-store.Game().Save(&game); result.Err != nil {
		t.Fatal("save game failed")
	}

	if result := <-store.Game().Delete(game.Id); result.Err != nil {
		t.Fatal("save game failed, delete")
	}
}

func TestGameStoreUpdate(t *testing.T) {
	Setup()

	game := model.Game{
		Id:       model.NewId(),
		IdBlack:  model.NewId(),
		IdWhite:  model.NewId(),
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

	if result := <-store.Game().Delete(game.Id); result.Err != nil {
		t.Fatal("update game failed, delete")
	}
}

func TestGameStoreGet(t *testing.T) {
	Setup()

	game := model.Game{
		Id:       model.NewId(),
		IdBlack:  model.NewId(),
		IdWhite:  model.NewId(),
		Board:    "",
		NumLines: 9,
		Turn:     0,
		Finished: false,
	}

	if result := <-store.Game().Save(&game); result.Err != nil {
		t.Fatal("get game failed, save")
	}

	if result := <-store.Game().Get(game.Id); result.Err != nil {
		t.Fatal("get game failed, store")
	} else if result.Data.(*model.Game).IdBlack != game.IdBlack {
		t.Fatal("get game failed, not equal")
	}

	if result := <-store.Game().Delete(game.Id); result.Err != nil {
		t.Fatal("update game failed, delete")
	}
}

func TestGameStoreGetAll(t *testing.T) {
	Setup()

	game1 := model.Game{
		Id:       model.NewId(),
		IdBlack:  model.NewId(),
		IdWhite:  model.NewId(),
		Board:    "",
		NumLines: 9,
		Turn:     0,
		Finished: false,
	}

	game2 := model.Game{
		Id:       model.NewId(),
		IdBlack:  model.NewId(),
		IdWhite:  model.NewId(),
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

	if result := <-store.Game().Delete(game1.Id); result.Err != nil {
		t.Fatal("get all games failed, delete")
	}

	if result := <-store.Game().Delete(game2.Id); result.Err != nil {
		t.Fatal("get all games failed, delete")
	}
}
