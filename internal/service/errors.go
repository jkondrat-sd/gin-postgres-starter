// Why: shared service errors let handlers map business failures to consistent HTTP responses.
// What to do: add new domain-level errors here when multiple services or handlers need to recognize them.
package service

import "errors"

var (
	ErrConflict     = errors.New("resource already exists")
	ErrInvalidLogin = errors.New("invalid email or password")
	ErrNotFound     = errors.New("resource not found")
)
