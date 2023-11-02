# Databases

The `database/sql` package in Go `provides a generic interface around SQL (or SQL-like) databases`. It allows you to interact with various SQL databases using the same functions, `regardless of the specific database backend`.

## sql.Open

The `sql.Open` function is used to `open a database specified by its driver name and a driver-specific data source name`. This function does not generate an error in the event of an incorrect password, driver, etc. Use `db.Ping` to check if the connection is successful.

## sql.Prepare

The `db.Prepare` function is used to `create a prepared statement for later queries or executions`. Prepared statements are advantageous when you need to execute the same or similar SQL statements repeatedly with high efficiency.

## QueryRow

The `QueryRow` method is used to `execute a query that is expected to return at most one row`.

## Scan

The `Scan` method is then used to  `read the values from the current row into variables`. This method expects a list of destination variables as arguments, which it fills with the values from the columns in the result set.

## Code example

!!! example

    ```go
    --8<-- "src/db/main.go"
    ```
