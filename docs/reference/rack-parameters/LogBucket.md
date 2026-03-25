---
title: "LogBucket"
description: "Specify an existing S3 bucket to receive Rack load balancer access logs."
---

# LogBucket

An existing S3 bucket to receive load balancer access logs and other Rack logs. If blank, the Rack creates its own log buckets.

When set, ALB access logs and other infrastructure logs are delivered to the specified S3 bucket. The bucket must already exist and must have the proper bucket policy to allow Elastic Load Balancing to write logs to it.

| Default value  | "" |

## Use Cases

- Centralizing logs from multiple Racks into a single S3 bucket for consolidated analysis
- Using an existing bucket that already has lifecycle policies, encryption, or replication configured
- Meeting compliance requirements that mandate logs be stored in a specific bucket with specific access controls

## Additional Information

The S3 bucket must be in the same AWS region as your Rack. It also needs a bucket policy that grants the Elastic Load Balancing service permission to write objects. See the AWS documentation on [Access Logs for Your Application Load Balancer](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-access-logs.html) for the required bucket policy.

```bash
$ convox rack params set LogBucket=my-centralized-log-bucket
```

## See Also

- [LogDriver](/reference/rack-parameters/LogDriver)
- [LogRetention](/reference/rack-parameters/LogRetention)
- [EnableS3Versioning](/reference/rack-parameters/EnableS3Versioning)
