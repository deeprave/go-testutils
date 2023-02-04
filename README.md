# Test support functions for golang

## Description

The `test` package contains a number of easily comprehensible utility functions that abbreviate
a lot of the boilerplate code associated with Go unit testing.

## Functions

The following functions resemble simple assertions that return a `bool` type indicate whether the tested 
value succeeds.
This allows for immediate exit from the test function on test failure.

### Equality

`ShouldBeEqual(*testing.T, interface{}, interface{}, ...options) bool`
:  compares two values which are expected to be indentical via deep comparison

`ShouldNotBeEqual(*testing.T, interface{}, interface{}, ...options) bool`
:  compares two values which are expected not to be indentical via deep comparison

### True/False

`ShouldBeTrue(*testing T, bool, format, ...args) bool`
:  tests if the result of an expression is true

`ShouldBeFalse(*testing T, bool, format, ...args) bool`
:  tests if the result of an expression is false

### Nil

`ShouldBeNil(*testing T, interface{}, format, ...args) bool`
:  tests if a value is nil

`ShouldNotBeNil(*testing T, interface{}, format, ...args) bool`
:  tests if a value is not nil

### In Array

`ShouldBeInArray(*testing.T, []V, V, ...options) bool`
:  tests if a value is in the provided array

`ShouldNotBeInArray(*testing.T, []V, V, ...options) bool`
:  tests if a value is not in the provided array

### Error

`ShouldBeError(*testing.T, error, format, ...args) bool`
:  error is expected to be non-null

`ShouldBeNoError(*testing.T, error, format, ...args) bool`
:  error is expected to be null
