# Variables

In Go, `variables and constants are strongly typed`. To declare a constant, you use the `const` keyword, and for variables, you use `var`.

Go can also `automatically infer variable types`, providing various syntaxes for declaring variables and constants, as shown below.

!!! example "Simple package example"

    ```bash title="run command"
    go run src/variables/*
    ```
    ```go
    --8<-- "src/variables/basic.go"
    ```
    ```bash title="output"
    Starting basic variables
    Print A const: I'm a constant
    Print B var: false
    NOTICE: B was infer as false

    Different ways to declare a variable:
    Full declaration: A
    Short hand: B
    ```

??? example "Checking Default Infered Types"

    ```bash title="run command"
    go run variables/infer_types.go 
    ```
    ```go
    --8<-- "src/variables/infer_types.go"
    ```
    ```bash title="output"
    bool false
    int 0
    string 
    float32 +0.000000e+000
    float64 +0.000000e+000
    ```

## Using type

In Go the `type` keyword is used to `define custom data types`. These custom types can be based on existing types, making it a versatile feature in the language.

!!! example "Using type"

    ```bash title="run command"
    go run variables/using_type.go
    ```
    ```go
    --8<-- "src/variables/using_type.go"
    ```
    ```bash title="output"
    custom_id value: e96d759e-dc59-4394-9feb-79c8bf4130c9
    ```
