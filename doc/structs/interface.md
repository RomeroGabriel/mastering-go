# Interface

Interfaces in Go are the `only abstract type and they are implicit interfaces`. In Go, interfaces are a powerful and flexible way to `define sets of methods that types must implement`. They allow you to write functions and methods that can work with various types, as long as those types satisfy the interface contract.

!!! warning "And fields?"
    Interfaces support methods but `not properties or fields`.

To declare an interface, you list the method signatures that types must implement. Any type that has methods matching these signatures implicitly implements the interface. `If a type has all the required methods, it implicitly satisfies the interface`. Interfaces promote polymorphism and decoupling of code.

??? example "Declaring and Using Interfaces"

    ```bash title="run command"
    $ go run src/interface/interfaces.go
    Area: 78.50
    Area: 24.00
    ```
    ```go
    --8<-- "src/interface/interfaces.go"
    ```

Since interfaces are implicitly implemented, `a concrete type does not need to declare that it implements an interface`. This behavior enables type safety and decoupling. `The interface needs to be known only by the client`, nothing needs to be done in the implementation.

In Go, interfaces specify what `callers need`. The client code defines the interface to specify what functionality it requires, `nothing is declared on the implementation to indicate that it meets the interface`. In this way, it's possible to `add a new logic provider` in the future and provide `executable documentation` to ensure that any type passed into the client will match the client's need.

`Interfaces can be shared also`, where it is possible to define standard interfaces like *io.Reader* and *io.Writer*. `Anything that implements both interfaces, will function correctly` whether where it has been reading or writing the data.

??? example "Embedding Interfaces"

    ```go
    type Reader interface {
        Read(p []byte) (n int, err error)
    }

    type Closer interface {
        Close() (err error)
    }

    type ReadCloser interface {
        Reader
        Closer
    }
    ```

## Method Set for Pointer and Values Receivers

The methods defined by an interface are the method set of the interface. For struct, the `method set is different depending on the pointer and values receivers functions`. The method set of a `pointer instance` contains the methods defined with both pointer and values receivers, while the method set of a `value instance` contains only the methods with value receivers.

!!! note "In Summary"
    Method set for pointer instance: ALL struct's methods.

    Method set for value instance: ONLY value receiver methods.

