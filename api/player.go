// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"
)

func InitPlayer() {
	BaseRoutes.Players.Handle("/create", ApiHandler(createPlayer)).Methods("POST")

	BaseRoutes.NeedPlayer.Handle("/update", ApiPlayerRequired(updatePlayer)).Methods("POST")
	BaseRoutes.NeedPlayer.Handle("/newpassword", ApiPlayerRequired(updatePassword)).Methods("POST")
	BaseRoutes.NeedPlayer.Handle("/newusername", ApiPlayerRequired(updateUsername)).Methods("POST")
	BaseRoutes.NeedPlayer.Handle("/newemail", ApiPlayerRequired(updateEmail)).Methods("POST")
	BaseRoutes.NeedPlayer.Handle("/get", ApiPlayerRequired(getPlayer)).Methods("GET")
	BaseRoutes.NeedPlayer.Handle("/games", ApiPlayerRequired(getPlayerGames)).Methods("GET")
}

func createPlayer(s *Session, w http.ResponseWriter, r *http.Request) {

}

func updatePlayer(s *Session, w http.ResponseWriter, r *http.Request) {

}

func updatePassword(s *Session, w http.ResponseWriter, r *http.Request) {

}

func updateUsername(s *Session, w http.ResponseWriter, r *http.Request) {

}

func updateEmail(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getPlayer(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getPlayerGames(s *Session, w http.ResponseWriter, r *http.Request) {

}
