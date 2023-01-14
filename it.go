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
	"testing"
)

//
// Imperatives
//

// Check type is constructor of imperative testing expressions.
//
//	it.Then(t).Should( ... )
type Check struct {
	t *testing.T // Note: intentionally hidden from clients
}

// Then creates assertion expression, it takes a reference to
// standard *testing.T interface to setup the evaluation context.
func Then(t *testing.T) *Check { return &Check{t} }

// Ok alias to Then
func Ok(t *testing.T) *Check { return &Check{t} }

// Must is imperative keyword.
// The assert definition is an absolute requirement.
// It terminates execution of tests if assert is failed.
func (check *Check) Must(errs ...error) *Check {
	check.t.Helper()

	for _, err := range errs {
		if err != nil {
			var e interface{ Passed() bool }
			ok := errors.As(err, &e)
			output := check.fatalf

			if ok && e.Passed() {
				output = check.noticef
			}
			output("must %s", err)
		}
	}

	return check
}

// MustNot is imperative keyword.
// The definition is an absolute prohibition
// It terminates execution of tests if assert is not failed.
func (check *Check) MustNot(errs ...error) *Check {
	check.t.Helper()

	for _, err := range errs {
		if err != nil {
			var e interface{ Passed() bool }
			ok := errors.As(err, &e)
			output := check.fatalf

			if !(ok && e.Passed()) {
				output = check.noticef
			}
			output("must not %s", err)
		}
	}

	return check
}

// Should is imperative keyword.
// The assert definition is a strongly recommended requirement.
// However it's violation do not block continuation of testing.
// The test fails if assert is failed.
func (check *Check) Should(errs ...error) *Check {
	check.t.Helper()

	for _, err := range errs {
		if err != nil {
			var e interface{ Passed() bool }
			ok := errors.As(err, &e)
			output := check.errorf

			if ok && e.Passed() {
				output = check.noticef
			}
			output("should %s", err)
		}
	}

	return check
}

// ShouldNot is imperative keyword.
// The assert definition is a strongly recommended prohibition.
// However it's violation do not block continuation of testing.
// The test fails if assert is not failed.
func (check *Check) ShouldNot(errs ...error) *Check {
	check.t.Helper()

	for _, err := range errs {
		if err != nil {
			var e interface{ Passed() bool }
			ok := errors.As(err, &e)
			output := check.errorf

			if !(ok && e.Passed()) {
				output = check.noticef
			}
			output("should not %s", err)
		}
	}

	return check
}

// May is an optional imperative constrain.
// Its violation do not impact on testing flow in anyway.
// The informative warning message is produced to console if assert is failed.
// Error message will be printed only if the test fails or the -test.v
func (check *Check) May(errs ...error) *Check {
	check.t.Helper()

	for _, err := range errs {
		if err != nil {
			var e interface{ Passed() bool }
			ok := errors.As(err, &e)
			output := check.warningf

			if ok && e.Passed() {
				output = check.noticef
			}
			output("may %s", err)
		}
	}

	return check
}

// MayNot is an optional imperative constrain.
// Its violation do not impact on testing flow in anyway.
// The informative warning message is produced to console if assert is not failed.
// Error message will be printed only if the test fails or the -test.v
func (check *Check) MayNot(errs ...error) *Check {
	check.t.Helper()

	for _, err := range errs {
		if err != nil {
			var e interface{ Passed() bool }
			ok := errors.As(err, &e)
			output := check.warningf

			if !(ok && e.Passed()) {
				output = check.noticef
			}
			output("may not %s", err)
		}
	}

	return check
}

// Skip is imperative keyword
// It ignores results of assert
func (check *Check) Skip(errs ...error) *Check {
	check.t.Helper()
	for _, err := range errs {
		check.debugf("skip %s", err)
	}
	return check
}

func (check *Check) fatalf(msg string, args ...any) {
	check.t.Helper()
	check.t.Logf("\033[31m%s\033[0m", fmt.Sprintf(msg, args...))
	panic(fmt.Errorf("stop testing"))
}

func (check *Check) errorf(msg string, args ...any) {
	check.t.Helper()
	check.t.Errorf("\033[31m%s\033[0m", fmt.Sprintf(msg, args...))
}

func (check *Check) warningf(msg string, args ...any) {
	check.t.Helper()
	check.t.Logf("\033[33m%s\033[0m", fmt.Sprintf(msg, args...))
}

func (check *Check) noticef(msg string, args ...any) {
	check.t.Helper()
	check.t.Logf("\033[32m%s\033[0m", fmt.Sprintf(msg, args...))
}

func (check *Check) debugf(msg string, args ...any) {
	check.t.Helper()
	check.t.Logf("%s", fmt.Sprintf(msg, args...))
}

// ok labels assert with success
type ok struct{ err error }

func passed(err error) *ok       { return &ok{err} }
func (e *ok) Error() string      { return e.err.Error() }
func (e *ok) As(target any) bool { return errors.As(e.err, target) }
func (e *ok) Unwarp() error      { return e.err }
func (e *ok) Passed() bool       { return true }
