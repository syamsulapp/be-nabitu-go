package repositorys

import (
	"be-nabitu-go/configs"
	"be-nabitu-go/database"
	"be-nabitu-go/helpers"
	"be-nabitu-go/models"
	"os"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

func RepositoryCreateProfile(c *fiber.Ctx) error {
	CreateProfile := new(models.Profile)
	if err := c.BodyParser(&CreateProfile); err != nil {
		return c.Status(422).JSON(helpers.ErrorResponse(c, err.Error()))
	}

	//validation
	ValidateCreateProfile := validator.New()
	ErrorValidateCreateProfile := ValidateCreateProfile.Struct(CreateProfile)
	if ErrorValidateCreateProfile != nil {
		return c.Status(422).JSON(helpers.ErrorValidation(c, "Data Tidak Lengkap", ErrorValidateCreateProfile.Error()))
	}

	InitTimeZone := configs.GetConfigTimeZone(time.Now())
	FormatTimeZone := InitTimeZone.Format(os.Getenv("GO_TIMEZONE_FORMAT"))
	Profile := models.Profile{
		Fullname:   CreateProfile.Fullname,
		Age:        CreateProfile.Age,
		Alamat:     CreateProfile.Alamat,
		Users_id:   CreateProfile.Users_id,
		Created_at: FormatTimeZone,
		Updated_at: FormatTimeZone,
	}

	err := database.GetDBMysql().Create(&Profile).Error
	if err != nil {
		return c.Status(500).JSON(helpers.ErrorResponse(c, err.Error()))
	} else {
		return c.Status(200).JSON(helpers.SuccessResponse(c, "Berhasil Create Profile", Profile))
	}
}
