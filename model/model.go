package model

import (
	"gogogo/ai"
	"html/template"
	"net/http"
	"database/sql"
	"github.com/mattn/sql-lite3"
)

type Board struct {
	numLines int;
	board [][]int;
	id int;
}

//Handler to create or load game
func gameHandler(w http.ResponseWriter, r *http.Request) {
	//Find gameID in database
	id := r.URL.Path[len("/game/"):];
	b := loadGame(id);
	//Send game via JSON
}

//Handler for moves
func moveHandler(w http.ResponseWriter, r *http.Request) {
	//Handle moves as needed
	player := r.URL.Path[len("/move/"):len("/move/")+2];
	x := r.URL.Path[len("/move/")+2:len("/move/")+4];
	y := r.URL.Path[len("/move/")+4:len("/move/")+6];
	//Process move
	//Refresh view
}

//Load game
func loadGame(id string) Board{
	//Gets gameID, and loads game
	b := new(Board);
	b.id = id;
}

func ServerStart() {
	http.HandleFunc("/game/", gameHandler);
	http.HandleFunc("/move/", moveHandler);
	http.ListenAndServe(":80", nil);
	//Disabling HTTPS for now
	//http.ListenAndServe(":443", nil);
}

//TODO:
//Database work