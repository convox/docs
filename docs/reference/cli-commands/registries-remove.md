---
title: "registries remove"
description: "Remove a private registry from the Rack."
---

# registries remove

Remove a private Docker registry from the Rack. The Rack will no longer be able to pull images from this registry.

## Syntax

```bash
$ convox registries remove <server>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox registries remove registry.example.com
Removing registry... OK
```

## See Also

- [registries](/reference/cli-commands/registries)
- [registries add](/reference/cli-commands/registries-add)
- [Private Registries](/integrations/private-registries)
