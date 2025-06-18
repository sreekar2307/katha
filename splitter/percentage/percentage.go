package percentage

import (
	"context"
	stdErrors "errors"
	"fmt"
	"github.com/sreekar2307/katha/model"

	"github.com/sreekar2307/katha/errors"
	"github.com/sreekar2307/katha/model/table"
	"github.com/sreekar2307/katha/splitter"
)

func NewPercentageSplitter() splitter.Splitter {
	return percentageSplitter{}
}

// percentageSplitter implements the Splitter interface for splitting expenses based on percentages.
type percentageSplitter struct{}

func (s percentageSplitter) Split(_ context.Context, splits []model.Split, expense table.Expense) ([]table.Ledger, error) {
	var (
		ledgers         []table.Ledger
		totalPercentage float64
	)

	for _, split := range splits {
		if split.Percentage <= 0 || split.Percentage > 100 {
			return nil, stdErrors.Join(errors.ErrInvalidSplitConfiguration,
				fmt.Errorf("split percentage %f is not gt 0 and lte 100", split.Percentage))
		}
		totalPercentage += split.Percentage

		if totalPercentage > 100 {
			return nil, errors.ErrInvalidSplitConfiguration
		}
		amount := (float64(expense.Amount) * split.Percentage) / 100
		if amount != float64(uint64(amount)) {
			return nil, stdErrors.Join(errors.ErrInvalidSplitConfiguration,
				fmt.Errorf("split percentage %f results in non-integer amount %f for expense amount %d",
					split.Percentage, amount, expense.Amount))
		}
		ledger := table.Ledger{
			ExpenseID:  expense.ID,
			LenderID:   expense.LenderID,
			BorrowerID: split.BorrowerID,
			Amount:     uint64(amount),
		}
		ledgers = append(ledgers, ledger)
	}
	if totalPercentage < 100 {
		return nil, stdErrors.Join(errors.ErrInvalidSplitConfiguration,
			fmt.Errorf("total split percentage is less than 100%%"))
	}
	return ledgers, nil
}
