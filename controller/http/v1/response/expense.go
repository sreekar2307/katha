package response

import "github.com/sreekar2307/khata/model/table"

type Expense struct {
	Amount      uint64  `json:"amountInPaise"`
	Description string  `json:"description"`
	Lender      User    `json:"lender"`
	Splits      []Split `json:"splits,omitempty"`
}

func NewExpense(
	expense table.Expense,
	lender table.User,
	ledgers []table.Ledger,
) Expense {
	splits := make([]Split, 0, len(ledgers))
	for _, ledger := range ledgers {
		splits = append(splits, NewSplit(NewUser(ledger.Borrower), ledger.Amount))
	}
	return Expense{
		Amount:      expense.Amount,
		Lender:      NewUser(lender),
		Description: expense.Description,
		Splits:      splits,
	}
}
