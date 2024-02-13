# Go's CLI

Go's Command Line Interface (CLI) `refers to the set of tools and utilities provided by the Go` programming language that developers can use directly from the terminal or command prompt. The Go CLI is designed to be straightforward and efficient, allowing developers to perform common tasks quickly and with minimal configuration. It's an integral part of the Go ecosystem, enabling rapid development and deployment of Go applications.

These tools facilitate various tasks such as `compiling code`, `managing dependencies`, `running tests`, and more. Here's a brief overview:

- [go build](#go-build): Compiles Go source files into executable binaries.
- [go get](#go-get): Downloads and installs packages and dependencies required for a Go project.
- [go mod](go_mod.md#go-mod): Manages module dependencies for Go projects, which includes tasks like initializing modules, adding requirements, and tidying up dependencies.
- [go env](#go-env): Lists all Go environment variables.
- go run: Compiles and runs Go programs without explicitly creating an executable file.
- go test: Executes unit tests within Go packages to ensure code correctness.
- go fmt: Automatically formats Go source code according to the official style guidelines.
- go vet: Analyzes Go source code for potential issues that may not necessarily prevent compilation but could lead to bugs or unidiomatic code.
- go doc: Generates documentation for Go packages in various formats.
- go list: Lists all the packages named by the import paths, one per line.

## go env

Lists all Go `environment variables`.

!!! abstract "go env"

    ```bash
    $ go env
    GO111MODULE=''
    GOARCH='amd64'
    GOBIN=''
    GOCACHE='/home/gabriel/.cache/go-build'
    GOENV='/home/gabriel/.config/go/env'
    GOEXE=''
    ...
    ```
    ??? example "full output"

        ``` bash
        GO111MODULE=''
        GOARCH='amd64'
        GOBIN=''
        GOCACHE='/home/gabriel/.cache/go-build'
        GOENV='/home/gabriel/.config/go/env'
        GOEXE=''
        GOEXPERIMENT=''
        GOFLAGS=''
        GOHOSTARCH='amd64'
        GOHOSTOS='linux'
        GOINSECURE=''
        GOMODCACHE='/home/gabriel/go/pkg/mod'
        GONOPROXY=''
        GONOSUMDB=''
        GOOS='linux'
        GOPATH='/home/gabriel/go'
        GOPRIVATE=''
        GOPROXY='https://proxy.golang.org,direct'
        GOROOT='/usr/local/go'
        GOSUMDB='sum.golang.org'
        GOTMPDIR=''
        GOTOOLCHAIN='auto'
        GOTOOLDIR='/usr/local/go/pkg/tool/linux_amd64'
        GOVCS=''
        GOVERSION='go1.21.1'
        GCCGO='gccgo'
        GOAMD64='v1'
        AR='ar'
        CC='gcc'
        CXX='g++'
        CGO_ENABLED='1'
        GOMOD='/dev/null'
        GOWORK=''
        CGO_CFLAGS='-O2 -g'
        CGO_CPPFLAGS=''
        CGO_CXXFLAGS='-O2 -g'
        CGO_FFLAGS='-O2 -g'
        CGO_LDFLAGS='-O2 -g'
        PKG_CONFIG='pkg-config'
        GOGCCFLAGS='-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build4097433368=/tmp/go-build -gno-record-gcc-switches'
        ```

!!! info "GOPATH variable"
    `GOPATH` is a variable that `defines the root of your workspace`. It stores your code base and all the files necessary for your development. `It also contains the binaries of your compilations and the libraries used.`

!!! info "GOMODCACHE variable"
    `GOMODCACHE` is related to `Go module management`, which is used for handling dependencies. It is `typically a subfolder within your GOPATH`.

## go get

To install packages in Go, you can use the `go get command followed by the package name`.

!!! tip "Module Initialized"
    `Make sure that you have initialized Go modules in your project` ([with go mod init](go_mod.md#go-mod-init)) `before running the go get command`. This ensures that the go.mod file exists in the root of your project directory.

It is also used to update and manage dependencies in your `GOPATH`. Some helpful subcommands you can use with go get:

1. `-u`: Update the package and install the latest version;
1. `-d`: Download the package's source code but do not install it.
1. `-t`: Also download the packages needed to run tests.
1. `-v`: Enable verbose output to see the package being fetched and installed.

!!! example

    ```bash
    $ go get github.com/google/uuid
    $ go get github.com/google/uuid@1.0.0 # version 1.0.0
    $ go get github.com/google/uuid@abcdef1234567890 # commit hash abcdef1234567890
    $ go get -u github.com/google/uuid
    $ go get -d github.com/google/uuid
    $ go get -t github.com/google/uuid
    $ go get -v github.com/google/uuid
    ```

## go build

The `go build` command is used to compile packages and dependencies `into an executable binary file`. This command compiles the Go source files in the current directory and creates an executable binary file.

Building a `single file` in Go can be done by simply using the `go build filename.go` command followed by the name of the file. This command will `generate an executable binary based on the contents of the specified filename.go source file`.

!!! example "Building a single file"

    ```bash
    go build filename.go
    ```

For `projects  that utilize the go.mod file` can be achieved by running the go build `command in the root directory of the project`. This command automatically builds the project, resolving dependencies defined in the `go.mod` file.

!!! example "Building a project"

    ```bash
    go build
    ```

When building the code is possible to configure behaviors to build the code. `Build variables` are environment variables that influence the behavior of the `go build` command. `They allow you to customize the build process`, define build constraints, and pass additional flags to the compiler. Here's a concise summary of Build Variables.

!!! abstract "Some of variables"
    1. `GOARCH`: Specifies the target architecture for the binary, such as amd64, arm, 386, etc.
    1. `GOOS`: Defines the target operating system, like linux, windows, darwin (macOS), etc.
    1. `GOROOT`: Sets the location of the Go standard library and tools.
    1. `GOPATH`: Specifies the workspace directory where Go looks for source code and dependencies.
    1. `CGO_ENABLED`: Controls whether cgo is enabled for cross-compilation. Set to 1 to enable, 0 to disable.
    1. `GO111MODULE`: Determines the module mode for the current build. Can be set to on, off, or auto.
    1. `GOFLAGS`: Allows passing additional flags to the go build command.
    1. `GOBIN`: Defines the directory where compiled executables are placed.
    1. `GOPROXY`: Specifies the URL of the proxy server for downloading modules.
    1. `GOMOD`: Used to point to a specific go.mod file when building a module.
    1. `GODEBUG`: Enables debugging output for certain runtime components.

In Go, `GOOS` is an environment variable that specifies the `target operating system for code compilation`. The `GOARCH` is an environment variable that specifies the `target architecture for code compilation`. To find out the possibilities for `GOOS` and `GOARCH`, use [go tool dist list](#go-tool-dist-list).

To find out the available values for operating system and platform, checkout [go tool dist list](#go-tool-dist-list)

!!! example

    ```bash
    GOOS=[target-OS] GOARCH=[target-architecture] go build
    ```

## go tool

`go tool` is a suite of command-line tools included with the Go programming language distribution. `These tools are designed to assist developers in various aspects of software development, including code analysis, testing, performance profiling, and more`. The `go tool` command itself is a `meta-command` that provides access to several subcommands, each serving a distinct purpose.

Here's a brief description of some commonly used `go tool` subcommands:

- `go tool compile`: Compiles Go source files into object files.
- `go tool link`: Links object files together to produce an executable binary.
- `go tool nm`: Lists the symbols from object files.
- `go tool objdump`: Displays information from object files.
- `go tool trace`: Visualizes execution traces of Go programs.
- `go tool pprof`: Analyzes and visualizes performance profiles of Go programs.
- `go tool cover`: Provides code coverage analysis for Go tests.
- `go tool vet`: Reports suspicious constructs in Go source code.
- `go tool yacc`: Generates parsers from yacc-like grammar files.
- `go tool pack`: Packs object files into archive files.
- `go tool cgo`: Handles C code generation and linking for programs that use cgo.

Each of these tools is specialized for a particular task and can be used individually or in combination with other tools to achieve complex development workflows. The `go tool` suite is an essential part of the Go ecosystem, providing powerful functionality that complements the core features of the language and its standard library.

### go tool dist list

List all operating systems and platforms are available for building executables.

!!! example

    ```bash
    $ go tool dist list
    android/386
    android/arm
    android/arm64
    ...
    ```

## go work

Go 1.18 adds workspace mode to Go, which lets you `work on multiple modules simultaneously` without having to edit go.mod files for each module. `Each module within a workspace is treated as a main module when resolving dependencies`. With Go workspaces, `you control all your dependencies using a go.work file in the root of your workspace directory`. The go.work file has use and replace directives that override the individual go.mod files, `so there is no need to edit each go.mod file individually`.

??? note "Before workpaces"
    Previously, `to add a feature to one module and use it in another module, you needed to either publish the changes to the first module, or edit the go.mod file of the dependent module with a replace directive for your local, unpublished module changes`. In order to publish without errors, you had to remove the replace directive from the dependent module’s go.mod file after you published the local changes to the first module.

!!! example "Creating a workspace"
    To create a workspace, you can use the `go work init` command with a list of module directories as space-separated arguments. The workspace `doesn’t need to contain the modules you’re working with`. The init command creates a `go.work` file that lists modules in the workspace. If you run go work init without arguments, the command creates an empty workspace.

    ```bash
    $ go work init module1 module2 ...
    ```

!!! example "Add modules to a workspace"
    To add modules to the workspace, you can run `go work use` or manually edit the `go.work` file.

    ```bash
    go work use [moddir]
    ```
    The `go work use -r` command can be used to recursively add directories in the argument directory with a `go.mod` file to your workspace

    ```bash
    go work use -r .
    ```

## Refereces

- [How To Build Go Executables for Multiple Platforms on Ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-20-04)
- [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)
