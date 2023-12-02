# Go CLI

## Go Environment Command

Lists all Go `environment variables`.

!!! info

    ```bash
    go env
    ```

??? example "go env output"

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

### GOPATH

`GOPATH` is a variable that `defines the root of your workspace`. It stores your code base and all the files necessary for your development. `It also contains the binaries of your compilations and the libraries used.`

### GOMODCACHE

`GOMODCACHE` is related to `Go module management`, which is used for handling dependencies. It is `typically a subfolder within your GOPATH`.

## Go Get

To install packages in Go, you can use the `go get command followed by the package name`.

`Make sure that you have initialized Go modules in your project (with go mod init) before running the go get command`. This ensures that the go.mod file exists in the root of your project directory.

!!! info

    ```bash
    go get [packages]
    ```

It is also used to update and manage dependencies in your GOPATH. Some helpful subcommands you can use with go get:

1. `-u`: Update the package and install the latest version;
1. `-d`: Download the package's source code but do not install it.
1. `-t`: Also download the packages needed to run tests.
1. `-v`: Enable verbose output to see the package being fetched and installed.

## Go Build

The `go build` command is used to compile packages and dependencies `into an executable binary file`. This command compiles the Go source files in the current directory and creates an executable binary file.

Building a `single file` in Go can be done by simply using the `go build filename.go` command followed by the name of the file. This command will `generate an executable binary based on the contents of the specified filename.go source file`. For `projects  that utilize the go.mod file` can be achieved by running the go build `command in the root directory of the project`. This command automatically builds the project, resolving dependencies defined in the go.mod file.

!!! info

    ```bash title="just one file"
    go build filename.go
    ```
    ```bash title="for projects"
    go build
    ```

### Variables to Build

In Go, `GOOS` is an environment variable that specifies the `target operating system for code compilation`. The `GOARCH` is an environment variable that specifies the `target architecture for code compilation`. To find out the possibilities for `GOOS` and `GOARCH`, use [go tool dist list](#go-tool-dist-list). Is good start a new

!!! info

    ```bash
    GOOS=[target-OS] GOARCH=[target-architecture] go build
    ```

## Go Tool

### Go Tool dist list

In order to find out what operating systems and platforms are available for building executables, you can use the dist tool.

!!! info

    ```bash
    go tool dist list
    ```
??? example

    android/386

    android/arm

    android/arm64

    darwin/amd64

    darwin/arm64

    ios/amd64

    ios/arm64

    linux/386

    linux/amd64

    linux/arm

    windows/386

    windows/amd64

    windows/arm

    ...

## Go Work

To create a workspace, you can use the `go work init` command with a list of module directories as space-separated arguments.

!!! example

    ```bash
    go work init module1 module2 ...
    ```

The `go work use -r` command can be used to recursively add directories in the argument directory with a go.mod file to your workspace

!!! example

    ```bash
    go work use -r
    ```

To add modules to the workspace, you can run `go work use` or manually edit the `go.work` file.

!!! example

    ```bash
    go work use [moddir]
    ```

## Refereces

- [How To Build Go Executables for Multiple Platforms on Ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-20-04)
