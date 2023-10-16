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

## Go Mod

### Go Mod Init

!!! info

    ```bash
    go mod init example.com/module
    ```

This command `initializes a new module with the specified name`. You can then `add dependencies to the module using the go get command and manage the module's requirements through the go.mod file`.

### Go Mod TIDY

This is a command that `ensures that the go.mod file matches the source code in the module`. `It adds any missing and removes any unused module dependencies`, keeping the go.mod file clean and ensuring that it accurately reflects the actual dependencies needed for your project.

!!! info

    ```bash
    go mod tydi
    ```

`Make sure that you have initialized Go modules in your project (with go mod init) before running the go mod tidy command`. This ensures that the go.mod file exists in the root of your project directory.

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
