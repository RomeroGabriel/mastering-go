# Basic of Functions

!!! example "Basic Syntax"

    ```go
    func funcName(param1, param2 int, param3 string) (int, error) {
        ....
    }
    ```

## Parameters

Functions can take zero or more parameters and are declared with their type. `When function parameters share the same type, you can declare the type once for all of them`, as param1 and param2 in the example above show.

### Call by Value

When you pass `parameters to a function, they are typically passed by copy`. This means that a copy of the value is made, and `any changes you make to the parameter within the function do not affect the original value outside of the function`. To change the parameters passed, use [pointers](../pointers.md#pointers) and [pointers receivers](../structs/structs.md#pointer-receivers).

??? example "Checking Mutations on Params"

    ```bash title="run command"
    $ go run src/function/basic.go
    Memory allocated for num1: 0xc0000120c8
        Memory allocated for num (inside print_double): 0xc0000120f0
        Double:  10
    After call print_double:  5
    ```
    ```go
    --8<-- "src/function/basic.go"
    ```

### Named and Optional Parameters

Go `doesn't have support` for named and optional parameters. To simulate named and optional parameters, it's necessary to `create a struct with the desired parameters` and pass the struct to the function.

### Variadic Param

Go allows you to define `functions that can accept a variable number of arguments`. You use the `...` notation before the type of the `last parameter` to define variadic parameter. The veriadic parameter is a `slice of the specified type` and can be used just like any other slice.

!!! danger ""
    The variadic parameter `must be the last or the only parameter` in the input parameter list.

??? example "Receiving a Variadic Parameter"

    ```bash title="run command"
    $ go run src/function/func_variadic.go
    numbers param:  [1 2 3 4 5 6 7 8 9]
    Sum Result : 45
    ```
    ```go
    --8<-- "src/function/func_variadic.go"
    ```

## Return

Functions can `return zero or more values`.

### Multiple Return Values

Functions in Go can indeed `return multiple values`. If a function return multiple values, `the function must return all values`.

!!! tip "Ignoring return value"
    The `_` is used to ignore a return value when the value is not necessary.

??? example "Multiple Return"

    ```bash title="run command"
    $ go run src/function/return.go
    result is: 3
    result is: 4
    result is even!!!
    ```
    ```go
    --8<-- "src/function/return.go"
    ```

### Named Return Values

`Functions can declare named return values, which act as variables`. This can make the code more readable by specifying what the function is returning. When return values are named, the variables are predeclared that can be used within the function to hold the return values.

However, besides that return values can be created, `it is not necessary to use them in the return statement`. Also, return values named can cause shadowing variables. These two scenarios can lead to bugs in the code.

!!! tip "Nameless and Named Returns"
    The `_` can be used for nameless return values, when just some of the return values needs to be named.

!!! Danger "Blank Returns"
    Using blank returns makes some changes to the function. When there’s invalid input, the `function returns immediately`. In this case, their zero values are returned. This makes it harder to understand data flow.

??? example "Named Returns"

    ```bash title="run command"
    $ go run src/function/return_named.go
    result:  6  is_even:  true
    result:  100  is_even:  false
    result:  6  is_even:  true
    ```
    ```go
    --8<-- "src/function/return_named.go"
    ```
