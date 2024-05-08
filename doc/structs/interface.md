# Interface

Interfaces in Go are the `only abstract type and they are implicit interfaces`. In Go, interfaces are a powerful and flexible way to `define sets of methods that types must implement`. They allow you to write functions and methods that can work with various types, as long as those types satisfy the interface contract.

!!! warning "And fields?"
    Interfaces support methods but `not properties or fields`.

To declare an interface, you list the method signatures that types must implement. Any type that has methods matching these signatures implicitly implements the interface. `If a type has all the required methods, it implicitly satisfies the interface`. Interfaces promote polymorphism and decoupling of code.

??? example "Declaring and Using Interfaces"

    ```bash title="run command"
    $ go run structs/interfaces.go
    Area: 78.50
    Area: 24.00
    ```
    ```go
    --8<-- "src/structs/interfaces.go"
    ```

## Interfaces Are Type-Safe Duck Typing

Since interfaces are **implicitly** implemented, `a concrete type does not need to declare that it implements an interface`. This behavior enables **type safety** and **decoupling**.

Go interfaces have the `flexibility to grow` and change over time changing implementations and also `specifying what the code depends on`. This keeps the **dynamic** and **static** type behavior. `The interface needs to be known only by the client`, nothing needs to be done in the implementation.

In Go, interfaces specify what callers need. The client code defines the interface to specify what functionality it requires. `Interfaces can be shared also`, where it is possible to define standard interfaces like *io.Reader* and *io.Writer*. `Anything that implements both interfaces, will function correctly` whether where it has been reading or writing the data.

## Embedding Interfaces

