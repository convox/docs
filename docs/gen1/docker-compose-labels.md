---
title: "Docker Compose Labels"
---

Convox uses [Docker Compose Labels](https://docs.docker.com/compose/compose-file/#/labels) to add metadata to containers. These labels serve as a convenient way to specify Convox-specific configuration alongside the rest of your more standard container configuration.

<pre>
labels:
  - <a href="#convoxagent">convox.agent</a>
  - <a href="#convoxbalancer">convox.balancer</a>
  - <a href="#convoxcron">convox.cron.&lt;task name&gt;</a>
  - <a href="#convoxdeployment">convox.deployment.maximum</a>
  - <a href="#convoxdeployment">convox.deployment.minimum</a>
  - <a href="#convoxdrainingtimeout">convox.draining.timeout</a>
  - <a href="#convoxenvironmentsecure">convox.environment.secure</a>
  - <a href="#convoxhealth">convox.health.path</a>
  - <a href="#convoxhealth">convox.health.port</a>
  - <a href="#convoxhealth">convox.health.timeout</a>
  - <a href="#convoxidle">convox.idle.timeout</a>
  - <a href="#convoxport">convox.port.&lt;number&gt;.protocol</a>
  - <a href="#convoxport">convox.port.&lt;number&gt;.proxy</a>
  - <a href="#convoxport">convox.port.&lt;number&gt;.secure</a>
  - <a href="#convoxstart">convox.start.shift</a>
</pre>

## convox.agent

The `convox.agent` label allows you to run one process of a service per instance in your cluster, scheduled continuously by a Lambda process. This common pattern is used by popular software like New Relic Infrastructure and DataDog to collect metrics from all of your cloud servers.

    labels:
      - convox.agent=true

See [github.com/convox-examples/dd-agent](https://github.com/convox-examples/dd-agent) for an example of an app you might run as an agent.

## convox.balancer

The `convox.balancer` label allows you to explicitly disable the creation of an load balancer, while still exposing ports in the container and on the host. This is particularly useful for micro-service architectures.

When set to `false`, Convox will not create a load balancer for the given process. As a result, the `convox.health` labels will be ignored and Convox will not monitor the health of your containers (containers that die will still be restarted, however). You may wish to implement your own health-checking system. Additionally, for a port mapping of `80:5000`, the host will expose port `80` instead of selecting a random port for a load balancer to communicate with. Ensuring there are no port conflicts between processes is an exercise left up to the user.

## convox.cron

The `convox.cron` label allows you to schedule recurring tasks for any of your apps. The following example would run a task named `myjob` at 6:30pm UTC every weekday.

    labels:
      - convox.cron.myjob=30 18 ? * MON-FRI bin/myjob

See our [scheduled tasks documentation](/docs/gen1/scheduled-tasks) for more.

## convox.deployment

The `convox.deployment` labels allow you to fine-tune how ECS manages your deployment.

    labels:
      - convox.deployment.minimum=100
      - convox.deployment.maximum=200

Both `minimum` and `maximum` are percentages relative to the desired count for a given process. If your application has a `web` process scaled to a desired count of 4, a `minimum` of 100 would instruct ECS to keep at least 4 `web` processes (or "tasks" in ECS terms) running throughout a deployment. A `maximum` of 200 would allow ECS to run up to 8 processes: 4 old processes and 4 new processes starting up, before ECS kills the old ones.

If you'd like a more in-depth explanation, see the ECS doc [Updating a Service](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/update-service.html).

## convox.draining.timeout

The `convox.draining.timeout` label allows you to specify the amount of time in seconds to allow a draining balancer to keep active connections open. After the timeout, the load balancer will close all connections to a deregistered or unhealthy instance. The minimum value is 1 and the maximum is 3600. The default value is 60. See the [AWS CloudFormation ConnectionDrainingPolicy document](http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html) for more.

    labels:
      - convox.draining.timeout=300

## convox.environment.secure

Setting this label to `true` for a specific service prevents environment variables from being stored in cleartext on that service's ECS Task Definition. This means that your application itself will need to download the environment variables file from S3, decrypt it using the Rack's KMS key, and (if desired) source the variables into the container environment. Please see the [Environment documentation](/docs/gen1/environment/#additional-security) for more info, examples, and tools.

    labels:
      - convox.environment.secure=true

## convox.health

During [rolling updates](/docs/gen1/rolling-updates), Convox will attempt to start a new process and check its health before stopping an old process. These labels allow you to customize the path on your app that will respond to the health checks, the port on which the app will listen for the health check, the number of seconds Convox should wait for a health check response before giving up and trying again, the number of seconds between health checks, and the number of successful or failed health checks before a process is considered healthy or unhealthy.

    labels:
      - convox.health.interval=10
      - convox.health.path=/health_check
      - convox.health.port=3001
      - convox.health.threshold.healthy=3
      - convox.health.threshold.unhealthy=4
      - convox.health.timeout=60

## convox.idle

To customize the [idle timeout of an ELB](http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html), set the `convox.idle.timeout` to a number of seconds between 1 and 3600. Convox defaults to 3600.

    labels:
      - convox.idle.timeout=3000

## convox.port

Use these labels to configure load balancer behavior on specific ports.

    labels:
      - convox.port.<number>.protocol=tls
      - convox.port.<number>.proxy=true
      - convox.port.<number>.secure=true

See our [load balancer documentation](/docs/gen1/load-balancers) for more.

## convox.start.shift

_See also: [the `--shift` flag](/docs/gen1/running-locally/#shifting-ports)_

Use the `convox.start.shift` label to offset the external ports of processes run by `convox start` by a given number. This allows multiple applications to run on one host without conflicts. A container configured to listen on host ports 80 and 443 could be shifted to listen on ports 1080 and 1443 with the following configuration:

    labels:
      - convox.start.shift=1000
