# Cobra CLI

Cobra is a library for creating powerful modern `CLI applications` in Go. It provides a `simple interface to create powerful CLI interfaces`, similar to git or docker. It has a robust command handling system, including support for nested commands, global flags, and more. The library also includes functionality for generating auto-completion scripts, help text, and usage instructions.

## Commands

At its core, Cobra is built on a `structure of commands, arguments, and flags`. Commands represent actions, arguments are things, and flags are modifiers for those actions. The pattern to follow is `APPNAME COMMAND ARG --FLAG`. For instance, in the command `hugo server --port=1313`, `server` is a command, and `port` is a flag.

### Start Project

!!! example

The `cobra-cli init` command is used to `initialize a new Cobra application`. This command sets up the initial application code structure for you, allowing you to immediately start benefiting from the features of Cobra. It also has the capability to apply the `license` you specify to your application.

    ```bash
    cobra-cli init
    ```

### Add a new Command

The add command in Cobra CLI is a method used to `add a new command to a Cobra application`. This command is `added to the root command of your application and can be accessed via the command line interface`.

!!! example

    ```bash
    cobra-cli add {command_name}
    ```

For each command added, a new file in `cmd` folder will be created.

??? example

    ```go
    --8<-- "src/cobra_cli/cmd/testCommand.go"
    ```

    In this example, `testCommand` is a new command that is defined with a short description, a long description, and a `function to execute when the command is run`. `The init function is called when the package is initialized`, and it adds the cmdAdd command to the root command of the application.

    Once you've defined the add command, you can access it via the command line by typing `yourApplicationName add`.

## References

1. [Cobra GitHub](https://github.com/spf13/cobra)

<!-- https://www.phind.com/search?cache=wcz892kngyj1ntr5s2vjk4eo -->