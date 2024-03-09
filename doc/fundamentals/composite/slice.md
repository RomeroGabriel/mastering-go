# Slice

In Go, a slice is a `more versatile and dynamic alternative to arrays`. **Slices are like views into an underlying array**, allowing you to work with a portion of an **`array without specifying a fixed size`**. Slices are reference types, that hold references to an underlying array, and **`if you assign one slice to another, both refer to the same array`**. Different from [literal types](../literals.md#literals), the default value for slices is `nil`.

!!! tip
    Using `[...]` makes an array. Using `[]` makes a slice.

??? example "Simple Use Cases"

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

## Compare Slices

A slice `isn't comparable`, resulting in a compile-time error when using `==` or `!=` with a slice. The only thing you can compare a slice with using `==` is `nil`. However, starting from Go version `1.21`, the [slices](#gos-package-for-slices) package in the standard library provides two functions for comparing slices. The `slices.Equal` function accepts two slices as input parameters and returns true if both slices have the `same length and all their elements are equal`. The `slices.EqualFunc` function offers the flexibility to pass a `custom comparison function`, allowing comparisons without the need for slice elements to be directly comparable.

!!! note "Elements in `slices.Equal`"
    Notably, this function mandates that the elements of the slice must be comparable.

## Operations - len, append, cap

!!! info
    Slices are one of the supported types by the [len](../built_in/functions.md#len) function.

The `cap` function, short for capacity, is used to `determine the capacity` of [slices](../composite/slice.md#slice) and [arrays](../composite/array.md#array) in Go. For slices, it **`returns the maximum number of elements that the slice can hold without resizing the underlying array`**. For arrays, it returns the length of the array.

The `append` function is used to `add elements to slices dynamically`. It takes a slice and one or more elements as input and `returns a new slice` containing the original elements plus the additional ones. If the capacity of the original slice is sufficient to accommodate the new elements, `append modifies the existing slice in place`. Otherwise, it creates a new underlying array with increased capacity and copies the elements over.

??? example "Using len, append and cap"

    ```bash title="run command"
    $ go run src/fundamentals/composite/slices_operations.go
    Initial mySlice:  [10 20 30 40 50 60]
    len and cap:  6 6

    Starting using append:
    Appending an element --------->
    mySliceAdded:  [10 20 30 40 50 60 70]
    len and cap:  7 12
    mySlice without changes:  [10 20 30 40 50 60]
    len and cap:  6 6

    Removing an element --------->
    mySlice after remove index 3:  [10 20 30 50 60]
    len and cap:  5 6
    mySliceAdded:  [10 20 30 40 50 60 70]

    Insert at index --------->
    mySliceAdded after insert 1000 at index 5:  [10 20 30 40 50 1000 60 70]
    len and cap:  8 12
    ```
    ```go
    --8<-- "src/fundamentals/composite/slices_operations.go"
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
