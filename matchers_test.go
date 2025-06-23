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

func TestStringPrefix(t *testing.T) {
	it.Ok(t).
		Should(it.String("abcdef").HavePrefix("abc")).
		ShouldNot(it.String("abcdef").HavePrefix("def"))
}

func TestStringSuffix(t *testing.T) {
	it.Ok(t).
		ShouldNot(it.String("abcdef").HaveSuffix("abc")).
		Should(it.String("abcdef").HaveSuffix("def"))
}

func TestStringContain(t *testing.T) {
	it.Ok(t).
		Should(it.String("abcdef").Contain("cde")).
		ShouldNot(it.String("abcdef").Contain("xxx"))
}

func TestSeqEmpty(t *testing.T) {
	type T struct{ string }
	seq := []T{{"a"}, {"b"}, {"c"}}

	it.Ok(t).
		Should(it.Seq([]T{}).BeEmpty()).
		ShouldNot(it.Seq(seq).BeEmpty())
}

func TestSeqEqual(t *testing.T) {
	type T struct{ string }
	seq := []T{{"a"}, {"b"}, {"c"}}

	it.Ok(t).
		Should(it.Seq(seq).Equal(seq...)).
		ShouldNot(it.Seq(seq).Equal(T{"a"})).
		ShouldNot(it.Seq(seq).Equal([]T{{"a"}, {"bz"}, {"c"}}...))
}

func TestSeqContain(t *testing.T) {
	type T struct{ string }
	seq := []T{{"a"}, {"b"}, {"c"}}

	it.Ok(t).
		Should(it.Seq(seq).Contain(T{"b"})).
		ShouldNot(it.Seq(seq).Contain(T{"x"})).
		Should(it.Seq(seq).Contain().AllOf(T{"b"}, T{"c"})).
		ShouldNot(it.Seq(seq).Contain().AllOf(T{"b"}, T{"x"})).
		Should(it.Seq(seq).Contain().OneOf(T{"x"}, T{"c"})).
		ShouldNot(it.Seq(seq).Contain().OneOf(T{"y"}, T{"x"}))
}

func TestMapHave(t *testing.T) {
	type T struct{ string }
	set := map[int]T{100: {"a"}, 200: {"b"}, 300: {"c"}}

	it.Ok(t).
		Should(it.Map(set).Have(100, T{"a"})).
		ShouldNot(it.Map(set).Have(101, T{"a"})).
		ShouldNot(it.Map(set).Have(101, T{"a"})).
		ShouldNot(it.Map(set).Have(200, T{"a"}))
}

func TestJson(t *testing.T) {
	type S []string
	type M map[string]any
	t.Run("Success", func(t *testing.T) {
		for pat, val := range map[string]any{
			`null`:                  nil,
			`"_"`:                   "foo",
			`"foo"`:                 "foo",
			`10`:                    10,
			`10.33`:                 10.33,
			`true`:                  true,
			`["foo", "bar"]`:        S{"foo", "bar"},
			`["_", "bar"]`:          S{"foo", "bar"},
			`["foo", "..."]`:        S{"foo", "bar", "foobar"},
			`{"foo": "bar"}`:        M{"foo": "bar"},
			`{"foo": "m/ar/"}`:      M{"foo": "bar"},
			`{"foo": "_"}`:          M{"foo": "bar"},
			`{"foo": ["bar"]}`:      M{"foo": S{"bar"}},
			`{"foo": ["_"]}`:        M{"foo": S{"bar"}},
			`{"foo": {"bar": "f"}}`: M{"foo": M{"bar": "f"}},
		} {
			mock := new(testing.T)
			it.Ok(mock).Should(
				it.Json(val).Equiv(pat),
			)
			it.Then(t).ShouldNot(it.Be(mock.Failed))
		}
	})

	t.Run("Failed", func(t *testing.T) {
		for pat, val := range map[string]any{
			`null`:                  "xfoo",
			`"foo"`:                 "xfoo",
			`"bar"`:                 100,
			`10`:                    100,
			`10.33`:                 100.33,
			`100`:                   "100",
			`true`:                  false,
			`false`:                 "false",
			`["foo"]`:               S{"foo", "xbar"},
			`["foo", "bar"]`:        S{"foo", "xbar"},
			`["_", "bar"]`:          S{"foo", "xbar"},
			`["_"]`:                 nil,
			`{"foo": "bar"}`:        M{"foo": "xbar"},
			`{"foo": "m/^bar/"}`:    M{"foo": "xbar"},
			`{"foo": "_"}`:          M{"xfoo": "bar"},
			`{"foo": ["bar"]}`:      M{"foo": S{"xbar"}},
			`{"foo": ["_"]}`:        M{"xfoo": S{"bar"}},
			`{"foo": {"bar": "f"}}`: M{"foo": M{"bar": "xf"}},
		} {
			mock := new(testing.T)
			it.Ok(mock).Should(
				it.Json(val).Equiv(pat),
			)
			it.Then(t).Should(it.Be(mock.Failed))
		}
	})
}
