// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPlayerIsValid(t *testing.T) {
	o := Player{
		Id:         NewId(),
		CreateAt:   GetMillis(),
		UpdateAt:   GetMillis(),
		DeleteAt:   GetMillis(),
		Username:   "aaaaa",
		Password:   "aaaaa",
		Email:      "a@a.com",
		AllowStats: true,
		Locale:     "",
	}

	if err := o.IsValid(); err != nil {
		t.Fatal(err)
	}

	o.Id = ""
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}

	o.Id = NewId()

	o.CreateAt = 0
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}

	o.CreateAt = GetMillis()

	o.UpdateAt = 0
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}

	o.UpdateAt = GetMillis()

	o.Username = ""
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}

	o.Username = strings.Repeat("a", MAX_USERNAME_LENGTH+1)
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}

	o.Username = "aaaa&*"
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}

	o.Username = "admin"
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}

	o.Username = "bbbbb"

	o.Email = strings.Repeat("@", 129)
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}

	o.Email = ""
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}

	o.Email = "This has no at"
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}

	o.Email = "a@a.com"

	o.Password = ""
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}

	o.Password = strings.Repeat("a", MAX_PASSWORD_LENGTH+1)
	if err := o.IsValid(); err == nil {
		t.Fatal("should be invalid")
	}
}

func TestPlayerPreSave(t *testing.T) {
	o := Player{
		Username: "AAAAA",
		Email:    "AAAAA@A.com"}
	o.PreSave()

	if len(o.Id) == 0 {
		t.Fatal("should not be empty")
	}

	if len(o.Username) == 0 {
		t.Fatal("should not be empty")
	}

	if len(o.Locale) == 0 {
		t.Fatal("should not be empty")
	}

	if o.Username != strings.ToLower(o.Username) {
		t.Fatal("should be lowercase")
	}

	if o.Email != strings.ToLower(o.Email) {
		t.Fatal("should be lowercase")
	}

	o.Locale = "AAAAA"
	o.PreSave()

	if o.Locale != strings.ToLower(o.Locale) {
		t.Fatal("should be lowercase")
	}

	if o.CreateAt == 0 {
		t.Fatal("should not be empty")
	}

	if o.UpdateAt == 0 {
		t.Fatal("should not be empty")
	}

	if o.CreateAt != o.UpdateAt {
		t.Fatal("should be same")
	}
}

func TestPlayerPreUpdate(t *testing.T) {
	o := Player{
		Username: "AAAAA",
		Email:    "AAAAA@A.com",
		Locale:   "EN",
	}

	o.PreUpdate()

	if o.Username != strings.ToLower(o.Username) {
		t.Fatal("should be lowercase")
	}

	if o.Email != strings.ToLower(o.Email) {
		t.Fatal("should be lowercase")
	}

	if o.Locale != strings.ToLower(o.Locale) {
		t.Fatal("should be lowercase")
	}

	if o.UpdateAt == 0 {
		t.Fatal("should not be empty")
	}
}

func TestPlayerToJson(t *testing.T) {
	player := Player{
		Id:         NewId(),
		CreateAt:   GetMillis(),
		UpdateAt:   GetMillis(),
		DeleteAt:   GetMillis(),
		Username:   "aaaaa",
		Password:   "aaaaa",
		Email:      "a@a.com",
		AllowStats: true,
		Locale:     "en",
	}

	json := player.ToJson()
	rPlayer := PlayerFromJson(strings.NewReader(json))

	if rPlayer.Id != player.Id {
		t.Fatal("Ids do not match")
	}
}

func TestPlayerFromJson(t *testing.T) {
	player := Player{
		Id:         NewId(),
		CreateAt:   GetMillis(),
		UpdateAt:   GetMillis(),
		DeleteAt:   GetMillis(),
		Username:   "aaaaa",
		Password:   "aaaaa",
		Email:      "a@a.com",
		AllowStats: true,
		Locale:     "en",
	}

	json := player.ToJson()
	rPlayer := PlayerFromJson(strings.NewReader(json))
	rJson := rPlayer.ToJson()

	if json != rJson {
		t.Fatal("JSONs do not match")
	}
}

func TestPlayerComparePassword(t *testing.T) {
	password := "password"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if !ComparePassword(string(hash), password) {
		t.Fatal("should be true")
	}

	password = "badpassword"
	if ComparePassword(string(hash), password) {
		t.Fatal("should be false")
	}
}

func TestPlayerIsValidUsername(t *testing.T) {
	username := ""
	if IsValidUsername(username) {
		t.Fatal("should be invalid")
	}

	username = strings.Repeat("a", MAX_USERNAME_LENGTH+1)
	if IsValidUsername(username) {
		t.Fatal("should be invalid")
	}

	username = "@!#$%&^"
	if IsValidUsername(username) {
		t.Fatal("should be invalid")
	}

	username = "admin"
	if IsValidUsername(username) {
		t.Fatal("should be invalid")
	}

	username = strings.Repeat("a", MIN_USERNAME_LENGTH)
	if !IsValidUsername(username) {
		t.Fatal("should be valid")
	}
}

func TestPlayerSanitize(t *testing.T) {
	player := Player{Password: "aaaaa"}
	player.Sanitize()

	if player.Password != "" {
		t.Fatal("should be empty")
	}
}
