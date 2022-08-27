//
// Copyright (C) 2019 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/it
//

/*

Package it implements a human-friendly syntax for assertions to
validates correctness of your code. It's style allows to write
BDD-like specifications: "X should Y", "A equals to B", etc.

Inspiration

There is a vision that change of code style in testing helps developers to
"switch gears". It's sole purpose to write unit tests assertions in natural
language. The library is inspired by features of http://www.scalatest.org
and tries to adapt similar syntax for Golang.

  It Should / actual / Be / expected /

  It Should X Equal Y
  It Should X Less Y
  It Should String X Contain Y
  It Should Seq X Equal Y¹, ... Yⁿ

The approximation of the style into Golang syntax:

  it.Then(t).
    Should(it.Equal(x, y)).
    Should(it.Less(x, y)).
    Should(it.String(x).Contain(y)).
    Should(it.Seq(x).Equal(/ ... /))


Syntax at Glance

Each assertion begins with phrase:
 it Then ...

Continues with one of imperative keyword as defined by RFC 2119
  Must, Must Not, Should, Should Not, May, May Not


Assertions with inline closure is a technique to define
arbitrary boolean expression.

  it.Then(t).
    Should(it.True(be > 1 && be < 10 && be != 5))

Intercept any failures in target features

  it.Then(t).
    ShouldNot(it.Fail(refToCodeBlock))

Matches equality and identity

  it.Then(t).
    Should(it.Equal(x, y))

Matches type

  it.Then(t).
    Should(it.SameAs(x, y))

Matches Order and Ranges

  it.Ok(t).
    Should(it.Less(x, 10)).
    Should(it.LessOrEqual(x, 3)).
    Should(it.Greater(x, 1)).
    Should(it.GreaterOrEqual(x, 3)).

*/
package it
