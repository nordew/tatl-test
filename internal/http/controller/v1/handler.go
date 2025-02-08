package v1

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/nordew/tatl-test/internal/service"
)

var (
	ErrInvalidAPIKey = errors.New("invalid API key")
	ErrInternalError = errors.New("internal server error")
	ErrAPIKeyMissing = errors.New("API key is missing")
)

type Handler struct {
	userService service.UserService
	authService service.AuthService
}

func NewHandler(
	userService service.UserService,
	authService service.AuthService,
) *Handler {
	return &Handler{
		userService: userService,
		authService: authService,
	}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/api")

	v1.Use(h.AuthMiddleware())
	v1.Get("/profile", h.getProfile)
}
