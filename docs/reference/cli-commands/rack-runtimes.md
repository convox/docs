---
title: "rack runtimes"
description: "List attachable runtime integrations."
---

# rack runtimes

List attachable runtime integrations. Shows available runtime environments that can be connected to the Rack for use by Apps.

## Syntax

```bash
$ convox rack runtimes
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox rack runtimes
ID      TITLE
lambda  AWS Lambda
fargate AWS Fargate
```

## See Also

- [rack runtime attach](/reference/cli-commands/rack-runtime-attach)
- [rack](/reference/cli-commands/rack)
