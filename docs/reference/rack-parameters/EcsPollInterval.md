---
title: "EcsPollInterval"
description: "Controls how frequently the Convox Rack polls ECS for service events."
---

# EcsPollInterval

How often (in seconds) to poll ECS for service updates (to inject into the app logs).

| Default value  | `1` |

## Use Cases

- Increase the polling interval if you experience ECS API rate limiting or throttling, especially on racks with many services.
- Reduce to sub-second if faster event visibility is needed (note: the minimum effective value is `1`).
- Set to a higher value (e.g., `5` or `10`) on large racks to reduce the number of ECS API calls and stay within AWS service limits.

## Additional Information

The Rack API periodically polls ECS to retrieve service events such as deployment updates, task start/stop events, and health check status changes. These events are then surfaced in your application logs via `convox logs`.

Longer intervals reduce the volume of API calls to ECS but may delay the appearance of service events in your logs. The default of `1` second provides near-real-time visibility.

```bash
$ convox rack params set EcsPollInterval=5
```

## See Also

- [EcsContainerStopTimeout](/reference/rack-parameters/EcsContainerStopTimeout)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
