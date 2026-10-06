---
title: "rack install"
description: "Install a new Rack on a cloud provider."
---

# rack install

Install a new Rack. The type specifies the cloud provider (e.g., `aws`). Parameters can be set during installation to configure the Rack infrastructure.

The CLI saves the new Rack's API host and password on this machine. [racks](/reference/cli-commands/racks) then lists the Rack under its name, and `convox switch <name>` or `-r <name>` selects it (see [switch](/reference/cli-commands/switch#racks-installed-with-the-cli)). If the CLI is not logged in yet, it also logs in to the new Rack.

## Syntax

```bash
$ convox rack install <type> [Parameter=Value]...
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--name` | `-n` | Rack name (default: `convox`) |
| `--raw` | | Raw output |
| `--version` | `-v` | Rack version |

## Example Usage

```bash
$ convox rack install aws -n production InstanceType=t3.medium
Preparing... OK
Installing...  98.91% 7m17s
Starting... OK, rack.production-1234567890.us-east-1.convox.site
```

## See Also

- [rack uninstall](/reference/cli-commands/rack-uninstall)
- [rack update](/reference/cli-commands/rack-update)
- [rack params set](/reference/cli-commands/rack-params-set)
- [Getting Started](/introduction/getting-started)
