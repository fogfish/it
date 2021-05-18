//
// Copyright (C) 2019 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/it
//

package it_test

import (
	"errors"
	"testing"

	"github.com/fogfish/it"
)

func TestNewExpr(t *testing.T) {
	it.Ok(t)
}

func TestExprAliases(t *testing.T) {
	mock := new(testing.T)

	it.Ok(mock).IfTrue(true)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).IfFalse(false)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).IfNil(nil)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).IfNotNil("xxx")
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If("xxx").Equal("xxx")
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If("xxx").NotEqual("yyy")
	it.Ok(t).If(mock.Failed()).Should().Equal(false)
}

func TestImperativeKeywords(t *testing.T) {
	mock := new(testing.T)

	it.Ok(mock).If(1).Must().Assert(
		func(be interface{}) bool { return be == 1 },
	)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(2).MustNot().Assert(
		func(be interface{}) bool { return be == 1 },
	)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(1).Should().Assert(
		func(be interface{}) bool { return be == 1 },
	)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(2).ShouldNot().Assert(
		func(be interface{}) bool { return be == 1 },
	)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(1).May().Assert(
		func(be interface{}) bool { return be == 1 },
	)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)
}

func TestShouldAssert(t *testing.T) {
	mock := new(testing.T)

	it.Ok(mock).If(1).Should().Assert(
		func(be interface{}) bool { return be == 1 },
	)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(1).Should().Assert(
		func(be interface{}) bool { return be == 2 },
	)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)
}

func TestShouldEqual(t *testing.T) {
	mock := new(testing.T)

	it.Ok(mock).If(1).Should().Equal(1)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(1).Should().Equal(2)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)
}

func TestShouldEquiv(t *testing.T) {
	mock := new(testing.T)
	a := "string"
	b := "string"

	type X struct {
		a int
		b *string
	}
	c := X{1, &a}
	d := X{1, &b}

	it.Ok(mock).If(nil).Should().Equiv(nil)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(a).Should().Equiv(b)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(&a).Should().Equiv(&b)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(&a).Should().Equiv(b)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(a).Should().Equiv(&b)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(c).Should().Equiv(d)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(&c).Should().Equiv(&d)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(c).Should().Equiv(&d)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(&c).Should().Equiv(d)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(a).Should().Equal(&b)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

	it.Ok(mock).If(&a).Should().Equal(b)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)
}

func TestShouldEq(t *testing.T) {
	mock := new(testing.T)

	it.Ok(mock).If(1).Should().Eq(1)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(1).Should().Eq(2)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)
}

func TestShouldBeA(t *testing.T) {
	mock := new(testing.T)

	it.Ok(mock).If(1).Should().Be().A(1)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(1).Should().Be().A(2)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)
}

func TestShouldBeEq(t *testing.T) {
	mock := new(testing.T)

	it.Ok(mock).If(1).Should().Be().Eq(1)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(1).Should().Be().Eq(2)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)
}

func TestShouldBeLike(t *testing.T) {
	mock := new(testing.T)

	it.Ok(mock).If(1).Should().Be().Like(0)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(1).Should().Be().Like("")
	it.Ok(t).If(mock.Failed()).Should().Equal(true)
}

func TestShouldBeLikeStruct(t *testing.T) {
	mock := new(testing.T)

	type A struct{ Name string }
	type B struct{ Name string }

	a := A{"test"}
	b := B{"test"}

	it.Ok(mock).If(a).Should().Be().Like(A{})
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(&a).Should().Be().Like(A{})
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(a).Should().Be().Like(&A{})
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(&a).Should().Be().Like(&A{})
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(b).Should().Be().Like(A{})
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

	it.Ok(mock).If(&b).Should().Be().Like(A{})
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

	it.Ok(mock).If(b).Should().Be().Like(&A{})
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

	it.Ok(mock).If(&b).Should().Be().Like(&A{})
	it.Ok(t).If(mock.Failed()).Should().Equal(true)
}

var pairs map[interface{}]interface{} = map[interface{}]interface{}{
	int(1):       int(10),
	int8(1):      int8(10),
	int16(1):     int16(10),
	int32(1):     int32(10),
	int64(1):     int64(10),
	uint(1):      uint(10),
	uint8(1):     uint8(10),
	uint16(1):    uint16(10),
	uint32(1):    uint32(10),
	uint64(1):    uint64(10),
	float32(1.0): float32(10.0),
	float64(1.0): float64(10.0),
	"abcdef":     "bcdef",
}

