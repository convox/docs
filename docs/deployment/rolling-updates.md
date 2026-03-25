---
title: "Rolling Updates"
description: "How Convox performs zero-downtime rolling deployments and automatic rollbacks."
---

# Rolling Updates

When a Release is promoted, new processes are gracefully rolled out into production.

If there are errors starting new processes, new processes are not [verified as healthy](/application/health-checks), or the rollout doesn't complete in 10 minutes, an automatic rollback is performed.

The number of new processes started at a time is configurable so large Apps can be fully rolled out in 10 minutes.

Rollouts and rollbacks do not cause any service downtime.

## Rollout

A rollout coordinates starting new processes in a way that maintains Service uptime and capacity. The basic flow is:

* Start 1 new process
* Verify new process is [healthy](/application/health-checks)
* Stop 1 old process
* Repeat for all processes in the formation

## Deployment Configuration

Rolling updates respect the deployment configuration to control the minimum number of healthy processes and maximum number of overall processes running at any one time during the update. These values are **percentages** of the desired count. The defaults are a minimum of 50% and a maximum of 200%.

You can configure these values in your `convox.yml`:

```yaml
services:
  web:
    deployment:
      minimum: 25
      maximum: 125
```

In this example, if the Service has a count of 4, a minimum of 25% means at least 1 process must remain healthy during the update, and a maximum of 125% means at most 5 processes can exist simultaneously.

> For Agent or Singleton Services, the defaults are 0% minimum and 100% maximum.

## Automatic Rollback

If there are errors starting new processes, new processes are not verified as healthy, or the rollout doesn't complete in 10 minutes, an automatic rollback is performed. This results in a `rollback` or `failed` state for the App:

```bash
$ convox apps info
Name       httpd
Status     rollback

$ convox apps info
Name       httpd
Status     failed
```

## Health Checks

For a rolling update to succeed, certain criteria must be met. For more information, see [Health Checks](/application/health-checks).

## See Also

- [Health Checks](/application/health-checks)
- [Releases](/deployment/releases)
- [Rolling Back](/deployment/rolling-back)
- [Services](/application/services)
