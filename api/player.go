// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/mux"
	"github.com/sekisoft/gogogo/model"
	"github.com/sekisoft/gogogo/store"
)

func InitPlayer() {
	l4g.Info("Initializing Player API")
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
	props := model.MapFromJson(r.Body)

	tokenId := props["token_id"]
	username := props["username"]
	password := props["password"]

	if len(password) == 0 {
		s.Err = model.NewLocError("login", "Blank password", nil, "")
		s.Err.StatusCode = http.StatusBadRequest
		return
	}

	if len(tokenId) > 0 && len(username) > 0 {
		if token, err := LoginByTokenId(tokenId, username); err == nil {
			s.Token = token
		} else {
			s.Err = err
			return
		}
	} else if len(username) > 0 && len(password) > 0 {
		if token, err := Login(username, password); err == nil {
			s.Token = token
		} else {
			s.Err = err
			return
		}
	} else {
		s.Err = NewInvalidParamError("login", "username/password")
		return
	}

	w.Write([]byte(s.Token.ToJson()))
}

func Login(username, password string) (*model.Token, *model.Error) {
	var player *model.Player
	var err *model.Error
	if player, err = GetPlayerByUsername(username); err != nil {
		return nil, err
	}

	if !model.ComparePassword(player.Password, password) {
		return nil, model.NewLocError("login", "Invalid username or password", nil, "")
	}

	token := &model.Token{
		PlayerId: player.Id,
	}

	if player.IsAdmin {
		token.Roles = "admin"
	}

	var result store.StoreResult
	if result = <-Srv.Store.Token().Save(token); result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.Token), nil
}

func LoginByTokenId(tokenId, username string) (*model.Token, *model.Error) {
	var token *model.Token
	if result := <-Srv.Store.Token().Get(tokenId); result.Err == nil {
		token = result.Data.(*model.Token)
	} else {
		return nil, result.Err
	}

	if token.IsExpired() {
		return nil, model.NewLocError("login", "Token expired", nil, "")
	}

	var player *model.Player
	var err *model.Error
	if player, err = GetPlayerByUsername(username); err != nil {
		return nil, err
	}

	if token.PlayerId != player.Id {
		return nil, model.NewLocError("login", "Invalid player", nil, "")
	}

	return token, nil
}

func logout(s *Session, w http.ResponseWriter, r *http.Request) {
	if err := Logout(s.Token); err != nil {
		s.Err = err
		return
	}

	w.Write([]byte("success"))
}

func Logout(token *model.Token) *model.Error {
	result := <-Srv.Store.Token().Delete(token.Id)

	return result.Err
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
