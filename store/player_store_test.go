// Copyright (c) 2016 SEkiSoft
// See License.txt

package store

import (
	"testing"

	"github.com/sekisoft/gogogo/model"
)

func TestPlayerStoreSave(t *testing.T) {
	Setup()

	player := model.Player{
		Id:         model.NewId(),
		Username:   "aaaaa",
		Password:   "aaaaa",
		Email:      "a@a.com",
		AllowStats: true,
		Locale:     "",
	}

	if result := <-store.Player().Save(&player); result.Err != nil {
		t.Fatal("save player failed")
	}

	if result := <-store.Player().Delete(player.Id); result.Err != nil {
		t.Fatal("save player failed, delete")
	}
}

func TestPlayerStoreDelete(t *testing.T) {
	Setup()

	player := model.Player{
		Id:         model.NewId(),
		Username:   "aaaaa",
		Password:   "aaaaa",
		Email:      "a@a.com",
		AllowStats: true,
		Locale:     "",
	}

	if result := <-store.Player().Save(&player); result.Err != nil {
		t.Fatal("delete player failed")
	}

	if result := <-store.Player().Delete(player.Id); result.Err != nil {
		t.Fatal("delete player failed, delete")
	}
}

func TestPlayerStoreUpdate(t *testing.T) {
	Setup()

	player := model.Player{
		Id:         model.NewId(),
		Username:   "aaaaa",
		Password:   "aaaaa",
		Email:      "a@a.com",
		AllowStats: true,
		Locale:     "",
	}

	if result := <-store.Player().Save(&player); result.Err != nil {
		t.Fatal("update player failed, save")
	}

	player.Email = "b@b.com"

	if result := <-store.Player().Update(&player); result.Err != nil {
		t.Fatal("update player failed, store")
	} else if result.Data.(*model.Player).UpdateAt == result.Data.(*model.Player).CreateAt {
		t.Fatal("update player failed, not equal")
	}

	if result := <-store.Player().Delete(player.Id); result.Err != nil {
		t.Fatal("update player failed, delete")
	}
}

func TestPlayerStoreGet(t *testing.T) {
	Setup()

	player := model.Player{
		Id:         model.NewId(),
		Username:   "aaaaa",
		Password:   "aaaaa",
		Email:      "a@a.com",
		AllowStats: true,
		Locale:     "",
	}

	if result := <-store.Player().Save(&player); result.Err != nil {
		t.Fatal("get player failed, save")
	}

	if result := <-store.Player().Get(player.Id); result.Err != nil {
		t.Fatal("get player failed, store")
	} else if result.Data.(*model.Player).Id == player.Id {
		t.Fatal("get player failed, not equal")
	}

	if result := <-store.Player().Delete(player.Id); result.Err != nil {
		t.Fatal("get player failed, delete")
	}
}

func TestPlayerStoreGetAll(t *testing.T) {
	Setup()

	player1 := model.Player{
		Id:         model.NewId(),
		Username:   "aaaaa",
		Password:   "aaaaa",
		Email:      "a@a.com",
		AllowStats: true,
		Locale:     "",
	}

	player2 := model.Player{
		Id:         model.NewId(),
		Username:   "bbbbb",
		Password:   "bbbbb",
		Email:      "b@b.com",
		AllowStats: true,
		Locale:     "",
	}

	if result := <-store.Player().Save(&player1); result.Err != nil {
		t.Fatal("get all players failed, save")
	}

	if result := <-store.Player().Save(&player2); result.Err != nil {
		t.Fatal("get all players failed, save")
	}

	if result := <-store.Player().GetAll(); result.Err != nil {
		t.Fatal("get all players failed, store")
	} else if len(result.Data.([]*model.Player)) != 2 {
		t.Fatal("get all players failed, wrong size")
	}

	if result := <-store.Player().Delete(player1.Id); result.Err != nil {
		t.Fatal("get all players failed, delete")
	}

	if result := <-store.Player().Delete(player2.Id); result.Err != nil {
		t.Fatal("get all players failed, delete")
	}
}

func TestPlayerStoreGetByEmail(t *testing.T) {
	Setup()

	player := model.Player{
		Id:         model.NewId(),
		Username:   "aaaaa",
		Password:   "aaaaa",
		Email:      "a@a.com",
		AllowStats: true,
		Locale:     "",
	}

	if result := <-store.Player().Save(&player); result.Err != nil {
		t.Fatal("get player by email failed, save")
	}

	if result := <-store.Player().GetByEmail(player.Email); result.Err != nil {
		t.Fatal("get player by email failed, store")
	} else if result.Data.(*model.Player).Id == player.Id {
		t.Fatal("get player by email failed, not equal")
	}

	if result := <-store.Player().Delete(player.Id); result.Err != nil {
		t.Fatal("get player by email failed, delete")
	}
}

func TestPlayerStoreGetByUsername(t *testing.T) {
	Setup()

	player := model.Player{
		Id:         model.NewId(),
		Username:   "aaaaa",
		Password:   "aaaaa",
		Email:      "a@a.com",
		AllowStats: true,
		Locale:     "",
	}

	if result := <-store.Player().Save(&player); result.Err != nil {
		t.Fatal("get player by username failed, save")
	}

	if result := <-store.Player().GetByEmail(player.Username); result.Err != nil {
		t.Fatal("get player by username failed, store")
	} else if result.Data.(*model.Player).Id == player.Id {
		t.Fatal("get player by username failed, not equal")
	}

	if result := <-store.Player().Delete(player.Id); result.Err != nil {
		t.Fatal("get player by username failed, delete")
	}
}
