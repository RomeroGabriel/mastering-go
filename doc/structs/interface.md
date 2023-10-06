# Interface

In Go, interfaces are a powerful and flexible way to `define sets of methods that types must implement`. `They allow you to write functions and methods that can work with various types, as long as those types satisfy the interface contract`. Interfaces support methods but `not properties or fields`.

To declare an interface, `you list the method signatures that types must implement`. `Any type that has methods matching these signatures implicitly implements the interface`. To implement an interface, `a type needs to provide method implementations that match the method signatures in the interface`. If a type has all the required methods, `it implicitly satisfies the interface`. Interfaces `promote polymorphism and decoupling of code`.

??? example

    ```bash title="run command"
    go run structs/interfaces.go
    ```
    ```go
    --8<-- "src/structs/interfaces.go"
    ```
    ```bash title="output"
    Area: 78.50
    Area: 24.00
    ```
