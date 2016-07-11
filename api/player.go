// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/davidlu1997/gogogo/model"
	"net/http"
	"strings"
)

func InitPlayer() {
	BaseRoutes.Players.Handle("/create", ApiHandler(createPlayer)).Methods("POST")

	BaseRoutes.NeedPlayer.Handle("/update", ApiPlayerRequired(updatePlayer)).Methods("POST")
	BaseRoutes.NeedPlayer.Handle("/newpassword", ApiPlayerRequired(updatePassword)).Methods("POST")
	BaseRoutes.NeedPlayer.Handle("/newusername", ApiPlayerRequired(updateUsername)).Methods("POST")
	BaseRoutes.NeedPlayer.Handle("/newemail", ApiPlayerRequired(updateEmail)).Methods("POST")
	BaseRoutes.NeedPlayer.Handle("/get", ApiPlayerRequired(getPlayer)).Methods("GET")
	BaseRoutes.NeedPlayer.Handle("/games", ApiPlayerRequired(getPlayerGames)).Methods("GET")
	BaseRoutes.NeedPlayer.Handle("/find", ApiPlayerRequired(findPlayer)).Methods("POST")
}

func createPlayer(s *Session, w http.ResponseWriter, r *http.Request) {
	player := model.PlayerFromJson(r.Body)

	if player == nil {
		s.SetInvalidParam("createPlayer", "player")
		return
	}

	data := r.URL.Query().Get("d")
	props := model.MapFromJson(strings.NewReader(data))
	player.Email = props["email"]

	registeredPlayer, err := CreatePlayer(player)
	if err != nil {
		s.Err = err
		return
	}

	w.Write([]byte(registeredPlayer.ToJson()))
}

func updatePlayer(s *Session, w http.ResponseWriter, r *http.Request) {
	player := model.PlayerFromJson(r.Body)

	if player == nil {
		s.SetInvalidParam("updatePlayer", "player")
		return
	}

	data := r.URL.Query().Get("d")
	props := model.MapFromJson(strings.NewReader(data))
	player.Email = props["email"]

	updatedPlayer, err := CreatePlayer(player)
	if err != nil {
		s.Err = err
		return
	}

	w.Write([]byte(updatedPlayer.ToJson()))
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

func findPlayer(s *Session, w http.ResponseWriter, r *http.Request) {

}

func CreatePlayer(player *model.Player) (*model.Player, *model.Error) {
	if result := <-Srv.Store.Player().Save(player); result.Err != nil {
		l4g.Error("Create player save error: %s", result.Err)
		return nil, result.Err
	} else {
		registeredPlayer := result.Data.(*model.Player)

		return registeredPlayer, nil
	}
}

func UpdatePlayer(player *model.Player) (*model.Player, *model.Error) {
	if result := <-Srv.Store.Player().Update(player); result.Err != nil {
		l4g.Error("Player update error: %s", result.Err)
		return nil, result.Err
	} else {
		updatedPlayer := result.Data.(*model.Player)

		return updatedPlayer, nil
	}
}
