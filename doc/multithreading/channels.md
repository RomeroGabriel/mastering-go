# Channels

Channels in Go are a powerful feature that `allows goroutines to communicate with each other and synchronize their execution`. They are `used to pass data between goroutines and can be thought of as pipes that connect concurrent goroutines`.

You can create a channel using the `make` function. `Channels are typed by the values they convey`. You can `send data into a channel` using the `channel <-` syntax and `receive data` from a channel using the `<-channel` syntax.

??? example

    ```bash title="run command"
    go run src/multithereading/basic_channel.go
    ```
    ```go
    --8<-- "src/multithereading/basic_channel.go"
    ```
    ```bash title="output"
    Hello Word!
    ```

## Buffered Channels

By default, sends and receives `block until both the sender and receiver are ready`. However, you can `create a buffered channel that can hold a certain number of values before blocking`. For example, to create a buffered channel that can hold up to 2 strings. In this case, the `send operation won't block until the channel's buffer is full, and the receive operation won't block until the channel's buffer is empty`.

??? example

    ```bash title="run command"
    go run src/multithereading/buffered_channel.go
    ```
    ```go
    --8<-- "src/multithereading/buffered_channel.go"
    ```
    ```bash title="output"
    message 1
    message 2
    ```

## Deadlock

If a goroutine is sending data on a channel, then `it is expected that some other goroutine should be receiving the data`. `If this does not happen, then the program will panic at runtime with a deadlock error`. Similarly, if a goroutine is waiting to receive data from a channel, then some other goroutine is expected to write data on that channel, else the program will panic.

## Range and Close

You can use the `range` keyword to `read values from a channel until it's closed`. When the channel is closed and drained, the range loop will terminate.

??? example

    ```bash title="run command"
    go run src/multithereading/range_channel.go
    ```
    ```go
    --8<-- "src/multithereading/range_channel.go"
    ```
    ```bash title="output"
    ping 1
    ping 2
    ```

## References

1. [Channels](https://golangbot.com/channels/)
1. [How to use Go channels](https://blog.logrocket.com/how-use-go-channels/)
