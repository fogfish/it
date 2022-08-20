package it

import (
	"fmt"
	"reflect"
	"runtime"
)

// Is assert logical expression to the truth
//   it.Should(it.True(myPredicate))
func Is(f func() bool) error {
	if !f() {
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		return fmt.Errorf("%s be true", name)
	}
	return nil
}

// IsNot assert logical expression to the false
//   it.Should(it.False(myPredicate))
func IsNot(f func() bool) error {
	if f() {
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		return fmt.Errorf("not %s be true", name)
	}
	return nil
}

// Nil asserts the variable for undefined clause
//   it.Ok(t).Should(it.Nil(x))
func Nil(x interface{}) error {
	if x != nil {
		return fmt.Errorf("not %v be defined", x)
	}
	return nil
}

// NotNil asserts the variable for undefined clause
//   it.Ok(t).Should(it.NotNil(x))
func NotNil(x interface{}) error {
	if x == nil {
		return fmt.Errorf("not be %v", x)
	}
	return nil
}

// SameAs matches type of x, y
//   it.Ok(t).Should(it.SameAs(x, y))
func SameAs[T any](x, y T) error {
	return nil
}

// Equal check equality (x = y) of two scalar variables
//   it.Ok(t).Should(it.Equal(x, y))
func Equal[T comparable](x, y T) error {
	if x != y {
		return fmt.Errorf("%v be equal to %v", x, y)
	}
	return nil
}

// NotEqual check not equality (x != y) of two scalar variables
//   it.Ok(t).Should(it.NotEqual(x, y))
func NotEqual[T comparable](x, y T) error {
	if x == y {
		return fmt.Errorf("not %v be equal to %v", x, y)
	}
	return nil
}

// Orderable type constraint
type Orderable interface {
	~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Less compares (x < y) of two scalar variables
//   it.Ok(t).Should(it.Less(x, y))
func Less[T Orderable](x, y T) error {
	if !(x < y) {
		return fmt.Errorf("%v be less than %v", x, y)
	}
	return nil
}

// Less compares (x <= y) of two scalar variables
//   it.Ok(t).Should(it.LessOrEqual(x, y))
func LessOrEqual[T Orderable](x, y T) error {
	if !(x <= y) {
		return fmt.Errorf("%v be less or equal to %v", x, y)
	}
	return nil
}

// Greater compares (x > y) of two scalar variables
//   it.Ok(t).Should(it.Greater(x, y))
func Greater[T Orderable](x, y T) error {
	if !(x > y) {
		return fmt.Errorf("%v be greater than %v", x, y)
	}
	return nil
}

// Greater compares (x >= y) of two scalar variables
//   it.Ok(t).Should(it.GreaterOrEqual(x, y))
func GreaterOrEqual[T Orderable](x, y T) error {
	if !(x >= y) {
		return fmt.Errorf("%v be greater or equal to %v", x, y)
	}
	return nil
}

// Callable type constraint
type Callable interface {
	~func() error | ~func()
}

// NotFail catches any errors caused by the function under the test.
//   it.Ok(t).Should(it.NotFail(refToCodeBlock))
func Failed[T Callable](f T) error {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	switch ff := any(f).(type) {
	case func() error:
		if err := ff(); err == nil {
			return fmt.Errorf("%s fail", name)
		}
	case func():
		err := func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("%v", r)
				}
			}()
			ff()
			return
		}()
		if err == nil {
			return fmt.Errorf("%s fail", name)
		}
	}
	return nil
}

// NotFailed catches any errors caused by the function under the test.
//   it.Ok(t).Should(it.NotFail(refToCodeBlock))
func NotFailed[T Callable](f T) error {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	switch ff := any(f).(type) {
	case func() error:
		if err := ff(); err != nil {
			return fmt.Errorf("not %s fail with %w", name, err)
		}
	case func():
		err := func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("%v", r)
				}
			}()
			ff()
			return
		}()
		if err != nil {
			return fmt.Errorf("not %s fail with %w", name, err)
		}
	}
	return nil
}

// NoError checks return values of function on the error cases
//   it.Ok(t).Should(it.NoError(callMySut()))
func NoError[A any](x A, err error) error {
	if err != nil {
		return fmt.Errorf("not fail with %w", err)
	}
	return nil
}
