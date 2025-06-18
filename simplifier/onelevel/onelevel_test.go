package onelevel

import (
	"context"
	"testing"

	"github.com/sreekar2307/khata/model/table"
	repoMocks "github.com/sreekar2307/khata/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOneLevelSimplifier_ComplexScenarios(t *testing.T) {
	tests := []struct {
		name     string
		ledgers  []table.Ledger
		userID   uint64
		expected map[uint64]map[uint64]uint64
	}{
		{
			name: "Scenario 1: Multiple transactions between three users",
			ledgers: []table.Ledger{
				{LenderID: 1, BorrowerID: 2, Amount: 100}, // A -> B: 100
				{LenderID: 2, BorrowerID: 1, Amount: 30},  // B -> A: 30
				{LenderID: 2, BorrowerID: 3, Amount: 50},  // B -> C: 50
				{LenderID: 3, BorrowerID: 1, Amount: 20},  // C -> A: 20
			},
			userID: 1,
			expected: map[uint64]map[uint64]uint64{
				1: {2: 70}, // A -> B: 70 (100 - 30)
				2: {3: 50}, // B -> C: 50
				3: {1: 20}, // C -> A: 20
			},
		},
		{
			name: "Scenario 2: Circular transactions",
			ledgers: []table.Ledger{
				{LenderID: 1, BorrowerID: 2, Amount: 100}, // A -> B: 100
				{LenderID: 2, BorrowerID: 3, Amount: 100}, // B -> C: 100
				{LenderID: 3, BorrowerID: 1, Amount: 100}, // C -> A: 100
			},
			userID: 1,
			expected: map[uint64]map[uint64]uint64{
				1: {2: 100}, // A -> B: 100
				2: {3: 100}, // B -> C: 100
				3: {1: 100}, // C -> A: 100
			},
		},
		{
			name: "Scenario 3: Multiple transactions with same users",
			ledgers: []table.Ledger{
				{LenderID: 1, BorrowerID: 2, Amount: 50}, // A -> B: 50
				{LenderID: 1, BorrowerID: 2, Amount: 30}, // A -> B: 30
				{LenderID: 2, BorrowerID: 1, Amount: 40}, // B -> A: 40
				{LenderID: 2, BorrowerID: 1, Amount: 20}, // B -> A: 20
			},
			userID: 1,
			expected: map[uint64]map[uint64]uint64{
				1: {2: 20}, // A -> B: 20 (80 - 60)
			},
		},
		{
			name: "Scenario 4: Zero net transactions",
			ledgers: []table.Ledger{
				{LenderID: 1, BorrowerID: 2, Amount: 100}, // A -> B: 100
				{LenderID: 2, BorrowerID: 1, Amount: 100}, // B -> A: 100
			},
			userID:   1,
			expected: map[uint64]map[uint64]uint64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock repository
			mockRepo := repoMocks.NewMockLedgerRepository(t)
			mockRepo.On(
				"GetUserInvolvedLedgers",
				mock.Anything,
				mock.Anything,
				tt.userID,
				uint64(0),
				0,
			).Return(tt.ledgers, nil)

			// Create simplifier
			simplifier := NewOneLevelSimplifier(nil, mockRepo)

			// Execute simplification
			result, err := simplifier.Simplify(context.Background(), tt.userID)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
