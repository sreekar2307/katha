package response

type Split struct {
	Borrower User   `json:"borrower"`
	Amount   uint64 `json:"amountInPaise"`
}

func NewSplit(borrower User, amount uint64) Split {
	return Split{
		Borrower: borrower,
		Amount:   amount,
	}
}
