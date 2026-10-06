---
title: "Logs"
description: "View, filter, and manage application and Rack logs in Convox, including retention settings and third-party routing."
---

# Logs

You can view the live logs for a Convox application using `convox logs`:

```bash
$ convox logs
2025-04-12T19:45:00Z service/web/4f0d9a3c1b2e7f6a8d905e3c8576b942 10.0.1.242 - - [12/Apr/2025:19:45:00 +0000] "GET / HTTP/1.1" 200 70 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
2025-04-12T19:45:00Z service/web/4f0d9a3c1b2e7f6a8d905e3c8576b942 10.0.1.242 - - [12/Apr/2025:19:45:00 +0000] "GET / HTTP/1.0" 200 70 0.0019
```

## Retention

By default, new applications will retain 7 days worth of logs.  You can control the retention window (in numbers of days) through the `LogRetention` app-level parameter.

```bash
$ convox apps params set LogRetention=3
```

To set an unlimited retention window, configure the parameter to be blank/empty.

```bash
$ convox apps params set LogRetention=
```

> **Warning:** Setting the retention window to a high/unlimited value will affect the performance/reliability of `convox logs` over the long term. Keep it at a smaller value and use [syslog](/deployment/syslogs) to export your logs for long-term archival and analysis.

## Additional Options

| Option | Description |
|---|---|
| `--filter=POST` | Return only the logs that match all the filters. Filters are case sensitive and non-alphanumeric terms must be inside double quotes. |
| `--since=20m` | Return logs starting this duration ago. Values are a duration like `10m` or `48h`. |
| `--no-follow` | Return the logs in the `--since` window and exit instead of following new logs. |

You can tie all these together to generate a report from the logs over the last 2 days, filtering by a specific term:

```bash
$ convox ps
ID            SERVICE  STATUS   RELEASE      STARTED     COMMAND
310481bf223f  web      running  RSPZQWVWGOP  2 days ago  bin/web
5e3c8576b942  web      running  RSPZQWVWGOP  2 days ago  bin/web

$ convox logs --filter=GET --since=48h --no-follow
2025-04-25T21:12:02Z service/web/4f0d9a3c1b2e7f6a8d905e3c8576b942 10.0.3.243 - - [25/Apr/2025:21:12:02 +0000] "GET / HTTP/1.1" 200 226
2025-04-25T21:12:03Z service/web/4f0d9a3c1b2e7f6a8d905e3c8576b942 10.0.3.243 - - [25/Apr/2025:21:12:02 +0000] "GET / HTTP/1.1" 200 226
2025-04-25T21:12:04Z service/web/4f0d9a3c1b2e7f6a8d905e3c8576b942 10.0.3.243 - - [25/Apr/2025:21:12:02 +0000] "GET / HTTP/1.1" 200 226
...
```

## AWS Logs

We include events from AWS services in the output of `convox logs` to help you understand what's going on behind the scenes.

```text
2025-01-15T14:34:08Z system/ecs aws/ecs (service production-myapp-ServiceWeb-1A2B3C4D5E6F-Service-7G8H9I0J1K2L) has started 1 tasks: (task 4f0d9a3c1b2e7f6a8d905e3c8576b942).
2025-01-15T14:35:12Z system/ecs aws/ecs (service production-myapp-ServiceWeb-1A2B3C4D5E6F-Service-7G8H9I0J1K2L) has reached a steady state.

2025-01-15T14:59:56Z system/cloudformation aws/cfm production-myapp UPDATE_IN_PROGRESS ServiceWeb
2025-01-15T15:02:30Z system/cloudformation aws/cfm production-myapp UPDATE_COMPLETE ServiceWeb
2025-01-15T15:02:34Z system/cloudformation aws/cfm production-myapp UPDATE_COMPLETE production-myapp
```

These events can be useful for identifying issues with a deployment or an App. For example, when your Rack is in a "converging" state, i.e. some instances or ECS Tasks haven't stabilized yet, there are often AWS events in the App/Rack logs that will show a Service crashing, a health check failing, or a placement error due to insufficient resources.

## Rack Logs

You can view the logs for a Convox Rack itself using the `convox rack logs` command:

```bash
$ convox rack logs
2025-01-15T14:59:57Z service/web/0b92eed79c1d4b6e8f0a1c2d3e4f5a6b id=a1b2c3d4e5f6 ns=api at=SystemGet method="GET" path="/system" response=200 elapsed=70.493
2025-01-15T15:16:15Z service/monitor/e378ddb167fd4c5d6e7f8a9b0c1d2e3f who="EC2/ASG" what="Terminating EC2 instance: i-02ce4f07da10a5333" why="At 2025-01-15T15:14:38Z a user request update of AutoScalingGroup constraints to min: 3, max: 1000, desired: 3 changing the desired capacity from 4 to 3.  At 2025-01-15T15:15:02Z an instance was taken out of service in response to a difference between desired and actual capacity, shrinking the capacity from 4 to 3.  At 2025-01-15T15:15:02Z instance i-02ce4f07da10a5333 was selected for termination."
```

## Routing Logs to a 3rd Party

See the dedicated section to Syslogs [here](/deployment/syslogs)...

## Disabling logging system

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
