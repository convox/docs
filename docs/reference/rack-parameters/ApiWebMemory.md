---
title: "ApiWebMemory"
description: "Configure the memory reserved by the Convox Rack API web process."
---

# ApiWebMemory

How much memory (in MB) should be reserved by the API web process. This controls the memory allocation for each Rack API container.

| Default value  | `256` |

## Use Cases

- Increasing memory if the API process is running out of memory on Racks with many applications or frequent deployments
- Reducing memory on small development Racks to free resources for application workloads

## Additional Information

This is an ECS hard memory limit. If the API web process exceeds this amount, ECS will terminate and restart the container. The total memory consumed by the Rack API is approximately `ApiCount * ApiWebMemory` plus the monitor process memory.

```bash
$ convox rack params set ApiWebMemory=512
```

## See Also

- [ApiCount](/reference/rack-parameters/ApiCount)
- [ApiCpu](/reference/rack-parameters/ApiCpu)
- [ApiMonitorMemory](/reference/rack-parameters/ApiMonitorMemory)
