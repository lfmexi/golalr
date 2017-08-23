package prattparser

import (
	"fmt"
)

// ParseError is an implementation of the error interface
type ParseError struct {
	value string
	err   error
}

// NewParseError returns a new *ParseError with a given value and a custom error
func NewParseError(value string, err error) *ParseError {
	return &ParseError{value, err}
}

func (e *ParseError) Error() string {
	custom := ""
	if e.err != nil {
		custom = e.err.Error()
	}
	return fmt.Sprintf("Parse error: unexpected token %v %v", e.value, custom)
}
