package storage

import (
	"context"

	"github.com/nordew/tatl-test/internal/models"
	"gorm.io/gorm"
)

type userStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) UserStorage {
	return &userStorage{db: db}
}

func (s *userStorage) Get(ctx context.Context, params *models.GetUserParams) ([]models.Profile, error) {
	var profiles []models.Profile

	db := s.db.WithContext(ctx).Table("app_user as u").
		Select("u.id, u.username, up.first_name, up.last_name, up.city, ud.school").
		Joins("JOIN user_profile as up ON u.id = up.user_id").
		Joins("JOIN user_data as ud ON u.id = ud.user_id")

	if params.Username != nil && *params.Username != "" {
		db = db.Where("u.username = ?", *params.Username)
	}

	if err := db.Scan(&profiles).Error; err != nil {
		return nil, err
	}

	return profiles, nil
}
