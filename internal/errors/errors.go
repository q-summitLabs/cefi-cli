package errors

import "errors"

var (
	ErrExit        = errors.New("exit requested")
	ErrTooManyArgs = errors.New("too many arguments for given command")
)
