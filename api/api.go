// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	Root       *mux.Router
	Players    *mux.Router
	NeedPlayer *mux.Router
	Games      *mux.Router
	NeedGame   *mux.Router
	Ai         *mux.Router
	AiNeedGame *mux.Router
	Admin      *mux.Router
}

var BaseRoutes *Routes

func InitApi() {
	BaseRoutes = &Routes{}
	BaseRoutes.Root = Srv.Router

	BaseRoutes.Players = Srv.Router.PathPrefix("/players").Subrouter()
	BaseRoutes.NeedPlayer = BaseRoutes.Players.PathPrefix("/{player_id:[A-Za-z0-9]+}").Subrouter()

	BaseRoutes.Games = BaseRoutes.NeedPlayer.PathPrefix("/games").Subrouter()
	BaseRoutes.NeedGame = BaseRoutes.Games.PathPrefix("/{game_id:[A-Za-z0-9]+}").Subrouter()

	BaseRoutes.Ai = Srv.Router.PathPrefix("/ai").Subrouter()
	BaseRoutes.AiNeedGame = BaseRoutes.Ai.PathPrefix("/{game_id:[A-Za-z0-9]+}").Subrouter()

	BaseRoutes.Admin = Srv.Router.PathPrefix("/admin/{admin_id:[A-za-z0-9+}").Subrouter()

	InitPlayer()
	InitGame()
	InitAdmin()
	InitAi()

	Srv.Router.Handle("/", http.HandlerFunc(Handle404))
}
