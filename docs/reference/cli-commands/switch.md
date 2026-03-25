---
title: "switch"
description: "Switch the active Rack or display the name of the current Rack."
---

# switch

Switch the current Rack context for all subsequent CLI commands. When called without an argument, displays the name of the currently active Rack.

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

## See Also

- [login](/reference/cli-commands/login)
- [racks](/reference/cli-commands/racks)
- [Getting Started](/introduction/getting-started)
