package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

var (
	ErrAPIKeyInvalid = errors.New("api key is invalid")
)

type authStorage struct {
	db *gorm.DB
}

func NewAuthStorage(db *gorm.DB) AuthStorage {
	return &authStorage{db: db}
}

func (s *authStorage) VerifyAPIKey(ctx context.Context, apiKey string) error {
	var count int64

	if err := s.db.WithContext(ctx).Table("auth").Where("`api-key` = ?", apiKey).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return ErrAPIKeyInvalid
	}

	return nil
}
