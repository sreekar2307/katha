package v1

import (
	"github.com/sreekar2307/katha/controller/http"
	"github.com/sreekar2307/katha/service"
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
