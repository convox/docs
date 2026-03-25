---
title: "EcsContainerStopTimeout"
description: "Sets a custom timeout for graceful ECS container shutdowns in a Convox Rack."
---

# EcsContainerStopTimeout

Custom timeout duration for stopping ECS containers on Rack instances. This parameter defines the time (in seconds) ECS waits after sending a `SIGTERM` before issuing a `SIGKILL`, allowing for graceful shutdowns.

By default, this value is unset, meaning ECS will use its default 30-second stop timeout or any custom configuration already set at the ECS level.

| Default value  | "" |
| Allowed values | Numerical values in seconds (e.g., `10`, `60`, `120`) |

## Use Cases

- Increase the timeout for applications that need additional time to drain active connections, flush caches, or complete in-flight requests during shutdown.
- Set a shorter timeout for stateless workers that can be stopped quickly, reducing deployment time.
- Tune shutdown behavior for applications with complex cleanup processes, such as closing database connections or persisting session state.

## Additional Information

When ECS stops a container, it first sends a `SIGTERM` signal to the container's main process. The container then has the configured timeout period to shut down gracefully before ECS sends a `SIGKILL` to force termination. If your application does not handle `SIGTERM`, it will be forcibly killed after the timeout expires.

This parameter maps to the `ECS_CONTAINER_STOP_TIMEOUT` ECS agent configuration variable and is written to `/etc/ecs/ecs.config` on each EC2 instance. Because it is an agent-level setting, it applies to **all containers on the instance**, not individual services. This parameter has no effect on Fargate tasks, which use the `stopTimeout` setting in the task definition instead.

For more details, see the [ECS Agent Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html) documentation.

```bash
$ convox rack params set EcsContainerStopTimeout=60
```

## See Also

- [EcsPollInterval](/reference/rack-parameters/EcsPollInterval)
- [Rack Parameters](/reference/rack-parameters) for a full list of available parameters
