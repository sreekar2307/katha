package factory

import (
	"github.com/sreekar2307/katha/errors"
	"github.com/sreekar2307/katha/model"
	"github.com/sreekar2307/katha/splitter"
	"github.com/sreekar2307/katha/splitter/amount"
	"github.com/sreekar2307/katha/splitter/equal"
	"github.com/sreekar2307/katha/splitter/percentage"
)

type factory struct{}

// SplitterFactory is an interface that defines a method to create a new Splitter based on the split type.
type SplitterFactory interface {
	NewSplitter(model.SplitType) (splitter.Splitter, error)
}

func NewFactory() SplitterFactory {
	return factory{}
}

func (f factory) NewSplitter(splitType model.SplitType) (splitter.Splitter, error) {
	switch splitType {
	case model.SplitTypes.Equal:
		return equal.NewEqualSplitter(), nil
	case model.SplitTypes.Percentage:
		return percentage.NewPercentageSplitter(), nil
	case model.SplitTypes.Amount:
		return amount.NewAmountSplitter(), nil
	}
	return nil, errors.ErrInvalidSplitType
}
