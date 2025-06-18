package expense

import (
	"context"
	stdErrors "errors"
	"fmt"

	"github.com/sreekar2307/khata/model"
	"github.com/sreekar2307/khata/model/table"
	"github.com/sreekar2307/khata/repository"
	"github.com/sreekar2307/khata/service"
	splitterFactory "github.com/sreekar2307/khata/splitter/factory"
	"gorm.io/gorm"
)

type expense struct {
	splitterFactory splitterFactory.SplitterFactory
	primaryDB       *gorm.DB
	repository      repository.Repository
}

func NewExpenseService(splitterFactory splitterFactory.SplitterFactory, primaryDB *gorm.DB, repository repository.Repository) service.Expense {
	return expense{
		splitterFactory: splitterFactory,
		primaryDB:       primaryDB,
		repository:      repository,
	}
}

func (e expense) AddExpense(
	ctx context.Context,
	splitType model.SplitType,
	splits []model.Split,
	expense table.Expense,
) (table.Expense, []table.Ledger, error) {
	splitter, err := e.splitterFactory.NewSplitter(splitType)
	if err != nil {
		return expense, nil, stdErrors.Join(err, fmt.Errorf("failed to create splitter for split type %s", splitType))
	}
	tx := e.primaryDB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := e.repository.CreateExpense(ctx, tx, &expense); err != nil {
		return expense, nil, stdErrors.Join(err, fmt.Errorf("failed to create expense"))
	}
	ledgers, err := splitter.Split(ctx, splits, expense)
	if err != nil {
		return expense, nil, stdErrors.Join(err, fmt.Errorf("failed to split expense"))
	}
	userIDs := make(map[uint64]bool)
	for _, ledger := range ledgers {
		userIDs[ledger.LenderID] = true
		userIDs[ledger.BorrowerID] = true
	}
	users, err := e.repository.GetUsersByIDs(ctx, tx, userIDs)
	if err != nil {
		return expense, nil, stdErrors.Join(err, fmt.Errorf("failed to get users by IDs"))
	}
	userForID := make(map[uint64]table.User)
	for _, user := range users {
		userForID[user.ID] = user
	}
	if err := e.repository.CreateLedgers(ctx, tx, &ledgers); err != nil {
		return expense, nil, stdErrors.Join(err, fmt.Errorf("failed to create ledgers"))
	}
	for i := range ledgers {
		ledgers[i].Lender = userForID[ledgers[i].LenderID]
		ledgers[i].Borrower = userForID[ledgers[i].BorrowerID]
	}
	return expense, ledgers, tx.Commit().Error
}
