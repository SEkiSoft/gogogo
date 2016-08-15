// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	"github.com/SEkiSoft/gogogo/model"
	"github.com/gorilla/mux"
)

func InitMove() {
	BaseRoutes.Moves.Handle("/get/{move_id:[A-Za-z0-9]+}", ApiHandler(getMove)).Methods("GET")
	BaseRoutes.Moves.Handle("/get_all", ApiHandler(getAllMoves)).Methods("GET")
	BaseRoutes.Moves.Handle("/get_game/{game_id:[A-Za-z0-9]+}", ApiHandler(getGameMoves)).Methods("GET")
	BaseRoutes.Moves.Handle("/get_player/{player_id:[A-Za-z0-9]+}", ApiHandler(getPlayerMoves)).Methods("GET")
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

func getPlayerMoves(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	playerId := params["player_id"]

	if result, err := GetPlayerMoves(playerId); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(model.MovesToJson(result)))
	}
}

func GetPlayerMoves(playerId string) ([]*model.Move, *model.Error) {
	if result := <-Srv.Store.Move().GetByPlayer(playerId); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.([]*model.Move), nil
	}
}
