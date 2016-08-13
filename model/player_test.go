// Copyright (c) 2016 David Lu
// See License.txt
package model

import (
	"strings"
	"testing"
)

func TestPlayerIsValid() {
	o := Player{
		Id:         NewId(),
		CreateAt:   GetMillis(),
		UpdateAt:   GetMillis(),
		DeleteAt:   GetMillis(),
		Username:   "aaaaa",
		Password:   "aaaaa",
		Email:      "a@a.com",
		AllowStats: true,
		Locale:     ""}

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

	o.Email = strings.repeat("@", 129)
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
}

func TestPreSave(t *testing.T) {
	o := Player{}
	o.PreSave()
	o.Etag()
}

func TestPlayerPreUpdate(t *testing.T) {
	o := Player{}
	o.PreUpdate()
}

func TestPlayerComparePassword(t *testing.T) {
	o := Player{}
	password := "password"
	hash := bcryt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if !o.ComparePassword(hash, password) {
		t.Fatal("should be true")
	}

	password := "badpassword"
	if o.ComparePassword(hash, password) {
		t.Fatal("should be false")
	}
}

func TestPlayerIsValidUsername(t *testing.T) {
	o := Player{}

	username = ""
	if o.IsValidUsername(username) {
		t.Fatal("should be invalid")
	}

	username = strings.repeat("a", MAX_USERNAME_LENGTH+1)
	if o.IsValidUsername(username) {
		t.Fatal("should be invalid")
	}

	username = "@!#$%&^"
	if o.IsValidUsername(username) {
		t.Fatal("should be invalid")
	}

	username = "admin"
	if o.IsValidUsername(username) {
		t.Fatal("should be invalid")
	}

	username = strings.repeat("a", MIN_USERNAME_LENGTH)
	if !o.IsValidUsername(username) {
		t.Fatal("should be valid")
	}
}
