package v1

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nordew/tatl-test/internal/models"
)

func (h *Handler) getProfile(c fiber.Ctx) error {
	username := c.Query("username")

	user, err := h.userService.Get(c.Context(), &models.GetUserParams{Username: &username})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get users"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
