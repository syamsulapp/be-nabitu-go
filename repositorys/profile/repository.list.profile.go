package repositorys

import (
	"be-nabitu-go/models"
	"be-nabitu-go/schemas"
	"net/http"

	"gorm.io/gorm"
)

type RepositoryProfileResult interface {
	ResultProfileRepository() (*[]models.Profiles, schemas.SchemaDatabaseError)
}

type repositoryProfileResults struct {
	db *gorm.DB
}

func NewRepositoryProfileResult(db *gorm.DB) *repositoryProfileResults {
	return &repositoryProfileResults{db: db}
}

func (r *repositoryProfileResults) ResultProfileRepository() (*[]models.Profiles, schemas.SchemaDatabaseError) {
	var profiles []models.Profiles
	db := r.db.Model(&profiles)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	resultProfiles := db.Debug().Find(&profiles)

	if resultProfiles.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &profiles, <-errorCode
	}
	return &profiles, <-errorCode
}
