# Templates

Templates are a powerful way to `separate the presentation layer from your code logic`. The `text/template` and `html/template` packages provide functionalities for `parsing and executing text and HTML templates`. Templates are `useful for rendering dynamic content in web applications or generating structured text documents`.

!!! info "html/template package"
    The `html/template` package provides additional HTML-specific features and ensures proper escaping for HTML content.

When you use `template.Must`, it helps `simplify error handling when parsing templates`. It panics if the template parsing fails, eliminating the need to explicitly check for errors. However, `this approach is suitable only during initialization`, and the application should terminate if the templates cannot be parsed properly.

!!! example

    ```bash title="run command"
    go run src/templates/template.go
    ```
    ```go
    --8<-- "src/templates/template.go"
    ```
    ```bash title="output"
    Hello, Gabriel
    ************
        <p>All Users:</p>
        <h1>Hello, Gabriel!</h1>
        <p>Age: 15</p>
        <h1>Hello, Cassio!</h1>
        <p>Age: 40</p>
            <p>Adult User</p>
        <h1>Hello, Yuri!</h1>
        <p>Age: 20</p>
            <p>Adult User</p>
        <h1>Hello, Lucas!</h1>
        <p>Age: 17</p>
    ```

## Compose Templates

You can compose templates by defining reusable parts and embedding them within other templates. This helps in creating modular and maintainable templates.

!!! example

    ```bash title="run command"
    go run src/templates/compose_template.go
    curl localhost:8080
    ```
    ```go
    --8<-- "src/templates/compose_template.go"
    ```
    ```bash title="output"
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Persons</title>
    </head>
    <body>
            <p>All Users:</p>
            <h1>Hello, Gabriel!</h1>
            <p>Age: 15</p>
            <h1>Hello, Cassio!</h1>
            <p>Age: 40</p>
                <p>Adult User</p>
            <h1>Hello, Yuri!</h1>
            <p>Age: 20</p>
                <p>Adult User</p>
            <h1>Hello, Lucas!</h1>
            <p>Age: 17</p>
            <div>
            <p>A nice footer here!</p>
        </div>
    </body>
    </html>
    ```
