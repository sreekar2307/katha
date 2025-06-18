package splitter

import (
	"context"
	"github.com/sreekar2307/katha/model"

	"github.com/sreekar2307/katha/model/table"
)

type Splitter interface {
	Split(context.Context, []model.Split, table.Expense) ([]table.Ledger, error)
}
