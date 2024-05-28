package handlers

import (
	"be-nabitu-go/helpers"
	services "be-nabitu-go/services/profile"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type callService struct {
	service services.ServiceGetAllProfile
}

func NewHandlerGetAllProfile(service services.ServiceGetAllProfile) *callService {
	return &callService{service: service}
}

func (h *callService) ResultProfileHandler(ctx *fiber.Ctx) error {
	res, err := h.service.ResultProfileService()

	if err.Type == "error_01" {
		helpers.APIResponse(ctx, "Profile Data Does not exits", err.Code, http.MethodPost, nil)
	}

	return ctx.JSON(helpers.APIResponse(ctx, "Result Profile Data Successfully", http.StatusOK, http.MethodPost, res))

}
