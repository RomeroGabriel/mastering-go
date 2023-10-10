# Structs

In Go, `structs are composite data types used to group together variables (fields) under a single data structure`. They are similar to classes in other languages, but `Go doesn't have classes in the traditional sense`.

You define a struct using the `type` keyword, followed by the `name of the struct` and a set of field declarations within curly braces. To `create an instance` of a struct, you can use the `struct's name followed by curly braces and provide values for its fields`.

??? example

    ```bash title="run command"
    go run structs/basic.go
    ```
    ```go
    --8<-- "src/structs/basic.go"
    ```
    ```bash title="output"
    Create object person, from Person struct: {Gabriel 100}
    ```

## Anonymous Structs

Anonymous Structs are structs without a defined name, often used for `temporary data structures`.

??? example

    ```bash title="run command"
    go run structs/anonymous.go
    ```
    ```go
    --8<-- "src/structs/anonymous.go"
    ```
    ```bash title="output"
    Create object from an anonymous structs : {10 20}
    X : 10
    Y : 20
    ```

## Composing Structs

In Go, you can `embed one struct within another`, creating a composition of structs. This is `similar to inheritance` in other languages.

??? example

    ```bash title="run command"
    go run structs/composing.go
    ```
    ```go
    --8<-- "src/structs/composing.go"
    ```
    ```bash title="output"
    Person: {Gabriel 100 {Corinthians} {Knicks}}
    Person name: Gabriel, Clube: {Corinthians}
    Person name: Gabriel, Clube: Corinthians
    Person name: Gabriel, Clube: Corinthians
    Person name: Gabriel, Clube: {Knicks}
    Person name: Gabriel, Clube: Corinthians
    ```
??? tip "Composing x Property, and Accessing Properties"
    In the given example, it's evident that the `Person struct is composed of a Team` and `has a property named PersonTeam of type Team`. The key distinction here is that when using the abbreviation `person.TeamName`, it returns the value of the composition, not the property.
    In Go, `you can directly access the properties of a composed struct`.

## Struct Methods

In Go, `structs can have methods associated with them`. These methods are `functions that operate on instances of the struct`, and they enable you to define behavior specific to the struct type.

??? example

    ```bash title="run command"
    go run structs/methods.go
    ```
    ```go
    --8<-- "src/structs/methods.go"
    ```
    ```bash title="output"
    Area: 78.5
    ```

### Pointer Receivers

`Pointer receivers can modify the state of the struct directly`. Using a pointer receiver allows you to `change the state` of an instance directly. When you pass a struct to a function, you're essentially `passing a copy of the struct`. Using pointers, you can `pass a reference to the struct`, allowing the function to modify the original data.

??? example

    ```bash title="run command"
    go run structs/methods_pointer.go
    ```
    ```go
    --8<-- "src/structs/methods_pointer.go"
    ```
    ```bash title="output"
    Area: 78.5
    Circle change? {5}
    Area: 78.5
    Circle change? {100}
    ```

### Pointers Constructors

In Go, when you create a constructor function for a struct, `it typically returns a reference to the newly created struct`. With this technique, `any modification made to the reference returned by the constructor will affect all places that use that reference. This can be quite powerful, especially when managing resources like database connections`.

??? example

    ```bash title="run command"
    go run structs/constructor_pointer.go
    ```
    ```go
    --8<-- "src/structs/constructor_pointer.go"
    ```
    ```bash title="output"
    Query 1 executed **nice connection!*** true
    Query 2 executed **nice connection!*** true
    Query 3 executed **nice connection!*** true
    Closing the DB conncetion  **nice connection!***
    Query 1 executed **nice connection!*** false
    Query 2 executed **nice connection!*** false
    Query 3 executed **nice connection!*** false
    ```