??? example
    Taken from this [link](https://github.com/learning-go-book-2e/ch07/blob/main/sample_code/method_set/main.go)

    ```go
    package main

    import (
        "fmt"
        "time"
    )

    type Counter struct {
        total       int
        lastUpdated time.Time
    }

    func (c *Counter) Increment() {
        c.total++
        c.lastUpdated = time.Now()
    }

    func (c Counter) String() string {
        return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
    }

    type Incrementer interface {
        Increment()
    }

    func main() {
        var myStringer fmt.Stringer
        var myIncrementer Incrementer
        pointerCounter := &Counter{}
        valueCounter := Counter{}

        myStringer = pointerCounter    // ok
        myStringer = valueCounter      // ok
        myIncrementer = pointerCounter // ok
        myIncrementer = valueCounter   // compile-time error!

        fmt.Println(myStringer, myIncrementer)
    }
    ```

## How Interfaces are Implemented, Nil Interfaces and Comparable

Interfaces are implemented as a `struct with two pointer fields`. One pointer is for the `value` and one is for the `type` of the value. `Both type and value must be nil for an interface to be considered nil`.

??? example "Interface Structure Image"
    <figure markdown="span">
    ![The Relationship Between Interfaces and Reflection](https://blog.gopheracademy.com/postimages/advent-2018/interfaces-and-reflect/interface.svg)
    <figcaption>Interface Structure. Taken from <a href="https://blog.gopheracademy.com/advent-2018/interfaces-and-reflect/" target="_blank">The Relationship Between Interfaces and Reflection</a></figcaption>
    </figure>

Interfaces are `comparable` and two instances of an interface type are `equal only if their types and their values are equal`. However, if comparing two instances of an interface that the types aren't comparable like a slice, this will trigger a panic at runtime.

!!! warning "Comparing Interfaces"
    Be careful when using `==` or `!=` with interfaces or using an `interface as a map key`, since it's easy to generate panic in runtime if the type is not comparable.

??? example "Comparing Interfaces"

    ```bash title="run command"
    $ go run src/interface/interface_compare.go
    true
    false
    false
    panic: runtime error: comparing uncomparable type main.DoubleIntSlice

    goroutine 1 [running]:
    main.DoublerCompare(...)
        /home/romerin/Projects/mastering-go/src/interface_compare.go:10
    main.main()
        /home/romerin/Projects/mastering-go/src/interface_compare.go:35 +0x190
    exit status 2
    ```
    ```go
    --8<-- "src/interface/interface_compare.go"
    ```

## Performance

As discussed in [Less Pointers, Less Work for Garbage Collector](../pointers.md#less-pointers-less-work-for-garbage-collector), reducing heap allocations improves performance by reducing the amount of work for the garbage collector. Using structs avoids heap allocation, and `invoking a function with parameters of interface types triggers heap allocation for each interface parameter`.

!!! tip "About Performance"
    Figuring out the trade-off between better abstraction and better performance should be done over the life of your program.

## Passing Interface to Concrete Type

Go provides two ways to see if a variable of an interface type has a specific concrete type or if the concrete type implements another interface. This is not the same thing than type conversion, `conversions change a value to a new type, while assertions reveal the type of the value stored in the interface`. Type assertions work only with interfaces while type conversions work with concrete types and interfaces.

!!! warning
    These techniques should be used when necessary. It's better to be clear with inputs and returns values.

### Type Assertions

Type assertions provide one way to `check if a variable of an interface type has a specific concrete type or if the concrete type implements another interface`. Since Go is very careful about concrete types, even if two types share an underlying type, a type assertion must match the type of the value stored in the interface.

If a type assertion gets wrong, the code panics. To avoid that, it's possible to use the `comma ok idiom`. All types of assertions are `checked at runtime`, so any error that can happen will be at runtime with panic if the comma ok idiom is not used.

??? example "Using type assertion"

    ```bash title="run command"
    $ go run src/interface/type_assertion.go
    21
    Error
    ```
    ```go
    --8<-- "src/interface/type_assertion.go"
    ```

### Type Switches

Type switches provide another way to check if an interface could be one of `multiple possible types`. It's possible to assign the variable being checked to another variable valid only within the switch or shadowing the variable. The type of the new variable depends on which case matches.

??? example "Using type switch"

    ```go
    func doThings(i any) {
        switch j := i.(type) {
        case nil:
        // j is any
        case int:
        // j is of type int
        case MyInt:
        // j is of type MyInt
        case io.Reader:
        // j is of type io.Reader
        case string:
        // j is a string
        case bool, rune:
        // i is either a bool or rune, so j is of type any
        default:
        // no idea what i is, so j is of type any
        }
        }
    ```

### Checking Others Interfaces Implementation - Optional Interfaces

Using type assertions and type switches can be used to check `if the concrete type behind the interface also implements another interface`. This allow to specify optional interfaces.

??? example "Checking multiple Interfaces"

    ```go
    // WriterTo is the interface that wraps the WriteTo method.
    //
    // WriteTo writes data to w until there's no more data to write or
    // when an error occurs. The return value n is the number of bytes
    // written. Any error encountered during the write is also returned.
    //
    // The Copy function uses WriterTo if available.
    type WriterTo interface {
        WriteTo(w Writer) (n int64, err error)
    }
    func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
        // If the reader has a WriteTo method, use it to do the copy.
        // Avoids an allocation and a copy.
        if wt, ok := src.(WriterTo); ok {
            return wt.WriteTo(dst)
        }
        // Similarly, if the writer has a ReadFrom method, use it to do the copy.
        if rt, ok := dst.(ReaderFrom); ok {
            return rt.ReadFrom(src)
        }
        // function continues...
    }
    ```

## Interfaces and Function Params

Go encourages the usage of small interfaces but functions in Go are first-class concepts. Both approaches can be used since an interface of only one method could replace a parameter of a function type.

`If it's a simple function, then a parameter of a function type is a good choice`. However, if the function is likely to `depend on many other functions or other states that are not specified in its input parameters`, use an interface parameter with a defined function.

### Use Case Example - HTTP Handler

The HTTP Handler interface is defined in this way:

!!! example

    ```go
    type Handler interface{
        ServerHTTP(http.ResponseWriter, *http.Request)
    }
    ```

The `ServerHTTP` function allows `functions to implement the interface`. Any function with the same signature then the ServerHTTP function can be used as http.Handler. This allows the implementation of HTTP handlers using functions, methods, or closures.
