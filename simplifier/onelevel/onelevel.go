package onelevel

import (
	"context"
	"fmt"
	"github.com/sreekar2307/khata/repository"
	"github.com/sreekar2307/khata/simplifier"
	"gorm.io/gorm"
	"maps"
)

type oneLevelSimplifier struct {
	primaryDB *gorm.DB
	repo      repository.LedgerRepository
}

func NewOneLevelSimplifier(db *gorm.DB, repo repository.LedgerRepository) simplifier.Simplifier {
	return oneLevelSimplifier{
		primaryDB: db,
		repo:      repo,
	}
}

func (o oneLevelSimplifier) Simplify(ctx context.Context, userID uint64) (map[uint64]map[uint64]uint64, error) {
	ledgers, err := o.repo.GetUserInvolvedLedgers(ctx, o.primaryDB, userID, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get user involved ledgers: %w", err)
	}
	moneyLent := make(map[uint64]map[uint64]uint64)
	for _, l := range ledgers {
		from, to := l.LenderID, l.BorrowerID

		// from -> to
		if _, ok := moneyLent[from]; !ok {
			moneyLent[from] = make(map[uint64]uint64)
		}
		moneyLent[from][to] += l.Amount
	}
	for partyA, borrowers := range moneyLent {
		for partyB, amount := range borrowers {
			if _, ok := moneyLent[partyB]; ok {
				if _, ok := moneyLent[partyB][partyA]; ok {
					// If both parties have sent money to each other, we need to settle the amounts
					if moneyLent[partyB][partyA] >= amount {
						// if partyB sent more or equal to partyA, we can settle the amount
						moneyLent[partyB][partyA] -= amount
						moneyLent[partyA][partyB] = 0
					} else {
						// if partyA sent more than partyB, we can settle the amount
						moneyLent[partyA][partyB] -= moneyLent[partyB][partyA]
						moneyLent[partyB][partyA] = 0
					}
				}
			}
		}
	}
	simplified := maps.Clone(moneyLent)
	for partyA, borrowers := range moneyLent {
		for partyB, amount := range borrowers {
			if amount == 0 {
				delete(simplified[partyA], partyB)
			}
		}
		if len(simplified[partyA]) == 0 {
			delete(simplified, partyA)
		}
	}
	return simplified, nil
}
