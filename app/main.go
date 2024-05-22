package main

import (
	"be-nabitu-go/routes"

	"github.com/gorilla/mux"
)

func main() {
	// setup mux router
	router := mux.NewRouter()

	//setup router profile
	profileRouter := router.PathPrefix("api/v1").Subrouter()
	routes.InitRouteProfile(profileRouter)
}
