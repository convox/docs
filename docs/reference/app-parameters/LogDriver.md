---
title: "LogDriver"
description: "Sets the log driver used by the application to send logs to CloudWatch, syslog, or disable logging entirely."
---

# LogDriver

Log driver for the application. When set to `Syslog`, you must also set [SyslogDestination](/reference/app-parameters/SyslogDestination). Set to blank (`""`) to disable logging entirely.

| Default value  | `CloudWatch` |
| Allowed values | `CloudWatch`, `Syslog`, `""` |

## Use Cases

- Keep as `CloudWatch` (default) for standard log management using AWS CloudWatch Logs and `convox logs`
- Set to `Syslog` when you need to forward logs to an external log aggregation service (e.g., Papertrail, Datadog, Loggly)
- Set to `""` (blank) to disable logging entirely for applications that handle their own log routing or generate excessive log volume

## Additional Information

When set to `CloudWatch`, logs are sent to a CloudWatch Logs log group and are accessible via `convox logs`. The retention period is controlled by the [LogRetention](/reference/app-parameters/LogRetention) parameter.

When set to `Syslog`, you must also configure [SyslogDestination](/reference/app-parameters/SyslogDestination) with the endpoint URL and [SyslogFormat](/reference/app-parameters/SyslogFormat) with the desired message format. Logs will not appear in `convox logs` when using syslog.

Setting the value to blank (`""`) disables all log collection. No CloudWatch log group is created and no syslog driver is configured.

The Rack also has a `LogDriver` parameter that controls logging for the Rack infrastructure itself. The app-level parameter only affects this application's services and timers.

```bash
$ convox apps params set LogDriver=Syslog
```

## See Also

- [SyslogDestination](/reference/app-parameters/SyslogDestination)
- [SyslogFormat](/reference/app-parameters/SyslogFormat)
- [LogRetention](/reference/app-parameters/LogRetention)
- [LogBucket](/reference/app-parameters/LogBucket)
- [Rack Parameter: LogDriver](/reference/rack-parameters/LogDriver)
- [Syslogs](/deployment/syslogs)
