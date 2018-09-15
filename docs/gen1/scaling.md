---
title: "Scaling"
---

Convox allows you to scale your application's concurrency, memory allocation, and the resources available in the underlying Rack.

### Scaling an application

#### Show current application scaling

```
$ convox scale
NAME  DESIRED  RUNNING  MEMORY
web   2        1        256
redis 1        1        256
```

#### Concurrency

```
$ convox scale web --count=4
NAME  DESIRED  RUNNING  MEMORY
web   2        1        256
```

#### Memory

```
$ convox scale web --memory=1024
NAME  DESIRED  RUNNING  MEMORY
web   2        1        1024
```

#### CPU

Each rack instance has 1,024 [cpu units](http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html#ECS-Type-ContainerDefinition-cpu) for every CPU core. This parameter specifies the minimum amount of CPU to reserve for a container. Containers share unallocated CPU units with other containers on the instance with the same ratio as their allocated amount.

```
$ convox scale web --cpu=1024
NAME  DESIRED  RUNNING  CPU
web   1        1        1024
```

#### Scaling down unused services

It's often convenient to run a service like Redis in a container locally. You can do so by defining a `redis` process in your `docker-compose.yml`. However, when you've deployed the app to your rack, you should use a hosted resource like ElastiCache. In this case, you can scale redis down and destroy the ELB which was created:

```
$ convox scale redis --count=-1
NAME  DESIRED  RUNNING  MEMORY
redis   -1        1        256
```

Note: If you scale this service back up, the ELB will be recreated, but will have a different domain name associated with it. If you want to scale a service down, but keep the ELB, you can set `--count=0`.

### Scaling the Rack

You can define both the type and count of instances being run in your Rack.

```
$ convox rack scale --type=m4.xlarge --count=3
Name     demo
Status   updating
Version  20160409181028
Count    3
Type     m4.xlarge
```
<div class="block-callout block-show-callout type-warning" markdown="1">
  The minimum instance count for a Rack is 3. See the [PR](https://github.com/convox/rack/pull/1395#issuecomment-261961713) for details.
</div>

#### Autoscale

Your Rack can scale its own instance count based on the needs of the containers it provisions. Autoscaling is enabled by default. To disable it, set the `Autoscale` parameter:

```
$ convox rack params set Autoscale=No
```

To monitor for autoscaling events, use `convox rack logs` with the `--filter` option.

```
$ convox rack logs --filter="autoscaleRack change="
```

##### Under the hood

Every minute, your Rack runs an autoscale calculation to determine how many instances you need in your cluster. This calculation involves ports, memory, and CPU required by your services. When appropriate, autoscale will update your Rack instance count via a CloudFormation stack update. Autoscale will not change your instance type.

During a deployment, the calculation gets more nuanced, since processes from an old release and a new release will temporarily run at the same time. This is known as a rolling deployment or [rolling update](/docs/gen1/rolling-updates). In ECS terms, this translates to having tasks from both the primary (new) deployment of each service and the active deployment (the one being replaced) of each service running at the same time. Autoscaling will take into account the number of instances needed to run the processes from both releases, i.e. the tasks in both primary and active ECS service deployments.

When a deployment finishes, the old ECS tasks get terminated, and autoscale scales the Rack back down to the original instance count. This scaling down happens gradually--one instance at a time, every 5 minutes--to give ECS time to rebalance tasks across the instances in your cluster.

##### Why does my Rack keep autoscaling?

If your Rack shows more autoscaling activity than expected, there are a few possible explanations.

First, note that any services with open ports you have running at scale of `n` will result in `n+1` instances. This is to allow for rolling deployments.

Autoscaling does not yet take the subtleties of [deployment minimum/maximum](/docs/gen1/docker-compose-labels/#convoxdeployment) into account. For example: a `web` service listening on port 80 with a scale count of 3 will still require 4 instances, even if it has `convox.deployment.minimum=50` set.

You can also look for anomalies in the Rack's autoscaling log events with `convox rack logs --filter=autoscale`.
