---
title: "Scaling"
---

Convox allows you to scale your application's concurrency, memory allocation, and the resources available in the underlying Rack.

### Scaling an application

```
$ convox scale
NAME  DESIRED  RUNNING  MEMORY
web   2        1        256
redis 1        1        256
```

#### Changing the number of processes for a service

```
$ convox scale web --count=4
NAME  DESIRED  RUNNING  MEMORY
web   2        1        256
```

#### Setting new values for CPU or Memory

```
$ convox scale web --memory=1024 --cpu=512
NAME  DESIRED  RUNNING  MEMORY
web   2        1        1024
```

<div class="block-callout block-show-callout type-warning" markdown="1">
Each rack instance has 1024 [cpu units](http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html#ECS-Type-ContainerDefinition-cpu) for every CPU core. This parameter specifies the minimum amount of CPU to reserve for a container. Containers share unallocated CPU units with other containers on the instance with the same ratio as their allocated amount.
</div>

#### Setting defaults

You can specify default values for a service in the [convox.yml](/docs/convox-yml):

```yaml
services:
  web:
    scale:
      count: 4
      cpu: 512
      memory: 1024
```

### Autoscaling

There are two dimension for scale on a Convox Rack:

* The number of instances (servers) running that provide capacity to launch containers
* The number of processes (containers) running for each service

Convox can autoscale in both of these dimensions.

#### Rack Autoscaling

Rack-level autoscaling is enabled by default when you install a new Rack. When Rack-level autoscaling is enabled, the number of instances currently running will continually adjust based on the current container workload.

This level of autoscaling can be adjusted with the Rack parameter `Autoscale`:

    $ convox rack params set Autoscale=Yes

#### Service Autoscaling

Service-level autoscaling is controlled in the `convox.yml`:

```
service:
  web:
    scale:
      count: 1-10
      targets:
        cpu: 70
        memory: 90
        requests: 200
```

Setting scale targets for a service will cause the service-level autoscaler to adjust the number of running processes for a particular service to try to meet the targets you define.

* `cpu`: Average CPU utilization (%) for all processes
* `memory`: Average Memory utilization (%) for all processes
* `requests:` Requests per minute per process

You can also use custom Cloudwatch metrics as a target for the service autoscaler:

```
service:
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

These settings would continually adjust the level of workers to keep the maximum number of messages waiting in the queue to 20.

You can define any or all of these targets for each service. The autoscaler will select the maximum number of processes required to meet all of the defined targets.
