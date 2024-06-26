# Maps

!!! info "Default Value"
    The zero value for a map is `nil`.

Maps in Go are a data structure used to store `key-value pairs` that implements a *hash table* data structure. The map type is written as `map[keyType]valueType`. Maps are `reference type` like [pointers](../../pointers.md#pointers) and [slices](slice.md#slice).

??? info "Hash Table"
    Hash tables are a type of data structure in which the address/ index value of the data element is generated from a hash function. This enables very fast data access as the index value behaves as a key for the data value.

    Go doesn’t require (or even allow) you to define your own hash algorithm or equality definition. Instead, the Go runtime that’s compiled into every Go program has code that implements hash algorithms for all types that are allowed to be keys.

The `key can be of any can be any comparable type`, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays. The `value may be any type at all`, including another map. `Maps are not comparable`. You have the option to verify if they are `equal to nil`, but you `cannot compare two maps` possess identical keys and values by employing `==`, or difference `!=`.

!!! warning "Key types not supported"
    Slices and Maps cannot be used as map keys, because equality is not defined on them.

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

## Compare Maps

A map `isn't comparable`, resulting in a compile-time error when using `==` or `!=` with another map. The only thing you can compare a slice with using `==` is `nil`. However, starting from Go version `1.21`, the [maps](#gos-package-for-maps) package in the standard library provides two functions for comparing slices. The `maps.Equal` function accepts two maps as input parameters and returns true if both two maps contain the same key/value pairs. The `maps.EqualFunc` function offers the flexibility to pass a `custom comparison function`.

## As Function Params and Return Values

`If a function takes a map argument, changes it makes to the elements of the slice will be visible to the caller`, analogous to passing a pointer to the underlying array. Passing a map to a function means that you are `copying a pointer`.

Maps as input parameters or return values `should be used carefully`. Beyond the `mutability problem`, on an API-design level, `maps are bad because nothing is explicit about the values held by the variable`, the only way to know the data meaning is to walk through the code. When possible, use a struct instead of a map.

??? example "Changing a Map Param"

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

## Operations - delete, clear

The `delete` built-in function deletes the element with the specified `key (m[key])` from the map. If m is nil or there is no such element, nothing happens.

Go 1.21 added a `clear` function that takes in a map and `deletes all entries`, resulting in an `empty map`.

??? example "Playing with Map"

    ```bash title="run command"
    $ go run src/fundamentals/data_types/maps_operations.go
    Initial myMap:  map[key1:1 key2:2 key3:3]
    key3 Value:  3
    Added key4:  map[key1:1 key2:2 key3:3 key4:4]

    Key1 exists:  1
    key5 does not exist

    After delete key1 (exist) and key5 (not exist):  map[key2:2 key3:3 key4:4]
    Len of myMap:  3
    Key: key4, Value: 4
    Key: key2, Value: 2
    Key: key3, Value: 3
    map[]
    ```
    ```go
    --8<-- "src/fundamentals/data_types/maps_operations.go"
    ```

## Maps as Sets

Verifying whether an element exists in a slice becomes progressively `slower as you add more elements to the slice`. Since a map operates based on a hash table, you can employ a `map to emulate certain features of a set`. In this approach, utilize the map's key to represent the type you intend to incorporate into the set, with a boolean or an empty struct serving as the value. Should you require set functionalities such as union, intersection, and subtraction, you have the option to either develop your own implementation or leverage one of the numerous third-party libraries offering such capabilities.

!!! tip "Using a Empty Struct as Value"
    Some people opt to utilize `struct{}` as the value when utilizing a map to emulate a set. The reason that an **empty struct consumes zero bytes**, whereas a boolean occupies one byte. However, employing struct{} may introduce clumsiness into your code. It necessitates a less intuitive assignment approach and mandates the `use of the comma-ok idiom for checking value presence in the set`. Unless you're dealing with exceedingly large sets, the difference in memory consumption is unlikely to be substantial enough to offset the drawbacks.

??? example "Maps as Set"

    ```bash title="run command"
    $ go run src/fundamentals/data_types/maps_as_sets.go
    10 is in the set
    8 is in the set
    ```
    ```go
    --8<-- "src/fundamentals/data_types/maps_as_sets.go"
    ```

## Go's Package for Maps

Package maps defines various functions useful with maps of any type. [Link here](https://pkg.go.dev/golang.org/x/exp/maps).

## References

- [Go maps in action](https://go.dev/blog/maps)
- [Effective Go - Maps](https://go.dev/doc/effective_go#maps)
