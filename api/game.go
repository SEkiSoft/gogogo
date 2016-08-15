// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	"github.com/SEkiSoft/gogogo/model"
	"github.com/gorilla/mux"
)

func InitGame() {
	BaseRoutes.Games.Handle("/create", ApiHandler(createGame)).Methods("POST")

	BaseRoutes.NeedGame.Handle("/update", ApiGameRequired(updateGame)).Methods("POST")
	BaseRoutes.NeedGame.Handle("/move", ApiGameRequired(makeMove)).Methods("POST")
	BaseRoutes.NeedGame.Handle("/stats", ApiGameRequired(getGameStats)).Methods("GET")
	BaseRoutes.NeedGame.Handle("/get", ApiGameRequired(getGame)).Methods("GET")
}

func createGame(s *Session, w http.ResponseWriter, r *http.Request) {
	game := model.GameFromJson(r.Body)

	if game == nil {
		s.SetInvalidParam("createGame", "game")
		return
	}

	game.PreSave()

	if result := <-Srv.Store.Game().Save(game); result.Err != nil {
		s.Err = result.Err
	} else {
		w.Write([]byte(result.Data.(*model.Game).ToJson()))
	}
}

func getGame(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["game_id"]

	if result, err := GetGame(id); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(result.ToJson()))
	}
}

func GetGame(id string) (*model.Game, *model.Error) {
	if result := <-Srv.Store.Game().Get(id); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.(*model.Game), nil
	}
}

func getGameStats(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	if result := <-Srv.Store.Game().Get(id); result.Err != nil {
		s.Err = result.Err
	} else {
		g := result.Data.(*model.Game)

		if stats := g.GetStats(); stats != nil {
			w.Write([]byte(stats.ToJson()))
		}
	}
}

func updateGame(s *Session, w http.ResponseWriter, r *http.Request) {
	game := model.GameFromJson(r.Body)

	if game == nil {
		s.SetInvalidParam("updateGame", "game")
		return
	}

	game.PreUpdate()

	if result := <-Srv.Store.Game().Update(game); result.Err != nil {
		s.Err = result.Err
	} else {
		w.Write([]byte(result.Data.(*model.Game).ToJson()))
	}
}

func makeMove(s *Session, w http.ResponseWriter, r *http.Request) {
	move := model.MoveFromJson(r.Body)

	if move == nil {
		s.SetInvalidParam("makeMove", "move")
		return
	}

	var game *model.Game
	var err *model.Error

	if game, err = GetGame(move.GameId); err != nil {
		s.Err = err
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
