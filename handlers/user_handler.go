package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiibs/fitness-app/models"
	"github.com/shiibs/fitness-app/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(router fiber.Router, userService service.UserService) {
	handler := &UserHandler{userService}
	userRouter := router.Group("/user")

	userRouter.Put("/", handler.CreateOrUpdateUser)
}

func (h *UserHandler) CreateOrUpdateUser(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	userRequest := new(models.User)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	user, err := h.userService.UpdateUser(userId, userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"user":   user,
	})
}
