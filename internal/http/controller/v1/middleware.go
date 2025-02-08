package v1

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/nordew/tatl-test/internal/storage"
)

const (
	HeaderApiKey = "Api-key"
)

func (h *Handler) AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		apiKey := c.Get(HeaderApiKey)
		if apiKey == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": ErrAPIKeyMissing,
			})
		}

		if err := h.authService.VerifyAPIKey(context.Background(), apiKey); err != nil {
			if errors.Is(err, storage.ErrAPIKeyInvalid) {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": ErrInvalidAPIKey,
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": ErrInternalError,
			})
		}

		return c.Next()
	}
}
