---
title: "rack resources info"
description: "Get information about a Rack Resource."
---

# rack resources info

Display detailed information about a Rack Resource, including its name, type, status, configurable options, connection URL, and any linked Apps.

## Syntax

```bash
$ convox rack resources info <resource>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack resources info shared-postgres
Name    shared-postgres
Type    postgres
Status  running
Options MultiAZ=true
URL     postgres://user:pass@host.example.com:5432/app
Apps    myapp, api
```

## See Also

- [rack resources](/reference/cli-commands/rack-resources)
- [rack resources url](/reference/cli-commands/rack-resources-url)
- [rack resources options](/reference/cli-commands/rack-resources-options)
- [rack resources update](/reference/cli-commands/rack-resources-update)
- [Rack and External Resources](/management/rack-resources)
