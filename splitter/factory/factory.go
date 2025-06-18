package factory

import (
	"github.com/sreekar2307/khata/errors"
	"github.com/sreekar2307/khata/model"
	"github.com/sreekar2307/khata/splitter"
	"github.com/sreekar2307/khata/splitter/amount"
	"github.com/sreekar2307/khata/splitter/equal"
	"github.com/sreekar2307/khata/splitter/percentage"
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
