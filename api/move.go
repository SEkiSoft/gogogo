// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	"github.com/SEkiSoft/gogogo/model"
	"github.com/gorilla/mux"
)

func InitMove() {
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

func GetMove(id string) (*model.Move, *model.Error) {
	if result := <-Srv.Store.Move().Get(id); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.(*model.Move), nil
	}
}

func getGameMoves(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameId := params["game_id"]

	if result, err := GetGameMoves(gameId); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(model.MovesToJson(result)))
	}
}

func GetGameMoves(gameId string) ([]*model.Move, *model.Error) {
	if result := <-Srv.Store.Move().GetByGame(gameId); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.([]*model.Move), nil
	}
}

func makeMove(s *Session, w http.ResponseWriter, r *http.Request) {
	move := model.MoveFromJson(r.Body)
	params := mux.Vars(r)
	gameId := params["game_id"]

	if move == nil {
		s.SetInvalidParam("makeMove", "move")
		return
	}

	if move.GameId != gameId {
		s.SetInvalidParam("makeMove", "move")
		return
	}

	var game *model.Game
	var err *model.Error

	if game, err = GetGame(move.GameId); err != nil {
		s.Err = err
		return
	}

	if !game.HasPlayer(s.Token.PlayerId) {
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
