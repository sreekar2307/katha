package percentage

import (
	"context"
	"testing"

	"github.com/sreekar2307/khata/errors"
	"github.com/sreekar2307/khata/model"
	"github.com/sreekar2307/khata/model/table"
	"github.com/stretchr/testify/assert"
)

func TestPercentageSplitter_Split(t *testing.T) {
	tests := []struct {
		name        string
		expense     table.Expense
		splits      []model.Split
		expected    []table.Ledger
		expectError error
	}{
		{
			name: "Valid split with equal percentages",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Percentage: 50},
				{BorrowerID: 3, Percentage: 50},
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
			name: "Valid split with different percentages",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Percentage: 70},
				{BorrowerID: 3, Percentage: 30},
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
			name: "Invalid percentage greater than 100",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Percentage: 150},
			},
			expected:    nil,
			expectError: errors.ErrInvalidSplitConfiguration,
		},
		{
			name: "Invalid percentage less than or equal to 0",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Percentage: 0},
			},
			expected:    nil,
			expectError: errors.ErrInvalidSplitConfiguration,
		},
		{
			name: "Total percentage less than 100",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Percentage: 30},
				{BorrowerID: 3, Percentage: 40},
			},
			expected:    nil,
			expectError: errors.ErrInvalidSplitConfiguration,
		},
		{
			name: "Total percentage greater than 100",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Percentage: 60},
				{BorrowerID: 3, Percentage: 50},
			},
			expected:    nil,
			expectError: errors.ErrInvalidSplitConfiguration,
		},
		{
			name: "Non-integer amount result",
			expense: table.Expense{
				Base: table.Base{
					ID: 1,
				},
				LenderID: 1,
				Amount:   100,
			},
			splits: []model.Split{
				{BorrowerID: 2, Percentage: 33.33},
				{BorrowerID: 3, Percentage: 66.67},
			},
			expected:    nil,
			expectError: errors.ErrInvalidSplitConfiguration,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitter := NewPercentageSplitter()
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
