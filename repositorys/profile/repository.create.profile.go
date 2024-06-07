package repositorys

import (
	"be-nabitu-go/database"
	"be-nabitu-go/models"

	"github.com/gofiber/fiber/v2"
)

func RepositoryCreateProfile(c *fiber.Ctx) error {
	var Profile []models.Profile
	database.GetDBMysql().Find(&Profile)
	return c.JSON(Profile)
}
