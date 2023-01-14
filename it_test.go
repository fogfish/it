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

	"github.com/fogfish/it/v2"
)

func TestNewExpr(t *testing.T) {
	it.Then(t)
}

func TestExprAliases(t *testing.T) {
	mock := new(testing.T)

	it.Then(mock).Should(it.Equal("xxx", "xxx"))
	it.Then(t).ShouldNot(it.Be(mock.Failed))

	it.Then(mock).ShouldNot(it.Equal("xxx", "yyy"))
	it.Then(t).ShouldNot(it.Be(mock.Failed))
}

func TestImperativeKeywords(t *testing.T) {
	success := func() error { return nil }
	failure := func() error { return errors.New("fail") }

	mock := new(testing.T)
	it.Then(mock).Must(success())
	it.Then(t).ShouldNot(it.Be(mock.Failed))

	it.Then(t).Should(it.Fail(
		func() { it.Ok(mock).Must(failure()) },
	))

	mock = new(testing.T)
	it.Then(mock).Should(success())
	it.Then(t).ShouldNot(
		it.Be(mock.Failed),
		it.Be(mock.Failed),
	)

	mock = new(testing.T)
	it.Then(mock).Should(failure())
	it.Then(t).Should(
		it.Be(mock.Failed),
		it.Be(mock.Failed),
	)

	mock = new(testing.T)
	it.Then(mock).May(success())
	it.Then(t).ShouldNot(it.Be(mock.Failed))

	mock = new(testing.T)
	it.Then(mock).May(failure())
	it.Then(t).ShouldNot(it.Be(mock.Failed))
}
