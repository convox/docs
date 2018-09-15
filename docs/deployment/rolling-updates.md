---
title: "Rolling Updates"
order: 400
---

When a Release is promoted, new processes are gracefully rolled out into production.

If there are errors starting new processes, new processes are not [verified as healthy](/docs/health-checks), or the rollout doesn't complete in 10 minutes, an automatic rollback is performed.

The number of new processes started at a time is configurable so large apps can be fully rolled out in 10 minutes.

Rollouts and rollbacks do not cause any service downtime.

#### Rollout

A rollout coordinates starting new processes in a way that maintains service uptime and capacity. The basic flow is:

* Start 1 new process
* Verify new process is [healthy](/docs/health-checks)
* Stop 1 old process
* Repeat for all processes in the formation

#### Automatic Rollback

If there are errors starting new processes, new processes are not verified as healthy, or the rollout doesn't complete in 10 minutes, an automatic rollback is performed. This will result in a [`rollback`](/docs/rack-statuses/#rollback) or `failed` state for the app:

```
$ convox apps info
Name       httpd
Status     rollback

$ convox apps info
Name       httpd
Status     failed
```

#### Health Checks

For a rolling update to succeed, certain criteria must be met. For more information, see [Health Checks](/docs/health-checks).
