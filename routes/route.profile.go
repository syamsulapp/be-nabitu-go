package routes

import (
	"be-nabitu-go/schemas"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var InitRouteProfile = func(router *mux.Router) {
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		indexProfile := schemas.SchemaIndexProfile{Message: "halo world"}

		loadIndexProfile, err := json.Marshal(indexProfile)
		if err != nil {
			panic(err.Error())
		}
		res.Write(loadIndexProfile)
	}).Methods("GET")
}
