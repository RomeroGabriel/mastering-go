# Wire

Wire is a powerful `code generation tool` designed to automate the process of connecting components through `dependency injection` in Golang. By representing `dependencies as function parameters, Wire encourages explicit initialization over the use of global variables`. Unlike some other dependency injection tools for Go, such as [dig](https://github.com/uber-go/dig), `Wire operates without relying on runtime state or reflection`. This characteristic not only ensures efficient code execution but also makes code written for `Wire compatible with manual initialization`.

## Key Features

!!! info "Code Generation"
    Wire operates as a code generator, `eliminating the need for calls to a runtime library`. This approach facilitates introspection of initialization and ensures accurate cross-references.

!!! info "Providers and Injectors"
    Wire introduces two core concepts - `providers` and `injectors`. `Providers are functions that produce values, while injectors call providers in dependency order`. This enables clean and organized initialization of components.

!!! info "Provider Sets"
    Providers can be grouped into provider sets, a convenient way to `manage and use multiple providers together when necessary`.

## Generating code

!!! warning "Build Constrains"
    To generate the injector, add the build constraint `//+build wireinject` to your code. A build constraint, also known as a `build tag`, is a condition under which a file should be included in the package. Build constraints are given by a line comment that begins.

To generate code, run Wire in the package directory with the `wire.go` present. The generated implementation of the injector will be saved in a file named `wire_gen.go`.

!!! example
    ```bash
    wire
    ```

    ```go title="wire"
    --8<-- "src/di/wire.go"
    ```
 
    ```go title="wire_gen"
    --8<-- "src/di/wire_gen.go"
    ```

## References

1. [Hire GitHub](https://github.com/google/wire/tree/main)
1. [Build Constraints](https://pkg.go.dev/go/build?utm_source=godoc#hdr-Build_Constraints)
1. [Compile-time Dependency Injection With Go Cloud's Wire](https://go.dev/blog/wire)
