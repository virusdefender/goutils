package call

import (
	"io"
	"testing"
)

func TestWithRecover(t *testing.T) {
	err := WithRecover(func() error {
		return nil
	})
	if err != nil {
		t.Fatalf("err should be nil")
	}

	err = WithRecover(func() error {
		return io.EOF
	})
	if err != io.EOF {
		t.Fatalf("err should be io.EOF")
	}

	err = WithRecover(func() error {
		panic("test panic")
	})
	if err == nil {
		t.Fatalf("err should not be nil")
	}
	if !IsPanicError(err) {
		t.Fatalf("err should be PanicError")
	}
	if err.(*PanicError).Ret.(string) != "test panic" {
		t.Fatalf("err.Ret should be 'test panic', got %v", err.(*PanicError).Ret)
	}
}

func TestWithRecover1(t *testing.T) {
	i, err := WithRecover1(func() (int, error) {
		return 1, nil
	})
	if i != 1 || err != nil {
		t.Fatalf("i should be 1 and err should be nil")
	}

	i, err = WithRecover1(func() (int, error) {
		return 0, io.EOF
	})
	if i != 0 || err != io.EOF {
		t.Fatalf("i should be 0 and err should be io.EOF")
	}

	i, err = WithRecover1(func() (int, error) {
		panic("test panic")
	})
	if i != 0 || err == nil {
		t.Fatalf("i should be 0 and err should not be nil")
	}
}

func TestWithRecover2(t *testing.T) {
	i, i1, err := WithRecover2(func() (int, int, error) {
		return 1, 2, nil
	})
	if i != 1 || i1 != 2 || err != nil {
		t.Fatalf("i should be 1, i1 should be 2 and err should be nil")
	}

	i, i1, err = WithRecover2(func() (int, int, error) {
		return 0, 0, io.EOF
	})
	if i != 0 || i1 != 0 || err != io.EOF {
		t.Fatalf("i should be 0, i1 should be 0 and err should be io.EOF")
	}

	i, i1, err = WithRecover2(func() (int, int, error) {
		panic("test panic")
	})
	if i != 0 || i1 != 0 || err == nil {
		t.Fatalf("i should be 0, i1 should be 0 and err should not be nil")
	}
}
