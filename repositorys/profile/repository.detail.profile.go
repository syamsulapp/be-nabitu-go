package repositorys

import (
	"be-nabitu-go/database"
	"be-nabitu-go/helpers"
	"be-nabitu-go/models"

	"github.com/gofiber/fiber/v2"
)

func RepositoryGetProfileByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var Profiles models.Profile
	result := database.GetDBMysql().Table("profile").Find(&Profiles, id)

	if result.RowsAffected != 0 {
		return c.Status(200).JSON(helpers.SuccessResponse(c, "Successfully Data Details", Profiles))
	}
	return c.Status(422).JSON(helpers.ErrorResponse(c, "Data Tidak Di Temukan"))

}
