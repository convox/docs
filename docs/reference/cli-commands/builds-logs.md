---
title: "builds logs"
description: "Get the logs for a Build, including Docker build output and errors."
---

# builds logs

Get the logs for a Build. Shows the Docker build output, layer caching information, and any errors that occurred during the Build. This is useful for debugging failed Builds.

## Syntax

```bash
$ convox builds logs <build>
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--app` | `-a` | App name |
| `--rack` | `-r` | Rack name |

## Example Usage

```bash
$ convox builds logs BABCDEF -a myapp
Running: docker build -t 9836064b .
Sending build context to Docker daemon  4.096kB
Step 1/4 : FROM node:18-alpine
 ---> 2a5b3c4d5e6f
Step 2/4 : WORKDIR /app
 ---> Using cache
Step 3/4 : COPY . .
 ---> 7f8a9b0c1d2e
Step 4/4 : CMD ["node", "server.js"]
 ---> 3e4f5a6b7c8d
Successfully built 3e4f5a6b7c8d
```

## See Also

- [builds info](/reference/cli-commands/builds-info)
- [builds export](/reference/cli-commands/builds-export)
- [builds](/reference/cli-commands/builds)
- [logs](/reference/cli-commands/logs)
- [Builds](/deployment/builds)
