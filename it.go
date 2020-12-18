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
	"testing"
)

// Expr is a term of assert expression
// See Ok keyword
type Expr struct {
	testing *testing.T
}

// Value is a term of assert expression
// See If keyword
type Value struct {
	expr  Expr
	value interface{}
}

type level int

const (
	lCRITICAL level = iota + 1
	lERROR
	lNOTICE
)

// Imperative is a term of assert expression
// See imperative keywords (Must, MustNot, Should, ShouldNot, May)
type Imperative struct {
	actual  Value
	success func(bool) bool
	level   level
}

// Be is a term of assert expression
// See Be keyword
type Be struct {
	imp Imperative
}

// Fail is a term is assert expression
// See Fail keyword
type Fail struct {
	imp Imperative
}

// Ok creates assertion expression, it takes a reference to
// standard *testing.T interface to setup the evaluation context.
func Ok(t *testing.T) Expr {
	return Expr{t}
}

// If stashes an actual value for evaluation in further statements
func (t Expr) If(x interface{}) Value {
	return Value{t, x}
}

//-----------------------------------------------------------------------------
//
// Imperative keyword(s)
//
//-----------------------------------------------------------------------------

// Must is imperative keyword.
// The assert definition is an absolute requirement.
// It terminates execution of tests if assert is false.
func (t Value) Must() Imperative {
	return Imperative{t, success, lCRITICAL}
}

// MustNot is imperative keyword.
// The assert definition is an absolute prohibition.
// It terminates execution of tests if assert is true.
func (t Value) MustNot() Imperative {
	return Imperative{t, successNot, lCRITICAL}
}

// Should is imperative keyword.
// The assert definition is a strongly recommended requirement.
// However it's violation do not block continuation of testing.
func (t Value) Should() Imperative {
	return Imperative{t, success, lERROR}
}

// ShouldNot is imperative keyword.
// The assert definition is prohibited.
// However it's violation do not block continuation of testing.
func (t Value) ShouldNot() Imperative {
	return Imperative{t, successNot, lERROR}
}

// May is an optional imperative constrain.
// Its violation do not impact on testing flow in anyway.
// The informative warning message is produced to console.
// Error message will be printed only if the test fails or the -test.v
func (t Value) May() Imperative {
	return Imperative{t, success, lNOTICE}
}

func success(x bool) bool    { return x }
func successNot(x bool) bool { return !x }

//-----------------------------------------------------------------------------
//
// Asserts
//
//-----------------------------------------------------------------------------

func (t Imperative) native() *testing.T {
	return t.actual.expr.testing
}

func (t Imperative) error(msg string, args ...interface{}) {
	t.actual.expr.testing.Helper()
	switch t.level {
	case lCRITICAL:
		panic(fmt.Sprintf(msg, args...))
	case lERROR:
		t.actual.expr.testing.Errorf(msg, args...)
	case lNOTICE:
		t.actual.expr.testing.Logf(msg, args...)
	}
}

func (t Imperative) value() interface{} {
	return t.actual.value
}

// Assert with user-defined functions is a technique to define
// arbitrary boolean expression. The stashed value is fed as
// argument to the function. It fails if the function return
// false.
//
//   it.Ok(t).If(x).
//     Should().Assert(func(be interface{}) bool {
//	     (be > 1) && (be < 10) && (be != 5)
//     })
//
func (t Imperative) Assert(f func(interface{}) bool) Expr {
	t.native().Helper()

	value := t.value()
	if !t.success(f(value)) {
		t.error("assert is failed, unexpected value %v", value)
	}
	return t.actual.expr
}

// Intercept catches any errors caused by the function under the test.
// The assert fails if expected failure do not match.
//
//   it.Ok(t).If(refToCodeBlock).
//      Should().Intercept(/* ... */)
func (t Imperative) Intercept(err error) Expr {
	t.native().Helper()

	switch f := t.value().(type) {
	case func() error:
		value := f()
		if !t.success(errors.Is(value, err)) {
			t.error("returns unexpected error %v, it requires %v", value, err)
		}
	case func():
		defer func() {
			switch value := recover().(type) {
			case error:
				if !t.success(errors.Is(value, err)) {
					t.error("returns unexpected error %v, it requires %v", value, err)
				}
			default:
				t.error("error type is expected, returns %T %v", value, value)
			}
		}()
		f()
	}

	return t.actual.expr
}

