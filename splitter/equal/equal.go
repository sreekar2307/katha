package equal

import (
	"context"
	stdErrors "errors"
	"fmt"
	"github.com/sreekar2307/katha/model"

	"github.com/sreekar2307/katha/errors"
	"github.com/sreekar2307/katha/model/table"
	"github.com/sreekar2307/katha/splitter"
)

func NewEqualSplitter() splitter.Splitter {
	return equalSplitter{}
}

// equalSplitter implements the Splitter interface for splitting expenses equally among users.
type equalSplitter struct{}

func (s equalSplitter) Split(_ context.Context, splits []model.Split, expense table.Expense) ([]table.Ledger, error) {
	var ledgers []table.Ledger
	if int(expense.Amount)%len(splits) != 0 {
		return nil, stdErrors.Join(errors.ErrInvalidSplitConfiguration,
			fmt.Errorf("expense amount %d is not divisible by number of splits %d", expense.Amount, len(splits)))
	}
	perUserAmount := expense.Amount / uint64(len(splits))
	for _, split := range splits {
		ledger := table.Ledger{
			ExpenseID:  expense.ID,
			LenderID:   expense.LenderID,
			BorrowerID: split.BorrowerID,
			Amount:     perUserAmount,
		}
		ledgers = append(ledgers, ledger)
	}
	return ledgers, nil
}
