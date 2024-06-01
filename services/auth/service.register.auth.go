package services

import (
	repositoryAuth "be-nabitu-go/repositorys/auth"

	"github.com/gofiber/fiber/v2"
)

func GetRegisterRepository(c *fiber.Ctx) error {
	return repositoryAuth.ResultGetRegisterRepository(c)
}

func ResultGetRegisterServices(c *fiber.Ctx) error {
	return GetRegisterRepository(c)
}
