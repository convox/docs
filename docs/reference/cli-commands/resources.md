---
title: "resources"
description: "List App-level Resources defined in convox.yml."
---

# resources

List the Resources associated with an App. These are the Resources defined in your `convox.yml` manifest, such as databases, caches, and queues.

## Syntax

```bash
$ convox resources
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox resources -a myapp
NAME      TYPE      URL
database  postgres  postgres://user:pass@host.example.com:5432/app
cache     redis     redis://host.example.com:6379/0
```

## See Also

- [resources info](/reference/cli-commands/resources-info)
- [resources proxy](/reference/cli-commands/resources-proxy)
- [resources url](/reference/cli-commands/resources-url)
- [rack resources](/reference/cli-commands/rack-resources)
- [Accessing Resources](/management/resources)
- [Resources](/application/resources)
