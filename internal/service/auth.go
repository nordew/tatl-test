package service

import (
	"context"
	"github.com/nordew/tatl-test/internal/storage"
	"log/slog"
)

type authService struct {
	authStorage storage.AuthStorage
	logger      *slog.Logger
}

func NewAuthService(
	authStorage storage.AuthStorage,
	logger *slog.Logger,
) AuthService {
	return &authService{
		authStorage: authStorage,
		logger:      logger,
	}
}

func (s *authService) VerifyAPIKey(ctx context.Context, apiKey string) error {
	const op = "authService.VerifyAPIKey"

	log := s.logger.With(
		slog.String("op", op),
	)

	if err := s.authStorage.VerifyAPIKey(ctx, apiKey); err != nil {
		log.Error("failed to verify API key", "error", err)
		return err
	}

	return nil
}
