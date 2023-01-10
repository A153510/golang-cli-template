# ${{ values.cliCommandName }}

${{ values.description }}

## Adding Subcommands

To add a subcommand use the `cobra-cli` in the root of the project.

``` bash
cobra-cli add subcmd
```

To add a subcommand to another subcommand run the following:

``` bash
# Using the subcommand created in the previous example as the parent
cobra-cli add subsubcmd -p 'subcmdCmd'
```
