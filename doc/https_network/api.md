# API

You can create an HTTP server using the `net/http` package, allowing you to handle HTTP requests and responses.

!!! example

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

## Examples

1. [Get Address - Basic API](../projects/get_br_zipcode.md#get-address---basic-api)
1. [Following the Standards for API](https://github.com/RomeroGabriel/go-api-standards)
