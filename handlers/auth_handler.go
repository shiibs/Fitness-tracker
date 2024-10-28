package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiibs/fitness-app/service"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(router fiber.Router, authservice service.AuthService) {
	handler := &AuthHandler{service: authservice}

	router.Post("/login", handler.Login)
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	loginRequest := new(struct {
		GoogleID string `json:"google_id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		ImageURL string `json:"image_url"`
	})

	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	user, token, err := h.service.HandleLogin(loginRequest.GoogleID, loginRequest.Name, loginRequest.Email, loginRequest.ImageURL)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Succesfully logged in",
		"data": &fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}
