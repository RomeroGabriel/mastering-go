# Context

The `context` package provides a `mechanism for carrying deadlines, cancelation signals, and request-scoped` values across API boundaries and between processes. It is designed to `facilitate the management of long-running processes`, such as handling HTTP requests, canceling operations, or setting timeouts.

!!! info "context.Context"
    This is the primary type representing an incoming request or task. It allows the propagation of deadlines, cancelation signals, and key-value pairs across API boundaries.

!!! info "context.WithCancel"
    This function returns a derived Context and a cancel function. `Calling the cancel function signals that the work is done, and all the operations started from this Context should be canceled`.

!!! info "context.WithDeadline"
    This function returns a derived Context and a cancel function. `It is similar to context.WithTimeout, but it allows you to specify an absolute deadline instead of a relative one`.

!!! info "context.WithTimeout"
    This function returns a derived Context and a cancel function. `It is commonly used to set a timeout for an operation. After the specified duration, the cancel function is automatically called`.

!!! info "context.WithValue"
    This function returns a derived Context with the provided key-value pair. `It is used to pass request-scoped values across API boundaries`. Normally, is passed metadata.

??? example "HTTP Server using all context possible"

    ```bash title="run command"
    01. go run src/context/main.go
    02. curl localhost:8080/get-context
    03. curl localhost:8080/get-context AND CTRL + D to cancell the request before 5s
    04. curl localhost:8080/get-context-cancel
    05. curl localhost:8080/get-context-cancel?number=2
    06. curl localhost:8080/get-context-cancel?number=7
    07. curl localhost:8080/get-context-deadline
    08. curl localhost:8080/get-context-deadline?number=2
    09. curl localhost:8080/get-context-deadline?number=7
    10. curl localhost:8080/get-context-timeout
    11. curl localhost:8080/get-context-timeout?number=2
    12. curl localhost:8080/get-context-timeout?number=7
    13. curl localhost:8080/get-context-value
    14. curl localhost:8080/get-context-value?number=2
    15. curl localhost:8080/get-context-value?number=7
    ```
    ```go
    --8<-- "src/context/main.go"
    ```
    ```bash title="output"
    1. 2023/10/30 18:28:01 Server is listening on port 8080
    2. Context completed successfully.
    3. 2023/10/30 18:30:15 Request cancelled by client
    4. ContextCancel completed successfully
    5. ContextCancel completed successfully
    6. ContextCancel Task canceled or timed out. Number cannot be odd
    7. ContextDeadLineHandler completed successfully.
    8. ContextDeadLineHandler completed successfully.
    9. ContextDeadLineHandler Task canceled or timed out.
    10. ContextTimeout completed successfully.
    11. ContextTimeout completed successfully.
    12. ContextTimeout Task canceled or timed out.
    13. ContextValue completed successfully.
    13.1 Number passed 0
    14. ContextValue completed successfully.
    14.1 Number passed 2
    15. ContextValue canceled or timed out. Number cannot be odd.
    15.1 Number passed 7
    ```
