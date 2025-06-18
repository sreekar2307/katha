package service

import (
	"context"
	"github.com/sreekar2307/katha/model"
	"github.com/sreekar2307/katha/model/table"
)

type Expense interface {
	AddExpense(
		context.Context,
		model.SplitType,
		[]model.Split,
		table.Expense,
	) (table.Expense, []table.Ledger, error)
}
