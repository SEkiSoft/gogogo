package model

import (
	"gogogo/ai"
	"fmt"
	"net/http"
)

//Main handler for requests
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have reached %s\n", r.URL.Path[1:]);
	//Find gameID in database
		//Send game via JSON if found
	//Otherwise create new game
		//Send new game via JSON
	//Deal with POST requests for moves
}

//Process board
func processBoard() {
	//Processes the current game state
	//Takes out pieces as needed
}

func ServerStart() {
	http.HandleFunc("/", Handler);
	http.ListenAndServe(":8080", nil);
}

//TODO:
//Database work