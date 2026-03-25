---
title: "registries"
description: "List private registries configured on the Rack."
---

# registries

List the private Docker registries configured on the Rack. Shows the server URL and username for each registered registry.

## Syntax

```bash
$ convox registries
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox registries
SERVER                        USERNAME
registry.example.com          deploy
docker.io                     myorg
```

## See Also

- [registries add](/reference/cli-commands/registries-add)
- [registries remove](/reference/cli-commands/registries-remove)
- [Private Registries](/integrations/private-registries)
