---
title: "Health Checks"
description: "Configure HTTP health checks that determine deployment success for Convox services exposing a port."
---

# Health Checks

The Convox deployment process depends on application Health Checks to determine success.

If the Process [exposes a port](/application/port) is it considered healthy after it passes a certain number of HTTP health checks.

## Definition

```yaml
services:
  web:
    health:
      grace: 5
      interval: 5
      path: /
      timeout: 4
```

### Options

| Option | Description | Default |
|--------|-------------|---------|
| `grace` | The amount of time in seconds to wait for a service to boot before beginning health checks. | Same as `interval` (5 seconds) |
| `interval` | The amount of time in seconds between health checks. | `5` |
| `path` | The HTTP endpoint the load balancer will use to determine the application's health. | `/` |
| `timeout` | The time in seconds after which no response means a failed health check. | `interval - 1` (4 seconds) |

## Common Failures

* The cluster does not have sufficient memory or CPU resources available to start a new process
* The cluster does not have sufficient instances where a new process port is not already reserved by an older release
* A process crashes immediately after starting due to a problem in the latest code
* A process takes too long to initialize its server and therefore fails a network health check
* Returning a status code outside the acceptable (200-399,401) range



## See Also

- [Rolling Updates](/deployment/rolling-updates)
- [Load Balancing](/networking/load-balancing)
