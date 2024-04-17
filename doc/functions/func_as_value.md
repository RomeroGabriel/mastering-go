# Deeper in Functions

## Variables of Type Function

Functions in Go are `values`. The function type is composed of the func keyword followed by the parameters and return values, creating the `function signature`. Functions with the same number and types of parameters and return values meet the function signature. `The default zero value for a function variable is nil`.

??? example "Using Func Type Variables"

    ```bash title="run command"
    $ go run src/function/value_func.go.go
    TEST1 A PARAM:  CALLING funcSignature
    RESULT:  10
    TEST2 b PARAM:  CALLING AGAIN funcSignature
    RESULT:  200
    ```
    ```go
    --8<-- "src/function/value_func.go.go"
    ```

## Function Type Declarations

The `type` keyword can be used to `define a function type`, just like is used to create a [struct](../structs/structs.md#structs). `Any function that meets the defined function signature type can be considered as a variable of that type`.

??? example "Creating a Function Type"

    ```go
    package main

    func main() {
        type FuncType func(int, int) int
        var mapFunc = map[string]FuncType{}
    }
    ```

## Anonymous Functions (Closures)

Go supports anonymous functions, `which can be assigned to variables or passed as arguments` to other functions. `They can capture and use variables from their surrounding scope`.

??? example

    ```bash title="run command"
    go run function/closures.go
    ```
    ```go
    --8<-- "src/function/closures.go"
    ```
    ```bash title="output"
    result1 6
    result1 7
    result1 8
    ```
