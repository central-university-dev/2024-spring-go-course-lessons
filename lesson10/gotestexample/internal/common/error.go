package common

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrValidation = errors.New("validation error")
	ErrTechnical  = errors.New("a technical error happened")
)
