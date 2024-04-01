# Basic of Functions

## Definition and Arguments

The `func` keyword introduces a function definition, followed by the function name and a list of formal parameters enclosed in parentheses. Functions can take zero or more parameters. Parameters are declared with their type. `When function parameters share the same type, you can declare the type once for all of them`.

When you pass `parameters to a function, they are typically passed by copy`. This means that a copy of the value is made, and `any changes you make to the parameter within the function do not affect the original value outside of the function`. To change the parameters passed, use [pointers](../pointers.md#pointers) and [pointers receivers](../structs/structs.md#pointer-receivers).

??? example

    ```bash title="run command"
    go run function/basic.go
    ```
    ```go
    --8<-- "src/function/basic.go"
    ```
    ```bash title="output"
    Double of 6: 12
    sum_result is: 10
    ```

## Parameters

### Call by Value

a

### Named and Optional Parameters

a

### Variadic Functions

Go allows you to define `functions that can accept a variable number of arguments`. You use the `...` notation before the type of the last parameter

??? example

    ```bash title="run command"
    go run function/func_variadic.go
    ```
    ```go
    --8<-- "src/function/func_variadic.go"
    ```
    ```bash title="output"
    Result 45
    ```

### Functions as Parameters

## Return

Functions in Go can indeed `return multiple values`. When a function is called, it can produce multiple results. `Functions can declare named return values`, which act as variables. This can make the code more readable by specifying what the function is returning.

??? example

    ```bash title="run command"
    go run function/return.go
    ```
    ```go
    --8<-- "src/function/return.go"
    ```
    ```bash title="output"
    result is: 3
    result1 is: 4
    result1 is even!!!
    ```

### Multiple Return Values

### Named Return Values

### Returning Functions