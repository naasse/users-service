package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    userReps "go-playpen/users-representations"
    restReps "go-playpen/rest-representations"
)

var Users []userReps.User

func root(w http.ResponseWriter, r *http.Request) {
    links := []restReps.Link {
        restReps.Link{Rel: "getUsers", Method: http.MethodGet, Uri: "/users"},
    }
    api := restReps.Api{Links: links}
    json.NewEncoder(w).Encode(api)
    fmt.Println("Endpoint Hit: API Root")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: Users List")
    json.NewEncoder(w).Encode(Users)
}

func handleRequests() {
    http.HandleFunc("/", root)
    http.HandleFunc("/users", getUsers)
    log.Fatal(http.ListenAndServe(":10001", nil))
}

func main() {
    Users = []userReps.User {
        userReps.User{FirstName: "Nathan", LastName: "Asselstine"},
    }
    handleRequests()
}

