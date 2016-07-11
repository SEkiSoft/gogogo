// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"
)

func InitGame() {
	BaseRoutes.Games.Handle("/create", ApiHandler(createGame)).Methods("POST")

	BaseRoutes.NeedGame.Handle("/update", ApiGameRequired(updateGame)).Methods("POST")
	BaseRoutes.NeedGame.Handle("/stats", ApiGameRequired(getGameStats)).Methods("GET")
	BaseRoutes.NeedGame.Handle("/get", ApiGameRequired(getGame)).Methods("GET")
}

func createGame(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getGame(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getGameStats(s *Session, w http.ResponseWriter, r *http.Request) {

}

func updateGame(s *Session, w http.ResponseWriter, r *http.Request) {

}
