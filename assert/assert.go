/*
一些用于测试中的断言函数
*/

package assert

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type TestingObj interface {
	Fatal(args ...any)
}

func getCaller() string {
	_, file, line, ok := runtime.Caller(3)
	if ok {
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	return "???"
}

func joinOutput(msg string, extraMsg ...string) string {
	if len(extraMsg) > 0 {
		return fmt.Sprintf("\nTest: %s\nReason: %s\nExtra Message: %s", getCaller(), msg, strings.Join(extraMsg, "; "))
	}
	return fmt.Sprintf("\nTest: %s\n%s", getCaller(), msg)
}

func Nil(t TestingObj, data any, extraMsg ...string) {
	if data != nil {
		msg := fmt.Sprintf("expect nil, got %v", data)
		t.Fatal(joinOutput(msg, extraMsg...))
	}
}

func NotNil(t TestingObj, data any, extraMsg ...string) {
	if data == nil {
		msg := fmt.Sprintf("expect not nil, got nil")
		t.Fatal(joinOutput(msg, extraMsg...))
	}
}

func True(t TestingObj, data bool, extraMsg ...string) {
	if !data {
		msg := fmt.Sprintf("expect true, got false")
		t.Fatal(joinOutput(msg, extraMsg...))
	}
}

func False(t TestingObj, data bool, extraMsg ...string) {
	if data {
		msg := fmt.Sprintf("expect false, got true")
		t.Fatal(joinOutput(msg, extraMsg...))
	}
}

func Equal(t TestingObj, actual any, expected any, extraMsg ...string) {
	if !reflect.DeepEqual(expected, actual) {
		msg := fmt.Sprintf("expect %v, got %v", expected, actual)
		t.Fatal(joinOutput(msg, extraMsg...))
	}
}

func ErrorStringEqual(t TestingObj, err error, expected string, extraMsg ...string) {
	if err == nil {
		msg := fmt.Sprintf("expect error, got nil")
		t.Fatal(joinOutput(msg, extraMsg...))
	}
	if err.Error() != expected {
		msg := fmt.Sprintf("expect error message is %s, got %s", expected, err.Error())
		t.Fatal(joinOutput(msg, extraMsg...))
	}
}

func NoError(t TestingObj, fn func() error) {
	err := fn()
	if err != nil {
		msg := fmt.Sprintf("expect no error, got %s", err.Error())
		t.Fatal(joinOutput(msg))
	}
}

func NoError1[T any](t TestingObj, fn func() (T, error)) {
	_, err := fn()
	if err != nil {
		msg := fmt.Sprintf("expect no error, got %s", err.Error())
		t.Fatal(joinOutput(msg))
	}
}

func NoError2[T any, T1 any](t TestingObj, fn func() (T, T1, error)) {
	_, _, err := fn()
	if err != nil {
		msg := fmt.Sprintf("expect no error, got %s", err.Error())
		t.Fatal(joinOutput(msg))
	}
}

func LengthEqual(t TestingObj, data interface{}, length int, extraMsg ...string) {
	v := reflect.ValueOf(data)
	if v.Len() != length {
		msg := fmt.Sprintf("expect length %d, got %d", length, v.Len())
		t.Fatal(joinOutput(msg, extraMsg...))
	}
}
