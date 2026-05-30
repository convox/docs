---
title: "ContainerInsights"
description: "Enable CloudWatch Container Insights on a Convox Rack's ECS cluster."
---

# ContainerInsights

Enables CloudWatch Container Insights on the Rack's ECS cluster. When set to `Yes`, the cluster reports CPU, memory, network, and task-level metrics to the CloudWatch Container Insights dashboards. The setting is applied in place on the existing cluster, which is never replaced.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

## Use Cases

- Fleet-wide CPU and memory utilization across every Service on the Rack from a single dashboard
- Per-task and per-Service metrics for capacity planning and right-sizing
- Network throughput and task lifecycle visibility without instrumenting each App

## Additional Information

When `ContainerInsights=Yes`, Convox sets the `containerInsights` cluster setting to `enabled` on the Rack's `AWS::ECS::Cluster` resource. This is an in-place property update, so the cluster and all running tasks are untouched while the setting flips.

```bash
$ convox rack params set ContainerInsights=Yes
```

Container Insights metrics begin flowing within a few minutes of the cluster update completing. View them in the CloudWatch console under Container Insights, or query them in the `ECS/ContainerInsights` metric namespace.

Container Insights covers Fargate workloads automatically. EC2-launched Services additionally require the CloudWatch agent running as a daemon to report Container Insights metrics. That agent is not installed by this parameter.

This parameter belongs to the `logging` parameter group. List the whole group at once:

```bash
$ convox rack params -g logging
```

### Cost

Container Insights is a paid CloudWatch feature. Expect roughly 15 to 25 USD per month per Rack, varying with task count and metric volume. Review the CloudWatch pricing page for current rates before enabling on a large Rack.

### Disable

Set `ContainerInsights=No` to revert the cluster setting to `disabled`. Metric ingestion stops and stored metrics age out per CloudWatch retention. Downgrading to a Rack version that predates this parameter has the same effect: the cluster setting reverts to the AWS default (disabled), which is downgrade safe.

## See Also

- [Application Monitoring](/management/application)
- [LogDriver](/reference/rack-parameters/LogDriver)
- [Logs](/management/logs)
- [Rack Parameters](/reference/rack-parameters)
