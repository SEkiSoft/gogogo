// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/mux"
	"github.com/sekisoft/gogogo/model"
)

func InitPlayer() {
	l4g.Info("Initializing Player API")
	BaseRoutes.Players.Handle("/update", ApiPlayerRequired(updatePlayer)).Methods("POST")
	BaseRoutes.Players.Handle("/games", ApiPlayerRequired(getPlayerGames)).Methods("GET")
	BaseRoutes.Players.Handle("/get/{username:[A-Za-z0-9]+}", ApiPlayerRequired(getPlayerByUsername)).Methods("GET")
	BaseRoutes.Players.Handle("/get_me", ApiPlayerRequired(getMe)).Methods("GET")
	BaseRoutes.NeedPlayer.Handle("/get", ApiPlayerRequired(getPlayer)).Methods("GET")
}

func updatePlayer(s *Session, w http.ResponseWriter, r *http.Request) {
	player := model.PlayerFromJson(r.Body)

	if player == nil {
		s.SetInvalidParam("updatePlayer", "player")
		return
	}

	if s.Token.PlayerId != player.Id {
		s.SetInvalidParam("updatePlayer", "player")
		return
	}

	updatedPlayer, err := UpdatePlayer(player)
	if err != nil {
		s.Err = err
		return
	}

	w.Write([]byte(updatedPlayer.ToJson()))
}

func UpdatePlayer(player *model.Player) (*model.Player, *model.Error) {
	result := <-Srv.Store.Player().Update(player)
	if result.Err != nil {
		return nil, result.Err
	}

	updatedPlayer := result.Data.(*model.Player)
	return updatedPlayer, nil
}

func getPlayer(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	playerId := params["player_id"]

	if result, err := GetPlayer(playerId); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(result.ToJson()))
	}
}

func getMe(s *Session, w http.ResponseWriter, r *http.Request) {
	if result, err := GetPlayer(s.Token.PlayerId); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(result.ToJson()))
	}
}

func GetPlayer(id string) (*model.Player, *model.Error) {
	result := <-Srv.Store.Player().Get(id)
	if result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.Player), nil
}

func getPlayerGames(s *Session, w http.ResponseWriter, r *http.Request) {
	playerId := s.Token.PlayerId

	if result, err := GetPlayerGames(playerId); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(model.GamesToJson(result)))
	}
}

func GetPlayerGames(id string) ([]*model.Game, *model.Error) {
	result := <-Srv.Store.Player().GetPlayerGames(id)
	if result.Err != nil {
		return nil, result.Err
	}

	return result.Data.([]*model.Game), nil
}

func getPlayerByUsername(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	if result, err := GetPlayer(username); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(result.ToJson()))
	}
}

func GetPlayerByUsername(username string) (*model.Player, *model.Error) {
	result := <-Srv.Store.Player().GetByEmail(username)
	if result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.Player), nil
}
