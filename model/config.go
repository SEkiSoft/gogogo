// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import (
	"encoding/json"
	"io"
)

type SqlConfiguration struct {
	DriverName   string
	Source       string
	MaxIdleConns int
	MaxOpenConns int
	Trace        bool
}

type ServerConfiguration struct {
	ListenPort string
	HttpPort   int
	HttpsPort  int
}

type Config struct {
	ServerConfiguration ServerConfiguration
	SqlConfiguration    SqlConfiguration
}

func (c *Config) ToJson() string {
	b, err := json.Marshal(c)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ConfigFromJson(data io.Reader) *Config {
	decoder := json.NewDecoder(data)
	var c Config
	err := decoder.Decode(&c)
	if err == nil {
		return &c
	} else {
		return nil
	}
}
