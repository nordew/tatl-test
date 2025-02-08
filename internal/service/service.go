package service

import (
	"context"

	"github.com/nordew/tatl-test/internal/models"
)

type (
	UserService interface {
		Get(ctx context.Context, params *models.GetUserParams) ([]models.Profile, error)
	}

	AuthService interface {
		VerifyAPIKey(ctx context.Context, apiKey string) error
	}
)
