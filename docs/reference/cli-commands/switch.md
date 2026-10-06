---
title: "switch"
description: "Switch the active Rack or display the name of the current Rack."
---

# switch

Switch the current Rack context for all subsequent CLI commands. When called without an argument, displays the name of the currently active Rack. The `-r` flag or the `CONVOX_RACK` environment variable selects a Rack for a single command and takes precedence over the switched Rack.

`convox switch` accepts any name that [racks](/reference/cli-commands/racks) lists, or part of one that matches a single Rack.

## Syntax

```bash
$ convox switch [rack]
```

## Flags

*This command has no flags.*

## Example Usage

```bash
$ convox switch myorg/production
Switched to myorg/production
```

```bash
$ convox switch
myorg/production
```

## Racks Installed with the CLI

A Rack installed with [rack install](/reference/cli-commands/rack-install) appears in [racks](/reference/cli-commands/racks) under its own name, without an organization prefix. Select it by that name with `convox switch`, `-r` or `CONVOX_RACK`. Commands for that Rack go directly to its API with the password saved at install time. When you are logged in to Console, the login stays in place, so Console Racks remain selectable by their `<org>/<rack>` name.

```bash
$ convox racks
NAME              STATUS
myorg/production  running
staging           running

$ convox switch staging
Switched to staging

$ convox switch myorg/production
Switched to myorg/production
```

| Case | Behavior |
|:-----|:---------|
| `-r` or `CONVOX_RACK` with the exact name of an installed Rack | Runs on that Rack |
| `-r` or `CONVOX_RACK` with any other name, including part of an installed Rack's name | Goes to the host you logged in to with [login](/reference/cli-commands/login) |
| `convox switch <name>` with the exact name of an installed Rack, when the `<org>/<rack>` name of a Console Rack contains `<name>` | Selects the installed Rack and prints a `NOTICE` on stderr naming those Console Racks; select one of them by its full `<org>/<rack>` name. Part of the name prints `ERROR: ambiguous rack name: <name>` instead |
| `CONVOX_HOST` set | Every command goes to that host, and `convox switch` to an installed Rack prints `ERROR: unset CONVOX_HOST to switch to <name>` without changing the selection |

Requires CLI version 20261005214736 or newer.

## See Also

- [login](/reference/cli-commands/login)
- [racks](/reference/cli-commands/racks)
- [Getting Started](/introduction/getting-started)
