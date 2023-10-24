# File Server

The `FileServer` function in the `net/http` package efficiently `serves static files from a specified directory`. It simplifies the process of serving a directory by providing a handler that serves HTTP requests with the contents of the specified directory.

!!! example
    In this example, the FileServer function `serves files from the static directory`. The `http.Dir` function `specifies the directory from which to serve files`.

    ```bash title="run command"
    go run src/http/file_server.go
    curl localhost:8080
    ```
    ```go
    --8<-- "src/http/file_server.go"
    ```
    ```bash title="output"
    <p>Corinthians</p>
    ```
