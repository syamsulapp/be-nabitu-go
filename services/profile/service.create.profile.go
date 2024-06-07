package services

import (
	repositorys "be-nabitu-go/repositorys/profile"

	"github.com/gofiber/fiber/v2"
)

func GetRepositoryCreateProfile(c *fiber.Ctx) error {
	return repositorys.RepositoryCreateProfile(c)
}

func ResultServiceCreateProfile(c *fiber.Ctx) error {
	return GetRepositoryCreateProfile(c)
}
