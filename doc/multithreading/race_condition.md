# Race Condition

In Go, a `race condition occurs when two or more goroutines access and modify the same data concurrently`, leading to unpredictable and often erroneous behavior. `Go provides several tools to prevent race conditions and ensure the integrity of shared data`, including Mutexes, Read-Write Locks, and Atomic Operations.

!!! tip "go run -race"
    The `-race` flag in Go is used to enable the `built-in data race detector`. `Data races occur when two or more goroutines access the same memory location concurrently, and at least one of the accesses is a write`. This can lead to unpredictable behavior and hard-to-debug issues. It's typically used for `testing and debugging`, not for running production code.

    ```bash title="run command"
    go test -race mypkg   // test the package
    go run -race mysrc.go // compile and run the program
    go build -race mycmd  // build the command
    go install -race mypkg // install the package
    ```

## Mutexes (Mutual Exclusion)

A Mutex, short for mutual exclusion, is the `most basic form of concurrency control` in Go. It `allows only one goroutine to access a critical section of code at a time`.

??? example
    The shared `counter` variable, `which multiple goroutines increment concurrently`. `The sync.Mutex ensures that only one goroutine can execute the increment function at a time by locking and unlocking the Mutex before and after the critical section`.

    ```bash title="run command"
    go run src/multithereading/mutex.go
    ```
    ```go
    --8<-- "src/multithereading/mutex.go"
    ```
    ```bash title="output"
    Counter:  1000
    ```

## Read-Write Locks

Read-Write Locks are a `type of lock that can be held by multiple readers but only one writer at a time`. They are useful when you have `many read operations and few write operations`.

In Go, the `sync.RWMutex` type provides a read-write lock that `allows multiple goroutines to read data simultaneously but ensures exclusive access for writing`. This is useful in scenarios where you have many read operations and few write operations.

??? example
    The `sync.RWMutex` protect access to the `data` map. `Multiple goroutines can read the data using RLock()`, allowing for concurrent reading. `When writing data, we use Lock() to ensure exclusive access`, preventing any concurrent reads or writes.

    ```bash title="run command"
    go run src/multithereading/rwmutex.go
    ```
    ```go
    --8<-- "src/multithereading/rwmutex.go"
    ```
    ```bash title="output"
    Read Data: value1
    Read Data: value1
    Read Data: value1
    Read Data: value1
    Read Data: value1
    Write Data: key1 -> newvalue1
    Write Data: key1 -> newvalue1
    ```

## Atomic Operations

While Mutexes and Read-Write Locks provide high-level abstractions for concurrency control, Go also offers `atomic operations for more fine-grained control over shared variables`. These operations `allow you to perform read-modify-write operations on variables atomically`, without the need for locks.

??? example
    The `atomic.AddInt32` atomically increment the counter variable. `Since atomic operations are executed without locks, they are highly efficient and are suitable for scenarios where fine-grained control is needed with minimal overhead`.

    ```bash title="run command"
    go run src/multithereading/atomic.go
    ```
    ```go
    --8<-- "src/multithereading/atomic.go"
    ```
    ```bash title="output"
    Counter:  1000
    ```

## Choosing the Right Concurrency Control Mechanism

When it comes to choosing the right concurrency control mechanism in Go, you need to consider the `specific requirements of your application`. `Mutexes are suitable for scenarios where only one goroutine should access a critical section at a time`. `Atomic operations are more efficient but can be more difficult to use correctly`. In some cases, a combination of these mechanisms may be beneficial.

## References

1. [Concurrency Control in Go: Mutexes, Read-Write Locks, and Atomic Operations](https://clouddevs.com/go/concurrency-control/)
1. [Introducing the Go Race Detector](https://go.dev/blog/race-detector)
1. [Data Race Detector](https://go.dev/doc/articles/race_detector)
