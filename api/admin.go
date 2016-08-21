// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	"github.com/SEkiSoft/gogogo/model"
)

func InitAdmin() {
	BaseRoutes.Admin.Handle("/get_games", ApiAdminRequired(getAllGames)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_players", ApiAdminRequired(getAllPlayers)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_moves", ApiAdminRequired(getAllMoves)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_stats", ApiAdminRequired(getAllStats)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_ai", ApiAdminRequired(getAi)).Methods("POST")
}

func getAllGames(s *Session, w http.ResponseWriter, r *http.Request) {
	if !s.IsAdmin() {
		err := model.NewLocError("Admin.getAllGames", "Unauthorized admin access", nil, "")
		err.StatusCode = http.StatusUnauthorized
		w.Write([]byte(err.ToJson()))
		return
	}

	if result, err := GetAllGames(); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(model.GamesToJson(result)))
	}
}

func GetAllGames() ([]*model.Game, *model.Error) {
	if result := <-Srv.Store.Game().GetAll(); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.([]*model.Game), nil
	}
}

func getAllPlayers(s *Session, w http.ResponseWriter, r *http.Request) {
	if !s.IsAdmin() {
		err := model.NewLocError("Admin.getAllPlayers", "Unauthorized admin access", nil, "")
		err.StatusCode = http.StatusUnauthorized
		w.Write([]byte(err.ToJson()))
		return
	}

	if result, err := GetAllPlayers(); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(model.PlayersToJson(result)))
	}
}

func GetAllPlayers() ([]*model.Player, *model.Error) {
	if result := <-Srv.Store.Player().GetAll(); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.([]*model.Player), nil
	}
}

func getAllMoves(s *Session, w http.ResponseWriter, r *http.Request) {
	if !s.IsAdmin() {
		err := model.NewLocError("Admin.getAllMoves", "Unauthorized admin access", nil, "")
		err.StatusCode = http.StatusUnauthorized
		w.Write([]byte(err.ToJson()))
		return
	}

	if result, err := GetAllMoves(); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(model.MovesToJson(result)))
	}
}

func GetAllMoves() ([]*model.Move, *model.Error) {
	if result := <-Srv.Store.Move().GetAll(); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.([]*model.Move), nil
	}
}

func getAllStats(s *Session, w http.ResponseWriter, r *http.Request) {
	if !s.IsAdmin() {
		err := model.NewLocError("Admin.getAllStats", "Unauthorized admin access", nil, "")
		err.StatusCode = http.StatusUnauthorized
		w.Write([]byte(err.ToJson()))
		return
	}

	if result := <-Srv.Store.Game().GetAll(); result.Err != nil {
		s.Err = result.Err
	} else {
		g := result.Data.([]*model.Game)
		stats := make([]*model.GameStats, len(g))

		for index, element := range g {
			stats[index] = element.GetStats()
		}
		w.Write([]byte(model.GameStatssToJson(stats)))
	}
}

func getAi(s *Session, w http.ResponseWriter, r *http.Request) {
	// TODO: Complete this method when AI becomes accessible
}
