package repository

import (
	"context"
	stdErrors "errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sreekar2307/katha/errors"
	"github.com/sreekar2307/katha/model/table"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r repository) CreateLedger(ctx context.Context, db *gorm.DB, ledger *table.Ledger) error {
	return db.WithContext(ctx).Clauses(clause.Returning{}).Create(ledger).Error
}

func (r repository) CreateLedgers(ctx context.Context, db *gorm.DB, ledgers *[]table.Ledger) error {
	if err := db.WithContext(ctx).Clauses(clause.Returning{}).CreateInBatches(&ledgers, 10).Error; err != nil {
		var pgErr *pgconn.PgError
		if stdErrors.As(err, &pgErr) && pgErr.Code == "23503" {
			return errors.ErrCheckIfDependencyExists
		} else if stdErrors.Is(err, gorm.ErrForeignKeyViolated) {
			return errors.ErrCheckIfDependencyExists
		}
		return err
	}
	return nil
}

func (r repository) GetUserInvolvedLedgers(
	ctx context.Context,
	db *gorm.DB,
	userID uint64,
	gtID uint64,
	limit int,
) ([]table.Ledger, error) {
	ledgers := make([]table.Ledger, 0)
	userInvolved := db.WithContext(ctx).
		Where(table.Ledger{
			LenderID: userID,
		}).
		Or(table.Ledger{
			BorrowerID: userID,
		})
	notSelf := db.WithContext(ctx).Where("lender_id != borrower_id")
	query := db.WithContext(ctx).
		Preload("Expense").
		Where(userInvolved).
		Where(notSelf)
	if gtID > 0 {
		query = query.Where("id > ?", gtID)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	if err := query.Find(&ledgers).Error; err != nil {
		return nil, fmt.Errorf("failed to get user involved ledgers: %w", err)
	}
	return ledgers, nil
}
