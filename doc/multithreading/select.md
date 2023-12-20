# Select

The `select` statement in Go is used to `handle multiple channel operations concurrently`. It blocks until one of the operations can proceed. `This is particularly useful when you have multiple goroutines communicating via channels and you want to handle their messages as they arrive`. You can also use the `default` case in a select statement to `prevent blocking`. `If none of the cases are ready`, the `default` case will be executed.

!!! example
    We then use a `select statement to receive data from either channel`. In this case, it will print the data received from the `number channel because this channel is ready first`.

    ```bash title="run command"
    go run src/multithereading/basic_select.go
    ```
    ```go
    --8<-- "src/multithereading/basic_select.go"
    ```
    ```bash title="output"
    Channel Data: 15
    ```
!!! example "Default"

    ```bash title="run command"
    go run src/multithereading/default_select.go
    ```
    ```go
    --8<-- "src/multithereading/default_select.go"
    ```
    ```bash title="output"
    ...
    2023/12/20 12:13:12 Default
    Number received: 1
    2023/12/20 12:13:14 Default
    Number received: 1
    2023/12/20 12:13:16 Default
    2023/12/20 12:13:17 Default
    Number received: 1
    2023/12/20 12:13:19 Default
    2023/12/20 12:13:20 Default
    2023/12/20 12:13:21 Default
    ```

!!! example "More Complex example"

    ```bash title="run command"
    go run src/multithereading/select.go
    ```
    ```go
    --8<-- "src/multithereading/select.go"
    ```
    ```bash title="output"
    ...
    Received from ch2: ID: 1 - Message 1 from Kafka
    Received from ch2: ID: 2 - Message 2 from Kafka
    Received from ch2: ID: 4 - Message 4 from Kafka
    Received from ch2: ID: 5 - Message 5 from Kafka
    Received from ch1: ID: 5 - Message 3 from RabbitMQ
    Received from ch1: ID: 6 - Message 6 from RabbitMQ
    Received from ch1: ID: 7 - Message 7 from RabbitMQ
    Received from ch1: ID: 8 - Message 8 from RabbitMQ
    Received from ch1: ID: 9 - Message 9 from RabbitMQ
    2023/12/20 12:05:49 timeout
    2023/12/20 12:05:51 timeout
    Received from ch1: ID: 10 - Message 10 from RabbitMQ
    Received from ch2: ID: 11 - Message 11 from Kafka
    2023/12/20 12:05:53 timeout
    Received from ch2: ID: 12 - Message 12 from Kafka
    Received from ch1: ID: 13 - Message 13 from RabbitMQ
    2023/12/20 12:05:57 timeout
    Received from ch1: ID: 14 - Message 14 from RabbitMQ
    Received from ch2: ID: 15 - Message 15 from Kafka
    2023/12/20 12:06:01 timeout
    ```

## References

1. [Go Select](https://www.programiz.com/golang/select)
1. [How To Effectively Utilize Golang Select In Your Projects](https://marketsplash.com/tutorials/go/golang-select/)
