---
title: "rack runtime attach"
description: "Attach a runtime integration to the Rack."
---

# rack runtime attach

Attach a runtime integration to the Rack. Connects an external runtime environment for use by Apps deployed on the Rack.

## Syntax

```bash
$ convox rack runtime attach <runtime>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack runtime attach lambda
OK
```

## See Also

- [rack runtimes](/reference/cli-commands/rack-runtimes)
- [rack](/reference/cli-commands/rack)
