package response

import "github.com/sreekar2307/katha/model/table"

type Ledger struct {
	Borrower User   `json:"borrower"`
	Lender   User   `json:"lender"`
	Amount   uint64 `json:"amountInPaise"`
}

func NewLedger(ledger table.Ledger) Ledger {
	return Ledger{
		Borrower: NewUser(ledger.Borrower),
		Lender:   NewUser(ledger.Lender),
		Amount:   ledger.Amount,
	}
}
