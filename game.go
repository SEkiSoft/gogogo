package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ReposonseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have reached %s\n", r.URL.Path[1:]);
}

func main() {
	http.HandleFunc("/", handler);
	http.ListenAndServe(":8080", nil);
}