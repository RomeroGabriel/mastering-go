# Array

In Go, an array is a **`fixed-size collection of elements of the same type`**. Arrays are declared with a specific size, and that **`size cannot be changed after the array is created`**. Arrays are useful when planning the detailed layout of memory and `sometimes can help avoid allocation`, but primarily they are a building block for [slices](../vars.md#slices). **`If you pass an array to a function, it will receive a copy of the array`**, not a pointer to it. The `size of an array is part of its type`, so the types [10]int and [20]int are distinct.

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/arrays.go
    0
    [0 10 20]
    Value stored in 0 index is 0
    Value stored in 1 index is 10
    Value stored in 2 index is 20
    Value stored in 0 index is apple
    Value stored in 1 index is banana
    Value stored in 2 index is cherry
    ```
    ```go
    --8<-- "src/fundamentals/data_types/arrays.go"
    ```

## Slicing an Array

In Go, slicing an array you to `create a new` [slices](../vars.md#slices) that references a portion of the original array. There are several ways to slice an array, depending on the range of indices you specify. The index format follows this pattern: **`a[start:stop]`**.

!!! Danger "Change the original array"
    Remember that slicing does not create a new array. **`It creates a new slice header that points to the same underlying array. Modifying the elements of the new slice will affect the original array`**, but changing the length of the new slice (e.g., appending elements) will not affect the original array unless you explicitly make a copy of the array or slice.

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/arrays_slicing.go
    ```
    ```go
    --8<-- "src/fundamentals/data_types/arrays_slicing.go"
    ```

## Convert Array -> Slice and Array <- Slice

If you have an array and you want to convert it to a slice, you don't actually need to do anything special. A slice is just a reference to an underlying array, along with a length. `So, when you pass an array to a function expecting a slice, Go automatically converts the array to a slice that covers the entire array`.

Converting a slice back to an array is a bit trickier because `slices can be shorter or longer than arrays`. **`If you know the size of the array you want to convert to, you can create a new array and copy the elements from the slice into it`**.

??? example "Array to Slice"

    ```bash title="run command"
    $ go run src/fundamentals/data_types/arrays_convert.go
    printing s:  [1 2 3]
    type of s:  []int
    Slice to Array [1 2 3]
    ```
    ```go hl_lines="14-15" linenums="1"
    --8<-- "src/fundamentals/data_types/arrays_convert.go"
    ```

??? example "Slice to Array knowing the Size"

    ```bash title="run command"
    $ go run src/fundamentals/data_types/arrays_convert.go
    printing s:  [1 2 3]
    type of s:  []int
    Slice to Array [1 2 3]
    ```
    ```go hl_lines="17-20" linenums="1"
    --8<-- "src/fundamentals/data_types/arrays_convert.go"
    ```

## References

- [Effective Go - Arrays](https://go.dev/doc/effective_go#arrays)
