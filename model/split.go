package model

type SplitType string

var SplitTypes = struct {
	Equal      SplitType
	Percentage SplitType
	Amount     SplitType
}{
	Equal:      "EQUAL",
	Percentage: "PERCENTAGE",
	Amount:     "AMOUNT",
}

var SplitTypeValues = []SplitType{
	SplitTypes.Equal,
	SplitTypes.Percentage,
	SplitTypes.Amount,
}

type Split struct {
	Percentage float64
	Amount     uint64
	LenderID   uint64
	BorrowerID uint64
}
