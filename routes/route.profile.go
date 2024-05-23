package routes

import (
	"be-nabitu-go/schemas"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var InitRouteProfile = func(router *mux.Router) {
	// define all handler profile
	// resultGetAllProfileRepository := repositorys.RepositoryGetAllProfile
	// resultGetAllProfileService := repositorys.RepositoryGetAllProfile
	// resultGetAllProfileHandler := repositorys.RepositoryGetAllProfile

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		indexProfile := schemas.SchemaIndexProfile{Message: "Welcome Nabitu API 1.0"}

		loadIndexProfile, err := json.Marshal(indexProfile)
		if err != nil {
			panic(err.Error())
		}
		res.Write(loadIndexProfile)
	}).Methods("GET")

}
