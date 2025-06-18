package equal

import (
	"context"
	"testing"

	"github.com/sreekar2307/khata/errors"
	"github.com/sreekar2307/khata/model"
	"github.com/sreekar2307/khata/model/table"
	"github.com/stretchr/testify/assert"
)

func TestEqualSplitter_Split(t *testing.T) {
	tests := []struct {
		name        string
		expense     table.Expense
		splits      []model.Split
		expected    []table.Ledger
		expectError error
	}{
		{
			name: "Valid equal split between two users",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2},
				{BorrowerID: 3},
			},
			expected: []table.Ledger{
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 2,
					Amount:     50,
				},
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 3,
					Amount:     50,
				},
			},
			expectError: nil,
		},
		{
			name: "Valid equal split between three users",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   90,
			},
			splits: []model.Split{
				{BorrowerID: 2},
				{BorrowerID: 3},
				{BorrowerID: 4},
			},
			expected: []table.Ledger{
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 2,
					Amount:     30,
				},
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 3,
					Amount:     30,
				},
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 4,
					Amount:     30,
				},
			},
			expectError: nil,
		},
		{
			name: "Amount not divisible by number of splits",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2},
				{BorrowerID: 3},
				{BorrowerID: 4},
			},
			expected:    nil,
			expectError: errors.ErrInvalidSplitConfiguration,
		},
		{
			name: "Single user split",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2},
			},
			expected: []table.Ledger{
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 2,
					Amount:     100,
				},
			},
			expectError: nil,
		},
		{
			name: "Zero amount expense",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   0,
			},
			splits: []model.Split{
				{BorrowerID: 2},
				{BorrowerID: 3},
			},
			expected: []table.Ledger{
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 2,
					Amount:     0,
				},
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 3,
					Amount:     0,
				},
			},
			expectError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitter := NewEqualSplitter()
			result, err := splitter.Split(context.Background(), tt.splits, tt.expense)

			if tt.expectError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.expectError)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
