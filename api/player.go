// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	"github.com/SEkiSoft/gogogo/model"
	"github.com/gorilla/mux"
)

func InitPlayer() {
	BaseRoutes.Players.Handle("/create", ApiHandler(createPlayer)).Methods("POST")
	BaseRoutes.Players.Handle("/login", ApiHandler(login)).Methods("POST")
	BaseRoutes.Players.Handle("/logout", ApiPlayerRequired(logout)).Methods("GET")
	BaseRoutes.Players.Handle("/update", ApiPlayerRequired(updatePlayer)).Methods("POST")
	BaseRoutes.Players.Handle("/games", ApiPlayerRequired(getPlayerGames)).Methods("GET")
	BaseRoutes.Players.Handle("/get/{username:[A-Za-z0-9]+}", ApiPlayerRequired(getPlayerByUsername)).Methods("GET")
	BaseRoutes.Players.Handle("/get_me", ApiPlayerRequired(getMe)).Methods("GET")
	BaseRoutes.NeedPlayer.Handle("/get", ApiPlayerRequired(getPlayer)).Methods("GET")
}

func createPlayer(s *Session, w http.ResponseWriter, r *http.Request) {
	player := model.PlayerFromJson(r.Body)

	if player == nil {
		s.SetInvalidParam("createPlayer", "player")
		return
	}

	registeredPlayer, err := CreatePlayer(player)
	if err != nil {
		s.Err = err
		return
	}

	w.Write([]byte(registeredPlayer.ToJson()))
}

func CreatePlayer(player *model.Player) (*model.Player, *model.Error) {
	if result := <-Srv.Store.Player().Save(player); result.Err != nil {
		return nil, result.Err
	} else {
		registeredPlayer := result.Data.(*model.Player)

		return registeredPlayer, nil
	}
}

func login(s *Session, w http.ResponseWriter, r *http.Request) {

}

func Login(username, hashedPassword string) (*model.Token, *model.Error) {
	return nil, nil
}

func logout(s *Session, w http.ResponseWriter, r *http.Request) {

}

func Logout(token *model.Token) *model.Error {
	return nil
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
	if result := <-Srv.Store.Player().Update(player); result.Err != nil {
		return nil, result.Err
	} else {
		updatedPlayer := result.Data.(*model.Player)

		return updatedPlayer, nil
	}
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
	if result := <-Srv.Store.Player().Get(id); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.(*model.Player), nil
	}
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
	if result := <-Srv.Store.Player().GetPlayerGames(id); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.([]*model.Game), nil
	}
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
	if result := <-Srv.Store.Player().GetByEmail(username); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.(*model.Player), nil
	}
}
