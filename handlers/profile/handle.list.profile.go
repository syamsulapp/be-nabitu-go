package handlers

import (
	services "be-nabitu-go/services/profile"

	"github.com/gofiber/fiber/v2"
)

func GetServiceAllProfile(c *fiber.Ctx) error {
	return services.ResultServiceGetAllProfile(c)
}

func ResultHandlerGetAllProfile(c *fiber.Ctx) error {
	return GetServiceAllProfile(c)
}