// Fail catches any errors caused by the function under the test.
// The assert fails if error is not nil.
func (t Imperative) Fail() Expr {
	t.native().Helper()

	switch f := t.value().(type) {
	case func() error:
		value := f()
		if !t.success(value != nil) {
			t.error("successful, it requires an error")
		}
	case func():
		defer func() {
			switch value := recover().(type) {
			case error:
				if !t.success(value != nil) {
					t.error("successful, it requires an error")
				}
			default:
				t.error("error type is expected, returns %T %v", value, value)
			}
		}()
		f()
	}

	return t.actual.expr

}

// Equal compares left and right sides. The assert fails if they are not equal.
//
//  it.Ok(t).If(1).
//    Should().Equal(1)
func (t Imperative) Equal(expect interface{}) Expr {
	t.native().Helper()

	value := t.value()
	if !t.success(eq(value, expect)) {
		t.error("%v not equal %v", value, expect)
	}
	return t.actual.expr
}

// Equiv compares equivalence (same value) of left and right sides.
// The assert fails if they are not equal.
//
//  it.Ok(t).If(1).
//    Should().Equal(1)
func (t Imperative) Equiv(expect interface{}) Expr {
	t.native().Helper()

	value := t.value()
	if !t.success(ev(value, expect)) {
		t.error("%v not equal %v", value, expect)
	}
	return t.actual.expr
}

// Eq is an alias of Equal
func (t Imperative) Eq(expect interface{}) Expr {
	return t.Equal(expect)
}

// Be creates comparison asserts
//   Should().Be()
func (t Imperative) Be() Be {
	return Be{t}
}

//-----------------------------------------------------------------------------
//
// Comparison asserts
//
//-----------------------------------------------------------------------------

// A matches expected values against actual
//   it.Ok(t).If(actual).Should().Be().A(expected)
func (t Be) A(expect interface{}) Expr {
	t.imp.native().Helper()

	value := t.imp.value()
	if !t.imp.success(eq(value, expect)) {
		t.imp.error("%v not equal %v", value, expect)
	}
	return t.imp.actual.expr
}

// Eq is an alias of A
func (t Be) Eq(expect interface{}) Expr {
	return t.A(expect)
}

// Like matches type of expected and actual values.
// The assert fails if types are different.
//   it.Ok(t).If(actual).Should().Be().Like(expected)
func (t Be) Like(expect interface{}) Expr {
	t.imp.native().Helper()

	value := t.imp.value()
	if !t.imp.success(kind(value, expect)) {
		t.imp.error("%v not same type like %v, type %T is required", value, expect, expect)
	}
	return t.imp.actual.expr
}

// Less compares actual against expected value.
//   it.Ok(t).If(actual).Should().Be().Less(expected)
func (t Be) Less(expect interface{}) Expr {
	t.imp.native().Helper()

	value := t.imp.value()
	if !kind(value, expect) {
		t.imp.error("%v not same type like %v, type %T is required", value, expect, expect)
	} else if eq(value, expect) {
		t.imp.error("%v is equal %v", value, expect)
	} else if !t.imp.success(less(value, expect)) {
		t.imp.error("%v is not less then %v", value, expect)
	}
	return t.imp.actual.expr
}

// LessOrEqual compares actual against expected value.
//   it.Ok(t).If(actual).Should().Be().LessOrEqual(expected)
func (t Be) LessOrEqual(expect interface{}) Expr {
	t.imp.native().Helper()

	value := t.imp.value()
	if !kind(value, expect) {
		t.imp.error("%v not same type like %v, type %T is required", value, expect, expect)
	} else if !t.imp.success(eq(value, expect)) && !t.imp.success(less(value, expect)) {
		t.imp.error("%v is not less or equal to %v", value, expect)
	}
	return t.imp.actual.expr
}

