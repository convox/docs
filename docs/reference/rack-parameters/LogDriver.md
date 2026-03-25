---
title: "LogDriver"
description: "Configure the log driver used by the Rack and services to send logs."
---

# LogDriver

Log driver used by the Rack and services to send logs. Defaults to CloudWatch. You must provide the [SyslogDestination](/reference/rack-parameters/SyslogDestination) when setting this to `Syslog`. Setting this to blank disables logging.

**Attention:** Disabling CloudWatch will impact `convox logs` and `convox rack logs`. Use the Syslog resource if you still want to use convox logs. See [Resource Syslog](/deployment/syslogs) for more information.

| Default value  | `CloudWatch` |
| Allowed values | `CloudWatch`, `Syslog`, "" |

## Use Cases

- Switching to `Syslog` to forward logs to an external log aggregation service such as Datadog, Papertrail, or Splunk
- Disabling logging entirely (setting to blank) in development or test environments to reduce CloudWatch costs
- Keeping the default `CloudWatch` for production environments where `convox logs` integration is needed

## Additional Information

When set to `CloudWatch` (the default), logs from all Rack services and applications are sent to CloudWatch Logs. This enables the `convox logs` and `convox rack logs` CLI commands.

When set to `Syslog`, logs are forwarded to the destination specified by the [SyslogDestination](/reference/rack-parameters/SyslogDestination) parameter. You must set `SyslogDestination` before or at the same time as switching to `Syslog`.

When set to blank, logging is disabled entirely. This can significantly reduce CloudWatch costs but means you lose the ability to view logs through the Convox CLI or CloudWatch console.

```bash
$ convox rack params set LogDriver=Syslog SyslogDestination=tcp+tls://logs.example.com:514
```

## See Also

- [LogRetention](/reference/rack-parameters/LogRetention)
- [LogBucket](/reference/rack-parameters/LogBucket)
- [SyslogDestination](/reference/rack-parameters/SyslogDestination)
- [SyslogFormat](/reference/rack-parameters/SyslogFormat)
- [Resource Syslog](/deployment/syslogs)
