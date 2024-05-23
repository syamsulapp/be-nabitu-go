package main

import (
	"be-nabitu-go/configs"
	"be-nabitu-go/database"
	"be-nabitu-go/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/qinains/fastergoding"
)

func main() {
	//hot reload
	fastergoding.Run()
	//config .env
	configs.InitConfigEnv()
	//config db
	database.InitConnectionDB(configs.InitConfigDb())
	// setup mux router
	router := mux.NewRouter()
	//setup router profile
	SubRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.InitRouteProfile(SubRouter)
	log.Println("Api is running with port:" + os.Getenv("GO_PORT"))
	log.Fatalln(http.ListenAndServe(os.Getenv("GO_PORT"), SubRouter))

}
