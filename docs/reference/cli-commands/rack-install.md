---
title: "rack install"
description: "Install a new Rack on a cloud provider."
---

# rack install

Install a new Rack. The type specifies the cloud provider (e.g., `aws`). Parameters can be set during installation to configure the Rack infrastructure.

## Syntax

```bash
$ convox rack install <type> [Parameter=Value]...
```

## Flags

| Flag | Short | Description |
|:-----|:------|:------------|
| `--name` | `-n` | Rack name |
| `--raw` | | Raw output |
| `--version` | `-v` | Rack version |

## Example Usage

```bash
$ convox rack install aws -n production region=us-east-1
Installing Rack production...
Creating IAM resources...
Creating VPC...
Creating EKS cluster...
Waiting for cluster to be ready...
Rack production installed successfully
```

## See Also

- [rack uninstall](/reference/cli-commands/rack-uninstall)
- [rack update](/reference/cli-commands/rack-update)
- [rack params set](/reference/cli-commands/rack-params-set)
- [Getting Started](/introduction/getting-started)
