package repository

import (
	"be-nabitu-go/database"
	"be-nabitu-go/helpers"
	"be-nabitu-go/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterRepository(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	database.GetDBMysql().Create(&user)

	return c.Status(200).JSON(helpers.SuccessResponse(c, "Berhasil Registrasi", user))

}
