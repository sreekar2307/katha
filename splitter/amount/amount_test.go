package amount

import (
	"context"
	"testing"

	"github.com/sreekar2307/khata/errors"
	"github.com/sreekar2307/khata/model"
	"github.com/sreekar2307/khata/model/table"
	"github.com/stretchr/testify/assert"
)

func TestAmountSplitter_Split(t *testing.T) {
	tests := []struct {
		name        string
		expense     table.Expense
		splits      []model.Split
		expected    []table.Ledger
		expectError error
	}{
		{
			name: "Valid split with equal amounts",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Amount: 50},
				{BorrowerID: 3, Amount: 50},
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
			name: "Valid split with different amounts",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Amount: 70},
				{BorrowerID: 3, Amount: 30},
			},
			expected: []table.Ledger{
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 2,
					Amount:     70,
				},
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 3,
					Amount:     30,
				},
			},
			expectError: nil,
		},
		{
			name: "Split amount exceeds expense amount",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Amount: 150},
			},
			expected:    nil,
			expectError: errors.ErrInvalidSplitConfiguration,
		},
		{
			name: "Zero amount split should be skipped",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Amount: 0},
				{BorrowerID: 3, Amount: 100},
			},
			expected: []table.Ledger{
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 3,
					Amount:     100,
				},
			},
			expectError: nil,
		},
		{
			name: "Multiple zero amount splits",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Amount: 0},
				{BorrowerID: 3, Amount: 0},
				{BorrowerID: 4, Amount: 100},
			},
			expected: []table.Ledger{
				{
					ExpenseID:  1,
					LenderID:   1,
					BorrowerID: 4,
					Amount:     100,
				},
			},
			expectError: nil,
		},
		{
			name: "Total split amount less than expense amount",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Amount: 30},
				{BorrowerID: 3, Amount: 40},
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
					Amount:     40,
				},
			},
			expectError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitter := NewAmountSplitter()
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
