// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"strings"
	"testing"
	"time"
)

func TestNewID(t *testing.T) {
	id := NewID()

	if len(id) != ID_LENGTH {
		t.Fatal("wrong id length")
	}
}

func TestHashPassword(t *testing.T) {
	password := "passwd"
	hash := HashPassword("passwd")

	if ComparePassword(hash, password) == false {
		t.Fatal("hash not the same")
	}
}

func TestGetMillis(t *testing.T) {
	curTime := GetMillis()
	time.Sleep(100 * time.Millisecond)
	nextTime := GetMillis()

	if nextTime-curTime < 100 {
		t.Fatal("space-time continuity borked")
	}
}

func TestStringInterfaceToJson(t *testing.T) {
	stringInterface := map[string]interface{}{
		"a":    "b",
		"test": 5,
		"hi":   true,
		"what": map[string]interface{}{
			"howdy": "yall",
		},
	}

	json := StringInterfaceToJson(stringInterface)

	if len(json) == 0 {
		t.Fatal("should not fail")
	}
}

func TestStringInterfaceFromJson(t *testing.T) {
	stringInterface := map[string]interface{}{
		"a":    "b",
		"test": 5,
		"hi":   true,
		"what": map[string]interface{}{
			"howdy": "yall",
		},
	}

	json := StringInterfaceToJson(stringInterface)

	returnedStringInterface := StringInterfaceFromJson(strings.NewReader(json))

	if returnedStringInterface["a"] != "b" {
		t.Fatal("should not fail")
	}
}

func TestIsLower(t *testing.T) {
	if IsLower("abc") == false {
		t.Fatal("should be true")
	}

	if IsLower("ABC") == true {
		t.Fatal("should be false")
	}
}

func TestIsValidEmail(t *testing.T) {
	if IsValidEmail("ABC@EMAIL.COM") == true {
		t.Fatal("should be false")
	}

	if IsValidEmail("abc@email.com") == false {
		t.Fatal("should be true")
	}

	if IsValidEmail("abc@") == true {
		t.Fatal("should be false")
	}

	if IsValidEmail("abc") == true {
		t.Fatal("should be false")
	}
}

func TestMapToJson(t *testing.T) {

}

func TestMapFromJson(t *testing.T) {

}
