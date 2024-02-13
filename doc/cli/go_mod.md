# go mod

`go mod` is a command in the Go programming language that `manages module dependencies` for Go projects. Introduced in Go 1.11, Go Modules aimed to improve the experience of developing Go applications by simplifying dependency management.

## go mod init

This command `initializes a new module with the specified path`. You can then `add dependencies to the module using the` [go get command](general.md#go-get) `and manage the module's requirements through the go.mod file`.

!!! example "Create a new module"

    ```bash
    go mod init {module-path}
    ```

!!! tip "Standard for modules name"
    `It's a good practice to use the location of the repository where Go tools can find the module's source code as the module path`, especially if you're planning to publish the module for others to use. `Be sure to specify a module path that won’t conflict with the module path of other modules`.

## go mod tidy

This is a command that `ensures that the go.mod file matches the source code in the module`. `It adds any missing and removes any unused module dependencies`, keeping the go.mod file clean and ensuring that it accurately reflects the actual dependencies needed for your project.

!!! example "Matching dependencies used in source code"

    ```bash
    $ go mod tidy
    ```

??? warning "Has a module initialize"
    `Make sure that you have initialized Go modules in your project (with go mod init) before running the go mod tidy command`. This ensures that the go.mod file exists in the root of your project directory. There more complex edits that can be found running `go help mod edit` or checking the [documentation](https://go.dev/ref/mod#go-mod-edit).

!!! example "The -e flag"

    The `-e` flag, causes `go mod tidy to attempt to proceed despite errors encountered while loading packages`.
    
    ```bash
    go mod tidy -e
    ```

## go mod edit

The `go mod edit` command allows you to `make changes to the module's requirements and dependencies`.

!!! example

    ```bash
    go mod edit -require=example.com/module@v1.2.3
    ```
    ```bash title="remove a requirement"
    go mod edit -droprequire=example.com/module
    ```

!!! warning "go mod edit -replace"
    The `go mod edit --replace` command in Go is used to add or update a replace directive in the go.mod file.

    ```bash
    go mod edit --replace=example.com/module=../local/module
    ```

    `HOWEVER, it's recommended to utilize workspaces, as go mod edit may encounter limitations outside the local environment`.

## References

- [Managing dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module)
