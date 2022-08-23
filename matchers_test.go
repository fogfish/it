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
