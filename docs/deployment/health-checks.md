---
title: "Health Checks"
order: 350
---

By default, for a deployment to succeed, it must pass certain "health checks."

If the Process does not [expose a port](/docs/port-mapping) it is considered healthy if it starts and doesn't immediately exit or crash.

If the Process [exposes a port](/docs/port-mapping) is it considered healthy after it:

* Registers behind a load balancer
* Passes a certain number of HTTP health checks

Common causes for not passing health checks are:

* The cluster does not have sufficient memory or CPU resources available to start a new process
* The cluster does not have sufficient instances where a new process port is not already reserved by an older release
* A process crashes immediately after starting due to a problem in the latest code
* A process takes too long to initialize its server and therefore fails a network health check
* Returning a status code outside the acceptable (200-399,401) range


#### Health Check Options

```yaml
services:
  web:
  health:
    grace: 5
    interval: 5
    path: /health
    timeout: 3
```
<table>
  <tr>
    <th>Label</th>
    <th>Notes</th>
  </tr>
  <tr>
    <td><code>grace</code></td>
    <td>The amount of time to wait for a service to boot before beginning health checks.</td>
  </tr>
  <tr>
    <td><code>interval</code></td>
    <td>The amount of time between health checks (default 5 seconds).</td>
  </tr>
  <tr>
    <td><code>path</code></td>
    <td>The HTTP endpoint the load balancer will use to determine the application's health.</td>
  </tr>
  <tr>
    <td><code>timeout</code></td>
    <td>The time in seconds after which no response means a failed health check (defaults to <code>interval</code> minus 1).</td>
  </tr>
</table>

## See also

* [Rolling Updates](/docs/rolling-updates/)
