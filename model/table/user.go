package table

type User struct {
	Base
	Email           string `gorm:"uniqueIndex;not null;type:varchar(255)"`
	PasswordHash    string `gorm:"not null;type:varchar(255)"`
	DefaultCurrency string `gorm:"not null;type:varchar(3);default:'INR'"`
	Password        string `gorm:"-"`
}
