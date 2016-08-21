// Copyright (c) 2016 SEkiSoft
// See License.txt

package store

import (
	"testing"

	"github.com/SEkiSoft/gogogo/utils"
)

var store Store

func TestSetUp() {
	if store == nil {
		utils.LoadConfig()

		store = NewSqlStore()
	}
}

func TestSqlStore(t *testing.T) {
	TestSetUp()

	if store == nil {
		t.Fatal("should not fail")
	}
}

func TestSqlStoreClose(t *testing.T) {
	TestSetUp()

	store.Close()

	result := <-store.Game().GetAll()

	if result.Err == nil {
		t.Fatal("should have errored")
	}
}
