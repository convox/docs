---
title: "App Parameters"
---

## Setting Parameters

Parameters can be set using the following command.

    convox apps params set Foo=Yes Baz=No

<div class="block-callout block-show-callout type-info" markdown="1">
  When a Rack update adds new app parameters they become available for each app after its next deploy.
</div>

### FargateServices

Set to `Yes` to run all services for this application in [Fargate](https://aws.amazon.com/fargate/).

### FargateTimers

Set to `Yes` to run all timers for this application in [Fargate](https://aws.amazon.com/fargate/).

### IamPolicy

Specify the ARN of a custom IAM policy to add the the Service's [Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) at runtime.

### InternalDomains

Set to `No` to disable the internal domain names at `convox.site` and `.convox` to from routing to this application. You can use this parameter if you are running out of available rules on your load balancer.

### LoadBalancerAlgorithm

Sets the routing algorithm used for the services within the application.  Defaults to `round_robin`, but you can also set to `least_outstanding_requests`.

### LogRetention

Set to the number of days you wish to retain logs for.  The default for new applications is `7`.  Setting the retention window to a high/unlimited value will affect the performance/reliability of `convox logs` over the long term.  It is recommended to keep it at a smaller value and use [syslog](/deployment/syslogs) to export your logs for long-term archival and analysis.

### RedirectHttps

Set to `No` to allow the app to listen on HTTP rather than the default behavior of having all HTTP requests automatically redirected to HTTPS.

### ScanImages

Uses the built-in ECR image scanning to help identify vulnerabilities on each image push.  Defaults to `Yes`, but if you wish to disable automatic image scanning, set this app parameter to `No`.

### TaskTags

Set to `Yes` to cause ECS tags to be propagated to the task level.
