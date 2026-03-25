---
title: "instances keyroll"
description: "Roll the SSH key on all instances."
---

# instances keyroll

Roll the SSH key on all instances. Generates a new key pair and distributes it across the Rack. This is a security best practice that should be performed periodically or after personnel changes.

## Syntax

```bash
$ convox instances keyroll
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |
| `--wait` | `-w` | Wait for completion |

## Example Usage

```bash
$ convox instances keyroll
Rolling instance key... OK
```

## See Also

- [instances](/reference/cli-commands/instances)
- [instances ssh](/reference/cli-commands/instances-ssh)
