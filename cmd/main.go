package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var r = mux.NewRouter()

	r.HandleFunc("/signup", signUp).Methods("POST")

	defer log.Fatal(http.ListenAndServe(":8010", r))
}

func signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("<h1> Hi from CodeLabx</h1>"))
}
