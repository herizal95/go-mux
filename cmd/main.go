package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/herizal95/gomux/config"
	"github.com/herizal95/gomux/controller/authcontroller"
)

func main() {

	config.ConnectDatabase()

	router := mux.NewRouter()

	router.HandleFunc("/login", authcontroller.Login).Methods("POST")
	router.HandleFunc("/register", authcontroller.Register).Methods("POST")
	router.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
