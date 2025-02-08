package storage

import (
	"context"

	"github.com/nordew/tatl-test/internal/models"
)

type (
	UserStorage interface {
		Get(ctx context.Context, params *models.GetUserParams) ([]models.Profile, error)
	}

	AuthStorage interface {
		VerifyAPIKey(ctx context.Context, apiKey string) error
	}
)
