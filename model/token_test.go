// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"strings"
	"testing"
)

func TestTokenToJson(t *testing.T) {
	token := Token{
		ID:       NewID(),
		PlayerID: NewID(),
		DeviceID: NewID(),
		Roles:    "",
		IsOAuth:  false,
	}

	json := token.ToJson()
	rtoken := TokenFromJson(strings.NewReader(json))

	if rtoken.ID != token.ID {
		t.Fatal("ids do not match")
	}
}

func TestTokenFromJson(t *testing.T) {
	token := Token{
		ID:       NewID(),
		PlayerID: NewID(),
		DeviceID: NewID(),
		Roles:    "",
		IsOAuth:  false,
	}

	json := token.ToJson()
	rtoken := TokenFromJson(strings.NewReader(json))
	rjson := rtoken.ToJson()

	if rjson != json {
		t.Fatal("jsons do not match")
	}
}

func TestTokenPreSave(t *testing.T) {
	token := Token{
		ID:             NewID(),
		PlayerID:       NewID(),
		DeviceID:       NewID(),
		CreateAt:       0,
		LastActivityAt: 0,
		Roles:          "",
		IsOAuth:        false,
	}

	token.PreSave()

	if token.CreateAt == 0 {
		t.Fatal("create at did not update")
	}
}

func TestTokenPreUpdate(t *testing.T) {
	token := Token{
		ID:             NewID(),
		PlayerID:       NewID(),
		DeviceID:       NewID(),
		CreateAt:       0,
		LastActivityAt: 0,
		Roles:          "",
		IsOAuth:        false,
	}

	token.PreUpdate()

	if token.LastActivityAt == 0 {
		t.Fatal("last activity at did not update")
	}
}

func TestTokenIsExpired(t *testing.T) {
	token := Token{
		ID:             NewID(),
		PlayerID:       NewID(),
		DeviceID:       NewID(),
		CreateAt:       0,
		LastActivityAt: 0,
		Roles:          "",
		IsOAuth:        false,
	}

	token.ExpiresAt = -1

	if token.IsExpired() == true {
		t.Fatal("should not expire")
	}

	token.ExpiresAt = 1

	if token.IsExpired() == false {
		t.Fatal("should have expired")
	}
}

func TestTokensToJson(t *testing.T) {
	token1 := Token{
		ID:             NewID(),
		PlayerID:       NewID(),
		DeviceID:       NewID(),
		CreateAt:       0,
		LastActivityAt: 0,
		Roles:          "",
		IsOAuth:        false,
	}

	token2 := Token{
		ID:             NewID(),
		PlayerID:       NewID(),
		DeviceID:       NewID(),
		CreateAt:       0,
		LastActivityAt: 0,
		Roles:          "",
		IsOAuth:        false,
	}

	tokens := []*Token{&token1, &token2}

	json := TokensToJson(tokens)

	if json == "[]" {
		t.Fatal("tokens to json failed")
	}
}
