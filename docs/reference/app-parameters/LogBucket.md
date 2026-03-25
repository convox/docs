---
title: "LogBucket"
description: "Specify an S3 bucket to receive application load balancer access logs."
---

# LogBucket

S3 bucket name for application load balancer access logs. When blank, access logging is disabled.

| Default value  | "" |

## Use Cases

- Set to an S3 bucket name when you need ALB access logs for compliance auditing or security analysis
- Set when troubleshooting load balancer routing issues and you need detailed request-level logs
- Leave blank (default) when you do not need ALB access logs and want to avoid additional S3 storage costs

## Additional Information

The specified S3 bucket must already exist and must have the correct bucket policy to allow the Elastic Load Balancing service to write logs. See the [AWS documentation on ALB access logs](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/enable-access-logging.html) for the required bucket policy.

The S3 bucket is also used for logging access to the application's settings bucket. Log files are written with the prefix `convox/logs/<stack-name>/s3/`.

```bash
$ convox apps params set LogBucket=my-alb-logs-bucket
```

## See Also

- [LogDriver](/reference/app-parameters/LogDriver)
- [LogRetention](/reference/app-parameters/LogRetention)
- [SyslogDestination](/reference/app-parameters/SyslogDestination)
- [Rack Parameters](/reference/rack-parameters)
