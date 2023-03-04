package routes

import (
	"github.com/gorilla/mux"
	"github.com/herizal95/gomux/controller/authcontroller"
)

func AuthenticationRouter(r *mux.Router) {

	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/login", authcontroller.Login).Methods("POST")
	router.HandleFunc("/register", authcontroller.Register).Methods("POST")
	router.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
}
