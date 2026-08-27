---
title: "FargateBuildMemory"
description: "Sets the memory allocation for Fargate-based builds in a Convox Rack."
---

# FargateBuildMemory

How much memory to reserve for the builder. Only used when [BuildMethod](/reference/rack-parameters/BuildMethod) is set to `fargate`.

| Default value  | "" |

## Use Cases

- Increase memory allocation for builds that consume large amounts of memory, such as JavaScript bundling, Java compilation, or Docker image builds with large contexts.
- Set to a lower value for simple builds to minimize Fargate costs.
- Tune alongside [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu) to match valid Fargate task size combinations.

## Additional Information

This parameter is only effective when [BuildMethod](/reference/rack-parameters/BuildMethod) is set to `fargate`. When using the default `ec2` build method, this parameter is ignored.

The value is specified in MB. Fargate requires specific CPU and memory combinations. For example, with `1024` CPU units (1 vCPU), valid memory values range from `2048` to `8192`. See [AWS Fargate task size](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#task_size) for the full list of valid combinations.

When left blank, the builder runs at `4096` MB.

```bash
$ convox rack params set FargateBuildMemory=4096
```

AWS enforces the pairing, so a memory value that is not legal for the current CPU value is rejected even though each value is valid on its own. Setting [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu) to `4096` while this parameter stays at `4096` fails with:

```
No Fargate configuration exists for given values: 4096 CPU, 4096 memory
```

The Rack rolls back cleanly and keeps running on its previous parameters, but the change does not apply. Set both values together to move between task sizes:

```bash
$ convox rack params set FargateBuildCpu=4096 FargateBuildMemory=8192
```

## See Also

- [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu)
- [FargateBuildVolumeSize](/reference/rack-parameters/FargateBuildVolumeSize)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
- [BuildMemory](/reference/rack-parameters/BuildMemory)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
