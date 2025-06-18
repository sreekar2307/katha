package simplifier

import (
	"context"
)

// Simplifier is an interface that defines a method for simplifying a list of ledgers.
type Simplifier interface {
	Simplify(context.Context, uint64) (map[uint64]map[uint64]uint64, error)
}
