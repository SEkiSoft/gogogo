// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"
)

func InitGame() {
	BaseRoutes.Games.Handle("/create", ApiHandler(createGame)).Methods("POST")
	BaseRoutes.Games.Handle("/get_player", ApiHandler(getGameByPlayer)).Methods("GET")
	BaseRoutes.Games.Handle("/get_two_player", ApiHandler(getGameByTwoPlayers)).Methods("GET")
	BaseRoutes.Games.Handle("/get_all", ApiHandler(getAllGames)).Methods("GET")
	BaseRoutes.Games.Handle("/get_all_finished", ApiHandler(getAllFinishedGames)).Methods("GET")

	BaseRoutes.NeedGame.Handle("/", ApiHandler(getGame)).Methods("GET")
	BaseRoutes.NeedGame.Handle("/move", ApiHandler(moveHandler)).Methods("POST")
	BaseRoutes.NeedGame.Handle("/update", ApiHandler(updateHandler)).Methods("POST")
}

func createGame(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getGame(s *Session, w http.ResponseWriter, r *http.Request) {

}

func moveHandler(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getGameByPlayer(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getGameByTwoPlayers(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getAllGames(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getAllFinishedGames(s *Session, w http.ResponseWriter, r *http.Request) {

}

func updateHandler(s *Session, w http.ResponseWriter, r *http.Request) {

}
