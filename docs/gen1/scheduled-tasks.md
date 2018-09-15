---
title: "Scheduled Tasks"
---

Convox can set up `cron`-like recurring tasks on any of your application processes. This can be useful for background work like data dumps, batch jobs, or even queueing other background jobs for a worker.

#### Configuring tasks

Scheduled tasks are configured in `docker-compose.yml` using labels in the `convox.cron` namespace. The label format is:

`convox.cron.<task name>=<cron expression> <command>`

- **task name** is a unique name for each task that you choose. Task names may not be reused within the same application process and may only contain alphanumeric characters and dashes.

- **cron expression** describes the schedule on which the task will be invoked. See "Cron expression format" below for more info.

- **command** is the command to be run in this process. Before configuring the task you can test that your command works by running `convox run <process name> <command>`.

Example: to run the command `bin/myjob` every hour on the `web` process, you would configure the label like this:

```yaml
web:
  labels:
    - convox.cron.myjob=0 * * * ? bin/myjob
```

#### Cron expression format

Cron expressions use the following format. All times are UTC.

```
.----------------- minute (0 - 59)
|  .-------------- hour (0 - 23)
|  |  .----------- day-of-month (1 - 31)
|  |  |  .-------- month (1 - 12) OR JAN,FEB,MAR,APR ...
|  |  |  |  .----- day-of-week (1 - 7) OR SUN,MON,TUE,WED,THU,FRI,SAT
|  |  |  |  |
*  *  *  *  *
```

<div class="block-callout block-show-callout type-info" markdown="1">
  Using both day-of-week and day-of-month in the same expression is not supported. One of these fields must be a `?`.<br/>
  <br/>
  The actual start time of a job may be skewed up to 10 seconds. This is done to prevent infrastructure throttling errors when lots of jobs start at the same time. If you need to ensure that jobs are started in a specific order, make sure to start them at least 10 seconds apart.
</div>

Some example expressions:

<table>
  <tr>
    <th>Expression</th>
    <th>Meaning</th>
  </tr>
  <tr>
    <td><code>* * * * ?</code></td>
    <td>Run every minute</td>
  </tr>
  <tr>
    <td><code>*/10 * * * ?</code></td>
    <td>Run every 10 minutes</td>
  </tr>
  <tr>
    <td><code>0 * * * ?</code></td>
    <td>Run every hour</td>
  </tr>
  <tr>
    <td><code>30 6 * * ?</code></td>
    <td>Run at 6:30am UTC every day</td>
  </tr>
  <tr>
    <td><code>30 18 ? * MON-FRI</code></td>
    <td>Run at 6:30pm UTC every weekday</td>
  </tr>
  <tr>
    <td><code>0 12 1 * ?</code></td>
    <td>Run at noon on the first day of every month</td>
  </tr>
  <tr>
    <td><code>0 0,12 * * ?</code></td>
    <td>Run at Midnight and Noon every day</td>
  </tr>
</table>

See the [Scheduled Events](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) AWS documentation for more details.

## Run options and persistence

The service a scheduled task is associated with does not necessarily need running containers all the time. The service can be [scaled down to `-1` or `0`](/docs/gen1/scaling/#scaling-down-unused-services), and the scheduled task will "wake it up," causing a container to be created for the scheduled to task to run in, and exiting when it finishes.