var wrongtypes map[interface{}]interface{} = map[interface{}]interface{}{
	int(1):       int8(10),
	int8(1):      int16(10),
	int16(1):     int32(10),
	int32(1):     int64(10),
	int64(1):     int(10),
	uint(1):      uint8(10),
	uint8(1):     uint16(10),
	uint16(1):    uint32(10),
	uint32(1):    uint64(10),
	uint64(1):    uint(10),
	float32(1.0): float64(10.0),
	float64(1.0): float32(10.0),
}

func TestShouldBeLess(t *testing.T) {
	for x, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(x).Should().Be().Less(y)
		it.Ok(t).If(mock.Failed()).Should().Equal(false)
	}

	for x, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().Less(x)
		it.Ok(t).If(mock.Failed()).Should().Equal(true)
	}

	for _, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().Less(y)
		it.Ok(t).If(mock.Failed()).Should().Equal(true)
	}

	for x, y := range wrongtypes {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().Less(x)
		it.Ok(t).If(mock.Failed()).Should().Equal(true)
	}
}

func TestShouldBeLessOrEqual(t *testing.T) {
	for x, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(x).Should().Be().LessOrEqual(y)
		it.Ok(t).If(mock.Failed()).Should().Equal(false)
	}

	for _, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().LessOrEqual(y)
		it.Ok(t).If(mock.Failed()).Should().Equal(false)
	}

	for x, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().LessOrEqual(x)
		it.Ok(t).If(mock.Failed()).Should().Equal(true)
	}

	for x, y := range wrongtypes {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().LessOrEqual(x)
		it.Ok(t).If(mock.Failed()).Should().Equal(true)
	}
}

func TestShouldBeGreater(t *testing.T) {
	for x, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().Greater(x)
		it.Ok(t).If(mock.Failed()).Should().Equal(false)
	}

	for _, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().Greater(y)
		it.Ok(t).If(mock.Failed()).Should().Equal(true)
	}

	for x, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(x).Should().Be().Greater(y)
		it.Ok(t).If(mock.Failed()).Should().Equal(true)
	}

	for x, y := range wrongtypes {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().Greater(x)
		it.Ok(t).If(mock.Failed()).Should().Equal(true)
	}

}

func TestShouldBeGreaterOrEqual(t *testing.T) {
	for x, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().GreaterOrEqual(x)
		it.Ok(t).If(mock.Failed()).Should().Equal(false)
	}

	for _, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().GreaterOrEqual(y)
		it.Ok(t).If(mock.Failed()).Should().Equal(false)
	}

	for x, y := range pairs {
		mock := new(testing.T)
		it.Ok(mock).If(x).Should().Be().GreaterOrEqual(y)
		it.Ok(t).If(mock.Failed()).Should().Equal(true)
	}

	for x, y := range wrongtypes {
		mock := new(testing.T)
		it.Ok(mock).If(y).Should().Be().GreaterOrEqual(x)
		it.Ok(t).If(mock.Failed()).Should().Equal(true)
	}

}

func TestShouldBeIn(t *testing.T) {
	mock := new(testing.T)
	it.Ok(mock).If(5).Should().Be().In(0, 10)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).If(5).Should().Be().In(20, 30)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

	it.Ok(mock).If(35).Should().Be().In(20, 30)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

	it.Ok(mock).If(5).Should().Be().In(uint(20), uint(30))
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

	it.Ok(mock).If(5).Should().Be().In(20, uint(30))
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

	it.Ok(mock).If(5).Should().Be().In(uint(20), 30)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)
}

func TestShouldInterceptPanic(t *testing.T) {
	mock := new(testing.T)
	err := errors.New("emergency")

	it.Ok(mock).
		If(func() { panic(err) }).
		Should().Intercept(err)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).
		If(func() { panic(err) }).
		Should().Fail()
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).
		If(func() { panic(errors.New("other")) }).
		Should().Intercept(err)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

	it.Ok(mock).
		If(func() {}).
		Should().Intercept(err)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

	it.Ok(mock).
		If(func() {}).
		Should().Fail()
	it.Ok(t).If(mock.Failed()).Should().Equal(true)
}

func TestShouldInterceptError(t *testing.T) {
	mock := new(testing.T)
	err := errors.New("error")

	it.Ok(mock).
		If(func() error { return err }).
		Should().Intercept(err)
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).
		If(func() error { return err }).
		Should().Fail()
	it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.Ok(mock).
		If(func() error { return nil }).
		Should().Intercept(err)
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

	it.Ok(mock).
		If(func() error { return nil }).
		Should().Fail()
	it.Ok(t).If(mock.Failed()).Should().Equal(true)

}
