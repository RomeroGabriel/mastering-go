# Migrations

The [migrate library](https://github.com/golang-migrate/migrate) is a Go package that provides a `framework for managing database schema changes`. It supports various databases such as PostgreSQL, MySQL, SQLite, and others. The library allows you to `write migration scripts in SQL, JSON, YAML, etc.`, and then apply these scripts to your database using the command line interface or programmatically within your Go application.

## Key Features

!!! note "Database Support"
    The library supports a wide range of databases including PostgreSQL, MySQL, SQLite, and others. This makes it versatile and suitable for different projects that may use different types of databases.

!!! note "Script Types"
    You can write your migration scripts in several formats including SQL, JSON, YAML, etc. This flexibility allows you to choose the format that best suits your project's needs.

!!! note "Command Line Interface"
    The library comes with a command line interface that allows you to manage your database schema without writing any Go code. You can use this interface to create new migrations, apply existing ones, rollback migrations, etc.

!!! note "Programmatic Usage"
    In addition to the command line interface, you can also use the library programmatically within your Go application. This gives you more control over how and when your migrations are applied.

## CLI

The CLI is a powerful tool that allows you to `manage your database schema directly from the terminal`. Here are some common commands you might use.

!!! example "Migrate up"
    This command applies all available migrations to the specified database.

    ```bash
    $ migrate -path /path/to/migrations -database postgres://localhost:5432/mydb?sslmode=disable up
    ```

!!! example "Migrate down"
    This command undoes the last batch of migrations.

    ```bash
    $ migrate -path /path/to/migrations -database postgres://localhost:5432/mydb?sslmode=disable down
    ```

!!! example "Migrate force"
    This command forces the migration version to be 1, effectively rolling back all migrations.

    ```bash
    $ migrate -path /path/to/migrations -database postgres://localhost:5432/mydb?sslmode=disable force 1
    ```

!!! example "Migrate version"
    This command prints the current migration version.

    ```bash
    $ migrate -path /path/to/migrations -database postgres://localhost:5432/mydb?sslmode=disable version
    ```

## Go - Programmatic Use

The Go language API provides more control and flexibility, allowing you to integrate database migrations into your Go applications.

!!! example
    ```go
    package main

    import (
        "database/sql"
        "log"

        "github.com/golang-migrate/migrate"
        _ "github.com/lib/pq"
    )

    func main() {
        db, err := sql.Open("postgres", "postgres://localhost:5432/mydb?sslmode=disable")
        if err != nil {
            log.Fatal(err)
        }

        driver, err := postgres.WithInstance(db, &postgres.Config{})
        if err != nil {
            log.Fatal(err)
        }

        m, err := migrate.NewWithDatabaseInstance(
            "file://migrations",
            "postgres", driver)
        if err != nil {
            log.Fatal(err)
        }

        if err := m.Up(); err != nil && err != migrate.ErrNoChange {
            log.Fatal(err)
        }
    }
    ```
