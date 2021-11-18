---
title: "Rack Parameters"
---

## Setting Parameters

Parameters can be set using the following command.

    convox rack params set Foo=bar

You can also set multiple parameters at once.

    convox rack params set Foo=bar Baz=qux

### Ami

Which [Amazon Machine Image](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_container_instance.html) should be used.

### ApiCount

How many Rack API containers to run. Setting this higher than 2 will guarantee better Rack API availability for mission critical clusters.

| Default value | `2` |

### ApiCpu

How much CPU should be reserved by the API web process.

| Default value | `128` |

### ApiMemory

How much memory should be reserved by the API web process.

| Default value | `128` |

### Autoscale

Autoscale rack instances. See our [Scaling doc](/docs/scaling#autoscale) for more information.

| Default value  | `Yes`       |
| Allowed values | `Yes`, `No` |

### AutoscaleExtra

The number of instances of extra capacity that autoscale should keep running.

| Default value  | `1` |

### AvailabilityZones

Override the default availability zones used in a Rack. Please note that updating this parameter once a Rack is installed will require choosing new values for the various SubnetCIDR parameters since two subnets can not exist with the same CIDR simultaneously.

| Default value | *<blank>* |

### BuildCpu

How much CPU should be allocated to builds.

| Default value | `0` |

### BuildInstance

EC2 instance type to create and use as the Rack's [dedicated build instance](/docs/builds/#dedicated-build-instance).

Note: the build instance will also use the [`InstanceBootCommand`](/docs/rack-parameters#instancebootcommand) and [`InstanceRunCommand`](/docs/rack-parameters#instanceruncommand) Rack params, if defined.

| Default value  | *<blank>*                                                        |
| Allowed values | [EC2 Instance Types](https://aws.amazon.com/ec2/instance-types/) |

See also the [InstanceType](#instancetype) Rack parameter.

### BuildMemory

Defines the amount of memory (in MB) that the instance should allocate to build containers for each build.

| Default value  | 1024 |

<div class="alert alert-info">
Getting build errors like <b>Starting build... ERROR: not enough memory available to start process</b>? You should either reduce this parameter, or change the <a href="#instancetype">InstanceType</a> parameter to an <a href="https://aws.amazon.com/ec2/instance-types/">instance type</a> with more memory.
</div>

<div class="alert alert-warning">
Note: If you set BuildMemory to an amount that's more than half of the total memory available to the build instance, you'll only be able to run one build at a time. If this value is too high, builds may fail.
</div>

### BuildVolumeSize

<div class="alert alert-info">
Getting errors like <b>No space left on device</b> on your builds (not your running applications)? You can extend the space on the device by increasing this parameter.
</div>

Default container disk size in GB.

| Default value | `100` |

### ClientId

Anonymous identifier.

| Default value  | `dev@convox.com` |

### ContainerDisk

<div class="alert alert-info">
Getting errors like <b>No space left on device</b> on your running applications (not builds)? You can extend the space on the device by increasing this parameter.
</div>

Default container disk size in GB.

| Default value | `10` |

### EcsPollInterval

How often (in seconds) to poll ECS for service updates(to inject into the app logs.

| Default value  | `1` |

### EncryptEbs

Enable encryption at rest for EBS volumes.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

### Encryption

Encrypt secrets with KMS.

| Default value    | `Yes`       |
| Permitted values | `Yes`, `No` |

### ExistingVpc

Existing VPC ID (if blank, a VPC will be created).

### HttpProxy

HTTP proxy for outbound HTTP connections (for network-restricted Racks).

Set this value to the hostname (or IP address) and port number of an HTTP proxy to use for the ECS agent to connect to the internet.

| Default value | *<blank>* |

For more information, see [HTTP Proxy Configuration](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/http_proxy_config.html) in the AWS docs.

### InstanceBootCommand

A single line of shell script to run (as root) as a cloud-init command early during instance boot.

For more information about using cloud-init with EC2, see the AWS doc [Running Commands on Your Linux Instance at Launch](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html#user-data-cloud-init). For cloud-init specifics, see "bootcmd" in the doc [Run commands on first boot](http://cloudinit.readthedocs.io/en/latest/topics/examples.html#run-commands-on-first-boot).

The `InstanceBootCommand` will also apply to any [build instance](/docs/rack-parameters#buildinstance) associated with the Rack.

| Default value | *<blank>* |

#### Example: increase virtual memory

Tell the host EC2 instance to set operating system limits on mmap to `262144` (useful for ElasticSearch):

```
$ convox rack params set 'InstanceBootCommand="sysctl -w vm.max_map_count=262144"'
Updating parameters... OK
```

(Note the surrounding single quotes in the above command.)

### InstanceCount

The number of EC2 instances in your Rack cluster.

| Default value | `3`    |
| Minimum value | `3`    |

### InstancePolicy

ARN of an additional IAM policy to add to the instance-level role.

| Default value | *<blank>* |

### InstanceRunCommand

A single line of shell script to run as a cloud-init command late during instance boot.

For more information about using cloud-init with EC2, see the AWS doc [Running Commands on Your Linux Instance at Launch](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html#user-data-cloud-init). For cloud-init specifics, see "runcmd" in the doc [Run commands on first boot](http://cloudinit.readthedocs.io/en/latest/topics/examples.html#run-commands-on-first-boot).

The `InstanceRunCommand` will also apply to any [build instance](/docs/rack-parameters#buildinstance) associated with the Rack.

| Default value | *<blank>* |

### InstanceType

The type of EC2 instance to run in your Rack cluster.

| Default value  | `t2.small`                                                       |
| Allowed values | [EC2 Instance Types](https://aws.amazon.com/ec2/instance-types/) |

See also the [BuildInstance](#buildinstance) Rack parameter.

### InstanceUpdateBatchSize

The number of instances to update in a batch.

| Default value | `1`    |
| Minimum value | `1`    |

### HighAvailability

<div class="alert alert-warning">
This parameter cannot be changed after the rack is created.
</div>

Whether to enable High Availability mode. This ensure propper resources redundancy to mitigate system failures.

| Default value  | `true`             |
| Allowed values | `true`, `false`    |

### Internal

Enable the internal load balancer for this Rack. See [Internal Services](/docs/internal-services)

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

### InternetGateway

If installing rack on existing VPC, you need to pass existing InternetGateway ID attached to the VPC. See [ExistingVPC](/reference/rack-parameters#existingvpc).

| Default value | *<blank>* |

### LoadBalancerIdleTimeout

The idle timeout value for the ALB, in seconds. The valid range is 1-4000 seconds.

| Default value  | `3600` |

### OnDemandMinCount

If using spot instances through the [SpotInstanceBid](#spotinstancebid) parameter, this configures the minimum number of on demand instances. This should be set to a value that will guarantee the minimum acceptable service availbility.

| Default value | `3`    |

### Key

SSH key name for access to cluster instances.

## Password

(REQUIRED) API HTTP password.

| Minimum length  | 1  |
| Maximum length  | 50 |

### Private

Have the Rack create non-publicly routable resources, i.e. in a private subnet. See our [Private Networking doc](/docs/private-networking/) for more information.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

### PrivateApi

Put Rack API Load Balancer in a private network, i.e. have the Rack API use an Internal ELB, making it unreachable from the internet.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

### PrivateApiSecurityGroup

Specify a custom security group that can connect to the Rack API when `PrivateApi=Yes`.

| Default value  | *<blank>* |

### PrivateBuild

Place only the build instances into a private network (unused if `Private` is `Yes`)

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

### RouterMitigationMode

Determines how the load balancer handles requests that might pose a security risk to your application.  See [here](https://aws.amazon.com/about-aws/whats-new/2020/08/application-and-classic-load-balancers-adding-defense-in-depth-with-introduction-of-desync-mitigation-mode/) for more information.

| Default value  | `defensive`                         |
| Allowed values | `defensive`, `monitor`, `strictest` |

### RouterSecurityGroup

Specify a custom security group to use for the Rack's router.

| Default value  | *<blank>* |

### SpotInstanceBid

A value, in dollars, that you want to pay for spot instances. If spot instances are available for the bid price, the Rack instances will use spot instances instead of on demand instances, resulting in significant cost savings. If the parameter is empty, spot instances will not be utilized. This should be used with the [OnDemandMinCount](#ondemandmincount) parameter to guarantee some on demand instances are running if spot instances are not available.

| Default value  | *<blank>* |

### SslPolicy

Specify an SSL policy for the primary Rack load balancer.

| Default value  | *<blank>* |
| Allowed values | [ELB SSL Policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) |

### Subnet0CIDR

Public Subnet 0 CIDR Block.

| Default value | `10.0.1.0/24` |

### Subnet1CIDR

Public Subnet 1 CIDR Block.

| Default value | `10.0.2.0/24` |

### Subnet2CIDR

Public Subnet 2 CIDR Block.

| Default value | `10.0.3.0/24` |

### SubnetPrivate0CIDR

Private Subnet 0 CIDR Block.

| Default value | `10.0.4.0/24` |

### SubnetPrivate1CIDR

Private Subnet 1 CIDR Block.

| Default value | `10.0.5.0/24` |

### SubnetPrivate2CIDR

Private Subnet 2 CIDR Block.

| Default value | `10.0.6.0/24` |

### SwapSize

Default swap volume size in GB. Set this value to 0 to disable swap.

| Default value | `5` |

### Tenancy

Dedicated hardware.

| Default value  | `default`              |
| Allowed values | `default`, `dedicated` |

### Version

(REQUIRED) Convox release version.

| Minimum length | 1 |

### VolumeSize

Default disk size (in gibibytes) of the EBS volume attached to each EC2 instance in the cluster.

| Default value | `50` |

### VPCCIDR

VPC CIDR Block. Note that changing this has no effect since VPC CIDR ranges cannot be changed after they're created.

| Default value | `10.0.0.0/16` |
