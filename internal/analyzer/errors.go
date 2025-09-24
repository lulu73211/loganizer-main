package analyzer

import (
	"errors"
	"fmt"
)

var (
	ErrFileUnavailable = errors.New("file unavailable")
	ErrParseFailure    = errors.New("parse failure")
)

type FileError struct {
	Path string
	Err  error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("%v: %s (%v)", ErrFileUnavailable, e.Path, e.Err)
}
func (e *FileError) Unwrap() error { return e.Err }

type ParseError struct {
	Line    int
	Snippet string
	Err     error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%v at line %d: %s", ErrParseFailure, e.Line, e.Snippet)
}
func (e *ParseError) Unwrap() error { return e.Err }
