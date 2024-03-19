# Code Blocks

In Go, variable declarations can occur in various contexts, including outside of functions, as function parameters, and as local variables within functions. `Each context where a declaration occurs is referred to as a block`. Variables, constants, types, and functions declared `outside of any functions` are grouped in the *`package block`*. Additionally, `import statements` define names for other packages, and these names reside within the *`file block`*. Everything declared `inside a function` it's in the *`function block`*.

## Shadowing Variables

Accessing an identifier defined in an `outer block` from within an *`inner block`* is possible in Go. When a declaration shares the same name as an identifier in an enclosing block, a **`shadowing variable`** is created, bearing the same name as the variable in the enclosing block. While the shadowing variable persists, `access to the shadowed variable is restricted`.

Using the `:=` syntax can inadvertently result in variable shadowing. It's essential to exercise caution when using `:= to avoid unintentionally shadowing variables` from an outer scope on the left-hand side unless deliberate shadowing is intended.

??? example "Creating a Showdowing Variable"

    ```bash title="run command"
    $ go run src/fundamentals/blocks.go
    10
    7
    10
    Redefining int type
    ```
    ```go
    --8<-- "src/fundamentals/blocks.go"
    ```

## The Universe Block

A weird concept in Go is the `universe block`. Despite Go being a compact language with only 25 keywords, built-in types (e.g., int and string), constants (e.g., true and false), and functions (e.g., make or close) *`are not keywords`*. Similarly, `nil is not classified as a keyword`. Rather than designating them as keywords, `Go treats these entities as predeclared identifiers and defines them within the universe block`, encompassing all other blocks. *`Since these names are declared in the universe block, they can potentially be overshadowed in other scopes`*.

!!! danger "Redefinig Any Identifiers"
    Care must be taken to avoid redefining any identifiers in the universe block. Accidental redefinitions can lead to unexpected behavior.
