# Maps

Go provides a built-in map type that implements a **hash table**. Maps in Go are a data structure used to store **`key-value pairs`**. The `key can be of any type for which the equality operator is defined`, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays. The value may be any type at all, including another map!

!!! warning "Key types not supported"
    Slices cannot be used as map keys, because equality is not defined on them.

Maps are **`reference types`**, that hold references to an underlying **hash table**. **`If you pass a map to a function or assign it to another variable, both refer to the same data structure`**. Since maps are reference type like pointers or slices to initialize a map, use the built in make function.

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/maps.go
    Not initalize map:  map[]
    After add elements:  map[key1:1 key2:2 key3:3]
    Value from key1:  1
    When key doesn't exist, default value is returned:  0
    ```
    ```go
    --8<-- "src/fundamentals/data_types/maps.go"
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
