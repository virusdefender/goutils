/*
错误处理，主要使用 Wrap 系列函数
*/

package errors

import (
	"errors"
	"fmt"
)

type wrapError struct {
	message string
	next    error
}

func (e *wrapError) Unwrap() error {
	return e.next
}

func (e *wrapError) Error() string {
	if e.next == nil {
		return e.message
	}
	return fmt.Sprintf("%s: %v", e.message, e.next)
}

func wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return &wrapError{
		message: message,
		next:    err,
	}
}

// Wrap returns a error annotating `err` with `message` and the caller's frame.
// Wrap returns nil if `err` is nil.
func Wrap(err error, message string) error {
	return wrap(err, message)
}

// Wrapf returns a error annotating `err` with `message` formatted and the caller's frame.
func Wrapf(err error, message string, args ...interface{}) error {
	return wrap(err, fmt.Sprintf(message, args...))
}

var (
	New = errors.New
	Is  = errors.Is
)
