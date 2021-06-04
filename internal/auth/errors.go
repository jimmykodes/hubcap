package auth

import "errors"

var (
	ErrCanceled        = errors.New("oauth canceled")
	ErrNoState         = errors.New("missing state")
	ErrMismatchedState = errors.New("mismatched state")
	ErrNoCode          = errors.New("missing code")
	ErrNoUsername      = errors.New("missing username/email")
)
