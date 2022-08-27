//
// Copyright (C) 2019 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/it
//

package it

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

//
// Assertions
//

// Be assert logical predicated to the truth
//   it.Should(it.Be(myPredicate))
func Be(f func() bool) error {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	assert := fmt.Errorf("predicate %s be true", fn)

	if !f() {
		return assert
	}
	return passed(assert)
}

// True assert results of logical predicated to be true
//  it.Should(it.True( cnt > 10 ))
func True(x bool) error {
	assert := errors.New("be true")

	if !x {
		return assert
	}
	return passed(assert)
}

// SameAs matches type of x, y
//   it.Should(it.SameAs(x, y))
func SameAs[T any](x, y T) error {
	return passed(fmt.Errorf("type of %v same as %T", x, y))
}

// Nil asserts the variable for the nil value
//   it.Should(it.Nil(x))
func Nil(x interface{}) error {
	if x != nil {
		return fmt.Errorf("value [%v] be defined", x)
	}
	return passed(fmt.Errorf("nil value"))
}

//
// Intercepts
//

// Callable type constraint for scope of Intercepts
type Callable interface {
	~func() error | ~func()
}

// Fail catches any errors caused by the function under the test.
//   it.Should(it.Fail(refToCodeBlock))
func Fail[T Callable](f T) FailIt {
	switch ff := any(f).(type) {
	case func() error:
		return failWithError(ff)
	case func():
		return failWithPanic(ff)
	default:
		panic("runtime error")
	}
}

//
func failWithError(f func() error) FailIt {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	assert := fmt.Errorf("%s return error", fn)

	err := f()
	if err == nil {
		return FailIt{assert, nil}
	}

	return FailIt{passed(assert), err}
}

//
func failWithPanic(f func()) FailIt {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	assert := fmt.Errorf("%s panic", fn)

	err := func() (err error) {
		defer func() {
			switch r := recover().(type) {
			case nil:
			case error:
				err = r
			default:
				err = fmt.Errorf("%v", r)
			}
		}()
		f()
		return
	}()

	if err == nil {
		return FailIt{assert, nil}
	}

	return FailIt{passed(assert), err}
}

// FailIt extend Fail assert
type FailIt struct {
	assert error
	status error
}

func (x FailIt) Error() string      { return x.assert.Error() }
func (x FailIt) As(target any) bool { return errors.As(x.assert, target) }

// With asserts failure to the expected error
//   it.Should(it.Fail(refToCodeBlock).With(&notFound))
func (x FailIt) With(y any) error {
	assert := fmt.Errorf("%s with %T", x.assert, y)

	if !errors.As(x.status, y) {
		return assert
	}

	return passed(assert)
}

// Contain asserts error string for expected term
//   it.Should(it.Fail(refToCodeBlock).Contain("not found"))
func (x FailIt) Contain(y string) error {
	assert := fmt.Errorf("%s contain %s", x.assert, y)

	if !strings.Contains(x.status.Error(), y) {
		return assert
	}

	return passed(assert)
}

// Error checks return values of function on the error cases
//   it.Should(it.Error(refToCodeBlock()))
func Error[A any](x A, err error) FailIt {
	assert := errors.New("return error")

	if err == nil {
		return FailIt{assert, nil}
	}

	assert = fmt.Errorf("return error [%w]", err)
	return FailIt{passed(assert), err}
}

//
// Equality and Identity
//

// Equal check equality (x = y) of two scalar variables
//   it.Should(it.Equal(x, y))
func Equal[T comparable](x, y T) error {
	assert := fmt.Errorf("%v be equal to %v", x, y)

	if x != y {
		return assert
	}
	return passed(assert)
}

// Equiv check equality (x ≈ y) of two non scalar variables
//   it.Should(it.Equiv(x, y))
func Equiv[T any](x, y T) error {
	assert := fmt.Errorf("%v be equivalent to %v", x, y)

	if !equal(x, y) {
		return assert
	}

	return passed(assert)
}

// Orderable type constraint
type Orderable interface {
	~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Less compares (x < y) two scalar variables
//   it.Should(it.Less(x, y))
func Less[T Orderable](x, y T) error {
	assert := fmt.Errorf("%v be less than %v", x, y)
	if !(x < y) {
		return assert
	}
	return passed(assert)
}

// Less compares (x <= y) two scalar variables
//   it.Should(it.LessOrEqual(x, y))
func LessOrEqual[T Orderable](x, y T) error {
	assert := fmt.Errorf("%v be less or equal to %v", x, y)

	if !(x <= y) {
		return assert
	}
	return passed(assert)
}

// Greater compares (x > y) of two scalar variables
//   it.Should(it.Greater(x, y))
func Greater[T Orderable](x, y T) error {
	assert := fmt.Errorf("%v be greater than %v", x, y)
	if !(x > y) {
		return assert
	}
	return passed(assert)
}

// Greater compares (x >= y) of two scalar variables
//   it.Should(it.GreaterOrEqual(x, y))
func GreaterOrEqual[T Orderable](x, y T) error {
	assert := fmt.Errorf("%v be greater or equal to %v", x, y)
	if !(x >= y) {
		return assert
	}
	return passed(assert)
}

//
// Non public asserts
//

func isNul(a interface{}) bool {
	return a == nil || (reflect.ValueOf(a).Kind() == reflect.Ptr && reflect.ValueOf(a).IsNil())
}

func equal(a, b interface{}) bool {
	// Note: reflect.DeepEqual uses type metadata to compare.
	//       It would fail if nil value of pointer type is compared to nil literal
	//       var v *MyType
	//       it.Then(t).If(v).Should().Equal(nil)
	if isNul(a) && isNul(b) {
		return true
	}

	return reflect.DeepEqual(a, b)
}
