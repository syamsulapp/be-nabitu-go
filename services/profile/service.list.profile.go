package services

import (
	"be-nabitu-go/models"
	repositorys "be-nabitu-go/repositorys/profile"
	"be-nabitu-go/schemas"
)

type ServiceGetAllProfile interface {
	ResultProfileService() (*[]models.Profiles, schemas.SchemaDatabaseError)
}

type callRepositories struct {
	repository repositorys.RepositoryProfileResult
}

func NewServiceProfileResult(repository repositorys.RepositoryProfileResult) *callRepositories {
	return &callRepositories{repository: repository}
}

func (s *callRepositories) ResultProfileService() (*[]models.Profiles, schemas.SchemaDatabaseError) {
	res, err := s.repository.ResultProfileRepository()
	return res, err
}
