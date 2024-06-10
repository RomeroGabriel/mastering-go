# Generics

Generics enables you to write `functions and data structures that work with different types while maintaining type safety`. You can define functions and structures that operate on various types without sacrificing the compiler's ability to catch type errors at compile time.

??? example "Basic Example using ints and floats"

    ```bash title="run command"
    $ go run src/generics/basic.go
    FindMax integer: 9
    FindMax float: 3.14
    ```
    ```go
    --8<-- "src/generics/basic.go"
    ```

## Generics Constraints

In Go, `type constraints are a way to specify requirements` for the type parameters used in generic functions or data structures. Constraints `ensure that the generic code can only be used with types that meet certain criteria`.

!!! note "Possibilities of Type Constraint"
    - Interfaces
    - any constriant
    - comparable constriant
    - Any other concrete type

??? example "Using Interface as Constraint"

    ```bash title="run command"
    $ go run src/generics/constraint.go
    FindMax integer: 9
    FindMax float: 3.14
    ```
    ```go
    --8<-- "src/generics/constraint.go"
    ```

### Comparable Constraint

In Go generics, you can use the `comparable constraint` to specify that a type should be comparable. This constraint ensures that the type can be used with comparison operators like `==` and `!=`.

??? example "Using comparable Constraint"

    ```bash title="run command"
    $ go run src/generics/comparable.go
    CheckNumbers:  false
    CheckNumbers:  true
    FindMax integer: -1
    FindMax integer: 3
    FindMax integer: -1
    ```
    ```go
    --8<-- "src/generics/comparable.go"
    ```

### Type Constraint- Specify Operators

In Go, generics can be used to `define a type element`, which is composed of one or more type terms within an interface. `Type elements specify which types can be assigned to a type parameter`. The list of concrete types is separated by `|`.

This way, it's possible to define `types that an operator supports`. For example, the modulus (%) operator is valid only for integers, so it's possible to create an integer-type element to be used with the modulus operator.

??? example "Integer Interface Constraint"

    ```go
    // All int possibilities
    type Integer interface{
        int | int8 | int16 | int32 | int64 |
        uint | uint8 | uint16 | uint32 | uint64 | uintptr
    }
    ```

!!! warning
    Interfaces with type elements are only valid as type constraints. It's an error to use them as the type for a variable, field, return, or parameter.

    Furthermore, both parameter value and return value should be valid for all the type terms in the type element. In the Integer interface example, it's not possible to return a T Integer value bigger than 1.000 since an 8-bit integer cannot hold this value.

### Tilde(~) - Closer Than Identical Types

The special tilde `~` symbol, often referred to as the `approximation constraint`, is used in Go generics to `indicate that a type should approximately match one of the specified types or their derived types`. `Without the tilde, the function or method constrained by generic types would only accept exact types declared in the constraints and not their derived types`. The tilde allows for a broader match, including derived types, making your code more flexible and accommodating related types.

??? example "Using Closer Type"

    ```bash title="run command"
    $ go run src/generics/approximation_constraint.go
    FindMax integer: 9
    FindMax float: 3.14
    ```
    ```go
    --8<-- "src/generics/approximation_constraint.go"
    ```

## Standard Library

Go has a [cmp package](https://pkg.go.dev/cmp), which defines the `Ordered` interface and functions.

## References

- [An Introduction to Generics in Go](https://blog.streamelements.com/an-introduction-to-generics-in-go-cc8cdae15ef2)
