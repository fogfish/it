# Changelog and Migration guide

## Release v2

v2 | v1 | Why
--- | --- | ---
`it.Then` | `it.Ok` | encourage BDD-style Given-When-Then.
`it.Then(t).Should(it.Equal(x, y))` | `it.Ok(t).If(x).Should().Equal(y)` | imperative keywords takes results of asserts as input. The new api style supports better DSL with generics, hence compile time type level checks.


<!--
IfTrue
IfFalse
IfNil
IfNotNil
-->
