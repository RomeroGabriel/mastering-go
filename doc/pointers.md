# Pointers

In Go, pointers are `variables that store memory addresses`. `They are used to reference values rather than holding the values themselves`. It allows you to `indirectly access and modify the value of the variable` it points to. Pointers are represented using the `*` symbol.

To `access the value that a pointer points to`, you can use the `*` symbol again. This is called `dereferencing the pointer`. The `&` operator is `used to obtain the memory address of a variable`. This memory address is also known as a pointer.

!!! example

    ```bash title="run command"
    go run src/structs/pointers.go
    ```
    ```go
    --8<-- "src/structs/pointers.go"
    ```
    ```bash title="output"
    Value of x:  42
    Value that ptr point to:  42
    Memory address of ptr:  0xc0000120f0
    Updated value of x:  100
    ```
