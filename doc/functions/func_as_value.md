# Deeper in Functions

## Variables of Type Function

Functions in Go are `values`. The function type is composed of the func keyword followed by the parameters and return values, creating the `function signature`. Functions with the same number and types of parameters and return values meet the function signature. `The default zero value for a function variable is nil`.

??? example "Using Func Type Variables"

    ```bash title="run command"
    $ go run src/function/value_func.go
    TEST1 A PARAM:  CALLING funcSignature
    RESULT:  10
    TEST2 b PARAM:  CALLING AGAIN funcSignature
    RESULT:  200
    ```
    ```go
    --8<-- "src/function/value_func.go"
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

## Anonymous Functions

Go supports anonymous functions, `which can be assigned to variables or passed as arguments` to other functions. Inner functions are **anonymous**, they `don't have a name`.

You can call an anonymous function immediately after the creation. This is common in two situations: [defer](defer.md#defer) statements and creating goroutines.

??? example "Creating and Calling Two Anonymous Functions"

    ```bash title="run command"
    $ go run src/function/anonymous.go
    Result1:  6
    Result1:  7
    Result1:  8
    Result2:  20
    ```
    ```go
    --8<-- "src/function/anonymous.go"
    ```

## Closures

`Closures are functions declared inside function`. This is a computer science term that means that `closures can access and modify variables declared in the outer function`. They can capture and use variables from their surrounding scope.

!!! danger "Shadow Variable"
    Using `:=` instead of `=` inside the closure creates a new variable inside the function scope.

Closures allow to limit of a function's scope. `When a function is called from only one other function, an inner function can be used`, reducing declarations at the package level. `Repeated code logic within a function can be encapsulated into a closure`.

??? example "Closures Example"

    ```bash title="run command"
    $ go run src/function/closures.go
    Inner Function -> closure1:  10
    closure1 ENDING:  30

    Inner Function -> closure2:  10
    Inner Function SHADOW DATA -> closure2:  30
    closure2 ENDING:  10
    ```
    ```go
    --8<-- "src/function/closures.go"
    ```
