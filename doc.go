//
// Copyright (C) 2019 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/it
//

// Package it implements a human-friendly syntax for assertions to
// validates correctness of your code. It's style allows to write
// BDD-like specifications: "X should Y", "A equals to B", etc.
//
// Inspiration
//
// This library is heavily inspired by features of ScalaTest,
// see http://www.scalatest.org. It tries to adapt similar syntax for Golang.
// There is a vision that change of code style in testing helps developers to
// "switch gears". It's sole purpose to write unit tests assertions in
// natural language.
//
//  It Ok If /* actual */ Should /* expected */
//
//  It Ok If f() Should Equal 5
//  It Ok If f() Should Not Less 5
//  It Ok If f() Must Less 5
//
//
// Syntax at Glance
//
// Each assertion begins with phrase:
//  it Ok If ...
//
// Continues with one of imperative keyword as defined by RFC 2119
//   Must, Must Not, Should, Should Not, May
//
//
// Assertions with user-defined functions is a technique to define
// arbitrary boolean expression.
//
//   it.Ok(t).If(three).
//     Should().Assert(func(be interface{}) bool {
//	     (be > 1) && (be < 10) && (be != 5)
//     })
//
// Intercept any failures in target features
//
//   it.Ok(t).If(refToCodeBlock).
//      Should().Intercept(/* ... */)
//
// Matches equality and identity
//
//  it.Ok(t).
//    If(one).Should().Equal(1).
//    If(one).Should().Be().A(1)
//
// Matches type
//
//  it.Ok(t).If(one).
//    Should().Be().Like(1)
//
// Matches Order and Ranges
//
//  it.Ok(t).
//    If(three).Should().Be().Less(10).
//    If(three).Should().Be().LessOrEqual(3).
//    If(three).Should().Be().Greater(1).
//    If(three).Should().Be().GreaterOrEqual(3).
//    If(three).Should().Be().In(1, 10)
//
package it
