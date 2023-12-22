# Testing

In Go, the `testing` package provides support for `automated testing` of Go packages. It is intended to be used in concert with the `go test` command, `which automates execution of any function of the form func TestXxx(*testing.T) where Xxx does not start with a lowercase letter`. The function name serves to identify the test routine. Within these functions, use the `Error`, `Fail` or related methods to signal failure.

`To write a new test suite, create a file that contains the TestXxx functions and give that file a name ending in _test.go`. `The file will be excluded from regular package builds but will be included when the "go test" command is run`. The test file can be in the same package as the one being tested, or in a corresponding package with the suffix `_test`.

!!! example

    ```go
    --8<-- "src/testing/tax.go"
    ```
    ```go
    --8<-- "src/testing/tax_test.go"
    ```

## Measure Test Coverage

In Go, you can measure test coverage using the command `go test -coverprofile` flag. This flag instructs the command to write a coverage profile to a file after running the tests. `The coverage profile includes information about which parts of your code were executed during the tests`.

!!! example
    This command will run all tests in the current directory and write the coverage profile to a file named `coverage.out`.

    ```bash
    go test -coverprofile=coverage.out
    ```
    ```bash title="output"
    PASS
    coverage: 80.0% of statements
    ok      tax     0.003s
    ```

Analyze the coverage profile using the `go tool cover` command. `This command can generate a report in different formats, such as text or HTML`.

??? example

    This command will generate an `HTML report` and write it to a file named `coverage.html`. You can then open this file in a web browser to view the coverage report.

    ```bash
    go tool cover -html=coverage.out -o coverage.html
    ```
    ```html
    --8<-- "src/testing/coverage.html"
    ```

## Benchmarking

You can also write `benchmark tests` in Go, which `measure the performance of a function or program`. It allows you to identify `potential areas for optimization` and `understand the impact of the changes you make to your code`. The `testing` package in Go provides support for benchmarking, which is done by creating `functions with a specific signature that takes a pointer to testing.B as its only argument`.

To `run` the benchmark, use the `go test` command with the `-bench` flag.

??? example

    This command will generate an `HTML report` and write it to a file named `coverage.html`. You can then open this file in a web browser to view the coverage report.

    ```bash
    go test -bench=.
    # OR go test -bench=. -count 5 -run=^# -benchtime=5s
    ```
    ```go
    --8<-- "src/testing/tax_bench.go"
    ```
    ```go
    --8<-- "src/testing/tax_bench_test.go"
    ```
    ```bash title="output"
    goos: linux
    goarch: amd64
    pkg: tax
    cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
    BenchmarkCalculateTax-8         1000000000               0.2814 ns/op
    BenchmarkCalculateTax2-8            1084           1109872 ns/op
    PASS
    ok      tax     1.629s
    ```

## Fuzzing

Fuzzing is a type of automated testing which `continuously manipulates inputs to a program to find bugs`. Fuzzing uses `coverage guidance to intelligently walk through the code being fuzzed to find and report failures to the user`. It can reach `edge cases` which humans often miss, making fuzz testing particularly valuable for `finding security exploits and vulnerabilities`.

    ```bash
    go test -fuzz=. -fuzztime=5s
    ```
    ```go
    --8<-- "src/testing/tax.go"
    ```
    ```go
    --8<-- "src/testing/tax_fuzz_test.go"
    ```
    ```bash title="output"
    fuzz: elapsed: 0s, gathering baseline coverage: 0/8 completed
    fuzz: elapsed: 0s, gathering baseline coverage: 8/8 completed, now fuzzing with 8 workers
    fuzz: elapsed: 3s, execs: 89077 (29688/sec), new interesting: 0 (total: 8)
    fuzz: elapsed: 5s, execs: 150973 (29952/sec), new interesting: 0 (total: 8)
    PASS
    ok      tax     5.073s
    ```

## Reference

1. [Go testing](https://pkg.go.dev/testing)
1. [How To Write Unit Tests in Go](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package)
1. [Go Fuzzing](https://go.dev/doc/security/fuzz/)
