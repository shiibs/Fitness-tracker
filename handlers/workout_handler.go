package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiibs/fitness-app/models"
	"github.com/shiibs/fitness-app/service"
)

type WorkoutHandler struct {
	workoutService service.WorkoutService
}

func NewWorkoutHandler(router fiber.Router, workoutService service.WorkoutService) {
	handler := &WorkoutHandler{workoutService}
	workoutRouter := router.Group("/workout")

	workoutRouter.Post("/", handler.LogWorkout)

}

func (h *WorkoutHandler) LogWorkout(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	workoutEntry := new(models.WorkoutEntry)

	if err := ctx.BodyParser(workoutEntry); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	workoutEntry.UserID = userId

	if err := h.workoutService.LogWorkoutEntry(workoutEntry); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status": "fail",
			"error":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":    "success",
		"workEntry": workoutEntry,
	})
}
