package handlers

import (
	services "be-nabitu-go/services/profile"

	"github.com/gofiber/fiber/v2"
)

func GetServiceProfileByID(c *fiber.Ctx) error {
	return services.ResultServiceGetProfileByID(c)
}

func ResultHandlerGetProfileByID(c *fiber.Ctx) error {
	return GetServiceProfileByID(c)
}
