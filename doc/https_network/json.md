# JSON

Working with JSON in Go is straightforward and can be achieved using the `encoding/json` package, which provides functionalities to encode Go data structures into JSON and decode JSON into Go data structures.

!!! quote ""
    Nice to know [Struts Tags](../structs/structs.md#structs-tags)

## Struct to JSON

Using `json.Marshal` is the simple way. The `json.NewEncoder` is used to create a new `JSON encoder that writes to an io.Writer interface, such as a file or network connection`. This allows you to easily encode Go data structures to JSON and write them to an external destination.

??? example

    ```bash title="run command"
    go run src/http/struct_to_json.go
    ```
    ```go
    --8<-- "src/http/struct_to_json.go"
    ```
    ```bash title="output"
    Using Marshal
    JSON Data: {"name":"Gabriel","age":25}
    ***
    Using NewEncoder
    ```

## JSON to Struct

Using `json.Unmarshal` is the simple way. The `json.NewDecoder` is used to create a new `JSON decoder that reads from an io.Reader interface, such as a file or network connection`. This enables you to easily decode JSON data from an external source into Go data structures.

??? example

    ```bash title="run command"
    go run src/http/json_to_struct.go
    ```
    ```go
    --8<-- "src/http/struct_to_json.go"
    ```
    ```bash title="output"
    Using Unmarshal
    Person: {Gabriel 25}
    ***
    Using NewDecoder
    Person2: {Gabriel 25}
    ```
