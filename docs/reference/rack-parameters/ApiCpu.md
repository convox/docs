---
title: "ApiCpu"
description: "Configure the CPU units reserved by the Convox Rack API web process."
---

# ApiCpu

How much CPU should be reserved by the API web process. The value is specified in ECS CPU units, where 1024 units equals one full vCPU.

| Default value  | `128` |

## Use Cases

- Increasing CPU allocation if the Rack API is responding slowly under heavy deployment activity
- Reducing CPU allocation on small development Racks to leave more capacity for application containers

## Additional Information

This value is an ECS CPU reservation. It does not hard-cap the API process to this amount of CPU -- it guarantees at least this much CPU is available when the instance is under contention.

```bash
$ convox rack params set ApiCpu=256
```

## See Also

- [ApiCount](/reference/rack-parameters/ApiCount)
- [ApiWebMemory](/reference/rack-parameters/ApiWebMemory)
- [ApiMonitorMemory](/reference/rack-parameters/ApiMonitorMemory)
