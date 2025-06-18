package errors

import stdErrors "errors"

var (
	ErrInvalidSplitType          = stdErrors.New("invalid split type")
	ErrInvalidSplitConfiguration = stdErrors.New("invalid split configuration")
	ErrInvalidSimplifier         = stdErrors.New("invalid simplifier")

	ErrInvalidUserCredentials  = stdErrors.New("invalid user credentials")
	ErrUserNotFound            = stdErrors.New("user not found")
	ErrCheckIfDependencyExists = stdErrors.New("check if dependency exists")
)
