# Strings

In Go, a string is a sequence of one or more characters (letters, numbers, symbols) that can be either a constant or a variable. `Strings are made up of Unicode and are immutable sequences, meaning they cannot be changed once created`.

Strings in Go are `similar to byte slices` in that they can be used in operations that work with byte slices. This means you can use functions like **`copy`** and **`append`** with strings.

The internal structure of a string in Go is essentially a **`byte slice wrapper`**. A string is `represented as a struct with two fields`: **elements**, which points to the underlying bytes, and **len**, which stores the number of bytes. This internal representation is crucial for understanding how strings are manipulated and stored in memory.

`Raw strings` are delimited with backquotes (\`) and can contain any character except a backquote. There’s no escape character in a raw string literal, `all characters are included as is`. When using a raw string literal, you write in multiline.

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/string.go
    Memory allocated for str1:  0xc000014070
    Memory allocated for str2:  0xc000014080
    Even though both have the same value, the memory address is different.

    Using append bytes:  Salve o Corinthians
    Using copy bytes:  Salve o
    Lorem ipsum dolor sit amet, consectetur adipiscing elit.Nunc ac urna ex. Mauris eu dolor nec orci aliquam consectetur vitae ut sapien.
        Quisque sed commodo dui. Cras pulvinar aliquet eleifend. Nulla dignissim arcu quis nunc facilisis dictum.
    Curabitur consequat, purus vel sodales dictum, dui metus lacinia lacus, in congue est.
    ```
    ```go
    --8<-- "src/fundamentals/data_types/string.go"
    ```

## Go's Package for String

Package strings defines various functions useful with maps of any type. [Link here](https://pkg.go.dev/strings).

## Runes

In Go, a rune is a data type used to `represent a Unicode code point/character`, which is a unique integer (which is an alias for the int32 type) value that corresponds to a character in the Unicode standard. The Unicode standard assigns a unique code point to almost every character from every writing system in the world, making it suitable for working with text in various languages and scripts.

### When to Use Runes

You should use runes when you need to `work with individual characters in a string and handle text that includes characters from different languages`, including non-ASCII characters. Common use cases include text processing, manipulation, and validation. Using runes is an efficient way to handle text in Go because `it’s optimized for Unicode character encoding`. It ensures that you can accurately represent and manipulate characters from various languages.

!!! example

    ```bash title="run command"
    $ go run src/fundamentals/data_types/runes.go
    Equal
    H e l l o ,     世 界 
    Hello,  世界
    ```
    ```go
    --8<-- "src/fundamentals/data_types/runes.go"
    ```

## References

- [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)
