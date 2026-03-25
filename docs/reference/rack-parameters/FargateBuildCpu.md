---
title: "FargateBuildCpu"
description: "Sets the CPU allocation for Fargate-based builds in a Convox Rack."
---

# FargateBuildCpu

How much CPU to reserve for the builder. Only used when [BuildMethod](/reference/rack-parameters/BuildMethod) is set to `fargate`.

| Default value  | "" |

## Use Cases

- Increase CPU allocation for builds that are CPU-intensive, such as compiling large codebases or running asset pipelines.
- Set to a lower value for lightweight builds to reduce Fargate costs.
- Tune alongside [FargateBuildMemory](/reference/rack-parameters/FargateBuildMemory) to find the optimal cost-performance balance for your build workloads.

## Additional Information

This parameter is only effective when [BuildMethod](/reference/rack-parameters/BuildMethod) is set to `fargate`. When using the default `ec2` build method, this parameter is ignored.

The value is specified in CPU units. Fargate supports specific CPU and memory combinations. Valid CPU values are `256` (0.25 vCPU), `512` (0.5 vCPU), `1024` (1 vCPU), `2048` (2 vCPU), and `4096` (4 vCPU). The available memory values depend on the CPU value selected. See [AWS Fargate task size](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#task_size) for valid combinations.

When left blank, Convox uses a default Fargate CPU allocation appropriate for general-purpose builds.

```bash
$ convox rack params set FargateBuildCpu=1024
```

## See Also

- [FargateBuildMemory](/reference/rack-parameters/FargateBuildMemory)
- [BuildMethod](/reference/rack-parameters/BuildMethod)
- [BuildCpu](/reference/rack-parameters/BuildCpu)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
