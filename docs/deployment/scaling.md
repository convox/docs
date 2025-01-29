---
title: "Scaling"
---

Convox allows you to scale your application's concurrency, memory allocation, and the resources available in the underlying Rack.

## Setting Initial Defaults

You can specify initial default values for a service in the [convox.yml](/docs/convox-yml):

```
services:
  web:
    scale:
      count: 4
      cpu: 512
      memory: 1024
```

These values are used upon the first deployment. Subsequent changes should be managed through the CLI or by configuring autoscaling.

## Scaling an Application

```
$ convox scale
NAME  DESIRED  RUNNING  MEMORY
web   2        1        256
redis 1        1        256
```

### Changing the Number of Processes for a Service

```
$ convox scale web --count=4
NAME  DESIRED  RUNNING  MEMORY
web   4        4        256
```

### Setting New Values for CPU or Memory

```
$ convox scale web --memory=1024 --cpu=512
NAME  DESIRED  RUNNING  MEMORY
web   4        4        1024
```

<div class="block-callout block-show-callout type-warning" markdown="1">
Each rack instance has 1024 [CPU units](http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html#ECS-Type-ContainerDefinition-cpu) per core. This parameter specifies the minimum CPU allocation for a container. Containers share unallocated CPU units with other containers in proportion to their assigned amounts.
</div>

## Autoscaling

There are two dimensions of scaling in a Convox Rack:

- **Rack-level scaling** (the number of instances available to run containers)
- **Service-level scaling** (the number of containers running per service)

Convox supports autoscaling for both.

### Rack Autoscaling

Rack-level autoscaling is enabled by default in new Racks. When enabled, the number of instances running will dynamically adjust based on container workload.

You can configure rack autoscaling using the `Autoscale` rack parameter:

```
$ convox rack params set Autoscale=Yes
```

### Service Autoscaling

Service-level autoscaling is controlled in the `convox.yml`:

```
services:
  web:
    scale:
      count: 1-10
      targets:
        cpu: 70
        memory: 90
        requests: 200
```

Setting scale targets for a service will cause the service autoscaler to dynamically adjust the number of running containers to meet the targets.

- `cpu`: Average CPU utilization (%) across all instances
- `memory`: Average Memory utilization (%) across all instances
- `requests`: Requests per minute per process

You can also use custom CloudWatch metrics to define scaling behavior:

```
services:
  worker:
    scale:
      count: 1-10
      targets:
        custom:
          AWS/SQS/ApproximateNumberOfMessagesVisible:
            aggregate: max
            value: 20
            dimensions:
              QueueName: myqueue
```

These settings will dynamically adjust the number of worker instances to maintain a maximum of 20 messages in the queue.

You can define one cooldown period for both scaling directions:

```
services:
  web:
    scale:
      cooldown: 120
```

Or specify separate values for scale-up and scale-down events:

```
services:
  web:
    scale:
      cooldown:
        down: 180
        up: 45
```

If no value is defined, the default cooldown period is 60 seconds.

## Manual Scaling and IgnoreManualScaleCount

When autoscaling is enabled, manually setting a desired count using:

```
$ convox scale serviceName --count=<number> -a appName
```

will temporarily adjust the instance count. The autoscaler will monitor the scaling targets and adjust the count again within 3–5 minutes, unless the conditions sustaining the manual count persist.

Manually scaling can be useful for scenarios such as:
- **Adjusting to a new desired count after deployment**
- **Pre-scaling a service in anticipation of a traffic spike** (e.g., a scheduled event, sale, or product launch)

If traffic sustains at the increased level, the autoscaler will continue scaling accordingly. However, if traffic subsides, the autoscaler will adjust back within the configured range.

**Important:**  
A manually set count will be used as the target count for the application's **next deployment**. If the manual scaling was done for temporary ramping rather than establishing a new baseline, you should **reset the desired count** after the traffic spike subsides.

### IgnoreManualScaleCount

`IgnoreManualScaleCount` is an **app parameter** that determines whether manually set desired counts are overridden by the autoscaler. When set to `Yes`, Convox will **not** use manually specified values for scaling. The service count will be **fully controlled by the autoscaler**, and the **active count at deployment time will be treated as the desired count**.

This parameter is set at the **app level** and can be configured using the following command:

```
$ convox apps params set IgnoreManualScaleCount=Yes -a appName
```

For full details, refer to the [App Parameters page](/reference/app-parameters#ignoremanualscalecount).

| Default value  | `No` |
| Allowed values | `Yes`, `No` |

## Summary

- **Rack Autoscaling** dynamically adjusts instance counts based on workload.
- **Service Autoscaling** scales service containers based on defined targets.
- **Manual Scaling** can temporarily override desired counts but is corrected by the autoscaler.
- **IgnoreManualScaleCount** ensures full autoscaler control by ignoring manually set values.

For best practices, define scaling ranges and targets in `convox.yml` to ensure dynamic and efficient scaling behavior.
