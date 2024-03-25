# Control Flow

## if

One notable distinction between if statements in Go and those in other programming languages is the capability to `declare variables scoped exclusively to the condition and both the if and else blocks`. This feature enables the creation of `variables accessible only within their intended scope`, enhancing code clarity and reducing potential side effects. Upon concluding the series of if/else statements, the variable's scope does not exist anymore, producing a `compilation error`.

??? example "Variable Spoce to If"

    ```go
    if n := rand.Intn(10); n == 0 {
        fmt.Println("That's too low")
    } else if n > 5 {
        fmt.Println("That's too big:", n)
    } else {
        fmt.Println("That's a good number:", n)
    }
    ```

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
