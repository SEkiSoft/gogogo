// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"strings"
	"testing"
)

func TestErrorToString(t *testing.T) {
	error := Error{
		ID:         NewID(),
		Message:    "test",
		RequestID:  NewID(),
		StatusCode: 500,
		Where:      "here",
		params:     nil,
	}

	str := error.ToString()

	if !strings.Contains(str, "500") {
		t.Fatal("should be true")
	}
}

func TestErrorToJson(t *testing.T) {
    error := Error{
		ID:         NewID(),
		Message:    "test",
		RequestID:  NewID(),
		StatusCode: 500,
		Where:      "here",
		params:     nil,
	}

    json := error.ToJson()
    rerror := ErrorFromJson(strings.NewReader(json))

    if rerror.ID != error.ID {
        t.Fatal("ids do not match")
    }
}

func TestErrorFromJson(t *testing.T) {
    error := Error{
		ID:         NewID(),
		Message:    "test",
		RequestID:  NewID(),
		StatusCode: 500,
		Where:      "here",
		params:     nil,
	}

    json := error.ToJson()
    rerror := ErrorFromJson(strings.NewReader(json))
    rjson := rerror.ToJson()

    if rjson != json {
        t.Fatal("jsons do not match")
    }
}
