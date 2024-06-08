package services

import (
	repositoryAuth "be-nabitu-go/repositorys/auth"

	"github.com/gofiber/fiber/v2"
)

func GetLoginRepository(c *fiber.Ctx) error {
	return repositoryAuth.LoginRepository(c)
}

func ResultGetLoginServices(c *fiber.Ctx) error {
	return GetLoginRepository(c)
}
