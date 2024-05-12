# It Should Be Tested.

The library implements a human-friendly syntax for assertions to validates correctness of your code. It's style allows to write BDD-like specifications: "X should Y", "A equals to B", etc.

[![Documentation](https://pkg.go.dev/badge/github.com/fogfish/it)](https://pkg.go.dev/github.com/fogfish/it)
[![Build Status](https://github.com/fogfish/it/workflows/test/badge.svg)](https://github.com/fogfish/it/actions/)
[![Git Hub](https://img.shields.io/github/last-commit/fogfish/it.svg)](http://travis-ci.org/fogfish/it)
[![Coverage Status](https://coveralls.io/repos/github/fogfish/it/badge.svg?branch=master)](https://coveralls.io/github/fogfish/it?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/fogfish/it)](https://goreportcard.com/report/github.com/fogfish/it)



## Inspiration

There is a vision that change of code style in testing helps developers to "switch gears". It's sole purpose to write unit tests assertions in natural language. The library is inspired by features of [ScalaTest](http://www.scalatest.org) and tries to adapt similar syntax for Golang. 

```
It Should /* actual */ Be /* expected */

It Should X Equal Y
It Should X Less Y
It Should String X Contain Y
It Should Seq X Equal Y¹, ... Yⁿ
```

The approximation of the style into Golang syntax:

```go
it.Then(t).
  Should(it.Equal(x, y)).
  Should(it.Less(x, y)).
  Should(it.String(x).Contain(y)).
  Should(it.Seq(x).Equal(/* ... */))
```

## Getting Started

- [It Should Be Tested.](#it-should-be-tested)
  - [Inspiration](#inspiration)
  - [Getting Started](#getting-started)
    - [Style](#style)
    - [Assertions](#assertions)
    - [Intercepts](#intercepts)
    - [Equality and identity](#equality-and-identity)
    - [Ordering](#ordering)
    - [String matchers](#string-matchers)
    - [Slices and Sequence matchers](#slices-and-sequence-matchers)
    - [Map matchers](#map-matchers)
  - [How To Contribute](#how-to-contribute)
    - [commit message](#commit-message)
    - [bugs](#bugs)
  - [License](#license)

The latest version of the library is available at its `main` branch. All development, including new features and bug fixes, take place on the `main` branch using forking and pull requests as described in contribution guidelines. The stable version is available via Golang modules.

1. Use `go get` to retrieve the library and add it as dependency to your application.

```bash
go get -u github.com/fogfish/it
```

2. Import it in your unit tests

```go
import (
  "github.com/fogfish/it"
)
```

See the [go doc](http://godoc.org/github.com/fogfish/it) for api spec.


### Style

The coding style is like standard Golang unit tests but assert are written as a chain of asserts in a specification style: "X should Y," "A must B," etc.

```go
func TestMyFeature(t *testing.T) {
  /* Given */
  /*  ...  */

  /* When  */
  /*  ...  */

  it.Then(t).
    Should(/* X Equal Y */).
    Should(/* String X Contain Y */)
}
```

The library support 3 imperative keyword `Must`, `Should` and `May` as defined by [RFC 2119](https://www.ietf.org/rfc/rfc2119.txt). Its prohibition variants `MustNot`, `ShouldNot` and `May`. 

Use the `Skip` imperative keyword to ignore the assert and its result.


### Assertions

A first order logic expression asserts the result of unit test.

```go
it.Then(t).
  // Inline logical predicates
  Should(it.True(x == y && x > 10)).
  // Use closed function
  Should(it.Be(func() bool { return x == y && x > 10})).
  // X should be same type as Y
  Should(it.SameAs(x, y)).
  // X should be nil
  Should(it.Nil(x))
```


### Intercepts

Intercept any failures in target code block. Intercepts supports actual panics and function that return of errors.

```go
func fWithPanic() {/* ... */}
func fWithError() error {/* ... */}

it.Then(t).
  // Intercept panic in the code block
  Should(it.Fail(fWithPanic)).
  // Intercept error in the code block
  Should(it.Fail(fWithError)).
```

Assert error for behavior to check the "type" of returned error

```go
var err interface { Timeout() }

it.Then(t).
  // Intercept panic in the code block and assert for behavior
  Should(it.Fail(fWithPanic).With(&err)).
  Should(it.Fail(fWithError).With(&err)).
  //  Intercept panic in the code block and match the error code
  Should(it.Fail(fWithError).Contain("error code"))
```

The `it.Fail` interceptor evaluates code block inside and it is limited to function that return single value. The `it.Error` interceptor captures returns of function.   

```go
func fNaryError() (string, error) {/* ... */}

it.Then(t).
  Should(it.Error(fNaryError())).
  Should(it.Error(fNaryError()).With(&err)).
  Should(it.Error(fNaryError()).Contain("error code"))
```


### Equality and identity

Match unit test results with equality constraint.

```go
it.Then(t).
  // X should be equal Y
  Should(it.Equal(x, y))
```

### Ordering

Compare unit test results with ordering constraint.

```go
it.Then(t).
  // X should be less than Y
  Should(it.Less(x, y)).
  // X should be less or equal to Y
  Should(it.LessOrEqual(x, y)).
  // X should be greater than Y
  Should(it.Greater(x, y)).
  // X should be greater or equal to Y
  Should(it.GreaterOrEqual(x, y)) 
```


### String matchers

```go
it.Then(t).
  // String X should have prefix X
  Should(it.String(x).HavePrefix(y)).
  // String X should have suffix X
  Should(it.String(x).HaveSuffix(y)).
  // String X should contain X
  Should(it.String(x).Contain(y)).
```

### Slices and Sequence matchers

```go
it.Then(t).
  // Seq X should be empty
  Should(it.Seq(x).BeEmpty(y)).
  // Seq X should equal Y¹, ... Yⁿ
  Should(it.Seq(x).Equal(y1, ..., yn))
  // Seq X should contain Y
  Should(it.Seq(x).Contain(y1, ..., yn))
  // Seq X should contain one of Y
  Should(it.Seq(x).Contain().OneOf(y1, ..., yn)).
  // Seq X should contain all of Y
  Should(it.Seq(x).Contain().AllOf(y1, ..., yn))
```


### Map matchers

```go
it.Then(t).
  // Map X should have key K with value Y
  Should(it.Map(X).Have(k, y)) 
```


## How To Contribute

The library is [MIT](LICENSE) licensed and accepts contributions via GitHub pull requests:

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


The build and testing process requires [Go](https://golang.org) version 1.13 or later.

**Build** and **run** service in your development console. The following command boots Erlang virtual machine and opens Erlang shell.

```bash
git clone https://github.com/fogfish/gurl
cd gurl
go test -cover
```

### commit message

The commit message helps us to write a good release note, speed-up review process. The message should address two question what changed and why. The project follows the template defined by chapter [Contributing to a Project](http://git-scm.com/book/ch5-2.html) of Git book.

### bugs

If you experience any issues with the library, please let us know via [GitHub issues](https://github.com/fogfish/it/issue). We appreciate detailed and accurate reports that help us to identity and replicate the issue. 

## License

[![See LICENSE](https://img.shields.io/github/license/fogfish/it.svg?style=for-the-badge)](LICENSE)
