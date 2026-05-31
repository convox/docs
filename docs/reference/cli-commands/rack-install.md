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
$ convox rack install aws -n production InstanceType=t3.medium
Installing Rack production...
Creating CloudFormation stack...
2025-01-15T12:00:00Z system/cloudformation aws/cfm production CREATE_IN_PROGRESS AWS::CloudFormation::Stack User Initiated
2025-01-15T12:02:00Z system/cloudformation aws/cfm production CREATE_IN_PROGRESS AWS::EC2::VPC
2025-01-15T12:05:00Z system/cloudformation aws/cfm production CREATE_IN_PROGRESS AWS::ECS::Cluster
2025-01-15T12:10:00Z system/cloudformation aws/cfm production CREATE_COMPLETE AWS::CloudFormation::Stack
Rack production installed successfully
```

## See Also

- [rack uninstall](/reference/cli-commands/rack-uninstall)
- [rack update](/reference/cli-commands/rack-update)
- [rack params set](/reference/cli-commands/rack-params-set)
- [Getting Started](/introduction/getting-started)
