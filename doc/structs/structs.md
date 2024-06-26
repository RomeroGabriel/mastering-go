# Structs

In Go, `structs are composite data types used to group together variables (fields) under a single data structure`. They are similar to classes in other languages, but `Go doesn't have classes in the traditional sense`. You define a struct using the `type` keyword, followed by the `name of the struct` and a set of field declarations within curly braces. To `create an instance` of a struct, you can use the `struct's name followed by curly braces and provide values for its fields`.

After declaring a struct type, you can define variables of that type either using the `var` keyword or employing the `:=` syntax. In the absence of a value assigned to the variable, it assumes the `zero value corresponding to the struct type`. A zero-value struct entails `each field being set to its respective zero value`. There is no difference between assigning an empty struct literal and not assigning a value at all. Both scenarios result in initializing all fields within the struct to their zero values.

??? example "Initialization and Zero-Values for Structs"

    ```bash title="run command"
    $ go run src/structs/basic.go
    { 0}
    { 0}
    Both empty structs are equal?  true
    Create object person, from Person struct: {Gabriel 100}
    Created using field order: {Gabriel2 200}
    ```
    ```go
    --8<-- "src/structs/basic.go"
    ```

## Anonymous Structs

Anonymous Structs are structs without a defined name, often used for `temporary data structures`. Anonymous structs prove useful in a two typical scenarios. Firstly, when `translating external` data into a struct or vice versa, such as dealing with JSON or Protocol Buffers. This process is commonly referred to as unmarshaling and marshaling data, respectively.

??? example "Using Anonymous Structs"

    ```bash title="run command"
    $ go run src/structs/anonymous.go
    Vector: {10 20}
    X : 10
    Y : 20
    Person: {Yuri Alberto 123}
    Name : Yuri Alberto
    Age : 123
    ```
    ```go
    --8<-- "src/structs/anonymous.go"
    ```

## Compare Structs

The comparability of a struct depends on its fields. `Structs composed entirely of comparable types are comparable`, while those containing slice, map, function, and channel fields are not.

Go lacks a mechanism that allow overriding to redefine equality, thereby enabling the use of == and != with incomparable structs. However, developers have the option to craft `custom functions for comparing structs as an alternative approach`.

## Composing Structs

**Go doesn't support inheritance** since the language encourages code reuse by `composition` and `promotion`. Structs can have `embedded fields`, that's fields without names assigned. Any fields or methods declared on an `embedded field are promoted to the containing struct and can be invoked directly`. Any type can be embedded.

For fields or methods with the `same name between structs`, it's necessary to use the embedded field's name.

??? example "Compose Structs and Consuming"

    ```bash title="run command"
    $ go run src/structs/composing.go
    Person Name: Gabriel, Age: 20, and Team: Corinthians
    Team Name: Corinthians, Location: São Paulo
    São Paulo
    ```
    ```go
    --8<-- "src/structs/composing.go"
    ```

!!! warning
    Embedding is not inheritance, it's a different feature that Go supports.

## Struct Methods

In Go, `structs can have methods associated with them`. These methods are `functions that operate on instances of the struct`, and they enable you to define behavior specific to the struct type.

Methods can be `defined just at the package block level`, while functions can be defined inside any block. Methods names `cannot be overloaded` also, and the methods must be declared in the `same package as their associated type`, it's not possible to add methods to types you don't control.

??? example "Simple Example of Methods"

    ```bash title="run command"
    $ go run structs/methods.go
    Area: 78.5
    ```
    ```go
    --8<-- "src/structs/methods.go"
    ```

### Pointer Receivers vs Value Receivers

Check out: [Use Pointers with Wisdom](../pointers.md#use-pointers-with-wisdom).

`Pointer receivers can modify the state of the struct directly`. Using a pointer receiver allows you to `change the state` of an instance directly. When using value receivers, you're essentially `passing a copy of the struct`. Using pointers, you can `pass a reference to the struct`, allowing the function to modify the original data.

??? example "Using Pointer and Value Receivers"

    ```bash title="run command"
    $ go run structs/methods_pointer.go
    Area: 78.5
    Circle change? {5}
    Area: 78.5
    Circle change? {100}
    ```
    ```go
    --8<-- "src/structs/methods_pointer.go"
    ```

### Pointers Constructors

In Go, when you create a constructor function for a struct, `it typically returns a reference to the newly created struct`. With this technique, `any modification made to the reference returned by the constructor will affect all places that use that reference. This can be quite powerful, especially when managing resources like database connections`.

??? example "A Simple Pointer Constructor"

    ```bash title="run command"
    $ go run structs/constructor_pointer.go
    Query 1 executed **nice connection!*** true
    Query 2 executed **nice connection!*** true
    Query 3 executed **nice connection!*** true
    Closing the DB conncetion  **nice connection!***
    Query 1 executed **nice connection!*** true
    Query 2 executed **nice connection!*** false
    Query 3 executed **nice connection!*** false
    ```
    ```go
    --8<-- "src/structs/constructor_pointer.go"
    ```

## Structs Tags

Struct tags are `annotations that can be added to the fields of a struct to provide additional information or metadata about the field`. Struct tags are `typically used to describe how the struct fields should be encoded or decoded to and from other formats, such as JSON, XML, or other data serialization formats`. They play a crucial role in mapping Go data structures to external formats, enabling seamless data interchange between different systems or languages.

??? example
    1. Struct tags are added to each field to `specify the corresponding JSON field names`. For example, the json:"name" tag indicates that the Name field should be encoded as "name" in JSON.
    1. The `omitempty` option in the json:"address,omitempty" tag indicates that the Address `field should be omitted from the JSON output if it is empty`.

    ```bash title="run command"
    go run src/structs/tags.go
    ```
    ```go
    --8<-- "src/structs/tags.go"
    ```
    ```bash title="output"
    JSON Data: {"name":"Gabriel","age":25}
    ```
