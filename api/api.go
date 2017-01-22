// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"github.com/gorilla/mux"
)

type Routes struct {
	Root       *mux.Router
	Players    *mux.Router
	NeedPlayer *mux.Router
	Games      *mux.Router
	NeedGame   *mux.Router
	Moves      *mux.Router
	Ai         *mux.Router
	AiNeedGame *mux.Router
	Admin      *mux.Router
	WebSocket  *WebSocketRouter
}

var BaseRoutes *Routes

func InitApi() {
	BaseRoutes = &Routes{}
	BaseRoutes.Root = Srv.Router.PathPrefix("/api").Subrouter()

	BaseRoutes.Players = BaseRoutes.Root.PathPrefix("/players").Subrouter()
	BaseRoutes.NeedPlayer = BaseRoutes.Players.PathPrefix("/{player_id:[A-Za-z0-9]+}").Subrouter()

	BaseRoutes.Games = BaseRoutes.NeedPlayer.PathPrefix("/games").Subrouter()
	BaseRoutes.NeedGame = BaseRoutes.Games.PathPrefix("/{game_id:[A-Za-z0-9]+}").Subrouter()

	BaseRoutes.Moves = BaseRoutes.NeedGame.PathPrefix("/move").Subrouter()

	BaseRoutes.Ai = BaseRoutes.Root.PathPrefix("/ai").Subrouter()
	BaseRoutes.AiNeedGame = BaseRoutes.Ai.PathPrefix("/{game_id:[A-Za-z0-9]+}").Subrouter()

	BaseRoutes.Admin = BaseRoutes.Root.PathPrefix("/admin").Subrouter()

	BaseRoutes.WebSocket = NewWebSocketRouter()

	InitPlayer()
	InitPlayerAuth()
	InitGame()
	InitMove()
	InitAdmin()
	InitAi()
	InitWebSocket()
}
