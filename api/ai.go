// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"
	//"github.com/SEkiSoft/gogogo/ai"
)

func InitAi() {
	BaseRoutes.Ai.Handle("/stats", ApiHandler(getAiStats)).Methods("GET")

	BaseRoutes.AiNeedGame.Handle("/move", ApiGameRequired(getAiMove)).Methods("GET")
}

func getAiStats(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getAiMove(s *Session, w http.ResponseWriter, r *http.Request) {

}
