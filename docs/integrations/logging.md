---
title: "Logging Integrations"
description: "Forward application and system logs from Convox to external logging services including Datadog, Papertrail, Loggly, and Mezmo."
---

# Logging Integrations

Convox can forward application and system logs to external logging services using syslog endpoints. Each logging integration configures a syslog resource on your Rack that streams logs in real time.

## Available Integrations

| Integration | Description |
|-------------|-------------|
| [Syslogs](/deployment/syslogs) | Configure syslog forwarding to any syslog-compatible endpoint |
| [Datadog](/integrations/datadog) | Forward logs and metrics to Datadog |
| [Papertrail](/integrations/papertrail) | Stream logs to Papertrail for search and alerting |
| [Loggly](/integrations/loggly) | Forward logs to Loggly for analysis |
| [Mezmo](/integrations/logdna) | Send logs to Mezmo (formerly LogDNA) |

## How It Works

Convox logging integrations use Rack-level syslog resources. When you create a syslog resource, all application and system logs from the Rack are forwarded to the configured endpoint. See [Syslogs](/deployment/syslogs) for the underlying configuration.

## See Also

- [Integrations Overview](/integrations/integrations)
- [Logs](/management/logs)
- [Debugging](/management/debugging)
