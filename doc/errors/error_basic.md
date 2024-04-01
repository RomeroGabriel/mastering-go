# Error Handling

In Go, error handling differs significantly from languages that use try/catch blocks. `Go intentionally does not have try/catch mechanisms because it encourages a different approach to handling errors`. Instead, functions in Go often `return multiple values, where the last value is typically an error`.

??? example

    ```bash title="run command"
    go run function/return_error.go
    ```
    ```go
    --8<-- "src/function/return_error.go"
    ```
    ```bash title="output"
    Result: 5
    Error: division by zero
    ```
