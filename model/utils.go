// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"bytes"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"io"
	"net/mail"
	"strings"
	"time"

	"github.com/pborman/uuid"
	"golang.org/x/crypto/bcrypt"
)

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

const (
	ID_LENGTH = 24
)

func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(ID_LENGTH)

	return b.String()
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func Etag(parts ...interface{}) string {
	etag := ""
	for _, part := range parts {
		etag += fmt.Sprintf(".%v", part)
	}

	return etag
}

func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func MapToJson(objmap map[string]string) string {
	if json, err := json.Marshal(objmap); err == nil {
		return string(json)
	}

	return ""
}

func MapFromJson(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)

	var objmap map[string]string
	if err := decoder.Decode(&objmap); err == nil {
		return objmap
	}

	return make(map[string]string)
}

func ArrayToJson(objmap []string) string {
	if json, err := json.Marshal(objmap); err == nil {
		return string(json)
	}

	return ""
}

func ArrayFromJson(data io.Reader) []string {
	decoder := json.NewDecoder(data)

	var objmap []string
	if err := decoder.Decode(&objmap); err == nil {
		return objmap
	}
	return make([]string, 0)
}

func StringInterfaceToJson(objmap map[string]interface{}) string {
	if json, err := json.Marshal(objmap); err == nil {
		return string(json)
	}

	return ""
}

func StringInterfaceFromJson(data io.Reader) map[string]interface{} {
	decoder := json.NewDecoder(data)

	var objmap map[string]interface{}
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]interface{})
	}

	return objmap
}

func StringToJson(s string) string {
	json, err := json.Marshal(s)
	if err == nil {
		return string(json)
	}

	return ""
}

func StringFromJson(data io.Reader) string {
	decoder := json.NewDecoder(data)

	var s string
	if err := decoder.Decode(&s); err == nil {
		return s
	}

	return ""
}

func IsLower(s string) bool {
	if strings.ToLower(s) == s {
		return true
	}

	return false
}

func IsValidEmail(email string) bool {

	if !IsLower(email) {
		return false
	}

	if _, err := mail.ParseAddress(email); err == nil {
		return true
	}

	return false
}
