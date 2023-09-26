# Packages

Every `.go file begins with the package command`, and the `package's name should match the directory in which the file resides`. The exceptions to this rule are 'package main' and 'package test,' which serve as the entry points for applications and tests, respectively. `Packages within the same directory should share the same package name. Anything within the same package can be accessed and utilized`.

!!! example "Simple package example"

    ```bash title="run command"
    go run src/packages/*
    ```
    ```go title="src/packages/nice_func.go"
    --8<-- "src/packages/nice_func.go"
    ```
    ```go title="src/packages/main.go"
    --8<-- "src/packages/main.go"
    ```
    ```bash title="output"
    Starting script
    Now, calling nice_func that is defined from the same file directory, on nice_func.go file
        I'm on the same directory than main, so my package has to be main
    ```

## References

- [Go Expert - FullCycle](https://goexpert.fullcycle.com.br/curso/)
