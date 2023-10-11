# Generics

Generics enables you to write `functions and data structures that work with different types while maintaining type safety`. You can define functions and structures that operate on various types without sacrificing the compiler's ability to catch type errors at compile time.

??? example

    ```bash title="run command"
    go run src/generics/basic.go
    ```
    ```go
    --8<-- "src/generics/basic.go"
    ```
    ```bash title="output"
    FindMax integer: 9
    FindMax float: 3.14
    ```

## Combine Multiple Types (Constraint)

In Go, `generic constraints are a way to specify requirements for the type parameters used in generic functions or data structures`. Constraints `ensure that the generic code can only be used with types that meet certain criteria`.

??? example

    ```bash title="run command"
    go run src/generics/constraint.go
    ```
    ```go
    --8<-- "src/generics/constraint.go"
    ```
    ```bash title="output"
    FindMax integer: 9
    FindMax float: 3.14
    ```

## Tilde(~) - Closer Than Identical Types

The special tilde (`~`) symbol, often referred to as the `approximation constraint`, is used in Go generics to `indicate that a type should approximately match one of the specified types or their derived types`. `Without the tilde, the function or method constrained by generic types would only accept exact types declared in the constraints and not their derived types`. The tilde allows for a broader match, including derived types, making your code more flexible and accommodating related types.

??? example

    ```bash title="run command"
    go run src/generics/approximation_constraint.go
    ```
    ```go
    --8<-- "src/generics/approximation_constraint.go"
    ```
    ```bash title="output"
    FindMax integer: 9
    FindMax float: 3.14
    ```

## Comparable

In Go generics, you can use the `comparable constraint` to specify that a type should be comparable. This constraint ensures that the type can be used with comparison operators like `==` and `!=`.

??? example

    ```bash title="run command"
    go run src/generics/comparable.go
    ```
    ```go
    --8<-- "src/generics/comparable.go"
    ```
    ```bash title="output"
    CheckNumbers:  false
    CheckNumbers:  true
    FindMax integer: -1
    FindMax integer: 3
    FindMax integer: -1
    ```

## References

- [An Introduction to Generics in Go](https://blog.streamelements.com/an-introduction-to-generics-in-go-cc8cdae15ef2)
