package domain

import "errors"

var (
	/* Validation errors */

	ErrEmptyInput          = errors.New(EmptyInput)
	ErrInvalidURL          = errors.New(InvalidUrl)
	ErrInvalidInput        = errors.New(InvalidInput)
	ErrURLNotFound         = errors.New(URLNotFound)
	ErrURLIsAlreadyShorted = errors.New(URLIsAlreadyShorted)
)
