# Variables and Constants

In Go, `variables and constants are strongly typed`. To declare a constant, you use the `const` keyword, and for variables, you use `var`. Go can also `automatically infer variable types`, providing various syntaxes for declaring variables and constants, as shown below.

??? example "Simple Vars Example"

    ```bash title="run command"
    $ go run src/fundamentals/vars/basic.go
    Starting basic variables
    Print A const:  I'm a constant
    Print B var:  false
    NOTICE: B was infer as false

    Different ways to declare a variable: 
    Full declaration:  A
    Short hand: B
    ```
    ```go
    --8<-- "src/fundamentals/vars/basic.go"
    ```

## Checking Default Infered Types

Go assigns a default zero value to any variable that is declared but not assigned a value. Having an explicit zero value makes code clearer and removes a source of bugs.

??? example "Default Infered Types"

    ```bash title="run command"
    $ go run src/fundamentals/vars/infer_types.go
    bool: false
    int: 0
    string: 
    float32: 0
    float64: 0
    ```
    ```go
    --8<-- "src/fundamentals/vars/infer_types.go"
    ```

## Explict Type Conversion

In Go, automatic type promotion between variables is not permitted. Whenever variable types do not align, explicit type conversion is necessary. Even when dealing with integers and floats of varying sizes, they must be converted to the same type before interaction. This approach ensures clarity regarding the intended data type.

Since all type conversions in Go are explicit, `you cannot treat another Go type as a boolean`. Unlike some languages where a nonzero number or a nonempty string can be interpreted as true, `Go doesn't support such implicit conversions`. No other type can be directly converted to a boolean, `either implicitly or explicitly`. If there's a need to convert from another data type to boolean, it's necessary to use one of the `comparison operators explicitly`.

??? example "Converting Types"

    ```bash title="run command"
    $ go run src/fundamentals/type_conversion.go
    Example: Converting Int and Float
    Sum1:  25.5
    Sum2:  25

    Example: Converting Int and Byte
    Sum3:  110
    Sum4:  110

    Right way to compare string
    Right way to compare int
    ```
    ```go
    --8<-- "src/fundamentals/type_conversion.go"
    ```

## Beyond Constants

In Go, `constants serve as identifiers for literals`, providing a means to assign meaningful names to values. These constants are restricted to holding values determinable by the compiler during compilation. Unlike some other languages, **`Go lacks built-in support for declaring values computed at runtime as immutable`**. Consequently, there are no immutable arrays, slices, maps, or structs in Go. Furthermore, there's no inherent mechanism to declare a field within a struct as immutable.

!!! note "Resume"
    1. Constant can only hold values determinable in compile-time
        - Numeric Literals
        - Boolean
        - Strings/Runes
    1. There are no ways to declare immutable values in runtime
    1. Arrays, slices, maps, or structs are not immutable
    1. No mechanism to declare struct fields as immutable

Constants in Go can be `typed` or `untyped`. An untyped constant behaves like a literal; it has no type but `defaults to a specific type` when no other type can be inferred. In contrast, a typed constant can only be assigned directly to a variable of its designated type. The decision to designate a constant as typed or untyped hinges on the rationale behind its declaration. Generally, `opting for an untyped constant provides greater flexibility`. However, `specific scenarios may necessitate the enforcement of a particular type for a constant`.

??? example "Typed vs Untyped"
    ```go
    const x = 10
    var y int = x
    var z float64 = x
    var d byte = x

    const typedX int = 10
    y = typedX
    ```

<!-- ## Using type

In Go the `type` keyword is used to `define custom data types`. These custom types can be based on existing types, making it a versatile feature in the language.

??? example "Using type"

    ```bash title="run command"
    go run variables/using_type.go
    ```
    ```go
    --8<-- "src/variables/using_type.go"
    ```
    ```bash title="output"
    custom_id value: e96d759e-dc59-4394-9feb-79c8bf4130c9
    ``` -->
