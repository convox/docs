---
title: "Rolling Updates"
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

## Deployment configuration

Rolling updates will respect the deployment configuration to control the minimum number of healthy processes and maximum number of overall processes to have running at any one time during the update.  This defaults to a minimum of 50% and a maximum of 200%.  This can be configured in your `convox.yml` for your service:

```yml
service:
  web:
    deployment:
      minimum: 25
      maximum: 125
```

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
