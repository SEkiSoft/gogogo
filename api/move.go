// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/mux"
	"github.com/sekisoft/gogogo/model"
)

func InitMove() {
	l4g.Info("Initializing Move API")
	BaseRoutes.Moves.Handle("/", ApiPlayerRequired(makeMove)).Methods("POST")
	BaseRoutes.Moves.Handle("/get/{move_id:[A-Za-z0-9]+}", ApiPlayerRequired(getMove)).Methods("GET")
	BaseRoutes.Moves.Handle("/get", ApiPlayerRequired(getGameMoves)).Methods("GET")
}

func getMove(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["move_id"]

	if result, err := GetMove(id); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(result.ToJson()))
	}
}

func GetMove(id string) (*model.Move, *model.AppError) {
	result := <-Srv.Store.Move().Get(id)
	if result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*model.Move), nil
}

func getGameMoves(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameID := params["game_id"]

	if result, err := GetGameMoves(gameID); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(model.MovesToJson(result)))
	}
}

func GetGameMoves(gameID string) ([]*model.Move, *model.AppError) {
	result := <-Srv.Store.Move().GetByGame(gameID)
	if result.Err != nil {
		return nil, result.Err
	}

	return result.Data.([]*model.Move), nil
}

func makeMove(s *Session, w http.ResponseWriter, r *http.Request) {
	move := model.MoveFromJson(r.Body)
	params := mux.Vars(r)
	gameID := params["game_id"]

	if move == nil {
		s.SetInvalidParam("makeMove", "move")
		return
	}

	if move.GameID != gameID {
		s.SetInvalidParam("makeMove", "move")
		return
	}

	var game *model.Game
	var err *model.AppError

	if game, err = GetGame(move.GameID); err != nil {
		s.Err = err
		return
	}

	if !game.HasPlayer(s.Token.PlayerID) {
		s.SetInvalidParam("makeMove", "move")
		return
	}

	if moveErr := move.IsValid(game); moveErr != nil {
		s.Err = moveErr
		return
	}

	if result := <-Srv.Store.Move().Save(move); result.Err != nil {
		s.Err = result.Err
	} else {
		w.Write([]byte(result.Data.(*model.Move).ToJson()))
	}
}
