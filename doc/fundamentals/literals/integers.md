# Integers

Integer literals are `base 10` by default in Golang, but different prefixes are used to indicate other bases:

1. 0b for binary (base 2)
1. 0o for octal (base 8)
1. 0x for hexadecimal (base 16)

!!! tip
    You can use either upper or lowercase letters for the prefix.

To make it easier to read longer integer literals, Go allows `you to put underscores in the middle of your literal`. The underscores do not affect the value of the number, and they can't be at the beginning or end of numbers and next to each other.

!!! example
    For example, group by thousands in base 10 (1_234).

## Integer Types

Go provides both signed and unsigned integers in a variety of sizes.

| Type name | Value range                                    |
|-----------|------------------------------------------------|
| int8      | –128 to 127                                    |
| int16     | –32768 to 32767                                |
| int32     | –2147483648 to 2147483647                      |
| int64     | –9223372036854775808 to 9223372036854775807    |
| uint8     | 0 to 255                                       |
| uint16    | 0 to 65535                                     |
| uint32    | 0 to 4294967295                                |
| uint64    | 0 to 18446744073709551615                      |

??? tip "Bit alias"
    A *byte* is an alias for **uint8**.

Go has `int` and `unit` as a special name. On a 32-bit CPU, int is a 32-bit signed/unsigned integer `like an int32`. On most 64-bit CPUs, int is a 64-bit signed/unsigned integer, `just like an int64`.

!!! danger "Compare ints"
    Because `int` isn’t consistent from platform to platform, it is a **compile-time error** to assign, compare, or perform mathematical operations between an int and an int32 or int64 without a type conversion.
