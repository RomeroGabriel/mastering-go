# Databases

The `database/sql` package in Go `provides a generic interface around SQL (or SQL-like) databases`. It allows you to interact with various SQL databases using the same functions, `regardless of the specific database backend`.

## sql.Open

The `sql.Open` function is used to `open a database specified by its driver name and a driver-specific data source name`. This function does not generate an error in the event of an incorrect password, driver, etc. Use `db.Ping` to check if the connection is successful.

## sql.Prepare

The `db.Prepare` function is used to `create a prepared statement for later queries or executions`. Prepared statements are advantageous when you need to execute the same or similar SQL statements repeatedly with high efficiency.

## Code example

!!! example

    ```go
    --8<-- "src/db/main.go"
    ```
