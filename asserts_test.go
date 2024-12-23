//
// Copyright (C) 2019 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/it
//

package it_test

import (
	"testing"

	"github.com/fogfish/it/v2"
)

func TestBe(t *testing.T) {
	it.Then(t).
		Should(it.Be(func() bool { return true })).
		ShouldNot(it.Be(func() bool { return false }))
}

func TestTrue(t *testing.T) {
	it.Then(t).
		Should(it.True(true)).
		ShouldNot(it.True(false))
}

type err string

func (e err) Error() string { return string(e) }
func (e err) Behavior()     {}

func TestFail(t *testing.T) {
	fWithPanic := func() { panic(err("func with panic")) }
	fNoPanic := func() {}
	fWithError := func() error { return err("func with error") }
	fNoError := func() error { return nil }

	it.Then(t).
		Should(it.Fail(fWithPanic)).
		ShouldNot(it.Fail(fNoPanic)).
		Should(it.Fail(fWithError)).
		ShouldNot(it.Fail(fNoError))
}

func TestFailWith(t *testing.T) {
	var thisIsErr interface{ Behavior() }
	var thisIsNotErr interface{ Timeout() }

	fWithPanic := func() { panic(err("func with panic")) }
	fWithError := func() error { return err("func with error") }

	it.Then(t).
		Should(it.Fail(fWithPanic).With(&thisIsErr)).
		Should(it.Fail(fWithPanic).Contain("with panic")).
		ShouldNot(it.Fail(fWithPanic).With(&thisIsNotErr)).
		ShouldNot(it.Fail(fWithPanic).Contain("with error")).
		Should(it.Fail(fWithError).With(&thisIsErr)).
		Should(it.Fail(fWithError).Contain("with error")).
		ShouldNot(it.Fail(fWithError).With(&thisIsNotErr)).
		ShouldNot(it.Fail(fWithError).Contain("with panic"))
}

func TestError(t *testing.T) {
	var thisIsErr interface{ Behavior() }
	var thisIsNotErr interface{ Timeout() }

	fWithError := func() (string, error) { return "", err("func with error") }
	fNoError := func() (int, error) { return 10, nil }

	it.Then(t).
		Should(it.Error(fWithError())).
		Should(it.Error(fWithError()).With(&thisIsErr)).
		ShouldNot(it.Error(fWithError()).With(&thisIsNotErr)).
		ShouldNot(it.Error(fNoError()))
}

func TestSameAs(t *testing.T) {
	it.Then(t).
		Should(it.SameAs("abc", "def"))
}

func TestType(t *testing.T) {
	it.Then(t).Should(
		it.TypeOf[string]("x"),
	)
}

func TestNil(t *testing.T) {
	it.Then(t).
		Should(it.Nil(nil)).
		ShouldNot(it.Nil("abc"))
}

func TestEqual(t *testing.T) {
	it.Then(t).
		Should(it.Equal(1, 1)).
		ShouldNot(it.Equal(1, 2))
}

func TestEquiv(t *testing.T) {
	type T struct{ A string }

	it.Then(t).
		Should(it.Equiv(T{"A"}, T{"A"})).
		Should(it.Equiv(&T{"A"}, &T{"A"})).
		ShouldNot(it.Equiv(T{"A"}, T{"B"})).
		ShouldNot(it.Equiv(&T{"A"}, &T{"B"}))
}

func TestLike(t *testing.T) {
	type T struct{ A string }

	it.Then(t).
		Should(it.Like(T{"A"}, T{"A"})).
		Should(it.Like(&T{"A"}, &T{"A"})).
		ShouldNot(it.Like(T{"A"}, T{"B"})).
		ShouldNot(it.Like(&T{"A"}, &T{"B"}))
}

func TestLess(t *testing.T) {
	it.Then(t).
		Should(it.Less(0, 1)).
		ShouldNot(it.Less(1, 1))
}

func TestLessOrEqual(t *testing.T) {
	it.Then(t).
		Should(it.LessOrEqual(0, 1)).
		Should(it.LessOrEqual(1, 1)).
		ShouldNot(it.LessOrEqual(1, 0))
}

func TestGreater(t *testing.T) {
	it.Then(t).
		Should(it.Greater(1, 0)).
		ShouldNot(it.Greater(0, 1))
}

func TestGreaterOrEqual(t *testing.T) {
	it.Then(t).
		Should(it.GreaterOrEqual(1, 0)).
		Should(it.GreaterOrEqual(1, 1)).
		ShouldNot(it.GreaterOrEqual(0, 1))
}
