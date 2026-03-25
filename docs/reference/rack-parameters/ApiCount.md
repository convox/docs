---
title: "ApiCount"
description: "Set the number of Rack API containers to run for improved availability."
---

# ApiCount

How many Rack API containers to run. Setting this higher than 2 will guarantee better Rack API availability for mission critical clusters.

| Default value  | `2` |

## Use Cases

- Increasing to 3 or more for production Racks where API downtime is unacceptable
- Reducing to 1 for development or staging Racks to save resources
- Scaling up temporarily during large deployment operations

## Additional Information

Each API container consumes the CPU and memory defined by [ApiCpu](/reference/rack-parameters/ApiCpu) and [ApiWebMemory](/reference/rack-parameters/ApiWebMemory). Ensure your instances have enough capacity to run the desired number of API containers alongside your application workloads.

```bash
$ convox rack params set ApiCount=3
```

## See Also

- [ApiCpu](/reference/rack-parameters/ApiCpu)
- [ApiWebMemory](/reference/rack-parameters/ApiWebMemory)
- [ApiMonitorMemory](/reference/rack-parameters/ApiMonitorMemory)
- [ApiRouter](/reference/rack-parameters/ApiRouter)
