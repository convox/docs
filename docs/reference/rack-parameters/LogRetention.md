---
title: "LogRetention"
description: "Configure the number of days to retain CloudWatch logs for the Convox Rack."
---

# LogRetention

Number of days to keep logs (blank for unlimited). This controls the retention period for CloudWatch Log Groups created by the Rack.

Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557, 2922, 3288, and 3653. See [Logs retention](/management/logs#retention) for more information.

| Default value  | `7` |

## Use Cases

- Increasing retention to 30 or 90 days for production environments where historical logs are needed for troubleshooting
- Setting to 365 or higher for compliance requirements that mandate long-term log storage
- Setting to blank for unlimited retention when combined with a separate log archival strategy
- Reducing retention to 1 or 3 days in development environments to minimize CloudWatch Logs costs

## Additional Information

The retention period applies to all CloudWatch Log Groups managed by the Rack, including both infrastructure logs and application logs.

CloudWatch Logs charges are based on the volume of data ingested and stored. Longer retention periods will increase your monthly CloudWatch costs. For long-term log archival at lower cost, consider exporting logs to S3 or using a [LogDriver](/reference/rack-parameters/LogDriver) set to `Syslog` with an external log aggregation service.

```bash
$ convox rack params set LogRetention=30
```

## See Also

- [LogDriver](/reference/rack-parameters/LogDriver)
- [LogBucket](/reference/rack-parameters/LogBucket)
- [Logs Management](/management/logs#retention)
