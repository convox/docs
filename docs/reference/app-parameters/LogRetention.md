---
title: "LogRetention"
description: "Defines how many days CloudWatch logs are retained for the application."
---

# LogRetention

Log retention period for the application's CloudWatch Log Groups.

| Default value  | `7` |

## Use Cases

- Increase to `30`, `60`, or `90` days when you need longer log history for debugging or compliance
- Set to blank (`""`) for unlimited retention when regulatory requirements mandate indefinite log storage
- Keep at `7` (default) to balance log availability with CloudWatch storage costs

## Additional Information

Setting this to a high or unlimited value may affect the performance of `convox logs`.
For long-term log storage, consider using [syslog](/deployment/syslogs).

This parameter only applies when [LogDriver](/reference/app-parameters/LogDriver) is set to `CloudWatch`. When the log driver is set to `Syslog` or blank, no CloudWatch log group is created and this parameter has no effect.

The value is passed directly to the CloudWatch Logs log group's `RetentionInDays` property. When blank, CloudWatch retains logs indefinitely (no expiration).

The Rack also has a `LogRetention` parameter that controls retention for Rack-level logs. The app-level parameter only affects this application's log group.

```bash
$ convox apps params set LogRetention=30
```

## See Also

- [LogDriver](/reference/app-parameters/LogDriver)
- [LogBucket](/reference/app-parameters/LogBucket)
- [SyslogDestination](/reference/app-parameters/SyslogDestination)
- [Rack Parameter: LogRetention](/reference/rack-parameters/LogRetention)
- [Syslogs](/deployment/syslogs)
