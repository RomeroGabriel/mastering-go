# Slice

## Using make

The [make](../built_in/functions.md#make) function allows you to create an empty slice that already has a `length or capacity specified`. When creating a slice using `make`, all elements are initialized to the default value of the slice's type.

One common beginner mistake is to try to populate those initial elements using append. Is this case, the value passed to append is placed at the end of the slice.

!!! danger "When capacity is lower than length"
    You can also specify an initial capacity with make. However, **`never specify a capacity that’s less than the length`**! It is a compile-time error to do so with a constant or numeric literal. If you use a variable to specify a capacity that’s smaller than the length, your program will `panic at runtime`.

??? example "How to use make with slice"

    ```bash title="run command"
    $ go run src/fundamentals/data_types/slice_make.go
    Notice that the value 10 is placed in the end:  [0 20 0 0 0 10]

    Creating slice is len 0 but with cap--->
    Notice that now value 10 is placed in the begin:  [10] 1 5
    [10 1 2 3 4]
    ```
    ```go
    --8<-- "src/fundamentals/data_types/slice_make.go"
    ```

In Go, a slice is a `more versatile and dynamic alternative to arrays`. **Slices are like views into an underlying array**, allowing you to work with a portion of an **`array without specifying a fixed size`**. Slices are reference types, that hold references to an underlying array, and **`if you assign one slice to another, both refer to the same array`**. Different from [literal types](../literals.md#literals), the default value for slices is `nil`. Each element in a slice is **`assigned to consecutive memory locations`**, which makes it quick to read or write these values.

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

## Operations - len, append, cap, clear

Slices are one of the supported types by the [len](../built_in/functions.md#len) function. The length of a slice is the `number of consecutive memory locations` that have been assigned a value.

The `cap` function, short for capacity, is used to `determine the capacity` of [slices](../composite/slice.md#slice) and [arrays](../composite/array.md#array) in Go. For slices, it **`returns the maximum number of elements that the slice can hold without resizing the underlying array`**. For arrays, it returns the length of the array.

The `append` function is used to `add elements to slices dynamically`. It takes a slice and one or more elements as input and `returns a new slice` containing the original elements plus the additional ones. If the capacity of the original slice is sufficient to accommodate the new elements, `append modifies the existing slice in place`. Otherwise, it creates a new underlying array with increased capacity and copies the elements over.

Go 1.21 added a `clear` function that takes in a slice and `sets all of the slice’s elements to their zero value`. The length of the slice remains unchanged.

!!! tip
    Check out the [Memory Allocation Section](#memory-allocation) to understand what happens when a slice's capacity is insufficient.

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
    mySliceAdded after clear:  [0 0 0 0 0 0 0 0] 8
    ```
    ```go
    --8<-- "src/fundamentals/composite/slices_operations.go"
    ```

## Memory Allocation

Remember that slices are based on arrays, so they can change in size dynamically. `If the underlying array runs out of capacity, a new larger array is allocated, and the data is copied over`.

When a slice runs out of capacity, `it involves allocating new memory and copying the existing data from the old memory to the new`. This process also necessitates garbage collection for the old memory. To optimize this operation, the Go runtime typically increases the slice's capacity more than just one element when it runs out of space. **`For slices with a capacity less than 256, the capacity is doubled`**. For larger slices, the growth rate is adjusted to **`(current_capacity + 768)/4`**, aiming for a 25% growth rate. This approach ensures that the growth rate gradually decreases as the slice's capacity increases, resulting in a more efficient memory management strategy.

This allows for more elements to be added in the future without frequent reallocations.

!!! danger "When a slice needs more memory..."
    **`In the case of very large slices, this can consume more memory than necessary`**.

!!! tip
    When you anticipate working with a large slice, it's a **`good practice to initialize the slice with a size closer to the maximum you expect to use to minimize unnecessary memory consumption`**.

!!! tip
    If you need to create a slice that’s independent of the original, see [Copying a Slice](#copying-a-slice).

??? example "Seeing slice changing capacity"

    ```bash title="run command"
    $ go run src/fundamentals/data_types/slices_memory.go
    mySlice: len=6 cap=6 [10 20 30 40 50 60]
    noneValue: len=0 cap=6 []
    firstFour: len=4 cap=6 [10 20 30 40]

    mySlice with 70:  [10 20 30 40 50 60 70]
    mySlice: len=7 cap=12 [10 20 30 40 50 60 70]
    firstFour: len=4 cap=6 [10 20 30 40]
    noneValue: len=0 cap=6 []
    ```
    ```go
    --8<-- "src/fundamentals/data_types/slices_memory.go"
    ```

## Slicing a Slice

Same thing that [slice an array](array.md#slicing-an-array).

## Copying a Slice

If is necessary to create a slice that’s independent of the original, use the built-in `copy` function. The copy built-in function copies elements from a `source slice into a destination slice`. Copy `returns the number of elements copied`.

!!! tip
    As a special case, it also will copy bytes from a string to a slice of bytes.

The function `copies as many values as it can from source to destination`, limited by whichever slice is smaller. The capacity of source and destination doesn’t matter, it’s the **length** that’s important.

??? example "Using Copy"

    ```bash title="run command"
    $ go run src/fundamentals/data_types/slices_copy.go
    Copying first 2 elements [1 2] 2
    Copying last 2 elements [3 4] 2
    Overlaping elements in source:  [2 3 4 4] 3

    Using arrays: 
    Copy from array [5 6]
    Copying all elements from source to array:  [1 2 3 4]
    ```
    ```go
    --8<-- "src/fundamentals/data_types/slices_copy.go"
    ```

## As Function Params

**`If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller`**, analogous to passing a pointer to the underlying array.

??? example "Slice Params"

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

## Two-dimensional Slices

Because slices are variable-length, it is possible to have each inner slice be a different length.

??? example "2D examples"

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
