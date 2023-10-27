package assert

import (
	"errors"
	"github.com/virusdefender/goutils/memory"
	"testing"
)

func TestAssert(t *testing.T) {
	Nil(t, nil)
	NotNil(t, 123)

	True(t, true)
	False(t, false)

	Equal(t, 123, 123)
	Equal(t, "123", "123")
	Equal(t, []int{1, 2, 3}, []int{1, 2, 3})
	Equal(t, map[int]int{1: 1}, map[int]int{1: 1})

	ErrorStringEqual(t, errors.New("test"), "test")

	NoError(t, func() error {
		return nil
	})
	NoError1(t, func() (int, error) {
		return 1, nil
	})

	LengthEqual(t, []int{1, 2, 3}, 3)
	LengthEqual(t, map[int]int{1: 1}, 1)
	LengthEqual(t, "123", 3)

	memory.PrintMemoryStat()
}

const testFailed = false

func TestAssertFailed(t *testing.T) {
	if !testFailed {
		t.SkipNow()
	}

	Nil(t, 123)
	Nil(t, 123, "case1")
}
