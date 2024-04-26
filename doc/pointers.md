# Pointers

!!! info "Default Value"
    The zero value for a pointer is `nil`.

## Resuming Pointers and How They Work

Pointers are `variables that store memory addresses where a value is stored`. Variables are stored in one or more contiguous memory locations (addresses). A pointer type **can be based on any type**.

Different types of variables take up different amounts of memory, `but pointers always have the same amount of memory, which is 4-byte`.

## Operators & and *

### The Address Operator - &

The `&` is the address operator that `returns the address where the value is stored`.

!!! Warning "Pointers can be based on any type but..."
    It's not possible to use an & before a primitive literal (numbers, booleans, and strings) or a constant since they don't have a memory address (exist only at compile time). When it is necessary a pointer to a primitive type, declare a variable and point to it.

    ??? example
        ```go
        x := &Foo{}
        var y string
        z := &y
        ```

### The Indirection Operator - *

The `*` is the indirection operator that `returns the value pointed`. Using this operator is called **dereferencing**.

## Mutability with Pointers

!!! quote "[MIT’s course on Software Construction](http://web.mit.edu/6.031/www/fa20/classes/08-immutability/)"
    Using mutable objects is just fine if you are using them entirely locally within a method, and with only one reference to the object.

The ability to choose between value and pointer parameter types helps in the Go's lack of support for immutability types. In Go, using pointers indicates that a parameter is mutable.

!!! note
    [Go is a call-by-value language](functions/basic_func.md#call-by-value).

---

`They are used to reference values rather than holding the values themselves`. It allows you to `indirectly access and modify the value of the variable` it points to. Pointers are represented using the `*` symbol.

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
