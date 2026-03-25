---
title: "App Parameters"
description: "Reference for all application-level parameters that configure Convox App behavior on AWS."
---

# App Parameters

## Setting Parameters

App parameters configure specific behaviors at the application level. Parameters can be set using the following command:

```bash
$ convox apps params set ParameterName=Value
```

> **Note:** When a Rack update adds new App parameters, they become available for each App after its next deployment.

## Parameters

| Parameter | Default | Description |
|:----------|:--------|:------------|
| [AutoMinorVersionUpgrade](/reference/app-parameters/AutoMinorVersionUpgrade) | `true` | Auto-upgrade minor versions for database resources |
| [CircuitBreaker](/reference/app-parameters/CircuitBreaker) | `No` | Enable ECS Circuit Breaker deployments |
| [EnableContainerReadonlyRootFilesystem](/reference/app-parameters/EnableContainerReadonlyRootFilesystem) | `No` | Enable read-only root filesystem for containers |
| [FargateServices](/reference/app-parameters/FargateServices) | `No` | Run services on AWS Fargate |
| [FargateTimers](/reference/app-parameters/FargateTimers) | `No` | Run timers on AWS Fargate |
| [IamPolicy](/reference/app-parameters/IamPolicy) | "" | Custom IAM policy ARN for the task role |
| [IgnoreManualScaleCount](/reference/app-parameters/IgnoreManualScaleCount) | `No` | Ignore manual scale counts during autoscaling |
| [InternalDomains](/reference/app-parameters/InternalDomains) | `Yes` | Enable internal domain routing |
| [Isolate](/reference/app-parameters/Isolate) | `No` | Run services in isolated network configuration |
| [LoadBalancerAlgorithm](/reference/app-parameters/LoadBalancerAlgorithm) | `round_robin` | Load balancer routing algorithm |
| [LoadBalancerGrpcSuccessCodes](/reference/app-parameters/LoadBalancerGrpcSuccessCodes) | `12` | GRPC health check success codes |
| [LoadBalancerSuccessCodes](/reference/app-parameters/LoadBalancerSuccessCodes) | `200-399,401` | HTTP health check success codes |
| [LogBucket](/reference/app-parameters/LogBucket) | "" | S3 bucket for ALB access logs |
| [LogDriver](/reference/app-parameters/LogDriver) | `CloudWatch` | Log driver for the application |
| [LogRetention](/reference/app-parameters/LogRetention) | `7` | Days to retain logs |
| [PlaceLambdaInVpc](/reference/app-parameters/PlaceLambdaInVpc) | `No` | Place Lambda functions inside the VPC |
| [Private](/reference/app-parameters/Private) | `No` | Deploy tasks to private subnets |
| [RackUrl](/reference/app-parameters/RackUrl) | `No` | Inject RACK_URL environment variable |
| [RedirectHttps](/reference/app-parameters/RedirectHttps) | `Yes` | Redirect HTTP to HTTPS |
| [ResourcePassword](/reference/app-parameters/ResourcePassword) | "" | Override password for embedded resources |
| [SlowStartDuration](/reference/app-parameters/SlowStartDuration) | `0` | Ramp-up period for new services (seconds) |
| [SyslogDestination](/reference/app-parameters/SyslogDestination) | "" | Syslog endpoint URL |
| [SyslogFormat](/reference/app-parameters/SyslogFormat) | `rfc5424` | Syslog message format |
| [TaskTags](/reference/app-parameters/TaskTags) | `No` | Propagate ECS tags to task level |

## See Also

- [Rack Parameters](/reference/rack-parameters)
- [App Statuses](/reference/app-statuses)
- [Scaling](/scaling/scaling)
- [Services](/application/services)
