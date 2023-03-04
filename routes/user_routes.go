package routes

import (
	"github.com/gorilla/mux"
	"github.com/herizal95/gomux/controller/data/barang"
	"github.com/herizal95/gomux/middleware"
)

func UserRoutes(r *mux.Router) {

	router := r.PathPrefix("/data").Subrouter()

	router.Use(middleware.Deserializer)

	router.HandleFunc("/barang", barang.GetAllBarang).Methods("GET")
}
