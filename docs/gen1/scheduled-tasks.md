---
title: "Scheduled Tasks"
description: "Gen 1 (End of Life): How to configure cron-like scheduled tasks for Gen 1 Convox applications using Docker Compose labels."
---

# Scheduled Tasks

> **This page documents Generation 1, which has reached End of Life.** Gen 1 apps use `docker-compose.yml`. For current documentation, see [Timers](/application/timers).

Convox can set up `cron`-like recurring tasks on any of your application processes. This can be useful for background work like data dumps, batch jobs, or even queueing other background jobs for a worker.

## Configuring tasks

Scheduled tasks are configured in `docker-compose.yml` using labels in the `convox.cron` namespace. The label format is:

`convox.cron.<task name>=<cron expression> <command>`

- **task name** is a unique name for each task that you choose. Task names may not be reused within the same application process. A name starts with a letter, contains only letters, digits and dashes, and is 4 to 30 characters long.

- **cron expression** describes the schedule on which the task will be invoked. See "Cron expression format" below for more info.

- **command** is the command to be run in this process. Before configuring the task you can test that your command works by running `convox run <process name> <command>`.

Example: to run the command `bin/myjob` every hour on the `web` process, you would configure the label like this:

```yaml
web:
  labels:
    - convox.cron.myjob=0 * * * ? bin/myjob
```

## Cron expression format

Cron expressions use the following format. All times are UTC.

```text
.----------------- minute (0 - 59)
|  .-------------- hour (0 - 23)
|  |  .----------- day-of-month (1 - 31)
|  |  |  .-------- month (1 - 12) OR JAN,FEB,MAR,APR ...
|  |  |  |  .----- day-of-week (1 - 7) OR SUN,MON,TUE,WED,THU,FRI,SAT
|  |  |  |  |
*  *  *  *  *
```

> **Note:** One of the day-of-month and day-of-week fields must be `?`. An expression cannot set both.
>
> Each run starts after a random delay of up to 10 seconds, so jobs scheduled for the same minute can start in any order. To run jobs in a specific order, schedule them in different minutes.

Some example expressions:

| Expression | Meaning |
|:--|:--|
| `* * * * ?` | Run every minute |
| `*/10 * * * ?` | Run every 10 minutes |
| `0 * * * ?` | Run every hour |
| `30 6 * * ?` | Run at 6:30am UTC every day |
| `30 18 ? * MON-FRI` | Run at 6:30pm UTC every weekday |
| `0 12 1 * ?` | Run at noon on the first day of every month |
| `0 0,12 * * ?` | Run at Midnight and Noon every day |

See the [Scheduled Events](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) AWS documentation for more details.

## Run options and persistence

The service a scheduled task is associated with does not necessarily need running containers all the time. The service can be [scaled down to `-1` or `0`](/gen1/scaling#scaling-down-unused-services), and the scheduled task will "wake it up," causing a container to be created for the scheduled to task to run in, and exiting when it finishes.

## See Also

- [Timers (Gen 2)](/application/timers)
- [Docker Compose Labels](/gen1/docker-compose-labels)
- [Scaling](/gen1/scaling)
