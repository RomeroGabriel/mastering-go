# Error Handling

In Go, error handling differs significantly from languages that use try/catch blocks. Go intentionally **does not have try/catch** mechanisms because it encourages a different approach to handling errors.

Functions in Go often `return multiple values, where the last value is typically an error`. This is a convention, but it is such a strong convention that it should never be broken. To `create` a new error, it is used the `New` function in the `errors` package passing a string. By convention, the `error message should not be capitalized nor should they end with punctuation or a newline`.

To check and handle errors, it's used a simple `if` statement checking the error variable is non-nil.

??? example "A Simple Returned Error"

    ```bash
    $ go run src/errors/return_error.go
    Result: 5
    Error: division by zero
    ```
    ```go
    --8<-- "src/errors/return_error.go"
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

## Sentinel Erros

Sentinal errors are errors that `signal that processing cannot continue because of a problem with the current state`. A specific value is used to signify the error. Sentinal errors are the few variables `declared at the package level`, with their names starting with `Err`, and should be `treated as read-only`, even Go didn't guarantee the immutability for those variables, but it is a bad practice change their value.

Examples of sentinel errors can be found in many standard libraries, like the `archive/zip` library for processing ZIP files. The `ErrFormat` variable is a sentinel error, indicating that the data passed is invalid. The `context.Canceled` is also a sentinel error.

??? example "Using th zip.ErrFormat"

    ```go
    func main() {
        data := []byte("fake data")
        fakeZipFile := bytes.NewReader(data)
        _, err := zip.NewReader(fakeZipFile, int64(len(data)))
        if err == zip.ErrFormat {
            fmt.Println("Error in zip file")
        }
    }
    ```

Sentinel errors are `part of the public API` for a module, bringing responsibility to be available and compatible with future releases. It's preferable to `reuse one of the existing ones` in the standard library or to define an error type that includes information about the error. `Use sentinel errors when there is an error condition that indicates a specific state has been reached in the application where no further processing is possible`.

## User Defined Errors

User-defined errors can be created by `implementing the error interface`. By using this approach, it is possible to `avoid string comparisons to determine error causes`.

Always use `error` as the return type for the error return value in the function signature. This allows `return different types of errors` from the function and the function's callers can choose not to depend on the specific error type.

??? example "A Simple Returned Error"

    ```bash
    $ go run src/errors/user-defined.go
    Result: 5
    Error: division by zero
    ```
    ```go
    --8<-- "src/errors/user-defined.go"
    ```
