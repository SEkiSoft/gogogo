// Copyright (c) 2016 David Lu
// See License.txt

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/davidlu1997/gogogo/api"
	"github.com/davidlu1997/gogogo/model"
	"github.com/davidlu1997/gogogo/utils"

	l4g "github.com/alecthomas/log4go"
)

func doLoadConfig() (err string) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(string)
		}
	}()
	utils.LoadConfig()
	return ""
}

func main() {
	if err := doLoadConfig(); err != "" {
		l4g.Exit("Unable to load configuration file: %s", err)
		return
	}

	l4g.Info("GoGoGo version %s", model.CurrentVersion)

	api.NewServer()
	api.InitApi()

	api.StartServer()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	api.StopServer()
}
