package model

import (
	"encoding/base32"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func NewID() string {
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
