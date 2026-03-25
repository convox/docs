---
title: "ScheduleRackScaleUp"
description: "Define a cron schedule for automatically starting the Convox Rack after a scheduled shutdown."
---

# ScheduleRackScaleUp

The recurring schedule for when the Rack will start back up. Use this parameter together with [ScheduleRackScaleDown](/reference/rack-parameters/ScheduleRackScaleDown) to turn the Rack on and off on a schedule. Both parameters must be set for scheduling to work.

| Default value  | "" |

The supported cron expression format consists of five fields separated by white spaces: `[Minute] [Hour] [Day_of_Month] [Month_of_Year] [Day_of_Week]`.

## Use Cases

- Starting development or staging Racks at the beginning of business hours each weekday
- Bringing QA environments online before the testing team begins work
- Scheduling Rack startup ahead of automated CI/CD pipelines that run on a schedule

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

- [ScheduleRackScaleDown](/reference/rack-parameters/ScheduleRackScaleDown)
- [InstanceCount](/reference/rack-parameters/InstanceCount)
- [Autoscale](/reference/rack-parameters/Autoscale)
