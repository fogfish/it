# Changelog and Migration guide

## Release v2

### New features

* Human friendly logging of passed and failed cases
* imperative keyword `Skip` to disable results of assertion
* better support of containers `String`, `Seq` and `Map`

### Breaking changes

* `it.Ok ⇒ it.Then` encourage BDD-style Given-When-Then.
* `it.Ok(t).If(x).Should().Equal(y) ⇒ it.Then(t).Should(it.Equal(x, y))` imperative keywords takes results of asserts as input. The new api style supports better DSL with generics, hence compile time type level checks
* `Assert(func(interface{}) bool) ⇒ it.Be(func() bool)` assert functions became inline closures
* `If(codeBlock).Should().Intercept(errors) ⇒ ShouldNot(it.Fail(codeBlock).With(errors))`
* `If(codeBlock).Should().Fail() ⇒ Should(it.Fail(codeBlock))`
* `If(x).Should().Equal(y) ⇒ it.Equal(x, y)` became type-safe assert on equality
* `If(x).Should().Eq(y) ⇒ it.Equal(x, y)` no support for abbreviations
* `If(x).Should().Equiv(y)` removed
* `If(x).Should().Be().A(x) ⇒ it.Equal(x, y)`
* `If(x).Should().Be().Like(y) ⇒ it.SameAs(x, y)` became a compile type assert
* `If(x).Should().Be().Less(x) ⇒ it.Less(x, y)` same migration pattern is applicable for `LessOrEqual`, `Greater`, `GreaterOrEqual`
* `If(x).Should().Be().In(from, to)` removed, use `it.Less` and `it.Greater` constraints
* `it.Ok(t).IfTrue(x)` removed together with other aliases `IfFalse`, `IfNil`, `IfNotNil`, `NotEqual`, `Equal`. 
