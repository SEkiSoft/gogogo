// Copyright (c) 2016 SEkiSoft
// See License.txt

package store

import (
	"testing"

	"github.com/SEkiSoft/gogogo/model"
)

func TestTokenStoreSave(t *testing.T) {
	TestSetUp()

	token := model.Token{
		Id:       model.NewId(),
		PlayerId: model.NewId(),
		DeviceId: model.NewId(),
		Roles:    "",
	}

	if result := <-store.Token().Save(&token); result.Err != nil {
		t.Fatal("save token failed")
	}

	if result := <-store.Token().Delete(token.Id); result.Err != nil {
		t.Fatal("save token failed, delete")
	}
}

func TestTokenStoreDelete(t *testing.T) {
	TestSetUp()

	token := model.Token{
		Id:       model.NewId(),
		PlayerId: model.NewId(),
		DeviceId: model.NewId(),
		Roles:    "",
	}

	if result := <-store.Token().Save(&token); result.Err != nil {
		t.Fatal("delete token failed, save")
	}

	if result := <-store.Token().Delete(token.Id); result.Err != nil {
		t.Fatal("delete token failed")
	}
}

func TestTokenStoreGet(t *testing.T) {
	TestSetUp()

	token := model.Token{
		Id:       model.NewId(),
		PlayerId: model.NewId(),
		DeviceId: model.NewId(),
		Roles:    "",
	}

	if result := <-store.Token().Save(&token); result.Err != nil {
		t.Fatal("get token failed, save")
	}

	if result := <-store.Token().Get(token.Id); result.Err != nil {
		t.Fatal("get token failed, store")
	} else if result.Data.(*model.Token).Id != token.Id {
		t.Fatal("get token failed, not equal")
	}

	if result := <-store.Token().Delete(token.Id); result.Err != nil {
		t.Fatal("get token failed, delete")
	}
}

func TestTokenStoreGetTokens(t *testing.T) {
	TestSetUp()

	token1 := model.Token{
		Id:       model.NewId(),
		PlayerId: model.NewId(),
		DeviceId: model.NewId(),
		Roles:    "",
	}

	token2 := model.Token{
		Id:       model.NewId(),
		PlayerId: token1.PlayerId,
		DeviceId: model.NewId(),
		Roles:    "",
	}

	if result := <-store.Token().Save(&token1); result.Err != nil {
		t.Fatal("get tokens failed, save")
	}

	if result := <-store.Token().Save(&token2); result.Err != nil {
		t.Fatal("get tokens failed, save")
	}

	if result := <-store.Token().GetTokens(token1.PlayerId); result.Err != nil {
		t.Fatal("get tokens failed, store")
	} else if len(result.Data.([]*model.Token)) != 2 {
		t.Fatal("get tokens failed, wrong size")
	}

	if result := <-store.Token().Delete(token1.Id); result.Err != nil {
		t.Fatal("get tokens failed, delete")
	}

	if result := <-store.Token().Delete(token2.Id); result.Err != nil {
		t.Fatal("get tokens failed, delete")
	}
}

func TestTokenStoreDeleteAll(t *testing.T) {
	TestSetUp()

	token1 := model.Token{
		Id:       model.NewId(),
		PlayerId: model.NewId(),
		DeviceId: model.NewId(),
		Roles:    "",
	}

	token2 := model.Token{
		Id:       model.NewId(),
		PlayerId: token1.PlayerId,
		DeviceId: model.NewId(),
		Roles:    "",
	}

	if result := <-store.Token().Save(&token1); result.Err != nil {
		t.Fatal("delete all tokens failed, save")
	}

	if result := <-store.Token().Save(&token2); result.Err != nil {
		t.Fatal("delete all tokens failed, save")
	}

	if result := <-store.Token().DeleteAll(token1.PlayerId); result.Err != nil {
		t.Fatal("delete all tokens failed, store")
	}

	if result := <-store.Token().GetTokens(token1.PlayerId); result.Err != nil {
		t.Fatal("delete all tokens failed, store")
	} else if len(result.Data.([]*model.Token)) != 0 {
		t.Fatal("delete all tokens failed")
	}
}

func TestTokenStoreGetAll(t *testing.T) {
	TestSetUp()

	token1 := model.Token{
		Id:       model.NewId(),
		PlayerId: model.NewId(),
		DeviceId: model.NewId(),
		Roles:    "",
	}

	token2 := model.Token{
		Id:       model.NewId(),
		PlayerId: model.NewId(),
		DeviceId: model.NewId(),
		Roles:    "",
	}

	if result := <-store.Token().Save(&token1); result.Err != nil {
		t.Fatal("get all tokens failed, save")
	}

	if result := <-store.Token().Save(&token2); result.Err != nil {
		t.Fatal("get all tokens failed, save")
	}

	if result := <-store.Token().GetAll(); result.Err != nil {
		t.Fatal("get all tokens failed, store")
	} else if len(result.Data.([]*model.Token)) != 2 {
		t.Fatal("get all tokens failed, wrong size")
	}

	if result := <-store.Token().Delete(token1.Id); result.Err != nil {
		t.Fatal("get all tokens failed, delete")
	}

	if result := <-store.Token().Delete(token2.Id); result.Err != nil {
		t.Fatal("get all tokens failed, delete")
	}
}
