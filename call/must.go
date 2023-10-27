/*
某些情况下要求操作必须成功（比如 init 中），Must 系列函数在发生错误的时候直接 panic
*/

package call

func Must(fn func() error) {
	err := fn()
	if err != nil {
		panic(err)
	}
}

func Must1[T any](fn func() (T, error)) T {
	r, err := fn()
	if err != nil {
		panic(err)
	}
	return r
}

func Must2[T any, T1 any](fn func() (T, T1, error)) (T, T1) {
	r, r1, err := fn()
	if err != nil {
		panic(err)
	}
	return r, r1
}
