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

## Reference

1. [Go testing](https://pkg.go.dev/testing)
