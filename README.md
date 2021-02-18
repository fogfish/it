# It Should Be Tested.

The library implements a human-friendly syntax for assertions to validates correctness of your code. It's style allows to write BDD-like specifications: "X should Y", "A equals to B", etc.

[![Documentation](https://pkg.go.dev/badge/github.com/fogfish/it)](https://pkg.go.dev/github.com/fogfish/it)
[![Build Status](https://github.com/fogfish/it/workflows/build/badge.svg)](https://github.com/fogfish/it/actions/)
[![Git Hub](https://img.shields.io/github/last-commit/fogfish/it.svg)](http://travis-ci.org/fogfish/it)
[![Coverage Status](https://coveralls.io/repos/github/fogfish/it/badge.svg?branch=master)](https://coveralls.io/github/fogfish/it?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/fogfish/it)](https://goreportcard.com/report/github.com/fogfish/it)
[![Maintainability](https://api.codeclimate.com/v1/badges/d685a9d909983da3d2da/maintainability)](https://codeclimate.com/github/fogfish/it/maintainability)


## Inspiration

This library is heavily inspired by features of [ScalaTest](http://www.scalatest.org). It tries to adapt similar syntax for Golang. There is a vision that change of code style in testing helps developers to "switch gears". It's sole purpose to write unit tests assertions in natural language.

```
It Ok If /* actual */ Should /* expected */

It Ok If x Should Equal 5
It Ok If x Should Not Less 5
It Ok If x Must Less 5
```

## Key features

* assertions with user defined functions
* intercept failures
* check equality
* makes comparison and range checks
* type checks
* match failures and error to given type
* imperative keyword as defined by RFC 2119
* composition/chaining of assertions


## Getting Started

The latest version of the library is available at its `master` branch. All development, including new features and bug fixes, take place on the `master` branch using forking and pull requests as described in contribution guidelines.

Import the library in your code

```go
import (
  "github.com/fogfish/it"
)

func TestMyFeature(t *testing.T) {
  it.Ok(t).
    If(myfeature()).Should().Equal(5).
    If(myfeature()).Should().Be().Less(10)
}

func myfeature() int {
  return 5
}
```

See the [go doc](http://godoc.org/github.com/fogfish/it) for api spec.


## Syntax at Glance

Each assertion begins with phrase: 

```go
func TestMyFeature(t *testing.T) {
  it.Ok(t).If(/*...*/)
}
```

Then, it continues with one of imperative keyword as defined by [RFC 2119](https://www.ietf.org/rfc/rfc2119.txt)  :
* `Must` the definition is an absolute requirement.
* `MustNot` the definition is an absolute prohibition.
* `Should` the definition is a strongly recommended requirement, however it's violation do not block continuation of testing.
* `ShouldNot` the definition is prohibited, however it's violation do not block continuation of testing.
* `May` is an optional constrain, its violation do not impact on testing flow in anyway.

The library provides a rich set of asserts 

```go
// Assertions with user-defined functions
it.Ok(t).If(three).
  Should().Assert(func(be interface{}) bool {
      (be > 1) && (be < 10) && (be != 5)
  })

// Intercept any failures in target features
it.Ok(t).If(refToCodeBlock).
  Should().Intercept(/* ... */)

// Matches equality and identity
it.Ok(t).
  If(one).Should().Equal(1).
  If(one).Should().Be().A(1)

// Matches type
it.Ok(t).If(one).
  Should().Be().Like(1)

// Matches Order and Ranges
it.Ok(t).
  If(three).Should().Be().Less(10).
  If(three).Should().Be().LessOrEqual(3).
  If(three).Should().Be().Greater(1).
  If(three).Should().Be().GreaterOrEqual(3).
  If(three).Should().Be().In(1, 10)
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
