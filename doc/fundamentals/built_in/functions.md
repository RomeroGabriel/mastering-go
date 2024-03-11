# Built-In Functions

## len

The `len` function returns the length of various types in Go, including [arrays](../composite/array.md#array), [slices](../composite/slice.md#slice), [maps](../composite/maps.md#maps), [strings](../literals/strings.md#strings), and [channels](../../multithreading/channels.md#channels). For arrays, slices, and strings, it returns the number of elements or bytes, while for maps, it returns the number of key-value pairs. For channels, it returns the number of elements currently in the channel buffer.

## make

The `make` built-in function allocates and initializes an object of type [slices](../composite/slice.md#slice), [maps](../composite/maps.md#maps), or [chan](../../multithreading/channels.md#channels). The first argument is a type, not a value. Make's return type is the same as the type of its argument, not a pointer to it.

??? note "For each type possible"
    1. Slice: The size specifies the `length`. The capacity of the slice is equal to its length. A second integer argument may be provided to specify a different capacity; it must be no smaller than the length. For example, make([]int, 0, 10) allocates an underlying array of size 10 and returns a slice of length 0 and capacity 10 that is backed by this underlying array.
    1. Map: An empty map is allocated with enough space to hold the specified number of elements. The size may be omitted, in which case a small starting size is allocated.
    1. Channel: The channel's buffer is initialized with the specified buffer capacity. If zero, or the size is omitted, the channel is unbuffered.

## References

- [Go Builtin](https://pkg.go.dev/builtin)
