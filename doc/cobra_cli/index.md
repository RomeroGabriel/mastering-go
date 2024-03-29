# Cobra CLI

Cobra is a library for creating powerful modern `CLI applications` in Go. It provides a `simple interface to create powerful CLI interfaces`, similar to git or docker. It has a robust command handling system, including support for nested commands, global flags, and more. The library also includes functionality for generating auto-completion scripts, help text, and usage instructions.

## Start Project

!!! example

The `cobra-cli init` command is used to `initialize a new Cobra application`. This command sets up the initial application code structure for you, allowing you to immediately start benefiting from the features of Cobra. It also has the capability to apply the `license` you specify to your application.

    ```bash
    cobra-cli init
    ```

## Commands

At its core, Cobra is built on a `structure of commands, arguments, and flags`. Commands represent actions, arguments are things, and flags are modifiers for those actions. The pattern to follow is `APPNAME COMMAND ARG --FLAG`. For instance, in the command `hugo server --port=1313`, `server` is a command, and `port` is a flag.

### Add a new Command

The add command in Cobra CLI is a method used to `add a new command to a Cobra application`. This command is `added to the root command of your application and can be accessed via the command line interface`.

!!! info

    ```bash
    cobra-cli add {command_name}
    ```

For each command added, a new file in `cmd` folder will be created.

??? example

    ```go linenums="1" hl_lines="14-23 37"
    --8<-- "src/cobra_cli/cmd/testCommand.go"
    ```

    In this example, `testCommand` is a new command that is defined with a short description, a long description, and a `function to execute when the command is run`. `The init function is called when the package is initialized`, and it adds the cmdAdd command to the root command of the application.

    Once you've defined the add command, you can access it via the command line by typing `yourApplicationName add`.

### Subcommands

`Subcommands are commands that are part of other commands`. They allow you to create complex command structures in your CLI application, making it easier to manage and navigate. To create a nested command in Cobra, you `first define the parent command, then add the nested command to it`.

!!! info

    ```bash
    $ cobra-cli add category
    $ cobra-cli add create -p 'categoryCmd'
    $ cobra-cli add list -p 'categoryCmd'
    ```

??? example

    ```go linenums="1" hl_lines="18-19"
    --8<-- "src/cobra_cli/cmd/category.go"
    ```
    ```go linenums="1" hl_lines="20-21"
    --8<-- "src/cobra_cli/cmd/create.go"
    ```
    ```go linenums="1" hl_lines="20-21"
    --8<-- "src/cobra_cli/cmd/list.go"
    ```
    ```bash linenums="1" hl_lines="10-11 18-21"
    $ go run main.go category
    A longer description that spans multiple lines and likely contains examples
    and usage of using your command.

    Usage:
    cobra_cli_example category [flags]
    cobra_cli_example category [command]

    Available Commands:
    create      A brief description of your command
    list        A brief description of your command

    Flags:
    -h, --help   help for category

    Use "cobra_cli_example category [command] --help" for more information about a command.
    $
    $ go run main.go category create
    create called
    $ go run main.go category list
    list called
    ```

## Arguments/Flags

Arguments in Cobra CLI are `values that can be passed to commands`. They are typically used to provide input data to the command. `Arguments are used to modify the behavior of a command`. They can be both `global` and `local`.

??? example

    ```go linenums="1" hl_lines="13 24-30 37-41"
    --8<-- "src/cobra_cli/cmd/testCommand.go"
    ```
    ```bash
    $ go run main.go testCommand -t "Salve o Corinthians"
    salve o corinthians
    $ go run main.go testCommand -t "Salve o Corinthians" -u
    SALVE O CORINTHIANS
    ```

### Local and Global Flags

??? example "Local"
    A flag associated with a specific command.

    ```go linenums="1" hl_lines="22"
    --8<-- "src/cobra_cli/cmd/create.go"
    ```
    ```bash linenums="1" hl_lines="9"
    $ go run main.go category create -h
    A longer description that spans multiple lines and likely contains examples
    and usage of using your command.

    Usage:
    cobra_cli_example category create [flags]

    Flags:
        --created_type string   The Category type to be created
    -h, --help                  help for create

    Global Flags:
        --name string   Name of the category
    ```

??? example "Global"

    ```go linenums="1" hl_lines="20"
    --8<-- "src/cobra_cli/cmd/category.go"
    ```
    ```bash linenums="1" hl_lines="15"
    $ go run main.go category
    A longer description that spans multiple lines and likely contains examples
    and usage of using your command.

    Usage:
    cobra_cli_example category [flags]
    cobra_cli_example category [command]

    Available Commands:
    create      A brief description of your command
    list        A brief description of your command

    Flags:
    -h, --help          help for category
        --name string   Name of the category

    Use "cobra_cli_example category [command] --help" for more information about a command.
    ```

A flag persistent across multiple commands.

## Hooks

Cobra does not natively support hooks, but you can achieve similar functionality using the `RunE`, `PreRun`, `PreRunE`, `PostRun`, and `PostRunE` functions. These functions allow you to specify what happens before and after a command runs.

??? example

    ```go linenums="1" hl_lines="14-22"
    --8<-- "src/cobra_cli/cmd/hookCmd.go"
    ```
    ```bash linenums="1"
    $ go run main.go hookCmd
    PRE RUN
    hookCmd called
    POST RUN
    ```

## References

1. [Cobra GitHub](https://github.com/spf13/cobra)
1. [Cobra](https://cobra.dev/)
