package main

import (
	"be-nabitu-go/routes"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	//config .env
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err.Error())
	}

	envPath := filepath.Join(dir, ".env")
	err = godotenv.Load(envPath)
	log.Fatal(err)

	// setup mux router
	router := mux.NewRouter()

	//setup router profile
	SubRouter := router.PathPrefix("api/v1").Subrouter()
	routes.InitRouteProfile(router)
	log.Println("Api is running with port:" + os.Getenv("GO_PORT"))
	log.Fatalln(http.ListenAndServe(":8085", SubRouter))

}
