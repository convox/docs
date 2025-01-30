---
title: "App Parameters"
---

## Setting Parameters

App parameters configure specific behaviors at the application level. Parameters can be set using the following command:

```
$ convox apps params set ParameterName=Value
```

<div class="block-callout block-show-callout type-info" markdown="1">
  When a Rack update adds new app parameters, they become available for each app after its next deployment.
</div>

## Available Parameters

- [AutoMinorVersionUpgrade](#autominorversionupgrade)
- [CircuitBreaker](#circuitbreaker)
- [EnableContainerReadonlyRootFilesystem](#enablecontainerreadonlyrootfilesystem)
- [FargateServices](#fargateservices)
- [FargateTimers](#fargatetimers)
- [IamPolicy](#iampolicy)
- [IgnoreManualScaleCount](#ignoremanualscalecount)
- [InternalDomains](#internaldomains)
- [LoadBalancerAlgorithm](#loadbalanceralgorithm)
- [LoadBalancerGrpcSuccessCodes](#loadbalancergrpcsuccesscodes)
- [LoadBalancerSuccessCodes](#loadbalancersuccesscodes)
- [LogRetention](#logretention)
- [RedirectHttps](#redirecthttps)
- [SlowStartDuration](#slowstartduration)
- [TaskTags](#tasktags)

---

### <a href="#autominorversionupgrade" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#autominorversionupgrade'); return false;">🔗</a> autominorversionupgrade <a name="autominorversionupgrade"></a>

Set to `false` to disable automatic minor version upgrades for any database resources in this app. Defaults to `true`.

---

### <a href="#circuitbreaker" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#circuitbreaker'); return false;">🔗</a> circuitbreaker <a name="circuitbreaker"></a>

Set to `Yes` to enable Circuit Breaker deployments in ECS. This allows failing deployments to roll back faster. However, if a deployment requires the Rack to scale up, it may trip the Circuit Breaker prematurely. Best used when sufficient capacity is available.

---

### <a href="#enablecontainerreadonlyrootfilesystem" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#enablecontainerreadonlyrootfilesystem'); return false;">🔗</a> EnableContainerReadonlyRootFilesystem <a name="enablecontainerreadonlyrootfilesystem"></a>

This parameter controls whether ECS containers operate with a **read-only root filesystem**, enhancing security by preventing modifications to critical system files.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

- `Yes`: Activates **read-only mode**, restricting write operations to the root filesystem. Applications must use external storage or explicitly writable paths.
- `No` (default): Containers have standard write access to the root filesystem.

**Potential Breaking Change:**  
Enabling this parameter may **disrupt applications** that write to the root filesystem. Ensure your application is compatible before enabling it. Use **external storage solutions** or **designated writable paths** to prevent unexpected failures.

**Example usage:**
```
$ convox apps params set EnableContainerReadonlyRootFilesystem=Yes
```

---

### <a href="#fargateservices" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#fargateservices'); return false;">🔗</a> fargateservices <a name="fargateservices"></a>

Set to `Yes` to run all services for this application in [Fargate](https://aws.amazon.com/fargate/).  
Set to `Spot` to run all services in [Fargate Spot](https://aws.amazon.com/blogs/aws/aws-fargate-spot-now-generally-available/).

---

### <a href="#fargatetimers" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#fargatetimers'); return false;">🔗</a> fargatetimers <a name="fargatetimers"></a>

Set to `Yes` to run all timers for this application in [Fargate](https://aws.amazon.com/fargate/).  
Set to `Spot` to run all timers in [Fargate Spot](https://aws.amazon.com/blogs/aws/aws-fargate-spot-now-generally-available/).

---

### <a href="#iampolicy" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#iampolicy'); return false;">🔗</a> iampolicy <a name="iampolicy"></a>

Specify the ARN of a custom IAM policy to attach to the Service's [Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) at runtime.  
If the service has the [Policies](/application/services) parameter set, this will not apply at the service level.

---

### <a href="#ignoremanualscalecount" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#ignoremanualscalecount'); return false;">🔗</a> ignoremanualscalecount <a name="ignoremanualscalecount"></a>

When autoscaling is enabled, this parameter controls whether manually set desired counts are ignored in favor of autoscaler-managed scaling.

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

If set to `Yes`, any manually set service counts using `convox scale` will be ignored, and the autoscaler will dynamically manage the service count based on the targets defined in `convox.yml`.  
Additionally, during deployments, the **active** instance count at the time of deployment will be used as the new desired count, ensuring the autoscaler fully controls scaling behavior without manual overrides.

**Example usage:**
```
$ convox apps params set IgnoreManualScaleCount=Yes
```

For more details on how this parameter interacts with scaling, see [Scaling](/deployment/scaling).

---

### <a href="#internaldomains" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#internaldomains'); return false;">🔗</a> internaldomains <a name="internaldomains"></a>

Set to `No` to disable internal domain names at `convox.site` and `.convox` from routing to this application.  
Use this parameter if you are running out of available rules on your load balancer.

---

### <a href="#loadbalanceralgorithm" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#loadbalanceralgorithm'); return false;">🔗</a> loadbalanceralgorithm <a name="loadbalanceralgorithm"></a>

Sets the routing algorithm used for the services within the application.  
Defaults to `round_robin`. Can also be set to `least_outstanding_requests`.

---

### <a href="#loadbalancergrpcsuccesscodes" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#loadbalancergrpcsuccesscodes'); return false;">🔗</a> loadbalancergrpcsuccesscodes <a name="loadbalancergrpcsuccesscodes"></a>

Specifies the GRPC codes that healthy targets must use when responding to a GRPC health check.

| Default value  | `12` |
| Allowed values | `0-99` |

You can specify multiple values (e.g., `12,13`) or a range (e.g., `10-99`).

---

### <a href="#loadbalancersuccesscodes" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#loadbalancersuccesscodes'); return false;">🔗</a> loadbalancersuccesscodes <a name="loadbalancersuccesscodes"></a>

Specifies the HTTP codes that healthy targets must return when responding to an HTTP health check.

| Default value  | `200-399,401` |
| Allowed values | `200-499` |

You can specify multiple values (e.g., `200,202`) or a range (e.g., `200-299`).

---

### <a href="#logretention" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#logretention'); return false;">🔗</a> LogRetention <a name="logretention"></a>

Defines how long logs should be retained.

| Default value  | `7` days |

Setting this to a high/unlimited value may affect the performance of `convox logs`.  
For long-term log storage, consider using [syslog](/deployment/syslogs).

---

### <a href="#redirecthttps" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#redirecthttps'); return false;">🔗</a> redirecthttps <a name="redirecthttps"></a>

Set to `No` to allow the app to listen on HTTP rather than automatically redirecting all HTTP requests to HTTPS.

---

### <a href="#slowstartduration" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#slowstartduration'); return false;">🔗</a> slowstartduration <a name="slowstartduration"></a>

Sets the ramp-up period during which a newly deployed service gradually receives an increasing share of traffic.

| Default value  | `0` seconds (disabled) |
| Allowed values | `30-900` seconds |

Once the duration expires, the full share of traffic will be directed to the new service.

---

### <a href="#tasktags" title="Copy link" onclick="navigator.clipboard.writeText(location.origin + location.pathname + '#tasktags'); return false;">🔗</a> tasktags <a name="tasktags"></a>

Set to `Yes` to propagate ECS tags to the task level.
