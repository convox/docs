---
title: "SyslogDestination"
description: "The syslog endpoint URL to send application logs to when LogDriver is set to Syslog."
---

# SyslogDestination

The syslog endpoint to send logs to when [LogDriver](/reference/app-parameters/LogDriver) is set to `Syslog`. Include the protocol prefix (e.g., `tcp+tls://logs.example.com:1234`).

| Default value  | "" |

## Use Cases

- Set when forwarding application logs to an external log aggregation service such as Papertrail, Datadog, or Loggly
- Set when your organization requires centralized log collection outside of AWS CloudWatch
- Leave blank when using CloudWatch (the default log driver)

## Additional Information

This parameter is only used when [LogDriver](/reference/app-parameters/LogDriver) is set to `Syslog`. If LogDriver is set to `CloudWatch` or blank, this value is ignored.

The destination must include the protocol prefix. Supported formats include:

- `tcp+tls://logs.example.com:1234` -- TCP with TLS encryption (recommended)
- `tcp://logs.example.com:1234` -- TCP without encryption
- `udp://logs.example.com:1234` -- UDP

```bash
$ convox apps params set LogDriver=Syslog SyslogDestination=tcp+tls://logs.example.com:1234
```

The syslog message format is controlled by the [SyslogFormat](/reference/app-parameters/SyslogFormat) parameter.

## See Also

- [LogDriver](/reference/app-parameters/LogDriver)
- [SyslogFormat](/reference/app-parameters/SyslogFormat)
- [LogRetention](/reference/app-parameters/LogRetention)
- [Syslogs](/deployment/syslogs)
