package response

import (
	"github.com/sreekar2307/khata/service"
)

type SimplifiedView struct {
	ID          uint64 `json:"id"`
	Lender      User   `json:"lender"`
	Borrower    User   `json:"borrower"`
	Description string `json:"description"`
	Amount      uint64 `json:"amountInPaise"` // Amount in paise
}

type Owe struct {
	Lender User   `json:"lender"`
	Amount uint64 `json:"amountInPaise"` // Amount in paise
}

type Lend struct {
	Borrower User   `json:"borrower"`
	Amount   uint64 `json:"amountInPaise"` // Amount in paise
}

func NewSimplifiedView(view service.SimplifiedView) SimplifiedView {
	return SimplifiedView{
		ID:          view.ID,
		Lender:      NewUser(view.Lender),
		Borrower:    NewUser(view.Borrower),
		Description: view.Description,
		Amount:      view.Amount,
	}
}
func NewOwe(owe service.Owes) Owe {
	return Owe{
		Lender: NewUser(owe.Lender),
		Amount: owe.Amount,
	}

}
func NewLend(lend service.Lends) Lend {
	return Lend{
		Borrower: NewUser(lend.Borrower),
		Amount:   lend.Amount,
	}
}
