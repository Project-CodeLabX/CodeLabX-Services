package main

import (
	"codelabx/models"
	"codelabx/repos"
	"encoding/json"
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

	var incomingUser models.User
	json.NewDecoder(r.Body).Decode(&incomingUser)

	if repos.UserExists(&incomingUser) {
		json.NewEncoder(w).Encode("user Exists all ready")
	} else {
		ret := repos.AddUser(&incomingUser)
		json.NewEncoder(w).Encode(&ret)
	}

	// w.Write([]byte("<h1> Hi from CodeLabx</h1>"))
}
