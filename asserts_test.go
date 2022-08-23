package it_test

import (
	"testing"

	"github.com/fogfish/it"
)

func TestBe(t *testing.T) {
	it.Then(t).
		Should(it.Be(func() bool { return true })).
		ShouldNot(it.Be(func() bool { return false }))
}

func TestTrue(t *testing.T) {
	it.Then(t).
		Should(it.True(10 == 10)).
		ShouldNot(it.True(10 > 10))
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
