// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"
)

func InitAdmin() {
	BaseRoutes.Admin.Handle("/get_games", ApiAdminRequired(getAllGames)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_players", ApiAdminRequired(getAllPlayers)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_moves", ApiAdminRequired(getAllMoves)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_stats", ApiAdminRequired(getAllStats)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_ai", ApiAdminRequired(getAi)).Methods("POST")
}

func getAllGames(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getAllPlayers(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getAllMoves(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getAllStats(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getAi(s *Session, w http.ResponseWriter, r *http.Request) {

}
