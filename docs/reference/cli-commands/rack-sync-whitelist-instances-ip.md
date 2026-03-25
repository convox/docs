---
title: "rack sync whitelist instances ip"
description: "Sync whitelisted instance IPs in the security group."
---

# rack sync whitelist instances ip

Sync whitelisted instance IPs in the security group. Ensures all current instance IPs are allowed in the Rack security group configuration, keeping access rules up to date with infrastructure changes.

## Syntax

```bash
$ convox rack sync whitelist instances ip
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack sync whitelist instances ip
OK
```

## See Also

- [rack](/reference/cli-commands/rack)
- [rack params](/reference/cli-commands/rack-params)
- [Rack Parameters](/reference/rack-parameters)
