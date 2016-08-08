// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"
)

func InitMove() {
	BaseRoutes.Moves.Handle("/get/{move_id:[A-Za-z0-9]+}", ApiHandler(getMove)).Methods("GET")
	BaseRoutes.Moves.Handle("/get_all", ApiHandler(getAllMoves)).Methods("GET")
	BaseRoutes.Moves.Handle("/get_game/{game_id:[A-Za-z0-9]+}", ApiHandler(getGameMoves)).Methods("GET")
	BaseRoutes.Moves.Handle("/get_player/{player_id:[A-Za-z0-9]+}", ApiHandler(getPlayerMoves)).Methods("GET")
}

func getMove(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	if result, err := GetMove(id); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(result.ToJson()))
	}
}

func GetMove(id string) (*model.Move, *model.Error) {
	if result := <-Srv.Store.Move().Get(id); result.Err != nil {
		return nil, result.Error
	} else {
		return result.Data.(*model.Move), nil
	}
}

func getAllMoves(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getGameMoves(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameId := params["game_id"]	

	if result, err := GetGameMoves(gameId); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(result.ToJson()))
	}
}

func GetGameMoves(gameId string) (*model.Move, *model.Error) {
	if result := <-Srv.Store.Move().GetByGame(gameId); result.Err != nil {
		return nil, result.Error
	} else {
		return result.Data.(*model.Move), nil
	}
}
