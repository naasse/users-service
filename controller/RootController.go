package controller

import (
    "net/http"
    "encoding/json"
    "go-playpen/users-service/constants"
    restReps "go-playpen/rest-representations"
    userReps "go-playpen/users-representations"
)

// Handle requests to the root endpoint
func GetRoot(w http.ResponseWriter, r *http.Request) {
    links := []restReps.Link {
        restReps.Link{Rel: userReps.GetUsersRel, Method: http.MethodGet, Uri: constants.GetUsersUri + "/"},
    }
    api := restReps.Api{Links: links}
    json.NewEncoder(w).Encode(api)
}

