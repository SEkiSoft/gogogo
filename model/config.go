// Copyright (c) 2016 David Lu
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

type GameConfiguration struct {
	ListenPort int
	HttpPort   int
	HttpsPort  int
}

type Config struct {
	GameConfiguration GameConfiguration
	SqlConfiguration  SqlConfiguration
}

func (c *Config) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ConfigFromJson(data io.Reader) *Config {
	decoder := json.NewDecoder(data)
	var o Config
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}
