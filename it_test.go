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

func TestZ(t *testing.T) {
	it.OkZ(t).
		// If(it.Equal(1, 3)).
		Should(it.Equal(3, 3))
	// Should(it.Nil("xxx")).
	// Should(it.NotNil(nil)).
	// Should(it.Less(8, 7)).
	// Should(it.NotFail(func() {
	// 	panic(fmt.Errorf("xx"))
	// }))
}

func TestNewExpr(t *testing.T) {
	it.OkZ(t)
}

func TestExprAliases(t *testing.T) {
	mock := new(testing.T)

	// it.Ok(mock).IfTrue(true)
	// it.Ok(t).If(mock.Failed()).Should().Equal(false)

	// it.Ok(mock).IfFalse(false)
	// it.Ok(t).If(mock.Failed()).Should().Equal(false)

	// it.Ok(mock).IfNil(nil)
	// it.Ok(t).If(mock.Failed()).Should().Equal(false)

	// it.Ok(mock).IfNotNil("xxx")
	// it.Ok(t).If(mock.Failed()).Should().Equal(false)

	it.OkZ(mock).If(it.Equal("xxx", "xxx"))
	it.OkZ(t).If(it.IsNot(mock.Failed))

	it.OkZ(mock).If(it.NotEqual("xxx", "yyy"))
	it.OkZ(t).If(it.IsNot(mock.Failed))
}

func TestImperativeKeywords(t *testing.T) {
	success := func() error { return nil }
	failure := func() error { return errors.New("fail") }

	mock := new(testing.T)
	it.OkZ(mock).Must(success())
	it.OkZ(t).If(it.IsNot(mock.Failed))

	it.OkZ(t).If(it.Failed(
		func() {
			it.OkZ(mock).Must(failure())
		},
	))

	mock = new(testing.T)
	it.OkZ(mock).Should(success())
	it.OkZ(t).If(it.IsNot(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).Should(failure())
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).May(success())
	it.OkZ(t).If(it.IsNot(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).May(failure())
	it.OkZ(t).If(it.IsNot(mock.Failed))
}

func TestShouldLogicalExpression(t *testing.T) {
	valid := func() bool { return true }
	notvalid := func() bool { return false }

	mock := new(testing.T)
	it.OkZ(mock).If(it.Is(valid))
	it.OkZ(t).If(it.IsNot(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.IsNot(valid))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.Is(notvalid))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.IsNot(notvalid))
	it.OkZ(t).If(it.IsNot(mock.Failed))
}

func TestShouldEqual(t *testing.T) {
	mock := new(testing.T)
	it.OkZ(mock).If(it.Equal(1, 1))
	it.OkZ(t).If(it.IsNot(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.NotEqual(1, 1))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.Equal(1, 2))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.NotEqual(1, 2))
	it.OkZ(t).If(it.IsNot(mock.Failed))
}

/*
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
*/

/*
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
*/

func TestShouldSameAs(t *testing.T) {
	mock := new(testing.T)
	it.OkZ(mock).If(it.SameAs(1, 0))
	it.OkZ(t).If(it.IsNot(mock.Failed))

	it.OkZ(mock).If(it.SameAs("abc", ""))
	it.OkZ(t).If(it.IsNot(mock.Failed))
}

/*
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
*/

/*
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
*/

func TestShouldBeLess(t *testing.T) {
	mock := new(testing.T)
	it.OkZ(mock).If(it.Less(0, 1))
	it.OkZ(t).If(it.IsNot(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.Less(1, 1))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.Less(1, 0))
	it.OkZ(t).If(it.Is(mock.Failed))
}

func TestShouldBeLessOrEqual(t *testing.T) {
	mock := new(testing.T)
	it.OkZ(mock).If(it.LessOrEqual(0, 1))
	it.OkZ(t).If(it.IsNot(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.LessOrEqual(1, 1))
	it.OkZ(t).If(it.IsNot(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.LessOrEqual(1, 0))
	it.OkZ(t).If(it.Is(mock.Failed))
}

func TestShouldBeGreater(t *testing.T) {
	mock := new(testing.T)
	it.OkZ(mock).If(it.Greater(0, 1))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.Greater(1, 1))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.Greater(1, 0))
	it.OkZ(t).If(it.IsNot(mock.Failed))
}

func TestShouldBeGreaterOrEqual(t *testing.T) {
	mock := new(testing.T)
	it.OkZ(mock).If(it.GreaterOrEqual(0, 1))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.GreaterOrEqual(1, 1))
	it.OkZ(t).If(it.IsNot(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).If(it.GreaterOrEqual(1, 0))
	it.OkZ(t).If(it.IsNot(mock.Failed))
}

/*
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
*/

func TestShouldFailedWithPanic(t *testing.T) {
	err := errors.New("emergency")

	mock := new(testing.T)
	it.OkZ(mock).Should(it.Failed(func() { panic(err) }))
	it.OkZ(t).If(it.IsNot(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).Should(it.Failed(func() {}))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).Should(it.NotFailed(func() { panic(err) }))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).Should(it.NotFailed(func() {}))
	it.OkZ(t).If(it.IsNot(mock.Failed))
}

func TestShouldFailedWithError(t *testing.T) {
	err := errors.New("emergency")

	mock := new(testing.T)
	it.OkZ(mock).Should(it.Failed(func() error { return err }))
	it.OkZ(t).If(it.IsNot(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).Should(it.Failed(func() error { return nil }))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).Should(it.NotFailed(func() error { return err }))
	it.OkZ(t).If(it.Is(mock.Failed))

	mock = new(testing.T)
	it.OkZ(mock).Should(it.NotFailed(func() error { return nil }))
	it.OkZ(t).If(it.IsNot(mock.Failed))
}

/*
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
*/
