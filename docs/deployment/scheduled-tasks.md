---
title: "Scheduled Tasks"
order: 450
---

Convox can set up `cron`-like recurring tasks on any of your application processes. This can be useful for background work like data dumps, batch jobs, or even queueing other background jobs for a worker.

#### Configuring timers

```yaml
timers:
  cleanup:
    command: bin/cleanup
    schedule: "0 3 * * ? *"
    service: web
```

You can think of a timer as issuing a `convox run` on the defined schedule. This timer would be equivalent to running `convox run web bin/cleanup` at 3AM every day.

#### Cron expression format

Cron expressions use the following format. All times are UTC.

```
.----------------- minute (0 - 59)
|  .-------------- hour (0 - 23)
|  |  .----------- day-of-month (1 - 31)
|  |  |  .-------- month (1 - 12) OR JAN,FEB,MAR,APR ...
|  |  |  |  .----- day-of-week (1 - 7) OR SUN,MON,TUE,WED,THU,FRI,SAT
|  |  |  |  |  .-- year (1970 - 2199)
|  |  |  |  |  |
*  *  *  *  *  *
```

<div class="block-callout block-show-callout type-info" markdown="1">
  Using both day-of-week and day-of-month in the same expression is not supported. One of these fields must be a `?`.<br/>
</div>

Some example expressions:

<table>
  <tr>
    <th>Expression</th>
    <th>Meaning</th>
  </tr>
  <tr>
    <td><code>* * * * ? *</code></td>
    <td>Run every minute</td>
  </tr>
  <tr>
    <td><code>*/10 * * * ? *</code></td>
    <td>Run every 10 minutes</td>
  </tr>
  <tr>
    <td><code>0 * * * ? *</code></td>
    <td>Run every hour</td>
  </tr>
  <tr>
    <td><code>30 6 * * ? *</code></td>
    <td>Run at 6:30am UTC every day</td>
  </tr>
  <tr>
    <td><code>30 18 ? * MON-FRI *</code></td>
    <td>Run at 6:30pm UTC every weekday</td>
  </tr>
  <tr>
    <td><code>0 12 1 * ? *</code></td>
    <td>Run at noon on the first day of every month</td>
  </tr>
  <tr>
    <td><code>0 0,12 * * ? *</code></td>
    <td>Run at Midnight and Noon every day</td>
  </tr>
</table>

See the [Scheduled Events](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) AWS documentation for more details.

### Examples

#### Dedicated Service

Two services, `web` is normally running, `timers` is not (scaled to 0). The `cleanup` timer will spawn a new process using the configuration of `timers` once per minute, run the command `bin/cleanup` inside it, and terminate on completion.

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
        schedule: */1 * * * ?
        service: timers

#### Existing Service

One service `web` is normally running. The `cleanup` timer will spawn a new process using the configuration of `web` one per minute, run the command `bin/cleanup` inside it, and terminate on completion.

    services:
      web:
        build: .
        command: bin/webserver
    timers:
      cleanup:
        command: bin/cleanup
        schedule: */1 * * * ?
        service: web
