# Maps

Maps in Go are a data structure used to store **`key-value pairs`** that implements a *hash table* data structure. The map type is written as **`map[keyType]valueType`**. **`The zero value for a map is nil`**.

??? info "Hash Table"
    Hash tables are a type of data structure in which the address/ index value of the data element is generated from a hash function. This enables very fast data access as the index value behaves as a key for the data value.

    Go doesn’t require (or even allow) you to define your own hash algorithm or equality definition. Instead, the Go runtime that’s compiled into every Go program has code that implements hash algorithms for all types that are allowed to be keys.

The `key can be of any can be any comparable type`, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays. The `value may be any type at all`, including another map. `Maps are not comparable`. You have the option to verify if they are `equal to nil`, but you `cannot compare two maps` possess identical keys and values by employing `==`, or difference `!=`.

!!! warning "Key types not supported"
    Slices and Maps cannot be used as map keys, because equality is not defined on them.

Maps are **`reference types`**, that hold references to an underlying **hash table**. **`If you pass a map to a function or assign it to another variable, both refer to the same data structure`**. Maps are reference type like `pointers` or `slices` to initialize a map, use the built in make function. Just like slices, `maps automatically grow as you add key-value pairs to them`.

## Declaration

You have the option to utilize a `var` declaration to initialize a `map variable set to its zero value`, which in this case is `nil`. A nil map is characterized by a length of 0 and `attempting to write to a nil map will result in a panic`. Alternatively, using `:=` allows you to create an empty map. If you possess an estimate of the number of key-value pairs intended for the map but lack precise values, you can employ `make` to create a map with a `default size`. Maps instantiated with make retain a length of 0 initially and `can expand beyond the initially specified size`.

??? example "Initialize Maps"

    ```bash title="run command"
    $ go run src/fundamentals/composite/maps.go
    Not initalize map:  map[]
    After add elements:  map[key1:1 key2:2 key3:3]
    Value from key1:  1
    When key doesn't exist, default value is returned :  0
    newMap:  map[Key1:12]
    fullFillMap:  map[Key1:12 Key2:22 Key3:32]
    ```
    ```go
    --8<-- "src/fundamentals/composite/maps.go"
    ```

## The Ok Idiom

In Go, the comma-ok idiom comes into play when you need to `distinguish between retrieving a value and receiving the zero value, as well as identifying a key not in the map`. Instead of assigning the result of a map read to a single variable, the comma-ok idiom involves assigning the outcomes of a map read to two variables. The first variable captures the `value associated with the key`, while the second variable, commonly named *ok*, is a `boolean indicator`. When *ok* is true, it signifies the `presence of the key in the map`. Conversely, when *ok* is false, it denotes the `absence of the key`.

??? example "Using Ok Idiom"

    ```bash title="run command"
    $ go run src/fundamentals/composite/maps_ok.go
    1 true
    2 true
    Key3 not exists in data:  0 false
    3 true
    ```
    ```go
    --8<-- "src/fundamentals/composite/maps_ok.go"
    ```

## As Function Params

**`If a function takes a map argument, changes it makes to the elements of the slice will be visible to the caller`**, analogous to passing a pointer to the underlying array.

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/maps_params.go
    Inital mySlice:  map[key1:Corinthians key2:Lakers]
    Memory allocated for mySlice: 0xc000050020
    Memory allocated for data: 0xc000050030
    After call func Slice:  map[key1:Corinthians key2:Lakers key3:Liverpool]
    ```
    ```go
    --8<-- "src/fundamentals/data_types/maps_params.go"
    ```

## Operations

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/maps_operations.go
    Initial myMap:  map[key1:1 key2:2 key3:3]
    key3 Value:  3
    Added key4:  map[key1:1 key2:2 key3:3 key4:4]

    Key1 exists:  1
    key5 does not exist

    After delete key1 (exist) and key5 (not exist):  map[key2:2 key3:3 key4:4]
    Len of myMap:  3
    Key: key2, Value: 2
    Key: key3, Value: 3
    Key: key4, Value: 4
    ```
    ```go
    --8<-- "src/fundamentals/data_types/maps_operations.go"
    ```

## Go's Package for Maps

Package maps defines various functions useful with maps of any type. [Link here](https://pkg.go.dev/golang.org/x/exp/maps).

## References

- [Go maps in action](https://go.dev/blog/maps)
- [Effective Go - Maps](https://go.dev/doc/effective_go#maps)
