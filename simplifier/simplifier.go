package simplifier

import (
	"context"
)

// Simplifier is an interface that defines a method to simplify lends and borrows for a given user ID.
type Simplifier interface {
	Simplify(context.Context, uint64) (map[uint64]map[uint64]uint64, error)
}
