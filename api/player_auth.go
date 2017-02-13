// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"
	"github.com/sekisoft/gogogo/model"
	"github.com/sekisoft/gogogo/store"
)

func InitPlayerAuth() {
	l4g.Info("Initializing Player Authentication API")
	BaseRoutes.Players.Handle("/create", ApiHandler(createPlayer)).Methods("POST")
	BaseRoutes.Players.Handle("/login", ApiHandler(login)).Methods("POST")
	BaseRoutes.Players.Handle("/logout", ApiPlayerRequired(logout)).Methods("GET")
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

func CreatePlayer(player *model.Player) (*model.Player, *model.AppError) {
	result := <-Srv.Store.Player().Save(player)
	if result.Err != nil {
		return nil, result.Err
	}

	registeredPlayer := result.Data.(*model.Player)

	return registeredPlayer, nil
}

func login(s *Session, w http.ResponseWriter, r *http.Request) {
	props := model.MapFromJson(r.Body)
	tokenID := props["token_id"]
	username := props["username"]
	password := props["password"]

	err := validateLoginProps(tokenID, username, password)

	if err != nil {
		s.Err = err
		return
	}

	if len(tokenID) > 0 {
		if token, err := LoginByTokenID(tokenID, username); err == nil {
			s.Token = token
		} else {
			s.Err = err
			return
		}
	} else if len(username) > 0 {
		if token, err := Login(username, password); err == nil {
			s.Token = token
		} else {
			s.Err = err
			return
		}
	}

	w.Write([]byte(s.Token.ToJson()))
}

func validateLoginProps(tokenID, username, password string) *model.AppError {

	if len(password) == 0 && len(tokenID) == 0 {
		err := model.NewAppError("login", "Invalid parameters", http.StatusBadRequest)
		err.StatusCode = http.StatusBadRequest
		return err
	}

	if len(username) == 0 {
		err := model.NewAppError("login", "Blank username", http.StatusBadRequest)
		err.StatusCode = http.StatusBadRequest
		return err
	}

	return nil
}

func Login(username, password string) (*model.Token, *model.AppError) {
	var player *model.Player
	var err *model.AppError
	if player, err = GetPlayerByUsername(username); err != nil {
		return nil, err
	}

	if !model.ComparePassword(player.Password, password) {
		return nil, model.NewAppError("login", "Invalid username or password", http.StatusBadRequest)
	}

	token := &model.Token{
		PlayerID: player.ID,
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

func LoginByTokenID(tokenID, username string) (*model.Token, *model.AppError) {
	var token *model.Token
	if result := <-Srv.Store.Token().Get(tokenID); result.Err == nil {
		token = result.Data.(*model.Token)
	} else {
		return nil, result.Err
	}

	if token.IsExpired() {
		return nil, model.NewAppError("login", "Token expired", http.StatusUnauthorized)
	}

	var player *model.Player
	var err *model.AppError
	if player, err = GetPlayerByUsername(username); err != nil {
		return nil, err
	}

	if token.PlayerID != player.ID {
		return nil, model.NewAppError("login", "Invalid player", http.StatusUnauthorized)
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

func Logout(token *model.Token) *model.AppError {
	result := <-Srv.Store.Token().Delete(token.ID)

	return result.Err
}
