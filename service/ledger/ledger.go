package ledger

import (
	"context"
	"fmt"
	"github.com/sreekar2307/katha/model/table"
	"github.com/sreekar2307/katha/repository"
	"github.com/sreekar2307/katha/service"
	"github.com/sreekar2307/katha/simplifier"
	"gorm.io/gorm"
)

type ledgerServ struct {
	simplifier simplifier.Simplifier
	primaryDB  *gorm.DB
	repository repository.Repository
}

func NewLedgerService(simplifier simplifier.Simplifier, primaryDB *gorm.DB, repository repository.Repository) service.Ledger {
	return ledgerServ{
		simplifier: simplifier,
		primaryDB:  primaryDB,
		repository: repository,
	}
}

func (l ledgerServ) GetUserInvolvedExpenses(ctx context.Context, userID, gtID uint64, limit int) ([]service.SimplifiedView, error) {
	ledgers, err := l.repository.GetUserInvolvedLedgers(ctx, l.primaryDB, userID, gtID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get user involved ledgers: %w", err)
	}
	if len(ledgers) == 0 {
		return nil, nil
	}
	partiesInvolved := make(map[uint64]bool)
	for _, ledger := range ledgers {
		partiesInvolved[ledger.LenderID] = true
		partiesInvolved[ledger.BorrowerID] = true
	}
	users, err := l.repository.GetUsersByIDs(ctx, l.primaryDB, partiesInvolved)
	if err != nil {
		return nil, fmt.Errorf("failed to get users by IDs: %w", err)
	}
	userForID := make(map[uint64]table.User)
	for _, user := range users {
		userForID[user.ID] = user
	}
	simplifiedViews := make([]service.SimplifiedView, 0)
	for _, ledger := range ledgers {
		simplifiedViews = append(simplifiedViews, service.SimplifiedView{
			ID:          ledger.ID,
			Lender:      userForID[ledger.LenderID],
			Borrower:    userForID[ledger.BorrowerID],
			Description: ledger.Expense.Description,
			Amount:      ledger.Amount,
		})
	}
	return simplifiedViews, nil
}

func (l ledgerServ) GetBalanceReport(ctx context.Context, userID uint64) ([]service.Owes, []service.Lends, error) {
	simplified, err := l.simplifier.Simplify(ctx, userID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to simplify ledgers: %w", err)
	}
	otherParties := make(map[uint64]bool)
	for lenderID, borrowers := range simplified {
		for borrowerID := range borrowers {
			otherParties[borrowerID] = true
		}
		otherParties[lenderID] = true
	}
	users, err := l.repository.GetUsersByIDs(ctx, l.primaryDB, otherParties)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get users by IDs: %w", err)
	}
	userForID := make(map[uint64]table.User)
	for _, user := range users {
		userForID[user.ID] = user
	}
	owes := make([]service.Owes, 0)
	lends := make([]service.Lends, 0)
	for lenderID, borrowers := range simplified {
		lender, _ := userForID[lenderID]
		for borrowerID, amount := range borrowers {
			borrower, _ := userForID[borrowerID]
			if lenderID == userID {
				lends = append(lends, service.Lends{
					Borrower: borrower,
					Amount:   amount,
				})
			} else if borrowerID == userID {
				owes = append(owes, service.Owes{
					Lender: lender,
					Amount: amount,
				})
			}
		}
	}
	return owes, lends, nil
}
