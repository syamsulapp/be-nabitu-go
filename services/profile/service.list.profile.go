package services

import (
	repositorys "be-nabitu-go/repositorys/profile"

	"github.com/gofiber/fiber/v2"
)

func GetRepositoryAllProfile(c *fiber.Ctx) error {
	return repositorys.RepositoryGetAllProfile(c)
}

func ResultServiceGetAllProfile(c *fiber.Ctx) error {
	return GetRepositoryAllProfile(c)
}
