// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"
	"strings"

	"github.com/davidlu1997/gogogo/model"
	"github.com/gorilla/mux"
)

func InitPlayer() {
	BaseRoutes.Players.Handle("/create", ApiHandler(createPlayer)).Methods("POST")

	BaseRoutes.NeedPlayer.Handle("/update", ApiPlayerRequired(updatePlayer)).Methods("POST")
	BaseRoutes.NeedPlayer.Handle("/get", ApiPlayerRequired(getPlayer)).Methods("GET")
	BaseRoutes.NeedPlayer.Handle("/games", ApiPlayerRequired(getPlayerGames)).Methods("GET")
	BaseRoutes.NeedPlayer.Handle("/find", ApiPlayerRequired(findPlayer)).Methods("POST")
	//Need to add endpoints for getByEmail and getByUsername 
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
	params := mux.Vars(r)
	playerId := params["player_id"]	

	if result, err := GetPlayer(playerId); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(result.ToJson(result)))
	}	
}

func GetPlayer(id string) (*model.Player, *model.Error) {
	if result := <-Srv.Store.Player().Get(id); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.(*model.Player), nil
	}
}

func getPlayerGames(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	playerId := params["player_id"]	

	if result, err := GetPlayerGames(playerId); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(model.GamesToJson(result)))
	}	
}

func GetPlayerGames(playerId string) ([]*model.Game, *model.Error) {
	if result := <-Srv.Store.Player().GetPlayerGames(playerId); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.([]*model.Game), nil
	}
}

func getEmail(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]	

	if result, err := GetPlayer(email); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(result.ToJson(result)))
	}	
}

func GetEmail(email string) (*model.Player, *model.Error) {
	if result := <-Srv.Store.Player().GetByEmail(email); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.(*model.Player), nil
	}
}

func getUsername(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["username"]	

	if result, err := GetPlayer(username); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(result.ToJson(result)))
	}	
}

func GetUsername(username string) (*model.Player, *model.Error) {
	if result := <-Srv.Store.Player().GetByEmail(username); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.(*model.Player), nil
	}
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

