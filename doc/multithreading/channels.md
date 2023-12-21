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

## Channel Directions

In Go, channels can be created with a `specific direction`, which can be either `send-only`, `receive-only`, or `bidirectional`. This feature allows you to `enforce certain restrictions on how a channel can be used`, which can help prevent errors and make your code easier to understand.

!!! example "Send-Only Channels"
    Syntax: `chan<-`.

    ```go
    ch := make(chan<- int) // Send-only channel
    ch <- 42 // Sending data
    ```

!!! example "Receive-Only Channels"
    Syntax: `<-chan`

    ```go
    ch := make(<-chan int) // Receive-only channel
    data := <-ch // Receiving data
    ```

??? example "Full Example"

    ```bash title="run command"
    go run src/multithereading/channel_directions.go
    ```
    ```go
    --8<-- "src/multithereading/channel_directions.go"
    ```
    ```bash title="output"
    Corinthians
    ```

The Bidirectional channel is denoted by the `chan` syntax, `without any direction indicator`.

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

??? example "Range with WaitGroup"

    ```bash title="run command"
    go run src/multithereading/range_waitgroup.go
    ```
    ```go
    --8<-- "src/multithereading/range_waitgroup.go"
    ```
    ```bash title="output"
    Publishing value :0
    Publishing value :1
    Subscribe value: :0
    Subscribe value: :1
    Publishing value :2
    Publishing value :3
    Subscribe value: :2
    Subscribe value: :3
    Publishing value :4
    Publishing value :5
    Subscribe value: :4
    Subscribe value: :5
    Publishing value :6
    Publishing value :7
    Subscribe value: :6
    Subscribe value: :7
    Publishing value :8
    Publishing value :9
    Subscribe value: :8
    Subscribe value: :9
    ```

## Examples

1. [Api Race Challenge](https://github.com/RomeroGabriel/go-api-race)

??? example "Fake Load Balancer"

    ```bash title="run command"
    go run src/multithereading/fake_loadbalancer.go
    ```
    ```go
    --8<-- "src/multithereading/fake_loadbalancer.go"
    ```
    ```bash title="output"
    ...
    Worked 9 received 1926
    Worked 903 received 1997
    Worked 975 received 1971
    Worked 765 received 1942
    Worked 274 received 1839
    Worked 763 received 1856
    Worked 914 received 1860
    Worked 915 received 1864
    Worked 702 received 1877
    Worked 981 received 1898
    Worked 797 received 1879
    Worked 774 received 1853
    Worked 798 received 1882
    Worked 320 received 1275
    Worked 434 received 1276
    Worked 108 received 1073
    Worked 530 received 1277
    ```

## References

1. [Channels](https://golangbot.com/channels/)
1. [How to use Go channels](https://blog.logrocket.com/how-use-go-channels/)
1. [Golang: Channel Directions](https://towardsdev.com/golang-channel-directions-607637e9edac)
