package main

import (
	"codelabx/auth"
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

	var incomingUser models.User //get incoming user
	json.NewDecoder(r.Body).Decode(&incomingUser)

	if repos.UserExists(&incomingUser) { //checking if identical user exists
		json.NewEncoder(w).Encode("user Exists all ready")
		return
	}
	res := repos.AddUser(&incomingUser) // adding user to db

	if res == 1 {
		// json.NewEncoder(w).Encode(&incomingUser)

		tokenStr := auth.CreateToken(&incomingUser) //creating token

		mp := map[string]string{"token": tokenStr, " message": "Welcome to CodeLabX"}
		json.NewEncoder(w).Encode(&mp) // sending token to user
		return
	} else {
		json.NewEncoder(w).Encode("Error Happened during process please try again.")
		return
	}
}
