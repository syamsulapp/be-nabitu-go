package repository

import (
	"be-nabitu-go/configs"
	"be-nabitu-go/database"
	"be-nabitu-go/helpers"
	"be-nabitu-go/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterRepository(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	initTimezone := configs.GetConfigTimeZone(time.Now())
	formatTimeZone := initTimezone.Format(os.Getenv("GO_TIMEZONE_FORMAT"))
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Username:   data["username"],
		Email:      data["email"],
		Password:   password,
		Created_at: formatTimeZone,
		Updated_at: formatTimeZone,
	}

	err := database.GetDBMysql().Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(helpers.ErrorResponse(c, err.Error()))
	} else {
		return c.Status(200).JSON(helpers.SuccessResponse(c, "Berhasil Registrasi", user))
	}
}
