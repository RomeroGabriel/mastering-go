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

Check out: [Use Pointers with Wisdom](../pointers.md#use-pointers-with-wisdom).

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
    Query 1 executed **nice connection!*** true
    Query 2 executed **nice connection!*** false
    Query 3 executed **nice connection!*** false
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
