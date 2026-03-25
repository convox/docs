---
title: "registries add"
description: "Add a private registry to the Rack."
---

# registries add

Add a private Docker registry to the Rack. Once added, the Rack can pull images from this registry during Builds and deployments.

## Syntax

```bash
$ convox registries add <server> <username> <password>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox registries add registry.example.com deploy s3cr3t
Adding registry... OK
```

## See Also

- [registries](/reference/cli-commands/registries)
- [registries remove](/reference/cli-commands/registries-remove)
- [Private Registries](/integrations/private-registries)
