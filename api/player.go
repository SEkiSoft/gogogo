// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"fmt"
	"net/http"

	"github.com/davidlu1997/gogogo/model"
	"github.com/davidlu1997/gogogo/store"
	"github.com/davidlu1997/gogogo/utils"
	"github.com/gorilla/mux"
)

func InitPlayer() {
	BaseRoutes.Players.Handle("/create", ApiHandler(createPlayer)).Methods("POST")
	BaseRoutes.Players.Handle("/update", ApiPlayerRequired(updatePlayer)).Methods("POST")
	BaseRoutes.Players.Handle("/newpassword", ApiPlayerRequired(updatePassword)).Methods("POST")
	BaseRoutes.Players.Handle("/newusername", ApiPlayerRequired(updateUsername)).Methods("POST")
	BaseRoutes.Players.Handle("/newemail", ApiPlayerRequired(updateEmail)).Methods("POST")
	BaseRoutes.Players.Handle("/counts", ApiHandler(getPlayerCounts)).Methods("GET")

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

func getPlayerCounts(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getPlayer(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getPlayerGames(s *Session, w http.ResponseWriter, r *http.Request) {

}
