package repository

import (
	"context"
	stdErrors "errors"
	"fmt"
	"github.com/sreekar2307/khata/errors"
	"github.com/sreekar2307/khata/model/table"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r repository) GetUserByEmail(ctx context.Context, db *gorm.DB, s string) (table.User, error) {
	var user table.User
	if err := db.WithContext(ctx).Where(table.User{
		Email: s,
	}).First(&user).Error; err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return table.User{}, errors.ErrUserNotFound
		}
		return table.User{}, fmt.Errorf("user with email %s not found", s)
	}
	return user, nil
}

func (r repository) CreateUser(ctx context.Context, db *gorm.DB, user *table.User) error {
	return db.WithContext(ctx).Clauses(clause.Returning{}).Create(user).Error
}

func (r repository) GetUsersByIDs(ctx context.Context, db *gorm.DB, m map[uint64]bool) ([]table.User, error) {
	var users []table.User
	userIDs := make([]uint64, 0, len(m))
	for id := range m {
		userIDs = append(userIDs, id)
	}
	if len(userIDs) == 0 {
		return nil, fmt.Errorf("no user IDs provided")
	}
	if err := db.WithContext(ctx).Where("id IN ?", userIDs).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
