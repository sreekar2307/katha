package v1

import (
	"github.com/sreekar2307/khata/controller/http"
	"github.com/sreekar2307/khata/service"
)

type controller struct {
	ExpenseService service.Expense
	UserService    service.User
	LedgerService  service.Ledger
}

func NewV1Controller(
	expenseService service.Expense,
	userService service.User,
	ledgerService service.Ledger,
) http.V1Controller {
	return &controller{
		ExpenseService: expenseService,
		LedgerService:  ledgerService,
		UserService:    userService,
	}
}
