# Changelog and Migration guide

## Release v2

* `it.Ok ⇒ it.Then` encourage BDD-style Given-When-Then.
* `it.Ok(t).If(x).Should().Equal(y) ⇒ it.Then(t).Should(it.Equal(x, y))` imperative keywords takes results of asserts as input. The new api style supports better DSL with generics, hence compile time type level checks.


<!--
IfTrue
IfFalse
IfNil
IfNotNil
-->
