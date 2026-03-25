---
title: "ScheduleRackScaleDown"
description: "Define a cron schedule for automatically shutting down the Convox Rack to save costs during off-hours."
---

# ScheduleRackScaleDown

The recurring schedule for when the Rack will be shut down. Use this parameter together with [ScheduleRackScaleUp](/reference/rack-parameters/ScheduleRackScaleUp) to turn the Rack on and off on a schedule. Both parameters must be set for scheduling to work.

| Default value  | "" |

The supported cron expression format consists of five fields separated by white spaces: `[Minute] [Hour] [Day_of_Month] [Month_of_Year] [Day_of_Week]`.

## Use Cases

- Shutting down development or staging Racks on evenings and weekends to reduce AWS costs
- Powering off QA environments after business hours when testing is not being performed
- Scheduling regular downtime for non-production Racks to enforce cost governance

## Additional Information

To turn your Rack off on weekends and back on during weekdays:

```bash
$ convox rack params set ScheduleRackScaleDown="0 18 * * 5" ScheduleRackScaleUp="0 9 * * 1"
```

In the example above, the Rack is configured to shut down every Friday (5th day) at 6pm UTC and start back up every Monday (1st day) at 9am UTC.

More details on the cron format can be found at [Crontab](https://crontab.org/) and [cron examples](https://crontab.guru/examples.html). You can see details about Scheduling Actions in the [AWS Auto Scaling documentation](https://docs.aws.amazon.com/autoscaling/ec2/userguide/schedule_time.html).

> **Note:** All schedule times are in UTC. Be sure to account for your local timezone offset when configuring schedules.

> **Warning:** Both `ScheduleRackScaleDown` and `ScheduleRackScaleUp` must be set for scheduling to function. Setting only one parameter will not produce the desired behavior.

## See Also

- [ScheduleRackScaleUp](/reference/rack-parameters/ScheduleRackScaleUp)
- [InstanceCount](/reference/rack-parameters/InstanceCount)
- [Autoscale](/reference/rack-parameters/Autoscale)
