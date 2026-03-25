---
title: "rack access"
description: "Generate temporary, scoped access credentials for a Rack."
---

# rack access

Generate a temporary, scoped `RACK_URL` credential for a Rack. The command issues a JWT-based URL with either `read` or `write` role access for a configurable duration (in hours). This is useful for troubleshooting, auditing, granting temporary access to team members, or providing credentials to CI/CD pipelines.

## Syntax

```bash
$ convox rack access
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--duration-in-hours` | | Duration in hours |
| `--rack` | `-r` | Rack name |
| `--role` | | Access role: `read` or `write` |

## Example Usage

```bash
$ convox rack access --role read --duration-in-hours 8
RACK_URL=https://jwt:eyJhbGciOi...@rack.example.org
```

## See Also

- [rack access key rotate](/reference/cli-commands/rack-access-key-rotate)
- [login](/reference/cli-commands/login)
- [racks](/reference/cli-commands/racks)
- [rack](/reference/cli-commands/rack)
- [Access Control](/console/access-control)
