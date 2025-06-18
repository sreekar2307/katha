package amount

import (
	"context"
	stdErrors "errors"
	"fmt"
	"github.com/sreekar2307/khata/model"

	"github.com/sreekar2307/khata/errors"
	"github.com/sreekar2307/khata/model/table"
	"github.com/sreekar2307/khata/splitter"
)

func NewAmountSplitter() splitter.Splitter {
	return amountSplitter{}
}

// amountSplitter implements the Splitter interface for splitting expenses based on specified amounts.
type amountSplitter struct{}

func (s amountSplitter) Split(_ context.Context, splits []model.Split, expense table.Expense) ([]table.Ledger, error) {
	var ledgers []table.Ledger
	for _, split := range splits {
		if split.Amount == 0 {
			continue
		}
		if split.Amount > expense.Amount {
			return nil, stdErrors.Join(errors.ErrInvalidSplitConfiguration,
				fmt.Errorf("split amount %d is greater than expense amount %d", split.Amount, expense.Amount))
		}
		expense.Amount -= split.Amount
		ledger := table.Ledger{
			ExpenseID:  expense.ID,
			LenderID:   expense.LenderID,
			BorrowerID: split.BorrowerID,
			Amount:     split.Amount,
		}
		ledgers = append(ledgers, ledger)
	}

	return ledgers, nil
}
