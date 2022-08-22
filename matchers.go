package it

import (
	"errors"
	"fmt"
	"strings"
)

//
// String
//

type String string

func (x String) HavePrefix(y string) error {
	assert := fmt.Errorf("string %s have prefix %s", x, y)

	if !strings.HasPrefix(string(x), y) {
		return assert
	}

	return passed(assert)
}

func (x String) HaveSuffix(y string) error {
	assert := fmt.Errorf("string %s have suffix %s", x, y)

	if !strings.HasSuffix(string(x), y) {
		return assert
	}

	return passed(assert)
}

func (x String) Contain(y string) error {
	assert := fmt.Errorf("string %s contain %s", x, y)

	if !strings.Contains(string(x), y) {
		return assert
	}

	return passed(assert)
}

//
// Sequence of elements
//

type SeqOf[A any] []A

func Seq[A any](seq []A) SeqOf[A] {
	return SeqOf[A](seq)
}

func (xs SeqOf[A]) BeEmpty() error {
	assert := fmt.Errorf("seq %v be empty", xs)

	if len(xs) != 0 {
		return assert
	}

	return passed(assert)
}

func (xs SeqOf[A]) Equal(ys ...A) error {
	if len(xs) != len(ys) {
		return fmt.Errorf("seq %v length be equal to %v", xs, ys)
	}

	for i, x := range xs {
		if !equal(x, ys[i]) {
			return fmt.Errorf("seq %dth element of %v be equal to %v", i, x, ys[i])
		}
	}

	return passed(fmt.Errorf("seq %v is equal to %v", xs, ys))
}

func (xs SeqOf[A]) Contain(ys ...A) SeqContainIt[A] {
	assert := fmt.Errorf("seq %v contain %v", xs, ys)

	for _, y := range ys {
		has := false
		for _, x := range xs {
			if equal(x, y) {
				has = true
				break
			}
		}
		if !has {
			return SeqContainIt[A]{assert, xs}
		}
	}

	return SeqContainIt[A]{passed(assert), xs}
}

// FailIt extend Fail assert
type SeqContainIt[A any] struct {
	assert error
	xs     []A
}

func (x SeqContainIt[A]) Error() string      { return x.assert.Error() }
func (x SeqContainIt[A]) As(target any) bool { return errors.As(x.assert, target) }

func (x SeqContainIt[A]) AllOf(ys ...A) error {
	assert := fmt.Errorf("seq %v contain all of %v", x.xs, ys)

	for _, y := range ys {
		has := false
		for _, x := range x.xs {
			if equal(x, y) {
				has = true
				break
			}
		}
		if !has {
			return assert
		}
	}

	return passed(assert)
}

func (x SeqContainIt[A]) OneOf(ys ...A) error {
	assert := fmt.Errorf("seq %v contain one of %v", x.xs, ys)

	for _, y := range ys {
		for _, x := range x.xs {
			if equal(x, y) {
				return passed(assert)
			}
		}
	}

	return assert
}

//
// Map of elements
//

type MapOf[K comparable, V any] map[K]V

func Map[K comparable, V any](val map[K]V) MapOf[K, V] {
	return MapOf[K, V](val)
}

func (xs MapOf[K, V]) Have(key K, y V) error {
	x, exists := xs[key]
	if !exists {
		return fmt.Errorf("map %v have key %v", xs, key)
	}

	assert := fmt.Errorf("key %v value %v of %T be equal to %v", key, x, (map[K]V)(xs), y)

	if !equal(x, y) {
		return assert
	}

	return passed(assert)
}
