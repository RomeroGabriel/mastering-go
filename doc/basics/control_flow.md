# Control Flow

Listing the fundamental and commonly used control flow structures in Go.

## for

??? example

    ```bash title="run command"
    go run src/control_flow/for.go
    ```
    ```go title="src/control_flow/for.go"
    --8<-- "src/control_flow/for.go"
    ```
    ```bash title="output"
    Basic for type 1:  0
    Basic for type 1:  1
    Basic for type 1:  2
    Basic for type 1:  3
    Basic for type 1:  4
    *********
    Basic for type 2:  2
    Basic for type 2:  3
    Basic for type 2:  4
    *********
    arr for, index:  0 value  1
    arr for, index:  1 value  2
    arr for, index:  2 value  3
    arr for, index:  3 value  4
    arr for, index:  4 value  5
    *********
    map for, key:  a value  1
    map for, key:  b value  2
    map for, key:  c value  3
    ```

## switch

??? example

    ```bash title="run command"
    go run src/control_flow/switch.go
    ```
    ```go title="src/control_flow/switch.go"
    --8<-- "src/control_flow/switch.go"
    ```
    ```bash title="output"
    Today is Monday.
    ```
