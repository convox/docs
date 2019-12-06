---
title: "AWS"
---

This page contains detailed information about how AWS is set up with convox. Some of this information may be useful when troubleshooting your apps' unexpected behavior, or simply to understand better what's underneath infrastructure looks like.

## ECS (EC2 launch type)

A [service](https://docs.convox.com/application/services)'s [memory limit](https://docs.convox.com/deployment/scaling) (set via `convox scale <service> --memory <memory>` [CLI](https://docs.convox.com/reference/cli-commands) command) translates to a <u>hard-only</u> memory limit on ECS task definition (see `memoryReservation` section on AWS's website [here](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_memory) to read more). This is true both for long-running processes, and for processes started via a [timer](https://docs.convox.com/application/timers).

If you were to execute a [one-off command](https://docs.convox.com/management/one-off-commands) however, the process created will have a <u>soft-only</u> memory limit with the same scale value. This _subtle_ difference may be important in understanding why sometimes a given command crashes when run either as a long-running process or a process started via a timer, but if executed via `convox run` it doesn't.

When hard-only limit is used, docker will kill your process if it were to run out of memory; whilst when soft-only memory is used your process can use all the memory available in the container instance, hence it may not crash even if the memory consumption is higher than the value set as scale value.
