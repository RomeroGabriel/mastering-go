# Defer

In Go, the defer statement is used to `schedule a function call to be executed just before the surrounding function returns`. It is often used to simplify functions that perform various clean-up actions. `defer is commonly used for tasks such as closing files, unlocking mutexes, or generally executing clean-up actions`. The deferred call's arguments are `evaluated immediately, but the function call is not executed until the surrounding function returns`. It's important to note that `defer statements are executed in Last In, First Out (LIFO) order`. Multiple defer statements will be executed in the reverse order of their definition.

??? example
    In this code, "Hello" is printed first, and then the rest is printed.  This is because the defer statement `defers the execution tatement until after the surrounding function main finishes`.

    ```bash title="run command"
    go run function/defer.go
    ```
    ```go
    --8<-- "src/function/defer.go"
    ```
    ```bash title="output"
    Hello
    Hi, testttting
    World
    ```
