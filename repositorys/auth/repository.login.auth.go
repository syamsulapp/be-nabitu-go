package repository

import (
	"be-nabitu-go/database"
	"be-nabitu-go/helpers"
	"be-nabitu-go/models"
	"be-nabitu-go/pkg"
	"time"

	jtoken "github.com/golang-jwt/jwt/v4"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginRepository(c *fiber.Ctx) error {
	LoginRequest := new(models.LoginRequest)
	if err := c.BodyParser(LoginRequest); err != nil {
		return c.Status(422).JSON(helpers.ErrorResponse(c, err.Error()))
	}

	var user models.User
	database.GetDBMysql().Where("username = ?", LoginRequest.Username).First(&user)
	if user.ID == 0 {
		return c.Status(422).JSON(helpers.ErrorResponse(c, "Wrong Username"))
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(LoginRequest.Password)); err != nil {
		return c.Status(422).JSON(helpers.ErrorResponse(c, "Wrong Password"))
	}

	day := time.Hour * 24
	claimsToken := jtoken.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(day * 1).Unix(),
	}
	//create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claimsToken)
	t, err := token.SignedString([]byte(pkg.Secret))
	if err != nil {
		return c.Status(500).JSON(helpers.ErrorResponse(c, err.Error()))
	}
	res := fiber.Map{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"token":    t,
	}
	return c.JSON(helpers.SuccessResponse(c, "Berhasil Login", res))
}
