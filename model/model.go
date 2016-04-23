package model

import (
	"fmt"
	"net/http"
)

//Main handler for requests
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have reached %s\n", r.URL.Path[1:]);
}

func ServerStart() {
	http.HandleFunc("/", Handler);
	http.ListenAndServe(":8080", nil);
}

