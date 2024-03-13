# Array

In Go, an array is a `fixed-size collection of elements of the same type`. Arrays are declared with a specific size, and that **`size cannot be changed after the array is created`**. Arrays are useful when planning the detailed layout of memory and `sometimes can help avoid allocation`, but primarily they are a building block for [slices](../vars.md#slices). The `size of an array is part of its type`, so the types `[10]int` and `[20]int` are distinct. **Arrays are equal if they are the same length and contain equal values**.

!!! info
    Arrays are rarely used directly in Go because of above reasons.

??? example "Declaration, Syntax and Use"

    ```bash title="run command"
    $ go run src/fundamentals/composite/arrays.go
    Indexes not set get default value:  0
    All values in the array:  [0 10 20]

    Value stored in 0 index is 0
    Value stored in 1 index is 10
    Value stored in 2 index is 20

    It's possible to change index value in array:  [0 600 20]

    myArray == secondArray?  true
    Other Declaration: [1 0 0 0 0 4 6 0 0 0 100 15]
    ```
    ```go
    --8<-- "src/fundamentals/composite/arrays.go"
    ```

## As Function Params

A function can receive an array by **`specifying the array's type and size`** in the function's parameter list. When you pass an array to a function, **`Go passes the entire array by value`**, which means the function receives a copy of the array.

If you want to pass an array to a function and have the function modify the original array, you would need to pass a pointer to the array instead. However, **`arrays in Go are not addressable`**, so you cannot take the address of an array directly. Instead, you would typically use a [slices](../vars.md#slices), which is a reference to an underlying array and can be passed by reference.

??? example "Array Params Side-Effect"

    ```bash title="run command"
    $ go run src/fundamentals/composite/array_params.go
    Inital myArray:  [1 2 3 4]
    Memory allocated for myArray: 0xc000018180
    Memory allocated for data: 0xc0000181e0
    After call func Slice:  [1 2 3 4]
    ```
    ```go
    --8<-- "src/fundamentals/composite/array_params.go"
    ```

## Slicing an Array

In Go, slicing an array you `create a new` [slices](../vars.md#slices) that references a portion of the original array. There are several ways to slice an array, depending on the range of indices you specify. The index format follows this pattern: **`a[start:stop]`**.

!!! Danger "Change the original array"
    Remember that slicing does not create a new array/slice. **`It creates a new slice header that points to the same underlying array. Both slices/arrays share the same memory, and changes to one are reflected in the other.`**. Avoid modifying slices after they have been sliced or if they were produced by slicing. Check out [How to Work with Subslice](#how-to-work-with-subslice).

??? example "Slicing Use Cases"

    ```bash title="run command"
    $ go run src/fundamentals/composite/arrays_slicing.go
    To create a copy of the array, use full slice omit both the start and end indices:  [10 20 30 40 50 60 70 80 90 100]
    Notice that the copyedArray is actually a slice variable: []int
    Address of myArray: 0xc00007c000
    Address of copyedArray: 0xc000010018
    Check that changing the copyedArray changes also the myArray!!!!!!! [10 20 30 40 50 60 70 80 90 10000] [10 20 30 40 50 60 70 80 90 10000]

    Getting 5 first elements, start to end:  [10 20 30 40 50]
    Getting from index 5 to end, from to end:  [60 70 80 90 10000]
    Getting between 2 and 5, between indexes:  [30 40 50]
    Go does not have built-in support for stepping through a slice like Python does. However, you can achieve this by using a loop or by manually selecting the desired indices.
    Stepped Slice:  [10 30 50 70 90]
    ```
    ```go
    --8<-- "src/fundamentals/composite/arrays_slicing.go"
    ```

### How to Work with Subslice

To simplify slice handling, you should either **`never use append with a subslice`** or ensure that append operations do not overwrite existing data by employing a *complete slice expression*. This approach explicitly `delineates the shared memory extent between the parent slice and the subslice`. A complete slice expression comprises three parts, with the third part denoting the final position within the parent slice's capacity that remains accessible to the subslice. To calculate the subslice's capacity, subtract the starting offset from this value.

!!! tip
    If you need to create a slice that’s independent of the original, , see [Copying a Slice](slice.md#copying-a-slice).

??? example "Complete Slice Expression"

    ```bash title="run command"
    $ go run src/fundamentals/composite/full_slice_expr.go
    x: [a b c d] 5 4
    y: [a b] 2 2
    z: [c d] 2 2

    x: [a b c d x] 5 5
    y: [a b i j k] 5 5
    z: [c d y] 4 3

    ```
    ```go
    --8<-- "src/fundamentals/composite/full_slice_expr.go"
    ```

## Convert Array -> Slice and Array <- Slice

If you have an array and you want to convert it to a slice, you don't actually need to do anything special. A slice is just a reference to an underlying array, along with a length. `So, when you pass an array to a function expecting a slice, Go automatically converts the array to a slice that covers the entire array`.

Converting a slice back to an array is a bit trickier because `slices can be shorter or longer than arrays`. **`If you know the size of the array you want to convert to, you can create a new array and copy the elements from the slice into it`**.

??? example "Array to Slice"

    ```bash title="run command"
    $ go run src/fundamentals/composite/arrays_convert.go
    printing s:  [1 2 3]
    type of s:  []int
    Slice to Array [1 2 3]
    ```
    ```go hl_lines="14-15" linenums="1"
    --8<-- "src/fundamentals/composite/arrays_convert.go"
    ```

??? example "Slice to Array knowing the Size"

    ```bash title="run command"
    $ go run src/fundamentals/composite/arrays_convert.go
    printing s:  [1 2 3]
    type of s:  []int
    Slice to Array [1 2 3]
    ```
    ```go hl_lines="17-20" linenums="1"
    --8<-- "src/fundamentals/composite/arrays_convert.go"
    ```

## References

- [Effective Go - Arrays](https://go.dev/doc/effective_go#arrays)
