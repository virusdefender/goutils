/*
基于泛型来实现自动 recover
*/

package call

import (
	"fmt"
	"runtime/debug"
)

type PanicError struct {
	Ret   interface{}
	Stack []byte
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("recovered panic: %v\nstack: %s", e.Ret, string(e.Stack))
}

func IsPanicError(err error) bool {
	_, ok := err.(*PanicError)
	return ok
}

func WithRecover(fn func() error) (err error) {
	defer func() {
		r := recover()
		// to simplify, before go 1.21, panic(nil) will not be handled
		if r != nil {
			err = &PanicError{
				Ret:   r,
				Stack: debug.Stack(),
			}
		}
	}()
	return fn()
}

func WithRecover1[T any](fn func() (T, error)) (T, error) {
	var r T
	var err error
	return r, WithRecover(func() error {
		r, err = fn()
		return err
	})
}

func WithRecover2[T any, T1 any](fn func() (T, T1, error)) (T, T1, error) {
	var r T
	var r1 T1
	var err error
	return r, r1, WithRecover(func() error {
		r, r1, err = fn()
		return err
	})
}
