# Functions as Value

## Declare Function as Type

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
