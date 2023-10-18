# Handle Files

## Create a File

To create a file, you can use the `os.Create` function along with the defer statement to close the file after writing to it.

??? example

    ```bash title="run command"
    go run src/handle_files/create_file.go
    ```
    ```go
    --8<-- "src/handle_files/create_file.go"
    ```

## Write into a File

Open the file in `append mode` and write to it using the `os.OpenFile` function.

??? example

    ```bash title="run command"
    go run src/handle_files/write_file.go
    ```
    ```go
    --8<-- "src/handle_files/write_file.go"
    ```

### File Permission Number

In the `os.OpenFile` function, the `perm argument is used to specify the file permission when creating or opening a file`. It is an integer value that represents the file mode and permission bits. The typical file permission values used in Unix-like systems are represented in octal notation. `The permission mode is composed of three digits, each of which represents the permissions for different user groups`.

1. `Owner Permissions`: The first digit represents the permissions for the file owne.
1. `Group Permissions`: The second digit represents the permissions for the group that the file belongs to.
1. `Other Permissions`: The third digit represents the permissions for everyone else (other users).

The digits in the mode are a sum of permission bits:

1. `4: Read (r)`
1. `2: Write (w)`
1. `1: Execute (x)`

And calculate the permissions numbers:

1. `6 (4 + 2) means read and write permissions (rw)`.
1. `4 (only 4) means read-only (r).`.
1. `7 (4 + 2 + 1) means read, write, and execute permissions (rwx).`.

??? example
    1. `0644`: Owner has read and write (6), group and others have read-only (4).
    1. `0755`: Owner has read, write, and execute (7), group and others have read and execute (5).
    1. `0600`: Owner has read and write (6), no permissions for group and others.

## Reading a File

To read from a file, you can use the `os.Open` function along with a `Scanner from the bufio package`.

??? example

    ```bash title="run command"
    go run src/handle_files/read_file.go
    ```
    ```go
    --8<-- "src/handle_files/read_file.go"
    ```
    ```bash title="output"
    Writing a string into the file
    Using bytes to write
    ```

## Delete a File

To delete a file, you can use the `os.Remove` function.

??? example

    ```bash title="run command"
    go run src/handle_files/delete_file.go
    ```
    ```go
    --8<-- "src/handle_files/delete_file.go"
    ```
