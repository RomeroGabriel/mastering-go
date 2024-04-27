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

### Check Out

??? example

    ```bash title="run command"
    $ go run src/structs/pointers.go
    Value of x:  42
    Value that ptr point to:  42
    Memory address of ptr:  0xc0000120f0
    Updated value of x:  100
    ```
    ```go
    --8<-- "src/structs/pointers.go"
    ```

## Mutability with Pointers

!!! quote "[MIT’s course on Software Construction](http://web.mit.edu/6.031/www/fa20/classes/08-immutability/)"
    Using mutable objects is just fine if you are using them entirely locally within a method, and with only one reference to the object.

The ability to choose between value and pointer parameter types helps in the Go's lack of support for immutability types. In Go, using pointers indicates that a parameter is mutable.

!!! note
    [Go is a call-by-value language](functions/basic_func.md#call-by-value).

## Use Pointers with Wisdom

Pointers can make it harder to understand data flow and can create extra work for the garbage collector. Instead of receiving a struct pointer to populate the struct, `have the function instantiate, and return the struct`.

??? example "Creating a Struct Obj"
    ```go
    func BadMakeFoo(f *Foo) error {
        f.Field1 = "Hi"
        f.Field2 = 20
        return nil
    }
    func GoodMakeFoo() (Foo, error) {
        f := Foo{
            Field1 = "Hi",
            Field2 = 20,
        }
        return f, nil
    }
    ```

`The only time you should use pointer parameters to modify a variable is when the function expects an interface`, such as the [JSON Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) function with an any parameter, where any must be a pointer.

Use a pointer type as a return type only if there is a state within the data that needs to be modified.

## Pointers Performance

`The time to pass a pointer into a function is constant for all data structures` since the size of a pointer is the same for all data types. Passing a value into a function takes longer as the data gets larger.

Returning is a little different. For data structures smaller than 10 Mb, returning pointers is slower than returning a value. For data structures larger than that, the performance advantage flips.

Besides these behaviors, these are very short times. `For the vast majority of cases, the difference won't affect the program's performance. In case of passing Mb of data between functions, consider using a pointer even if the data is meant to be immutable`.
