package table

type Ledger struct {
	Base
	Lender     User    `gorm:"foreignKey:LenderID;references:ID"`
	LenderID   uint64  `gorm:"not null;index"`
	Borrower   User    `gorm:"foreignKey:BorrowerID;references:ID"`
	BorrowerID uint64  `gorm:"not null;index"`
	Amount     uint64  `gorm:"not null"`
	Expense    Expense `gorm:"foreignKey:ExpenseID;references:ID"`
	Currency   string  `gorm:"not null;type:varchar(3);default:'INR'"`
	ExpenseID  uint64  `gorm:"not null"`
}
