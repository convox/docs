---
title: "App Parameters"
---

## Setting Parameters

Parameters can be set using the following command.

    convox apps params set Foo=Yes Baz=No

### FargateServices

Set to `Yes` to run all services for this application in [Fargate](https://aws.amazon.com/fargate/).

### FargateTimers

Set to `Yes` to run all timers for this application in [Fargate](https://aws.amazon.com/fargate/).

### IamPolicy

Specify the ARN of a custom IAM policy to add the the Service's [Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) at runtime.

### InternalDomains

Set to `No` to disable the internal domain names at `convox.site` and `.convox` to from routing to this application. You can use this parameter if you are running out of available rules on your load balancer.
