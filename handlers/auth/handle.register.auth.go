package handlers

import (
	servicesAuth "be-nabitu-go/services/auth"

	"github.com/gofiber/fiber/v2"
)

func GetServiceAuthRegister(c *fiber.Ctx) error {
	return servicesAuth.ResultGetRegisterServices(c)
}

func ResultHandleAuthRegister(c *fiber.Ctx) error {
	return GetServiceAuthRegister(c)
}
