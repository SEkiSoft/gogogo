// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"
	"github.com/braintree/manners"
	"github.com/davidlu1997/gogogo/store"
	"github.com/davidlu1997/gogogo/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Store  store.Store
	Router *mux.Router
}

var Srv *Server

func NewServer() {
	l4g.Info("Creating new server...")

	Srv = &Server{}
	Srv.Store = store.NewSqlStore()
	Srv.Router = mux.NewRouter()
	Srv.Router.NotFoundHandler = http.HandlerFunc(Handle404)
}

func StartServer() {
	l4g.Info("Starting server...")
	l4g.Info("Listening on %s.", utils.Cfg.ServerConfiguration.ListenPort)

	var handler http.Handler = Srv.Router

	go func() {
		err := manners.ListenAndServe(utils.Cfg.ServerConfiguration.ListenPort, handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler))
		if err != nil {
			l4g.Critical("Server start critical failure!")
		}
	}()
}

func StopServer() {
	l4g.Info("Stopping server...")

	manners.Close()
	Srv.Store.Close()

	l4g.Info("Server stopped.")
}
