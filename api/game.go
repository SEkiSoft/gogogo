// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"
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
		s.SetInvalidParam("createGame", "game");
		return
	}

	game.PreSave();

	if result := <-Srv.Store.Game().Save(game); result.Err != nil {
		s.Err = result.Err
	} else {
		w.Write([]byte(result.Data.(*model.Game).ToJson()))
	}
}

func getGame(s *Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	if result := <-Srv.Store.Game().Get(id); result.Err != nil {
		s.Err = result.Err
	} else {
		w.Write([]byte(result.Data.(*model.Game).ToJson()))
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

}

func makeMove(s *Session, w http.ResponseWriter, r *http.Request) {

}
