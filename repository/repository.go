package repository

import (
	"context"
	"github.com/sreekar2307/khata/model/table"
	"gorm.io/gorm"
)

type repository struct{}

func NewRepository() Repository {
	return repository{}
}

type ExpenseRepository interface {
	CreateExpense(context.Context, *gorm.DB, *table.Expense) error
}

type LedgerRepository interface {
	CreateLedger(context.Context, *gorm.DB, *table.Ledger) error
	CreateLedgers(context.Context, *gorm.DB, *[]table.Ledger) error
	GetUserInvolvedLedgers(context.Context, *gorm.DB, uint64, uint64, int) ([]table.Ledger, error)
}

type UserRepository interface {
	GetUserByEmail(context.Context, *gorm.DB, string) (table.User, error)
	CreateUser(context.Context, *gorm.DB, *table.User) error
	GetUsersByIDs(context.Context, *gorm.DB, map[uint64]bool) ([]table.User, error)
}

type Repository interface {
	ExpenseRepository
	LedgerRepository
	UserRepository
}
