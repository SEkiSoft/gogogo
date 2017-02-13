// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import (
	"net/http"
	//"github.com/sekisoft/gogogo/ai"
	l4g "github.com/alecthomas/log4go"
)

func InitAi() {
	l4g.Info("Initializing AI API")
	BaseRoutes.Ai.Handle("/stats", ApiHandler(getAiStats)).Methods("GET")

	BaseRoutes.AiNeedGame.Handle("/move", ApiHandler(getAiMove)).Methods("GET")
}

func getAiStats(s *Session, w http.ResponseWriter, r *http.Request) {

}

func getAiMove(s *Session, w http.ResponseWriter, r *http.Request) {

}
