---
title: "SyslogFormat"
description: "Sets the syslog message format used when sending logs to the configured SyslogDestination."
---

# SyslogFormat

Syslog message format for log delivery to [SyslogDestination](/reference/app-parameters/SyslogDestination).

| Default value  | `rfc5424` |
| Allowed values | `rfc5424`, `rfc3164` |

## Use Cases

- Keep as `rfc5424` (default) for modern syslog receivers that support the RFC 5424 structured format
- Change to `rfc3164` if your syslog receiver only supports the legacy BSD syslog format
- Adjust to match the format expected by your specific log aggregation service

## Additional Information

This parameter only takes effect when [LogDriver](/reference/app-parameters/LogDriver) is set to `Syslog` and a valid [SyslogDestination](/reference/app-parameters/SyslogDestination) is configured.

The `rfc5424` format includes structured data fields and is the recommended format for modern syslog infrastructure. The `rfc3164` format is the legacy BSD syslog format and may be required by older syslog receivers.

```bash
$ convox apps params set SyslogFormat=rfc3164
```

## See Also

- [SyslogDestination](/reference/app-parameters/SyslogDestination)
- [LogDriver](/reference/app-parameters/LogDriver)
- [LogRetention](/reference/app-parameters/LogRetention)
- [Syslogs](/deployment/syslogs)
