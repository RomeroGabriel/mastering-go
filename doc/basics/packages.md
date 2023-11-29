# Packages and Modules

## Packages

Every `.go file begins with the package command`, and the `package's name should match the directory in which the file resides`. The exceptions to this rule are `package main` and `package test` which serve as the entry points for applications and tests, respectively. `Packages within the same directory should share the same package name. Anything within the same package can be accessed and utilized`.

??? example "Simple package example"

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

### Installing Packages

See [Go Get](cli/general.md#go-get) and [Go Mod TIDY](cli/go_mod.md#go-mod-tidy).

## Modules

In Go, a `module is a collection of related Go packages that are versioned together as a single unit`. Modules are the key component of the Go module system, `introduced to manage package dependencies and versioning`.

A Go module typically consists of one or more Go packages, `along with a go.mod file that defines the module's requirements and dependencies`. The go.mod file keeps track of the module's dependencies, versions, and the other modules it requires.

To create a new module in Go, you can use the [go mod init](cli/go_mod.md#go-mod-init) command followed by the name of the module.

!!! note "Public and Private Access"
    In Go, the concept of `public and private access is determined by the first letter's case of a variable or function name`. If the first letter of an identifier is in `uppercase`, it is `exported and can be accessed from other packages`. On the other hand, if the first letter is `lowercase`, the identifier is considered `unexported and can only be accessed within the same package`.

    This convention helps maintain encapsulation and ensures that package-level variables and functions are used only where intended.

??? example

    ```bash title="run command"
    cd src/modules/
    go run main.go
    ```
    ```go title="src/modules/mathpackage/mathpackage.go"
    --8<-- "src/modules/mathpackage/mathpackage.go"
    ```
    ```go title="src/modules/car_package/car_package.go"
    --8<-- "src/modules/car_package/car_package.go"
    ```
    ```go title="src/modules/main.go"
    --8<-- "src/modules/main.go"
    ```
    ```bash title="output"
        Accessing sum function. Private to mathpackage
    Result of SumInt:  71
        Accessing sum function. Private to mathpackage
    Result of SumFloat:  20.94
    ****
    Starting explore Car Module
    Full Car name:  Corolla Giant car
    Full Car Years Old:  3
    ```

## References

- [Go Expert - FullCycle](https://goexpert.fullcycle.com.br/curso/)
