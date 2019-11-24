//
// Copyright (C) 2019 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/it
//

package it

import (
	"reflect"
	"testing"
)

type expr struct {
	testing *testing.T
}

type actual struct {
	expr  expr
	value interface{}
}

type should struct {
	actual actual
}

type be struct {
	should should
}

type fail struct {
	should should
}

//
func Ok(t *testing.T) expr {
	return expr{t}
}

//
func (t expr) If(x interface{}) actual {
	return actual{t, x}
}

//
func (t actual) Should() should {
	return should{t}
}

//
// Should
//

func (t should) Assert(f func(interface{}) bool) {
	this(t).Helper()

	if !f(actualValue(t)) {
		this(t).Errorf("assert is failed")
	}
}

func (t should) Equal(expect interface{}) {
	this(t).Helper()

	value := actualValue(t)
	if !reflect.DeepEqual(value, expect) {
		this(t).Errorf("%v not equal %v", value, expect)
	}
}

func (t should) Be() be {
	return be{t}
}

func (t should) Fail() fail {
	return fail{t}
}

//
func (t be) A(expect interface{}) {
	this(t).Helper()

	value := actualValue(t)
	if !reflect.DeepEqual(value, expect) {
		this(t).Errorf("%v not equal %v", value, expect)
	}
}

func (t be) Like(expect interface{}) {
	this(t).Helper()

	value := actualValue(t)
	if !kind(value, expect) {
		this(t).Errorf("%v not same type like %v, type %T is required", value, expect, expect)
	}
}

//
//
func (t be) Less(expect interface{}) {
	this(t).Helper()

	value := actualValue(t)
	if !kind(value, expect) {
		this(t).Errorf("%v not same type like %v, type %T is required", value, expect, expect)
		return
	}
	if reflect.DeepEqual(value, expect) {
		this(t).Errorf("%v is equal %v", value, expect)
		return
	}
	if !less(value, expect) {
		this(t).Errorf("%v is not less then %v", value, expect)
	}
}

func (t be) LessOrEqual(expect interface{}) {
	this(t).Helper()

	value := actualValue(t)
	if !kind(value, expect) {
		this(t).Errorf("%v not same type like %v, type %T is required", value, expect, expect)
		return
	}
	if !reflect.DeepEqual(value, expect) {
		if !less(value, expect) {
			this(t).Errorf("%v is not less or equal to %v", value, expect)
		}
	}
}

func (t be) Greater(expect interface{}) {
	this(t).Helper()

	value := actualValue(t)
	if !kind(value, expect) {
		this(t).Errorf("%v not same type like %v, type %T is required", value, expect, expect)
		return
	}
	if reflect.DeepEqual(value, expect) {
		this(t).Errorf("%v is equal %v", value, expect)
		return
	}
	if less(value, expect) {
		this(t).Errorf("%v is not greater then %v", value, expect)
	}
}

func (t be) GreaterOrEqual(expect interface{}) {
	this(t).Helper()

	value := actualValue(t)
	if !kind(value, expect) {
		this(t).Errorf("%v not same type like %v, type %T is required", value, expect, expect)
		return
	}

	if !reflect.DeepEqual(value, expect) {
		if less(value, expect) {
			this(t).Errorf("%v is not greater or equal to %v", value, expect)
		}
	}
}

func (t be) In(a, b interface{}) {
	this(t).Helper()

	value := actualValue(t)
	if !kind(value, a) {
		this(t).Errorf("%v not same type like %v, type %T is required", value, a, a)
		return
	}
	if !kind(value, b) {
		this(t).Errorf("%v not same type like %v, type %T is required", value, b, b)
		return
	}

	if less(value, a) {
		this(t).Errorf("%v is not greater then %v", value, a)
		return
	}
	if !less(value, b) {
		this(t).Errorf("%v is not less then %v", value, b)
	}
}

//
//
func (t fail) With(expect error) {
	this(t).Helper()

	f := actualValue(t).(func() error)
	value := f()
	if !reflect.DeepEqual(value, expect) {
		this(t).Errorf("returns unexpected error %v, it requires %v", value, expect)
	}
}

//
func actualValue(t interface{}) interface{} {
	switch v := t.(type) {
	case should:
		return v.actual.value
	case be:
		return v.should.actual.value
	case fail:
		return v.should.actual.value
	}
	return nil
}

func this(t interface{}) *testing.T {
	switch v := t.(type) {
	case should:
		return v.actual.expr.testing
	case be:
		return v.should.actual.expr.testing
	case fail:
		return v.should.actual.expr.testing
	}
	return nil
}

func kind(x, y interface{}) bool {
	xKind := reflect.ValueOf(x).Kind()
	yKind := reflect.ValueOf(y).Kind()
	return xKind == yKind
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
