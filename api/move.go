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

}

func getGameMoves(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getPlayerMoves(s *Session, w http.ResponseWriter, r *http.Request) {

}
