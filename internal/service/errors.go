package service

import "errors"

var (
	ErrConflict     = errors.New("resource already exists")
	ErrInvalidLogin = errors.New("invalid email or password")
	ErrNotFound     = errors.New("resource not found")
)
