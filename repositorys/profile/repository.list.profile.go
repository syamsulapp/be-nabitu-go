package repositorys

import (
	"be-nabitu-go/models"
	"be-nabitu-go/schemas"
	"encoding/json"
	"net/http"
)

func RepositoryGetAllProfile(res http.ResponseWriter, req *http.Request) {
	var getAllProfile []*schemas.Profile
	//query all data
	models.M_Get_All_Profile()
	//set header content type
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	//encode to json
	err := json.NewEncoder(res).Encode(getAllProfile)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
