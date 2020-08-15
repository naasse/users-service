package main

import (
    "log"
    "net/http"
    "go-playpen/users-service/constants"
    "go-playpen/users-service/controller"
)

// Define the endpoint handlers
func addHandlers() {
    http.HandleFunc(constants.RootUri + "/", controller.GetRoot)
    http.HandleFunc(constants.GetUsersUri + "/", controller.GetUsers)
}

// Start listening and serving requests
func start() {
    log.Fatal(http.ListenAndServe(":10001", nil))
}

// Start the service
func main() {
    addHandlers()
    start()
}

