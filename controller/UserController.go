package controller

import (
    "net/http"
    "encoding/json"
    userReps "go-playpen/users-representations"
)

// TODO: These are mocked Users.
// Next step will be to investigate persistence and ORM
var Users = []userReps.User {
    userReps.User{FirstName: "Nathan", LastName: "Asselstine"},
}

// Handle requests to get the list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Users)
}

