# Error Handling

In Go, error handling differs significantly from languages that use try/catch blocks. Go intentionally **does not have try/catch** mechanisms because it encourages a different approach to handling errors.

Functions in Go often `return multiple values, where the last value is typically an error`. This is a convention, but it is such a strong convention that it should never be broken. To `create` a new error, it is used the `New` function in the `errors` package passing a string. By convention, the `error message should not be capitalized nor should they end with punctuation or a newline`.

To check and handle errors, it's used a simple `if` statement checking the error variable is non-nil.

??? example "A Simple Returned Error"

    ```bash
    $ go run src/function/return_error.go
    Result: 5
    Error: division by zero
    ```
    ```go
    --8<-- "src/function/return_error.go"
    ```

## fmt.Errorf

An alternative way to create an error is using the `fmt.Errorf` function. This function lets you `include runtime information` in the error message by using the fmt.Printf verbs to format the string message.

??? example "Using fmt.Errorf"

    ```go
    func divide(a, b float64) (float64, error) {
        if b == 0 {
            return 0, fmt.Errorf("division by zero %d", b)
        }
        return a / b, nil
    }
    ```

## Error Interface

The built-in interface error defines a single method `Error`. Anything that implements this interface is considered an error.

!!! note "Error interface"

    ```go
    type error interface{
        Error() string
    }
    ```
