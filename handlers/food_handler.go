package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiibs/fitness-app/models"
	"github.com/shiibs/fitness-app/service"
)

type FoodHandler struct {
	foodService service.FoodService
}

func NewFoodHandler(router fiber.Router, foodService service.FoodService) {
	handler := &FoodHandler{foodService}
	foodHadler := router.Group("/food")

	foodHadler.Post("/", handler.LogFoodEntry)
}

func (h *FoodHandler) LogFoodEntry(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	foodEntry := new(models.FoodEntry)

	if err := ctx.BodyParser(foodEntry); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	foodEntry.UserID = userId

	if err := h.foodService.LogFoodEntry(foodEntry); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":    "success",
		"foodEntry": foodEntry,
	})
}
