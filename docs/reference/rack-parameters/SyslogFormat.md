---
title: "SyslogFormat"
description: "Set the syslog message format used when forwarding Convox Rack logs to a syslog destination."
---

# SyslogFormat

Syslog format (lowercase) to send to [SyslogDestination](/reference/rack-parameters/SyslogDestination). See [Docker Syslog logging driver](https://docs.docker.com/config/containers/logging/syslog/) and [RFC 5424](https://www.rfc-editor.org/rfc/rfc5424#section-6) for format details.

| Default value  | `rfc5424` |
| Allowed values | `rfc5424`, `rfc3164` |

## Use Cases

- Keeping the default `rfc5424` format for modern syslog receivers that support structured data
- Switching to `rfc3164` for legacy syslog servers that do not support RFC 5424
- Matching the format expected by your centralized logging platform

## Additional Information

The syslog format determines how log messages are structured when sent to the [SyslogDestination](/reference/rack-parameters/SyslogDestination). The two most common formats are:

- **rfc5424** -- The modern syslog format that supports structured data, message IDs, and more precise timestamps. This is the default and recommended format.
- **rfc3164** -- The legacy BSD syslog format with simpler structure. Use this if your syslog receiver does not support RFC 5424.

This parameter only applies when [LogDriver](/reference/rack-parameters/LogDriver) is set to `Syslog`.

```bash
$ convox rack params set SyslogFormat=rfc3164
```

## See Also

- [SyslogDestination](/reference/rack-parameters/SyslogDestination)
- [LogDriver](/reference/rack-parameters/LogDriver)
- [LogRetention](/reference/rack-parameters/LogRetention)
