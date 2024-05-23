package models

import (
	"be-nabitu-go/database"
	"be-nabitu-go/schemas"
)

func M_Get_All_Profile() {
	var getAllProfile []*schemas.Profile
	database.GetDBMysql().Find(getAllProfile)
}
