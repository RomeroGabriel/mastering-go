# Control Flow

## If Statement

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

In Go for is the **`only looping`** keyword in the language. The for loop can be used in four formats:

!!! note "A Complete For"
    The if statement has three parts, separated by semicolons: `initialization`, `comparison`, and `increment`. Go allows you to `leave out one or more of the three parts` of the for statement.

    ??? example "Ways to Create a Complete For"

        ```go
        for i := 0; i <= 9; i ++ {
            fmt.Println(i)
        }

        for i := 0; i <= 9; {
            fmt.Println(i)
            i++
        }
        
        i := 0
        for ; i <=9; i++ {
            fmt.Println(i)
        }
        ```

!!! note "A Condition-Only For"
    It's a for loop with just the `comparison` part.

    ??? example "Just Comparison for"

        ```go
        i := 1
        for i < 100 {
            fmt.Println(i)
            i++
        }
        ```

!!! note "An infinite For"
    A for loop without any part.

    ??? example "A infinite for"

        ```go
        for {
            fmt.Println("Hello")
        }
        ```

!!! note "For-Range"
    A for loop `iterating over elements in some of Go’s built-in types`. With this for, the first `variable is the position` in the data structure being iterated. Each time the for-range loop iterates with the variable data, it `copies the value` from the original variable to the value variable. `Modifying the value variable will not modify the value in the compound type`.

    ??? example "Iterating over built-in types"

        ```go
        data := []int{1, 2, 3, 4, 5, 6, 7}
        for i, v := range data {
            fmt.Println(i, v)
        }
        for _, v := range data {
            fmt.Println(v)
        }

        // Just the keys
        dicData := map[string]int{"Key1": 1, "Key2": 2}
        for k := range uniqueNames {
            fmt.Println(k)
        }

        for k, v := range m {
            fmt.Println(k, v)
        }

        strData := "hello"
        for i, r := range strData {
            fmt.Println(i, r, string(r))
        }
        ```

### break and continue keywords

In Go, `break` and `continue` keywords exist and behave similarly to other languages like C, and JavaScript. The `break` stops the loop, while the `continue` skips to the next iteration.

### Labeling For Statements

Labels in Go are used in `nested for loops and want to skip over an iteration of an outer loop`. Labels are indented to the same level as the surrounding function. Labels are always indented to the same level as the braces for the block.

??? example "Using Labels"

    ```bash title="run command"
    go run src/control_flow/labels.go
    Start sample:  hello
    h
    e
    l
    Start sample:  apple_π!
    a
    p
    p
    l
    ```
    ```go
    --8<-- "src/control_flow/labels.go"
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
