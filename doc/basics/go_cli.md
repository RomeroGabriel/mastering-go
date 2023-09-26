# Go CLI

## Go Environment Command

Lists all Go `environment variables`.

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
