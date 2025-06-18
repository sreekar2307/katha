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

// Ledger is an interface that defines methods for retrieving user expenses and balance reports.
type Ledger interface {
	// GetUserInvolvedExpenses retrieves a list of expenses that the user is involved in, starting from a given ID
	// and limited to a specified number of results.
	GetUserInvolvedExpenses(ctx context.Context, userID uint64, gtId uint64, limit int) ([]SimplifiedView, error)
	// GetBalanceReport retrieves a balance report for the user, showing amounts owed and lent by the user.
	GetBalanceReport(ctx context.Context, userID uint64) ([]Owes, []Lends, error)
}
