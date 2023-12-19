# Go Routines

Goroutines are one of the most important aspects of the Go programming language. `They are the smallest unit of execution and can continue their work alongside the main goroutine, creating concurrent execution`. They are `lightweight`, `cheap to create`, and can be used to effectively utilize `multi-core CPUs`. To create a goroutine is simple, just need to add the keyword `go` in front of the function you want to `run concurrently`.

??? example "Basic Example"

    ```bash title="run command"
    go run src/multithereading/basic.go
    ```
    ```go
    --8<-- "src/multithereading/basic.go"
    ```
    ```bash title="output"
    Just an annoying print
    0 from Corinthians
    0 from Knicks
    1 from Knicks
    Just an annoying print
    1 from Corinthians
    2 from Corinthians
    2 from Knicks
    Just an annoying print
    3 from Corinthians
    Just an annoying print
    3 from Knicks
    4 from Knicks
    4 from Corinthians
    5 from Knicks
    5 from Corinthians
    6 from Knicks
    6 from Corinthians
    7 from Knicks
    7 from Corinthians
    8 from Knicks
    8 from Corinthians
    9 from Knicks
    9 from Corinthians
    ```

## sync.WaitGroup

A `sync.WaitGroup` in Go is `used to wait for a collection of goroutines to finish`. It's a way to ensure that all goroutines have completed their execution before the program continues. This is particularly useful when you have multiple goroutines performing tasks concurrently and `you need to wait for all of them to finish before proceeding`.

Before starting a goroutine, you call the `Add` method on the WaitGroup to `increment the counter`. This tells the `WaitGroup` to `wait for one more goroutine to finish`. Inside the goroutine, when the work is done, you call the `Done` method on the `WaitGroup to decrement the counter`. Finally, you call the `Wait` method on the WaitGroup to `block the current goroutine until the counter is zero`.

??? example

    ```bash title="run command"
    go run src/multithereading/wait_group.go
    ```
    ```go
    --8<-- "src/multithereading/wait_group.go"
    ```
    ```bash title="output"
    Just an annoying print
    0 from Corinthians
    0 from Knicks
    1 from Knicks
    Just an annoying print
    1 from Corinthians
    2 from Corinthians
    2 from Knicks
    Just an annoying print
    3 from Corinthians
    Just an annoying print
    3 from Knicks
    4 from Knicks
    4 from Corinthians
    5 from Knicks
    5 from Corinthians
    6 from Knicks
    6 from Corinthians
    7 from Knicks
    7 from Corinthians
    8 from Knicks
    8 from Corinthians
    9 from Knicks
    9 from Corinthians
    ```

If you need to set a `timeout` for the WaitGroup, you can use a `select` statement with a `time.After` function.

??? example "With Timeout"

    ```bash title="run command"
    go run src/multithereading/wait_group_timeout.go
    ```
    ```go
    --8<-- "src/multithereading/wait_group_timeout.go"
    ```
    ```bash title="output"
    Just an annoying print
    0 from Corinthians
    1 from Corinthians
    Just an annoying print
    2 from Corinthians
    Just an annoying print
    3 from Corinthians
    Just an annoying print
    4 from Corinthians
    5 from Corinthians
    6 from Corinthians
    7 from Corinthians
    8 from Corinthians
    2023/12/18 18:43:22 Hit timeout
    ```

## When to Use Goroutines

Goroutines are useful when `one task can be split into different segments to perform better`. Any work that can utilize a multi-core CPU should be well optimized using goroutines. `Running background operations` in a program might also be a use case for a goroutine.

## Real-Life Use Cases

Some real-life use cases of goroutines include `reading a huge file and processing it` for exception or error messages, and `posting multiple API calls in different threads when they are not dependent on each other`.

## Goroutines x Normal Threads/Process

Goroutines and threads are `both used for concurrent execution of code`, but they have `significant differences in terms of scheduling, communication, execution speed, infrastructure dependency, stack management, latency during program execution, resource control, and local storage management`.

`Goroutines are extremely cheap when compared to threads`. `They are only a few kilobytes in stack size` and the stack can grow and shrink according to the needs of the application. `There might be only one thread in a program with thousands of Goroutines`. If any Goroutine in that thread blocks, then another OS thread is created and the remaining Goroutines are moved to the new OS thread.

!!! note "Scheduling Management"

    `Goroutines are managed by the Go runtime`, which uses a technique known as [m:n scheduling](https://en.wikipedia.org/wiki/Thread_(computing)#M:N_(hybrid_threading)), `where m goroutines are executed using n operating system threads using multiplexing`. This means that `the Go scheduler is not invoked periodically by a hardware timer`, but implicitly by certain Go language constructs. `Because it doesn’t need a switch to kernel context, rescheduling a goroutine is much cheaper than rescheduling a thread`.

    ??? example
        For example, when a goroutine calls `time.Sleep` or `blocks` in a [channel](channels.md#channels) or mutex operation, the scheduler puts it to sleep and runs another goroutine until it is time to wake the first one up.

    On the other hand, `threads are managed by the operating system`. Every few milliseconds, a hardware timer interrupts the processor, which causes a kernel function called the scheduler to be invoked. `This function suspends the currently executing thread and saves its registers in memory`, looks over the list of threads and decides which one should run next, restores that thread’s registers from memory, then resumes the execution of that thread. `Because OS threads are scheduled by the kernel, passing control from one thread to another requires a full context switch, which is a slow operation`.

!!! note "Communication Medium"
    Goroutines enhance communication through the use of a [channel](channels.md#channels) and `sync package` which provides a wait group function. `Threads do not have a clear communication medium`. In multiple threads executing, `communication is made through memory location`.

!!! note "Infrastructure Dependency"
    `Goroutines are not hardware dependent`, meaning they can be executed independently of any infrastructure. `Threads, are ardware dependent`.

!!! note "Stack Size"
    Goroutines are executed in a `stack of 2kb (kilobytes)`, which `grows gradually` and is `destroyed once execution is done/completed`. Threads also execute in a stack, but they require at least a minimum of `1 megabyte` to execute, and `stack size is fixed`. Thus, stack management is easier with goroutines compared to threads.

!!! note "Latency During Program Execution"
    `Goroutines communicate with each other through` [channels](channels.md#channels), thus low latency is experienced from one channel to another. In threads, since there is no communication medium between one thread to another, communication takes place with high latency.

## References

1. [Goroutines in Golang](https://golangdocs.com/goroutines-in-golang)
1. [Goroutines](https://golangbot.com/goroutines/)
1. [Goroutines vs Threads in Golang](https://www.golinuxcloud.com/goroutine-vs-threads-golang/)
