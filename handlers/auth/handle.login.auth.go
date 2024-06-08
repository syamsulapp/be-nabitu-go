package handlers

import (
	services "be-nabitu-go/services/auth"

	"github.com/gofiber/fiber/v2"
)

func GetServiceAuthLogin(c *fiber.Ctx) error {
	return services.ResultGetLoginServices(c)
}

func ResultHandleAuthLogin(c *fiber.Ctx) error {
	return GetServiceAuthLogin(c)
}
