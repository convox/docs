---
title: "Logs"
description: "View, filter, and manage application and Rack logs in Convox, including retention settings and third-party routing."
---

# Logs

You can view the live logs for a Convox application using `convox logs`:

```bash
$ convox logs
2025-04-12 19:45:00 i-0234d285 service/web:RSPZQWVWGOP/5e3c8576b942 : 10.0.1.242 - - [12/Apr/2025:19:45:00 +0000] "GET / HTTP/1.1" 200 70 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
2025-04-12 19:45:00 i-0234d285 service/web:RSPZQWVWGOP/5e3c8576b942 : 10.0.1.242 - - [12/Apr/2025:19:45:00 +0000] "GET / HTTP/1.0" 200 70 0.0019
```

### Retention

By default, new applications will retain 7 days worth of logs.  You can control the retention window (in numbers of days) through the `LogRetention` app-level parameter.

```bash
$ convox apps params set LogRetention=3
```

To set an unlimited retention window, configure the parameter to be blank/empty.

```bash
$ convox apps params set LogRetention=
```

> **Warning:** Setting the retention window to a high/unlimited value will affect the performance/reliability of `convox logs` over the long term. Keep it at a smaller value and use [syslog](/deployment/syslogs) to export your logs for long-term archival and analysis.

### Additional Options

| Option | Description |
|---|---|
| `--filter=POST` | Return only the logs that match all the filters. Filters are case sensitive and non-alphanumeric terms must be inside double quotes. |
| `--since=20m` | Return logs starting this duration ago. Values are a duration like `10m` or `48h`. |

You can tie all these together to generate a report from the logs over the last 2 days, filtering by a specific term:

```bash
$ convox ps
ID            NAME  RELEASE      SIZE  STARTED     COMMAND
310481bf223f  web   RSPZQWVWGOP  256   2 days ago  bin/web
5e3c8576b942  web   RSPZQWVWGOP  256   2 days ago  bin/web

$ convox logs --filter=GET --since=48h
2025-04-25T21:12:02Z service/web:RSPZQWVWGOP/5e3c8576b942 10.0.3.243 - - [25/Apr/2025:21:12:02 +0000] "GET / HTTP/1.1" 200 226
2025-04-25T21:12:03Z service/web:RSPZQWVWGOP/5e3c8576b942 10.0.3.243 - - [25/Apr/2025:21:12:02 +0000] "GET / HTTP/1.1" 200 226
2025-04-25T21:12:04Z service/web:RSPZQWVWGOP/5e3c8576b942 10.0.3.243 - - [25/Apr/2025:21:12:02 +0000] "GET / HTTP/1.1" 200 226
...
```

### AWS Logs

We include events from AWS services in the output of `convox logs` to help you understand what's going on behind the scenes.

```text
2025-01-15T14:34:08Z system/aws/ecs service="legit-cv-soulshake-net-ServiceWopr-1TNE86TXRESN5" task="46a3ea28-4868-4620-b43c-5e4af6732219" status="STOPPED" container="wopr"
2025-01-15T14:34:09Z system/aws/ecs service="legit-cv-soulshake-net-ServiceWopr-1TNE86TXRESN5" task="46a3ea28-4868-4620-b43c-5e4af6732219" status="RUNNING" container="wopr"

2025-01-15T14:59:56Z system/aws/cfm resource="LaunchConfiguration" status="UPDATE_IN_PROGRESS" reason="Resource creation Initiated"
2025-01-15T14:59:56Z system/aws/cfm resource="LaunchConfiguration" status="UPDATE_COMPLETE" reason=""
2025-01-15T15:00:04Z system/aws/cfm resource="Instances" status="UPDATE_IN_PROGRESS" reason=""
```

These events can be useful for identifying issues with a deployment or an app. For example, when your Rack is in a "converging" state, i.e. some instances or ECS Tasks haven't stabilized yet, there are often AWS events in the app/Rack logs that will show a service crashing, a health check failing, or a placement error due to insufficient resources.

### Rack Logs

You can view the logs for a Convox Rack itself using the `convox rack logs` command:

```bash
$ convox rack logs
2025-01-15T14:59:57Z service/web:20250110143201/0b92eed79c1d ns=provider.aws at=fetchLogs start=1736949597065 events=0 state=success end=1736949597066 elapsed=225.020
2025-01-15T15:16:15Z service/web:20250110143201/e378ddb167fd who="EC2/ASG" what="Terminating EC2 instance: i-02ce4f07da10a5333" why="At 2025-01-15T15:14:38Z a user request update of AutoScalingGroup constraints to min: 3, max: 1000, desired: 3 changing the desired capacity from 4 to 3.  At 2025-01-15T15:15:02Z an instance was taken out of service in response to a difference between desired and actual capacity, shrinking the capacity from 4 to 3.  At 2025-01-15T15:15:02Z instance i-02ce4f07da10a5333 was selected for termination."
```

### Routing Logs to a 3rd Party

See the dedicated section to Syslogs [here](/deployment/syslogs)...

### Disabling logging system

You can disable the logging system by setting `LogDriver` as empty:

```bash
$ convox rack params set LogDriver=""
```

It will not create a CloudWatch LogGroup, existing LogGroups will not be deleted. Be aware that disabling it, convox logs and convox rack logs will stop working.

## See Also

- [Syslogs](/deployment/syslogs)
- [Logging Integrations](/integrations/logging)
- [Debugging](/management/debugging)
- [Application Monitoring](/management/application)
- [Datadog Integration](/integrations/datadog)
