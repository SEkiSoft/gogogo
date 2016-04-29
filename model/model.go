package model

import (
	"gogogo/ai"
	"html/template"
	"net/http"
	"database/sql"
	"github.com/mattn/sql-lite3"
)

//Handler to create or load game
func gameHandler(w http.ResponseWriter, r *http.Request) {
	//Find gameID in database
		//Send game via JSON if found
	//Otherwise create new game
		//Send new game via JSON
}

//Handler for moves
func moveHandler(w http.ResponseWriter, r *http.Request) {
	//Handle moves as needed
	//Refresh view
}

func ServerStart() {
	http.HandleFunc("/game/", gameHandler);
	http.HandleFunc("/game/move/", moveHandler);
	http.ListenAndServe(":8080", nil);
}

//TODO:
//Database work