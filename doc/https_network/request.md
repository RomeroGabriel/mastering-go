# HTTP Request

## http.Get

Use `http.Get for simple GET requests` where you don't need to customize the headers extensively.

??? example

    ```bash title="run command"
    go run src/http/basic_http.go
    ```
    ```go
    --8<-- "src/http/basic_http.go"
    ```
    ```bash title="output"
    GOOGLE PAGE IN STRING
    ```

## http.Post

The `http.Post` function simplifies the process of sending POST requests. It allows you to `specify the URL, the content type, and the request body`. It's suitable for `simple POST requests` where you don't need to customize the headers extensively.

??? example

    ```bash title="run command"
    go run src/http/http_post.go
    ```
    ```go
    --8<-- "src/http/http_post.go"
    ```
    ```bash title="output"
    Status: 200 OK
    ```

## http.Client

The `http.Client object allows customization of the HTTP request behavior`, such as setting timeouts, defining transport properties, and managing the client's behavior.

??? example "http.Client with a 3-second Timeout"

    ```bash title="run command"
    go run src/http/http_client.go
    ```
    ```go
    --8<-- "src/http/http_client.go"
    ```
    ```bash title="output"
    {
        "userId": 1,
        "id": 1,
        "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
        "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
    }
    ```

## http.NewRequest

This function `creates a new HTTP request`. It allows `control over the request headers, methods, and request bodies`. You can set custom headers, perform more type requests (such as PUT and DELETE), and attach a request body. It's particularly useful for making custom or more complex HTTP requests.

??? example "Post using Client and NewRequest"

    ```bash title="run command"
    go run src/http/http_newRequest.go
    ```
    ```go
    --8<-- "src/http/http_newRequest.go"
    ```
    ```bash title="output"
    Status: 200 OK
    ```

## http.NewRequestWithContext

A function that creates a `new HTTP request with a specific context`. This function is particularly `useful when you need to set a deadline, cancel a request, or pass a request-specific value across API boundaries`.

??? example "Post using Client and NewRequestWithContext"

    ```bash title="run command"
    go run src/http/http_NewRequestWithContext.go
    ```
    ```go
    --8<-- "src/http/http_NewRequestWithContext.go"
    ```
    ```bash title="output"
    200 OK
    ```

## Examples

1. [Get Address - Basic API](../projects/get_br_zipcode.md#get-address---basic-api)