// Greater compares actual against expected value.
//   it.Ok(t).If(actual).Should().Be().Greater(expected)
func (t Be) Greater(expect interface{}) Expr {
	t.imp.native().Helper()

	value := t.imp.value()
	if !kind(value, expect) {
		t.imp.error("%v not same type like %v, type %T is required", value, expect, expect)
	} else if eq(value, expect) {
		t.imp.error("%v is equal %v", value, expect)
	} else if t.imp.success(less(value, expect)) {
		t.imp.error("%v is not greater then %v", value, expect)
	}
	return t.imp.actual.expr
}

// GreaterOrEqual compares actual against expected value.
//   it.Ok(t).If(actual).Should().Be().Greater(expected)
func (t Be) GreaterOrEqual(expect interface{}) Expr {
	t.imp.native().Helper()

	value := t.imp.value()
	if !kind(value, expect) {
		t.imp.error("%v not same type like %v, type %T is required", value, expect, expect)
	} else if !t.imp.success(eq(value, expect)) && t.imp.success(less(value, expect)) {
		t.imp.error("%v is not greater or equal to %v", value, expect)
	}
	return t.imp.actual.expr
}

// In checks that actual value fits to range
//   it.Ok(t).If(actual).Should().Be().In(from, to)
func (t Be) In(a, b interface{}) Expr {
	t.imp.native().Helper()

	value := t.imp.value()
	if !kind(value, a) {
		t.imp.error("%v not same type like %v, type %T is required", value, a, a)
	} else if !kind(value, b) {
		t.imp.error("%v not same type like %v, type %T is required", value, b, b)
	} else if t.imp.success(less(value, a)) {
		t.imp.error("%v is not greater then %v", value, a)
	} else if !t.imp.success(less(value, b)) {
		t.imp.error("%v is not less then %v", value, b)
	}
	return t.imp.actual.expr
}

//-----------------------------------------------------------------------------
//
// private helpers
//
//-----------------------------------------------------------------------------

func kind(x, y interface{}) bool {
	xKind := reflect.ValueOf(x).Kind()
	yKind := reflect.ValueOf(y).Kind()
	return xKind == yKind
}

func isNul(a interface{}) bool {
	return a == nil || (reflect.ValueOf(a).Kind() == reflect.Ptr && reflect.ValueOf(a).IsNil())
}

func eq(a, b interface{}) bool {
	// Note: reflect.DeepEqual uses type metadata to compare.
	//       It would fail if nil value of pointer type is compared to nil literal
	//       var v *MyType
	//       it.Ok(t).If(v).Should().Equal(nil)
	if isNul(a) && isNul(b) {
		return true
	}

	return reflect.DeepEqual(a, b)
}

func ev(a, b interface{}) bool {
	if isNul(a) || isNul(b) {
		return reflect.DeepEqual(a, b)
	}

	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	if va.Kind() == reflect.Ptr && vb.Kind() == reflect.Ptr {
		if !va.Elem().Type().Comparable() {
			return false
		}
		return reflect.DeepEqual(va.Elem().Interface(), vb.Elem().Interface())
	}

	if va.Kind() == reflect.Ptr {
		if !va.Elem().Type().Comparable() {
			return false
		}
		return reflect.DeepEqual(va.Elem().Interface(), vb.Interface())
	}

	if vb.Kind() == reflect.Ptr {
		if !vb.Elem().Type().Comparable() {
			return false
		}
		return reflect.DeepEqual(va.Interface(), vb.Elem().Interface())
	}

	return reflect.DeepEqual(a, b)
}

func less(x, y interface{}) bool {
	switch a := x.(type) {
	//
	case int:
		return a < y.(int)
	case int8:
		return a < y.(int8)
	case int16:
		return a < y.(int16)
	case int32:
		return a < y.(int32)
	case int64:
		return a < y.(int64)
	//
	case uint:
		return a < y.(uint)
	case uint8:
		return a < y.(uint8)
	case uint16:
		return a < y.(uint16)
	case uint32:
		return a < y.(uint32)
	case uint64:
		return a < y.(uint64)
	//
	case float32:
		return a < y.(float32)
	case float64:
		return a < y.(float64)
	//
	case string:
		return a < y.(string)
	}
	return false
}
