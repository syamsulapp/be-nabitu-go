package helpers

import (
	"be-nabitu-go/schemas"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(ctx *fiber.Ctx, Message string) schemas.ErrorResponse {
	return schemas.ErrorResponse{
		Message: Message,
	}
}

func ErrorValidation(ctx *fiber.Ctx, Message string, Errors interface{}) schemas.ErrorValidations {
	return schemas.ErrorValidations{
		Message: Message,
		Error:   Errors,
	}
}

func SuccessResponse(ctx *fiber.Ctx, Message string, Data interface{}) schemas.SuccessResponse {
	return schemas.SuccessResponse{
		Message: Message,
		Data:    Data,
	}

}
