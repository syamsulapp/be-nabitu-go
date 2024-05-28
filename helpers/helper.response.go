package helpers

import (
	"be-nabitu-go/schemas"

	"github.com/gofiber/fiber/v2"
)

func APIResponse(ctx *fiber.Ctx, Message string, StatusCode int, Method string, Data interface{}) schemas.SchemaResponses {

	return schemas.SchemaResponses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}

}
