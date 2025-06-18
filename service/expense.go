package service

import (
	"context"
	"github.com/sreekar2307/khata/model"
	"github.com/sreekar2307/khata/model/table"
)

// Expense is an interface that defines methods for managing expenses, including adding expenses
type Expense interface {
	// AddExpense adds a new expense with the given split type, splits, and expense details.
	AddExpense(
		context.Context,
		model.SplitType,
		[]model.Split,
		table.Expense,
	) (table.Expense, []table.Ledger, error)
}
