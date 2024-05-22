package main

import (
	"be-nabitu-go/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// setup mux router
	router := mux.NewRouter()

	//setup router profile
	SubRouter := router.PathPrefix("api/v1").Subrouter()
	routes.InitRouteProfile(router)

	log.Fatalln(http.ListenAndServe(os.Getenv("APP_PORT"), SubRouter))
}
