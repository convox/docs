---
title: "BuildMethod"
description: "Choose between EC2 and Fargate for the Convox Rack build process."
---

# BuildMethod

Build process type for the Rack. Controls whether builds run on dedicated EC2 instances or Fargate. `ec2` uses a dedicated build instance; `fargate` uses a Fargate task.

| Setting | Value |
|:--|:--|
| Default value  | `ec2`            |
| Allowed values | `ec2`, `fargate` |

## Use Cases

- Using `ec2` (default) for builds that need large disk space, fast I/O, or Docker layer caching across builds
- Switching to `fargate` to eliminate the cost of a continuously running build instance when builds are infrequent
- Using `fargate` when you want builds to run in an isolated, ephemeral environment for security reasons

## Additional Information

When using `ec2`, build resource allocation is controlled by [BuildCpu](/reference/rack-parameters/BuildCpu), [BuildMemory](/reference/rack-parameters/BuildMemory), and [BuildInstance](/reference/rack-parameters/BuildInstance).

When using `fargate`, the builder runs at 1024 CPU units and 4096 MB unless you override it. [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu) and [FargateBuildMemory](/reference/rack-parameters/FargateBuildMemory) set the builder size, and [FargateBuildVolumeSize](/reference/rack-parameters/FargateBuildVolumeSize) sets its disk. All three are optional.

By default, Fargate builds do not reuse Docker layers across builds. To cache layers across Fargate builds, enable the persistent [BuildCache](/reference/rack-parameters/BuildCache) parameter.

```bash
$ convox rack params set BuildMethod=fargate
```

> **Warning:** Generation 1 Apps cannot build on a Fargate builder. `BuildMethod` applies to the whole Rack, so there is no per-App opt out. A gen1 build on a Rack with `BuildMethod=fargate` fails with `exec: "docker": executable file not found in $PATH`, because the Fargate builder image is kaniko and gen1 builds require the `docker` binary. The Build is marked `failed`, no Release is created, and the App keeps serving its current Release. Keep `BuildMethod=ec2` on Racks that still run gen1 Apps.

### Downgrading a Fargate Rack

Set `FargateBuildCpu` and `FargateBuildMemory` explicitly before downgrading a Rack that builds on Fargate. Older Rack releases render a blank CPU value as empty, ECS refuses the task definition, and the update rolls back with `Invalid 'cpu' setting for task`. The rollback is clean: the Rack stays on the version it was running, its parameters are unchanged, and no App is disturbed.

```bash
$ convox rack params set FargateBuildCpu=1024 FargateBuildMemory=4096
```

## See Also

- [BuildCache](/reference/rack-parameters/BuildCache)
- [BuildInstance](/reference/rack-parameters/BuildInstance)
- [BuildCpu](/reference/rack-parameters/BuildCpu)
- [BuildMemory](/reference/rack-parameters/BuildMemory)
- [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu)
- [FargateBuildMemory](/reference/rack-parameters/FargateBuildMemory)
- [FargateBuildVolumeSize](/reference/rack-parameters/FargateBuildVolumeSize)
- [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize)
