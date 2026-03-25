---
title: "SyslogDestination"
description: "Configure the syslog endpoint for forwarding Convox Rack and service logs."
---

# SyslogDestination

Syslog address destination. You need to pass the protocol to be used, e.g. `tcp+tls://logsX.syslog.com:1234`.

| Default value  | "" |

## Use Cases

- Forwarding logs to a centralized logging platform such as Papertrail, Datadog, or Splunk via syslog
- Sending logs to an internal syslog collector for compliance or audit purposes
- Using syslog as the primary logging mechanism when CloudWatch is not desired

## Additional Information

This parameter is required when [LogDriver](/reference/rack-parameters/LogDriver) is set to `Syslog`. The destination must include the protocol prefix. Supported formats include:

- `tcp+tls://host:port` -- Syslog over TCP with TLS encryption (recommended)
- `tcp://host:port` -- Syslog over TCP without encryption
- `udp://host:port` -- Syslog over UDP

The format of the log messages sent to this destination is controlled by the [SyslogFormat](/reference/rack-parameters/SyslogFormat) parameter.

```bash
$ convox rack params set SyslogDestination=tcp+tls://logs.example.com:514
```

## See Also

- [SyslogFormat](/reference/rack-parameters/SyslogFormat)
- [LogDriver](/reference/rack-parameters/LogDriver)
- [LogRetention](/reference/rack-parameters/LogRetention)
