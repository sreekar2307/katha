package splitter

import (
	"context"
	"github.com/sreekar2307/katha/model"

	"github.com/sreekar2307/katha/model/table"
)

// Splitter is an interface that defines a method to split expenses based on the provided split config
// and expense details.
type Splitter interface {
	Split(context.Context, []model.Split, table.Expense) ([]table.Ledger, error)
}
