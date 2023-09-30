# Variables and Data Types

In Go, `variables and constants are strongly typed`. To declare a constant, you use the `const` keyword, and for variables, you use `var`.

Go can also `automatically infer variable types`, providing various syntaxes for declaring variables and constants, as shown below.

??? example "Simple vars example"

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

## Arrays

In Go, an array is a `fixed-size collection of elements of the same type`. Arrays are declared with a specific size, and that `size cannot be changed after the array is created`.

??? example

    ```bash title="run command"
    go run variables/arrays.go
    ```
    ```go
    --8<-- "src/variables/arrays.go"
    ```
    ```bash title="output"
    0
    [0 10 20]
    Value stored in 0 index is 0
    Value stored in 1 index is 10
    Value stored in 2 index is 20
    Value stored in 0 index is apple
    Value stored in 1 index is banana
    Value stored in 2 index is cherry
    ```

## Slices

In Go, a slice is a `more versatile and dynamic alternative to arrays`. Slices are like views into an underlying array, allowing you to work with a portion of an a`rray without specifying a fixed size`. Slices allow you to perform operations like appending elements, slicing (creating sub-slices), and growing dynamically.

Slices also have a `length and capacity` (the maximum number of elements it can hold without reallocation). Use the `len()` and `cap()` functions to get these values.

??? example

    ```bash title="run command"
    go run variables/slices.go
    ```
    ```go
    --8<-- "src/variables/slices.go"
    ```
    ```bash title="output"
    Full slice:  [10 20 30 40 50 60]
    First 3 elements:  [10 20 30]
    my_slice        = len=6 cap=6 [10 20 30 40 50 60]
    none_value      = len=0 cap=6 []
    first_four      = len=4 cap=6 [10 20 30 40]
    last_two        = len=4 cap=4 [30 40 50 60]
    my_slice with 70:  [10 20 30 40 50 60 70]
    my_slice        = len=7 cap=12 [10 20 30 40 50 60 70]
    ```

!!! warning
    Remember that slices are based on arrays, so they can change in size dynamically. `If the underlying array runs out of capacity, a new larger array is allocated, and the data is copied over`.
    When added 70 in the example aboce, notice that the `slice's capacity increases to 12 instead of 7`. This happens because Go, when it needs to increase capacity, `doubles the original size`. Doubling the capacity of the underlying array allows for more elements to be added in the future without frequent reallocations.
    `In the case of very large slices, this can consume more memory than necessary`. So, when you anticipate working with a large slice, it's a `good practice to initialize the slice with a size closer to the maximum you expect to use to minimize unnecessary memory consumption`.
