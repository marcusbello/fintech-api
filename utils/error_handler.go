package utils

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")

	ErrIncorrectPassword = errors.New("incorrect password")

	ErrInsufficientFunds = errors.New("insufficient funds")

	ErrBadRequest = errors.New("bad request")

	ErrDocuments = errors.New("error creating or reading document")
)
