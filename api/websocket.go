// Copyright (c) 2016 SEkiSoft
// See License.txt

package api

import "net/http"

func InitWebSocket() {
	BaseRoutes.WebSocket.Handle("/websocket", ApiPlayerRequired(connect)).Methods("GET")
}

func connect(s *Session, w http.ResponseWriter, r *http.Request) {
	// TODO
}
