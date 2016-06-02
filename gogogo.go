// Copyright (c) 2016 David Lu
// See License.txt

package main

import (
	"github.com/davidlu1997/gogogo/api"
	"github.com/davidlu1997/gogogo/model"
	"github.com/davidlu1997/gogogo/store"
	"github.com/davidlu1997/gogogo/utils"
)

func doLoadConfig(filename string) (err string) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(string)
		}
	}()
	utils.LoadConfig(filename)
	return ""
}

func main() {
	if err := doLoadConfig(flagConfigFile); err != "" {
		l4g.Exit("Unable to load configuration file: %s", errstr)
		return
	}

	l4g.Info("GoGoGo version %s", model.CurrentVersion)

	api.NewServer()
	api.InitApi()

	api.StartServer()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscallSIGTERM)
	<-class

	api.StopServer()
}
