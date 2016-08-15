// Copyright (c) 2016 SEkiSoft
// See License.txt

package utils

import (
	"encoding/json"
	"os"

	"github.com/davidlu1997/gogogo/model"
)

var Cfg *model.Config = &model.Config{}
var CfgFileLocation string = "config/config.json"

func LoadConfig() {
	file, err := os.Open(CfgFileLocation)
	if err != nil {
		panic("Cannot open config file")
	}

	decoder := json.NewDecoder(file)
	config := model.Config{}

	err = decoder.Decode(&config)
	if err != nil {
		panic("Error reading from config file")
	}

	Cfg = &config
}
