// Copyright (c) 2016 David Lu
// See License.txt

package api

import (
	"net/http"
)

func InitAi() {
	BaseRoutes.Ai.Handle("/stats", ApiHandler(getAiStats)).Methods("GET")

	BaseRoutes.AiNeedGame.Handle("/get_move", ApiGameRequired(getMove)).Methods("GET")
}

func getAiStats(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getMove(s *Session, w http.ResponseWriter, r *http.Request) {

}
