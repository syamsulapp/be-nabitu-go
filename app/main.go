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
	var (
		DB_HOSTS    = os.Getenv("DB_HOSTS")
		DB_USERNAME = os.Getenv("DB_USERNAME")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_DATABASE = os.Getenv("DB_DATABASE")
	)
	database.InitConnectionDB(configs.InitConfigDb(DB_USERNAME, DB_PASSWORD, DB_HOSTS, DB_DATABASE))
	// setup mux router
	router := mux.NewRouter()
	//setup router profile
	SubRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.InitRouteProfile(SubRouter)
	log.Println("Api is running with port:" + os.Getenv("GO_PORT"))
	log.Fatalln(http.ListenAndServe(os.Getenv("GO_PORT"), SubRouter))

}
