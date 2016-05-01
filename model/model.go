package model

import (
	"gogogo/ai"
	"html/template"
	"net/http"
	"database/sql"
	"github.com/mattn/sql-lite3"
)

type Board struct {
	numLines uint;
	board [][]uint;
	id string;
}

const idLen int = 8;

//Handler to create or load game
func gameHandler(w http.ResponseWriter, r *http.Request) {
	//Find gameID in database
	id := r.URL.Path[len("/game/"):len("/game/")+idLen];
	b := loadGame(id);
	//Send game via JSON
}

//Handler for moves
func moveHandler(w http.ResponseWriter, r *http.Request) {
	//Handle moves as needed
	idx := len("/move/") + idLen;
	id := r.URL.Path[len("/move/"):len("/move/")+idLen];
	player := r.URL.Path[idx:idx+2];
	x := r.URL.Path[idx+2:idx+4];
	y := r.URL.Path[idx+4:idx+6];
	//Process move
	//Write move to DB
	//Rely on client to refresh view
}

//AI queries
func aiHandler(w http.ResponseWriter, r *http.Request) {
	//Get id and player
	idx := len("/ai/") + idLen;
	id := r.URL.Path[len("/ai/"):len("/ai/")+idLen];
	player := r.URL.Path[idx:idx+2];

	//Call AI
	x, y, gg := ai.NextMove(loadGame(id), player);

	//Send moves to client
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
	http.HandleFunc("/ai/", aiHandler);
	http.ListenAndServe(":80", nil);
	//Disabling HTTPS for now
	//http.ListenAndServe(":443", nil);
}

//TODO:
//Database work