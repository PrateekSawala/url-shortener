package domain

import "errors"

var (
	/* Validation errors */

	ErrEmptyInput          = errors.New(EmptyInput)
	ErrInvalidURL          = errors.New(InvalidUrl)
	ErrURLNotFound         = errors.New(URLNotFound)
	ErrURLIsAlreadyShorted = errors.New(URLIsAlreadyShorted)
)
