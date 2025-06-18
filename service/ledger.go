package service

import (
	"context"
	"github.com/sreekar2307/katha/model/table"
)

type SimplifiedView struct {
	ID          uint64
	Lender      table.User
	Borrower    table.User
	Description string
	Amount      uint64
}

type Owes struct {
	Lender table.User
	Amount uint64
}

type Lends struct {
	Borrower table.User
	Amount   uint64
}

type Ledger interface {
	GetUserInvolvedExpenses(ctx context.Context, userID uint64, gtId uint64, limit int) ([]SimplifiedView, error)
	GetBalanceReport(ctx context.Context, userID uint64) ([]Owes, []Lends, error)
}
