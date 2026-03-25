---
title: "rack resources types"
description: "List available Resource types for the Rack."
---

# rack resources types

List the kinds of Resources that can be created on the Rack. Use this to discover which Resource types are supported by your Rack provider.

## Syntax

```bash
$ convox rack resources types
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack resources types
TYPE
memcached
mysql
postgres
redis
sns
sqs
```

## See Also

- [rack resources create](/reference/cli-commands/rack-resources-create)
- [rack resources options](/reference/cli-commands/rack-resources-options)
- [Rack and External Resources](/management/rack-resources)
