---
title: "rack params"
description: "Display Rack parameters and their current values."
---

# rack params

Display Rack parameters and their current values. Parameters control infrastructure, networking, scaling, and security settings for the Rack.

## Syntax

```bash
$ convox rack params
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack params
BuildInstance
CertDuration          2160
HighAvailability      true
IdleTimeout           3600
InstanceCount         3
InstanceType          t3.medium
Internal              false
NodeDisk              20
Ssl                   true
```

## See Also

- [rack params set](/reference/cli-commands/rack-params-set)
- [apps params](/reference/cli-commands/apps-params)
- [Rack Parameters](/reference/rack-parameters)
- [rack](/reference/cli-commands/rack)
