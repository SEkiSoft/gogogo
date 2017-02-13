// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"
	"github.com/sekisoft/gogogo/model"
)

func InitAdmin() {
	l4g.Info("Initializing Admin API")
	BaseRoutes.Admin.Handle("/get_games", ApiAdminRequired(getAllGames)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_players", ApiAdminRequired(getAllPlayers)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_moves", ApiAdminRequired(getAllMoves)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_stats", ApiAdminRequired(getAllStats)).Methods("POST")
	BaseRoutes.Admin.Handle("/get_ai", ApiAdminRequired(getAi)).Methods("POST")
}

func getAllGames(s *Session, w http.ResponseWriter, r *http.Request) {
	if !s.IsAdmin() {
		err := model.NewAppError("Admin.getAllGames", "Unauthorized admin access", http.StatusUnauthorized)
		err.StatusCode = http.StatusUnauthorized
		http.Error(w, err.Error(), err.StatusCode)
		return
	}

	if result, err := GetAllGames(); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(model.GamesToJson(result)))
	}
}

func GetAllGames() ([]*model.Game, *model.AppError) {
	result := <-Srv.Store.Game().GetAll()
	if result.Err != nil {
		return nil, result.Err
	}

	return result.Data.([]*model.Game), nil
}

func getAllPlayers(s *Session, w http.ResponseWriter, r *http.Request) {
	if !s.IsAdmin() {
		err := model.NewAppError("Admin.getAllGames", "Unauthorized admin access", http.StatusUnauthorized)
		err.StatusCode = http.StatusUnauthorized
		http.Error(w, err.Error(), err.StatusCode)
		return
	}

	if result, err := GetAllPlayers(); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(model.PlayersToJson(result)))
	}
}

func GetAllPlayers() ([]*model.Player, *model.AppError) {
	result := <-Srv.Store.Player().GetAll()
	if result.Err != nil {
		return nil, result.Err
	}

	return result.Data.([]*model.Player), nil
}

func getAllMoves(s *Session, w http.ResponseWriter, r *http.Request) {
	if !s.IsAdmin() {
		err := model.NewAppError("Admin.getAllGames", "Unauthorized admin access", http.StatusUnauthorized)
		err.StatusCode = http.StatusUnauthorized
		http.Error(w, err.Error(), err.StatusCode)
		return
	}

	if result, err := GetAllMoves(); err != nil {
		s.Err = err
	} else {
		w.Write([]byte(model.MovesToJson(result)))
	}
}

func GetAllMoves() ([]*model.Move, *model.AppError) {
	result := <-Srv.Store.Move().GetAll()
	if result.Err != nil {
		return nil, result.Err
	}

	return result.Data.([]*model.Move), nil
}

func getAllStats(s *Session, w http.ResponseWriter, r *http.Request) {
	if !s.IsAdmin() {
		err := model.NewAppError("Admin.getAllGames", "Unauthorized admin access", http.StatusUnauthorized)
		err.StatusCode = http.StatusUnauthorized
		http.Error(w, err.Error(), err.StatusCode)
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
