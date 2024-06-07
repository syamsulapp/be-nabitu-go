package handlers

import (
	services "be-nabitu-go/services/profile"

	"github.com/gofiber/fiber/v2"
)

func GetServiceCreateProfile(c *fiber.Ctx) error {
	return services.ResultServiceCreateProfile(c)
}

func ResultHandlerCreateProfile(c *fiber.Ctx) error {
	return GetServiceCreateProfile(c)
}
