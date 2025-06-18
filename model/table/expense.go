package table

type Expense struct {
	Base
	Amount      uint64 `gorm:"not null"`
	Currency    string `gorm:"not null;type:varchar(3);default:'INR'"`
	Lender      User   `gorm:"foreignKey:LenderID;references:ID"`
	LenderID    uint64 `gorm:"not null;index"`
	Description string `gorm:"text;null"`
}
