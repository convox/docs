---
title: "racks"
description: "List all available Racks."
---

# racks

List all available Racks with their current status. When you are logged in to Console, Console Racks from every organization you have access to are listed as `<org>/<rack>`. Racks installed on this machine with [rack install](/reference/cli-commands/rack-install) are listed under their own name, and show `unknown` when the CLI cannot reach them. Select any listed Rack with [switch](/reference/cli-commands/switch).

## Syntax

```bash
$ convox racks
```

## Flags

*This command has no flags.*

## Example Usage

```bash
$ convox racks
NAME                  STATUS
myorg/production      running
myorg/staging         running
myorg/development     updating
sandbox               running
```

## See Also

- [rack](/reference/cli-commands/rack)
- [switch](/reference/cli-commands/switch)
- [rack access](/reference/cli-commands/rack-access)
- [login](/reference/cli-commands/login)
