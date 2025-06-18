package repository

import (
	"context"
	"github.com/sreekar2307/khata/model/table"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r repository) CreateExpense(ctx context.Context, db *gorm.DB, expense *table.Expense) error {
	return db.WithContext(ctx).Clauses(clause.Returning{}).Create(expense).Error
}
