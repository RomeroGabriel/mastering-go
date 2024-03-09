# Slice

In Go, a slice is a `more versatile and dynamic alternative to arrays`. **Slices are like views into an underlying array**, allowing you to work with a portion of an **`array without specifying a fixed size`**. Slices are **`reference types`**, that hold references to an underlying array, and **`if you assign one slice to another, both refer to the same array`**.

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/slices.go
    Value in index 2:  30
    All values in the slice:  [10 20 30 40 50 60]

    Value stored in 0 index is 10
    Value stored in 1 index is 20
    Value stored in 2 index is 30
    Value stored in 3 index is 40
    Value stored in 4 index is 50
    Value stored in 5 index is 60

    It is possible to change index value in array in both arrays:  [10 600 30 40 50 60] [10 600 30 40 50 60]
    See that mySlice didn't change mySlice:  [10 999 30 40 50 60] [10 600 30 40 50 60]
    ```
    ```go
    --8<-- "src/fundamentals/data_types/slices.go"
    ```

## As Function Params

**`If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller`**, analogous to passing a pointer to the underlying array.

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/slices_params.go
    Inital mySlice:  [1 2 3 4]
    Memory allocated for mySlice: 0xc000010018
    Memory allocated for data: 0xc000010048
    After call func Slice:  [90 2 3 4]
    Notice that append elements didn't change mySlice, just change value in index
    ```
    ```go
    --8<-- "src/fundamentals/data_types/slices_params.go"
    ```

## Slicing a Slice

Same thing that [slice an array](array.md#slicing-an-array).

## Operations

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/slices_operations.go
    Initial mySlice:  [10 20 30 40 50 60]
    Appending an element --------->
    mySliceAdded:  [10 20 30 40 50 60 70]
    mySlice without changes:  [10 20 30 40 50 60]

    Removing an element --------->
    mySlice after remove index 3:  [10 20 30 50 60]
    mySliceAdded:  [10 20 30 40 50 60 70]

    Insert at index --------->
    mySlice after remove:  [10 20 30 40 50 1000 60 70]
    mySliceAdded:  [10 20 30 40 50 1000 60 70]
    ```
    ```go
    --8<-- "src/fundamentals/data_types/slices_operations.go"
    ```

## Memory Allocation

Remember that slices are based on arrays, so they can change in size dynamically. `If the underlying array runs out of capacity, a new larger array is allocated, and the data is copied over`.

!!! danger "When a slice needs more memory..."
    **`When a slice needs more capacity, Go doubles the original size. In the case of very large slices, this can consume more memory than necessary`**.

Doubling the capacity of the underlying array allows for more elements to be added in the future without frequent reallocations.

!!! tip
    When you anticipate working with a large slice, it's a **`good practice to initialize the slice with a size closer to the maximum you expect to use to minimize unnecessary memory consumption`**.

Slices also have a `length and capacity` (the maximum number of elements it can hold without reallocation). Use the `len()` and `cap()` functions to get these values.

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/slices_memory.go
    mySlice: len=6 cap=6 [10 20 30 40 50 60]
    noneValue: len=0 cap=6 []
    firstFour: len=4 cap=6 [10 20 30 40]
    last_two: len=4 cap=4 [30 40 50 60]
    mySlice with 70:  [10 20 30 40 50 60 70]
    mySlice: len=7 cap=12 [10 20 30 40 50 60 70]
    ```
    ```go
    --8<-- "src/fundamentals/data_types/slices_memory.go"
    ```

    When added 70 in the example above, notice that the `slice's capacity increases to 12 instead of 7`.

## Two-dimensional Slices

Because slices are variable-length, it is possible to have each inner slice be a different length.

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/2d-slices.go
    ```
    ```go
    --8<-- "src/fundamentals/data_types/2d-slices.go"
    ```

## Go's Package for Slices

Package slices defines various functions useful with slices of any type. [Link here](https://pkg.go.dev/slices).

## References

- [Effective Go - Slices](https://go.dev/doc/effective_go#slices)
