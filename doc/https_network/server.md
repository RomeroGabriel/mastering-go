# HTTP Server

You can create an HTTP server using the `net/http` package, allowing you to handle HTTP requests and responses.

!! example

    ```bash title="run command"
    go run src/http/basic_server.go
    curl localhost:8080
    curl localhost:8080/anonymous_func
    ```
    ```go
    --8<-- "src/http/basic_server.go"
    ```
    ```bash title="output"
    Server is running on http://localhost:8080
    Hello World!
    Hello World from anonymous_func!
    ```
