package services

import (
	repositorys "be-nabitu-go/repositorys/profile"

	"github.com/gofiber/fiber/v2"
)

func GetRepositoryProfileByID(c *fiber.Ctx) error {
	return repositorys.RepositoryGetProfileByID(c)
}

func ResultServiceGetProfileByID(c *fiber.Ctx) error {
	return GetRepositoryProfileByID(c)
}
