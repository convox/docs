---
title: "AWS Infrastructure Details"
description: "Details about the underlying AWS infrastructure Convox provisions, including ECS memory limits."
---

# AWS Infrastructure Details

This page contains detailed information about how AWS infrastructure is configured by Convox. This may be useful when troubleshooting unexpected App behavior or understanding what the underlying infrastructure looks like.

## ECS Memory Limits (EC2 Launch Type)

A [Service](/application/services)'s [memory limit](/scaling/scaling) (set via `convox scale <service> --memory <memory>`) translates to a **hard** memory limit on the ECS task definition. This applies both to long-running processes and to processes started via a [Timer](/application/timers).

If you execute a [one-off command](/management/one-off-commands), the process created has a **soft** memory limit instead. This difference may be important in understanding why a given command crashes when run as a long-running process or Timer, but succeeds via `convox run`.

When a hard limit is used, Docker kills the process if it exceeds its memory allocation. When a soft limit is used, the process can use all available memory on the container instance, so it may not crash even if memory consumption exceeds the configured scale value.

## See Also

- [Scaling](/scaling/scaling)
- [One-Off Commands](/management/one-off-commands)
- [Rack Parameters](/reference/rack-parameters)
- [Services](/application/services)
