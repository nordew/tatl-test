package service

import (
	"context"
	"log/slog"

	"github.com/nordew/tatl-test/internal/models"
	"github.com/nordew/tatl-test/internal/storage"
)

type userService struct {
	userStorage storage.UserStorage
	log         *slog.Logger
}

func NewUserService(
	userStorage storage.UserStorage,
	log *slog.Logger,
) UserService {
	return &userService{
		userStorage: userStorage,
		log:         log,
	}
}

func (s *userService) Get(ctx context.Context, params *models.GetUserParams) ([]models.Profile, error) {
	const op = "userService.Get"

	log := s.log.With(
		slog.String("op", op),
		slog.Any("params", params),
	)

	users, err := s.userStorage.Get(ctx, params)
	if err != nil {
		log.Error("failed to get users", "error", err)
		return nil, err
	}

	return users, nil
}
