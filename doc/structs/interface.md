# Interface

Interfaces in Go are the `only abstract type and they are implicit interfaces`. In Go, interfaces are a powerful and flexible way to `define sets of methods that types must implement`. They allow you to write functions and methods that can work with various types, as long as those types satisfy the interface contract.

!!! warning "And fields?"
    Interfaces support methods but `not properties or fields`.

To declare an interface, you list the method signatures that types must implement. Any type that has methods matching these signatures implicitly implements the interface. `If a type has all the required methods, it implicitly satisfies the interface`. Interfaces promote polymorphism and decoupling of code.

??? example "Declaring and Using Interfaces"

    ```bash title="run command"
    $ go run src/interface/interfaces.go
    Area: 78.50
    Area: 24.00
    ```
    ```go
    --8<-- "src/interface/interfaces.go"
    ```

## Interfaces Are Type-Safe Duck Typing

Since interfaces are **implicitly** implemented, `a concrete type does not need to declare that it implements an interface`. This behavior enables **type safety** and **decoupling**.

Go interfaces have the `flexibility to grow` and change over time changing implementations and also `specifying what the code depends on`. This keeps the **dynamic** and **static** type behavior. `The interface needs to be known only by the client`, nothing needs to be done in the implementation.

In Go, interfaces specify what callers need. The client code defines the interface to specify what functionality it requires. `Interfaces can be shared also`, where it is possible to define standard interfaces like *io.Reader* and *io.Writer*. `Anything that implements both interfaces, will function correctly` whether where it has been reading or writing the data.

## Embedding Interfaces

!!! example

    ```go
    type Reader interface {
        Read(p []byte) (n int, err error)
    }

    type Closer interface {
        Close() (err error)
    }

    type ReadCloser interface {
        Reader
        Closer
    }
    ```

## Performance

As discussed in [Less Pointers, Less Work for Garbage Collector](../pointers.md#less-pointers-less-work-for-garbage-collector), reducing heap allocations improves performance by reducing the amount of work for the garbage collector. Using structs avoids heap allocation, and `invoking a function with parameters of interface types triggers heap allocation for each interface parameter`.

!!! tip "About Performance"
    Figuring out the trade-off between better abstraction and better performance should be done over the life of your program.

## How Interfaces are Implemented, Nil Interfaces and Comparable

Interfaces are implemented as a `struct with two pointer fields`. One pointer is for the `value` and one is for the `type` of the value. `Both type and value must be nil for an interface to be considered nil`.

Interfaces are `comparable` and two instances of an interface type are `equal only if their types and their values are equal`. However, if comparing two instances of an interface that the types aren't comparable like a slice, this will trigger a panic at runtime.

!!! warning "Comparing Interfaces"
    Be careful when using `==` or `!=` with interfaces or using an `interface as a map key`, since it's easy to generate panic in runtime if the type is not comparable.

??? example "Comparing Interfaces"

    ```bash title="run command"
    $ go run src/interface/interface_compare.go
    true
    false
    false
    panic: runtime error: comparing uncomparable type main.DoubleIntSlice

    goroutine 1 [running]:
    main.DoublerCompare(...)
        /home/romerin/Projects/mastering-go/src/interface_compare.go:10
    main.main()
        /home/romerin/Projects/mastering-go/src/interface_compare.go:35 +0x190
    exit status 2
    ```
    ```go
    --8<-- "src/interface/interface_compare.go"
    ```

## Type Assertions

Type assertions provide one way to `check if a variable of an interface type has a specific concrete type or if the concrete type implements another interface`. Since Go is very careful about concrete types, even if two types share an underlying type, a type assertion must match the type of the value stored in the interface.

If a type assertion gets wrong, the code panics. To avoid that, it's possible to use the `comma ok idiom`. All types of assertions are `checked at runtime`, so any error that can happen will be at runtime with panic if the comma ok idiom is not used.

??? example "Using type assertion"

    ```bash title="run command"
    $ go run src/interface/type_assertion.go
    21
    Error
    ```
    ```go
    --8<-- "src/interface/type_assertion.go"
    ```

### Type Assertion vs Type Conversion

Type assertion and type conversion are **not the same thing**. `Conversions change a value to a new type, while assertions reveal the type of the value stored in the interface`. Type assertions work only with interfaces while type conversions work with concrete types and interfaces.

## Type Switches

Type switches provide another way to check if an interface could be one of `multiple possible types`. It's possible to assign the variable being checked to another variable valid only within the switch or shadowing the variable. The type of the new variable depends on which case matches.

??? example "Using type switch"

    ```go
    func doThings(i any) {
        switch j := i.(type) {
        case nil:
        // j is any
        case int:
        // j is of type int
        case MyInt:
        // j is of type MyInt
        case io.Reader:
        // j is of type io.Reader
        case string:
        // j is a string
        case bool, rune:
        // i is either a bool or rune, so j is of type any
        default:
        // no idea what i is, so j is of type any
        }
        }
    ```
