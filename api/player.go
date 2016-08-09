// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"
	"strings"

	"github.com/davidlu1997/gogogo/model"
)

func InitPlayer() {
	BaseRoutes.Players.Handle("/create", ApiHandler(createPlayer)).Methods("POST")

	BaseRoutes.NeedPlayer.Handle("/update", ApiPlayerRequired(updatePlayer)).Methods("POST")
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

func getPlayer(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getPlayerGames(s *Session, w http.ResponseWriter, r *http.Request) {

}

func findPlayer(s *Session, w http.ResponseWriter, r *http.Request) {

}

func CreatePlayer(player *model.Player) (*model.Player, *model.Error) {
	if result := <-Srv.Store.Player().Save(player); result.Err != nil {
		return nil, result.Err
	} else {
		registeredPlayer := result.Data.(*model.Player)

		return registeredPlayer, nil
	}
}

func UpdatePlayer(player *model.Player) (*model.Player, *model.Error) {
	if result := <-Srv.Store.Player().Update(player); result.Err != nil {
		return nil, result.Err
	} else {
		updatedPlayer := result.Data.(*model.Player)

		return updatedPlayer, nil
	}
}

func GetPlayer(id string) (*model.Player, *model.Error) {

}
