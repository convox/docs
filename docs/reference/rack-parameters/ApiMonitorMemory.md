---
title: "ApiMonitorMemory"
description: "Configure the memory reserved by the Convox Rack API monitor process."
---

# ApiMonitorMemory

How much memory (in MB) should be reserved by the API monitor process. The monitor process handles background Rack operations such as monitoring ECS services and managing instance health.

| Default value  | `128` |

## Use Cases

- Increasing memory if the monitor process is being OOM-killed on busy Racks with many services
- Reducing memory on small development Racks to leave more capacity for application workloads

## Additional Information

This is an ECS hard memory limit. If the monitor process exceeds this amount, ECS will terminate and restart the container. If you see unexpected monitor restarts, consider increasing this value.

```bash
$ convox rack params set ApiMonitorMemory=256
```

## See Also

- [ApiWebMemory](/reference/rack-parameters/ApiWebMemory)
- [ApiCpu](/reference/rack-parameters/ApiCpu)
- [ApiCount](/reference/rack-parameters/ApiCount)
