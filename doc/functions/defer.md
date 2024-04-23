# Defer

In Go, the defer statement is used to `schedule a function call to be executed after the return statement`. It is often used to simplify functions that perform various clean-up actions. `defer is commonly used for tasks such as closing files, unlocking mutexes, or generally executing clean-up actions`. The deferred call's arguments are `evaluated immediately, but the function call is not executed until the surrounding function returns`.

It's important to note that `defer statements are executed in Last In, First Out (LIFO) order`. Multiple defer statements will be executed in the reverse order of their definition.

??? example "Following Execution Flow with Defer"

    ```bash title="run command"
    $ go run src/function/defer.go
    Exiting:  30
    2º:  20
    1º:  10
    ```
    ```go
    --8<-- "src/function/defer.go"
    ```

## Using and Changing Return Values

It's possible for a deferred function to `examine and modify the return values` of the surrounding function, using named return values. Defer can be used to add contextual information to an error returned.

??? example "Changing If Error Happens"

    If there is no error, commit the changes. `If committing the changes an error happens, change the err variable and roll back`.

    ```go linenums="1" hl_lines="8"
    func InsertData(ctx context.Context, db *sql.DB, value1 string) (err error) {
        tx, err := db.BeginTx(ctx, nil)
        if err != nil {
            return nil
        }
        defer func() {
            if err == nil {
                err = tx.Commit()
            }

            if err != nil {
                tx.Rollback()
            }
        }()
        _, err = tx.ExecContext(ctx, "INSERT ...", value1)
        if err != nil {
            return err
        }
        ...
        return nil
    }
    ```

## Return Closures to Clean Up

A pattern in Go is for a function that allocates resources to `return a closure that cleans up` the used resource. This way it is explicit that the resource needs to be closed, and since Go doesn't allow unused variables, it's necessary to consume the function.

??? example "Return a Closer Function"

    ```go
    func getFile(name string) (*os.Filem func(), error) {
        file, err := os.Open(name)
        if err != nil {
            return nil, nil, err
        }

        return file, func(){
            file.Close()
        }, nil
    }

    f, closer, err := getFile("text.txt")
    if err != nil {
        panic(err)
    }
    defer closer()
    ```
