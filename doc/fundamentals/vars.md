# Variables and Data Types

In Go, `variables and constants are strongly typed`. To declare a constant, you use the `const` keyword, and for variables, you use `var`. Go can also `automatically infer variable types`, providing various syntaxes for declaring variables and constants, as shown below.

!!! example "Simple vars example"

    ```bash title="run command"
    $ go run src/fundamentals/vars/basic.go
    Starting basic variables
    Print A const:  I'm a constant
    Print B var:  false
    NOTICE: B was infer as false

    Different ways to declare a variable: 
    Full declaration:  A
    Short hand: B
    ```
    ```go
    --8<-- "src/fundamentals/vars/basic.go"
    ```

!!! example "Checking Default Infered Types"

    ```bash title="run command"
    $ go run src/fundamentals/vars/infer_types.go
    bool: false
    int: 0
    string: 
    float32: 0
    float64: 0
    ```
    ```go
    --8<-- "src/fundamentals/vars/infer_types.go"
    ```

## Maps

Maps in Go are a versatile data structure used to store `key-value pairs`. They are `similar to dictionaries or hash tables` in other programming languages. A key feature of maps is their ability to provide fast and efficient lookups for values based on their associated keys.

Remember that maps are `reference types`, so when you pass a map to a function or assign it to another variable, you're working with a reference to the same underlying data structure.

??? example

    ```bash title="run command"
    go run variables/maps.go
    ```
    ```go
    --8<-- "src/variables/maps.go"
    ```
    ```bash title="output"
    map[]
    map[]
    Init map_works_salary: map[Cassio:3000 Gabriel:100 Renato Augusto:3000]
    Gabriel salary: 100
    Yuri Alberto salary: 3000
    After delete Gabriel: map[Cassio:3000 Renato Augusto:3000 Yuri Alberto:3000]

    Player: Renato Augusto, Salaray: 3000
    Player: Cassio, Salaray: 3000
    Player: Yuri Alberto, Salaray: 3000
    ```

## Using type

In Go the `type` keyword is used to `define custom data types`. These custom types can be based on existing types, making it a versatile feature in the language.

??? example "Using type"

    ```bash title="run command"
    go run variables/using_type.go
    ```
    ```go
    --8<-- "src/variables/using_type.go"
    ```
    ```bash title="output"
    custom_id value: e96d759e-dc59-4394-9feb-79c8bf4130c9
    ```
