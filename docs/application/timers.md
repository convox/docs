---
title: "Timers"
description: "Schedule recurring cron-like tasks on Convox application services using timer definitions in convox.yml."
---

# Timers

Convox can set up `cron`-like recurring tasks on any of your application processes. This can be useful for background work like data dumps, batch jobs, or even queueing other background jobs for a worker.

## Definition

```yaml
timers:
  cleanup:
    command: bin/cleanup
    schedule: "0 3 * * ? *"
    service: web
    policies:
      - arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess
```

### Attributes

| Attribute | Description | Required |
|-----------|-------------|----------|
| `command` | The command to run. | Yes |
| `schedule` | A cron expression defining when the timer runs (UTC). | Yes |
| `service` | The Service whose configuration (image, environment, resources) the timer uses. | Yes |
| `policies` | A list of IAM policy ARNs to attach to the timer's task role. If omitted, the timer inherits the policies of its parent `service`. | No |

You can think of a timer as issuing a `convox run` on the defined schedule. This timer would be equivalent to running `convox run web --detach bin/cleanup` at 3AM every day.

#### Cron expression format

Cron expressions use the following format. All times are UTC.

```text
.----------------- minute (0 - 59)
|  .-------------- hour (0 - 23)
|  |  .----------- day-of-month (1 - 31)
|  |  |  .-------- month (1 - 12) OR JAN,FEB,MAR,APR ...
|  |  |  |  .----- day-of-week (1 - 7) OR SUN,MON,TUE,WED,THU,FRI,SAT
|  |  |  |  |  .-- year (1970 - 2199)
|  |  |  |  |  |
*  *  *  *  *  *
```

> **Note:** Using both day-of-week and day-of-month in the same expression is not supported. One of these fields must be a `?`.

Some example expressions:

| Expression | Meaning |
|---|---|
| `* * * * ? *` | Run every minute |
| `*/10 * * * ? *` | Run every 10 minutes |
| `0 * * * ? *` | Run every hour |
| `30 6 * * ? *` | Run at 6:30am UTC every day |
| `30 18 ? * MON-FRI *` | Run at 6:30pm UTC every weekday |
| `0 12 1 * ? *` | Run at noon on the first day of every month |
| `0 0,12 * * ? *` | Run at Midnight and Noon every day |

See the [Scheduled Events](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) AWS documentation for more details.

## Examples

### Dedicated Service

Two services, `web` is normally running, `timers` is not (scaled to 0). The `cleanup` timer will spawn a new process using the configuration of `timers` once per minute, run the command `bin/cleanup` inside it, and terminate on completion.

```yaml
services:
  web:
    build: .
    command: bin/webserver
  timers:
    build: ./timers
    scale: 0
timers:
  cleanup:
    command: bin/cleanup
    schedule: "*/1 * * * ?"
    service: timers
```

### Existing Service

One service `web` is normally running. The `cleanup` timer will spawn a new process using the configuration of `web` one per minute, run the command `bin/cleanup` inside it, and terminate on completion.

```yaml
services:
  web:
    build: .
    command: bin/webserver
timers:
  cleanup:
    command: bin/cleanup
    schedule: "*/1 * * * ?"
    service: web
```

## See Also

- [Services](/application/services)
- [One-off Commands](/management/one-off-commands)
- [Running Migrations](/deployment/running-migrations)
- [convox.yml Reference](/application/convox-yml)
