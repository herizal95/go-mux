package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/herizal95/gomux/config"
	"github.com/herizal95/gomux/routes"
)

func main() {

	config.ConnectDatabase()

	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()

	routes.AuthenticationRouter(router)
	routes.UserRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
