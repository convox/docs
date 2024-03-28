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

Override the default availability zones used in a Rack. Please note that updating this parameter once a Rack is installed will require setting `MaxAvailabilityZones` to the new AZs quantity you are choosing.

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

### CpuCredits

The credit option for CPU usage of a T instance.

| Allowed values  | `standard`, `unlimited` |

### DisableALBPort80

Disable exposing 80 port on ALB

| Default value  | `No`       |
| Allowed values  | `Yes`, `No` |

### DynamoDbTableDeletionProtectionEnabled

Determines if a dynamodb table is protected from deletion. When enabled, the table cannot be deleted by any user or process. This setting is disabled by default.

| Default value  | `false`       |
| Allowed values  | `true`, `false` |

### DynamoDbTablePointInTimeRecoveryEnabled

Indicates whether point in time recovery is enabled or disabled on the dynamodb table.

| Default value  | `false`       |
| Allowed values  | `true`, `false` |


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

Existing VPC-ID from AWS, if blank a VPC will be created. Additional paramater [InternetGateway](/reference/rack-parameters#internetgateway) must be set to use **ExistingVPC**.

| Default value    |*<blank>*|
| Permitted values |VPC ID|

### EnableS3Versioning

Enable s3 bucket versioning. This affects all the buckets created for this rack.

| Default value  | `Suspended`       |
| Allowed values  | `Enabled`, `Suspended` |

### EnableContainerReadonlyRootFilesystem

Enable container readonly root filesystem. Enabling this will remove write access to the root filesystem.

| Default value  | `No`       |
| Allowed values  | `Yes`, `No` |

### EnableSharedEFSVolumeEncryption

This will enable AWS KMS encryption on the default shared EFS volume used for application [volumes](/application/volumes).

| Default value  | `false`       |
| Allowed values  | `true`, `false` |

<div class="alert alert-warning">
<b>Important:</b> Enabling <code>EnableSharedEFSVolumeEncryption</code> will recreate the EFS volume and <strong>all</strong> application's shared volume data will be lost. To preserve data, it is crucial to follow these steps:
<ul>
<li><b>Backup:</b> Use AWS Backup or a similar tool to create a snapshot of the existing Amazon EFS volume, ensuring all current data is securely copied.</li>
<li><b>Restore:</b> After enabling encryption, restore your data from the backup snapshot to the new encrypted EFS volume.</li>
</ul>
</div>




### HighAvailability

<div class="alert alert-warning">
This parameter cannot be changed after the rack is created.
</div>

Whether or not enable High Availability mode, choose between failure resiliency and cost efficiency. This ensure proper resources redundancy to mitigate system failures.

If HighAvailability is set to true, the [InstanceCount](#instancecount) is used as initial cluster size. If false, the `NoHaInstanceCount` is used as initial cluster size. Both can be scaled to 1000 instances.

| Default value  | `true`          |
| Allowed values | `true`, `false` |

### HttpProxy

HTTP proxy for outbound HTTP connections (for network-restricted Racks).

Set this value to the hostname (or IP address) and port number of an HTTP proxy to use for the ECS agent to connect to the internet.

| Default value | *<blank>* |

For more information, see [HTTP Proxy Configuration](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/http_proxy_config.html) in the AWS docs.

The `HttpProxy`param will not change how the apps access the internet, the traffic won't go through the proxy. Only the instances requests will use the proxy for outbound connections, if you want your apps to use the proxy, you have to configure it on the convox.yml:

```
services:
  web:
    build: .
    port: 3000
    environment:
      - http_proxy=10.0.1.124:8888
      - https_proxy=10.0.1.124:8888
      - HTTP_PROXY=10.0.1.124:8888
      - HTTPS_PROXY=10.0.1.124:8888
      - NO_PROXY=169.254.170.2
```
### IMDSHttpPutResponseHopLimit

Specifies the maximum number of network hops that PUT response packets are allowed to travel from the EC2 instance metadata service (IMDS) to the requesting instance. This parameter is particularly relevant when `IMDSHttpTokens` is set to `required`, ensuring enhanced security by enforcing IMDSv2.

| Default value  | `1`            |
| Allowed values  | Numerical values (e.g., `1`, `2`) |

<div class="alert alert-info">
When <code>IMDSHttpTokens</code> is set to <code>required</code>, some configurations may require increasing the <code>IMDSHttpPutResponseHopLimit</code> to ensure proper functionality. If encountering connectivity issues with applications requiring IMDSv2, consider setting <code>IMDSHttpPutResponseHopLimit</code> to <code>2</code>. This adjustment helps facilitate necessary communications with the instance metadata service.
</div>

<div class="alert alert-warning">
Note: Adjusting the <code>IMDSHttpPutResponseHopLimit</code> above the default value should be done with understanding of your network topology and the security implications. Always verify that changes do not compromise your instance's security posture.
</div>

### IMDSHttpTokens

Set how your instances will access the instance metadata. You can set EC2 instances to use only v2 by setting IMDSHttpTokens as 'required', see [Configure the instance metadata options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-options.html#configuring-IMDS-new-instances).

| Default value  | `optional`             |
| Allowed values | `optional`, `required` |

### InstancesIpToIncludInWhiteListing
To auto include build and instances ips to whitelist when rack is public and whitelist is enabled.

| Default value  | `Both`             |
| Allowed values | `Both` , `Build`, `Workload`, `None`` |

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

The number of EC2 instances in your Rack cluster. This parameter is only used for clusters with HighAvailablity = true.

| Default value | `3` |
| Minimum value | `3` |

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

| Default value | `1` |
| Minimum value | `1` |

### Internal

Enable the internal load balancer for this Rack. See [Internal Services](/docs/internal-services)

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

### InternetGateway

If installing rack on existing VPC, you need to pass existing InternetGateway ID attached to the VPC. See [ExistingVPC](/reference/rack-parameters#existingvpc).

| Default value | *<blank>* |

### Key

SSH key name for access to cluster instances.

| Default value  | *<blank>* |

### LoadBalancerIdleTimeout

The idle timeout value for the ALB, in seconds. The valid range is 1-4000 seconds.

| Default value  | `3600` |

### LogDriver

Log driver used by the rack and services to send logs. Default to CloudWatch. You must provide the SyslogDestination when setting as Syslog. It disable logs if blank.

**Attention!!** Disabling CloudWatch will impact `convox logs` and `convox rack logs`. Use Syslog resource if you still want to use convox logs, see [Resource Syslog](/deployment/syslogs)

| Default value  | `CloudWatch` |
| Allowed values | `CloudWatch`, `Syslog`, *<blank>* |

### LogRetention

Number of days to keep logs (blank for unlimited). Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 2192, 2557, 2922, 3288, and 3653. See [Logs retention](/management/logs#retention).

| Default value  | `7` |

### MaintainTimerState

To maintain the state of timer if it's disabled/enabled in AWS console event rule. After deploying the timer, if you disable the timer in the AWS console event rule, this will keep it disabled unless it is recreated. By default it is not maintained explicitly and the behaviour of the timer state change(if you disable event rule in the AWS console) will be depened on the cloudformation.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

### MaxAvailabilityZones

The maximum number of Availability Zones that the cluster should use.

| Default value | `3` |
| Allowed values | `2`, `3` |

### NoHaInstanceCount

The number of EC2 instances in your non High Availability Rack cluster. It's only used for non high available clusters.

| Default value | `1` |
| Minimum value | `1` |

### OnDemandMinCount

If using spot instances through the [SpotInstanceBid](#spotinstancebid) parameter, this configures the minimum number of on demand instances. This should be set to a value that will guarantee the minimum acceptable service availability. You must set it even if you using the HighAvailability as `false`, as this will be used to create the minimum on demand instances.

| Default value | `3` |

### Password

(REQUIRED) API HTTP password.

| Minimum length  | 1  |
| Maximum length  | 50 |

### PlaceLambdaInVpc

Place convox related lambdas in vpc if rack is private.

| Default value  | `No`        |
| Allowed values | `Yes`, `No` |

### InternalRouterSuffix

Suffix for internal router domain

| Default value  | `-rti`        |

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


### PruneOlderImagesInHour

To prune docker images older than this specified hours.

| Default value  | `96`        |

### PruneOlderImagesCronRunFreq

Cron frequecy to prune docker older images.

| Default value  | `daily`        |
| Allowed values | `hourly`, `daily`, `weekly` |

### RouterMitigationMode

Determines how the load balancer handles requests that might pose a security risk to your application.  See [here](https://aws.amazon.com/about-aws/whats-new/2020/08/application-and-classic-load-balancers-adding-defense-in-depth-with-introduction-of-desync-mitigation-mode/) for more information.

| Default value  | `defensive`                         |
| Allowed values | `defensive`, `monitor`, `strictest` |

### RouterSecurityGroup

Specify a custom security group to use for the Rack's router.

| Default value  | *<blank>* |


### ScheduleRackScaleDown & ScheduleRackScaleUp

Use ScheduleRackScaleDown & ScheduleRackScaleUp if you want to turn the rack on/off based on a schedule. Keep in mind that both parameters need to be set.
To turn your rack off on weekends and back on during weekdays you can use:

```
convox rack params set ScheduleRackScaleDown="0 18 * * 5" ScheduleRackScaleUp="0 9 * * 1"
```

The supported cron expression format consists of five fields separated by white spaces: [Minute] [Hour] [Day_of_Month] [Month_of_Year] [Day_of_Week]. In the example above it's configured to shutdown every Friday (5th day) at 6pm (UTC). More details on the CRON format can be found in [Crontab](http://crontab.org/) and [examples](https://crontab.guru/examples.html).

You can see details about the Scheduling Actions on AWS [doc](https://docs.aws.amazon.com/autoscaling/ec2/userguide/schedule_time.html).

| Default value  | *<blank>* |

### SpotInstanceBid

A value, in dollars, that you want to pay for spot instances. If spot instances are available for the bid price, the Rack instances will use spot instances instead of on demand instances, resulting in significant cost savings. If the parameter is empty, spot instances will not be utilized. This must be used with the [OnDemandMinCount](#ondemandmincount) parameter to guarantee some on demand instances are running if spot instances are not available (even if the HighAvailability is `false`, if not set will use the default).

| Default value  | *<blank>* |

### SpotFleetAllowedInstanceTypes

Comma-separated list of allowed instance types in the Spot Fleet. It can not be used with SpotFleetExcludedInstanceTypes, it takes precedent over it. The following are examples: m5.8xlarge, c5*.*, m5a.*, r*, *3*. This parameter can be used only when [SpotFleet](#spotfleetmaxprice) is enabled.  

| Default value  | *<blank>* |

### SpotFleetExcludedInstanceTypes

Comma-separated list of excluded instance types in the Spot Fleet. . It can not be used with SpotFleetAllowedInstanceTypes. The following are examples: m5.8xlarge, c5*.*, m5a.*, r*, *3*. This parameter can be used only when [SpotFleet](#spotfleetmaxprice) is enabled.

| Default value  | *<blank>* |

### SpotFleetAllocationStrategy

The Spot Fleet allocation strategy. This parameter can be used only when [SpotFleet](#spotfleetmaxprice) is enabled.

| Allowed values | `lowestPrice`, `diversified`, `capacityOptimized` |

| Default value | `lowestPrice` |
  
### SpotFleetMaxPrice

The maximum price for instances in the Spot Fleet per hour. It will try to launch instances untill it crosses the price even if target [InstanceCount](#instancecount) or [NoHaInstanceCount](#nohainstancecount) is not fullfilled. Setting this parameter will enable spotfleet which will use the AWS Spot request to fullfill the instance demand and will be manage by the spot request(not autoscaling group). Currently this has only single zone support even if you set **HighAvailability** to `true`. **SpotFleetMinOnDemandCount** will be used to lauch ondemand instances along with these spot instances.

### SpotFleetMinMemoryMiB

Spot fleet's min memory in MiB. This parameter can be used only when [SpotFleet](#spotfleetmaxprice) is enabled.

| Default value | `1000` |

### SpotFleetMinVcpuCount

Spot fleet's min vcpu count. This parameter can be used only when [SpotFleet](#spotfleetmaxprice) is enabled.

| Default value | `0` |

### SpotFleetMinOnDemandCount

Spot fleet's minimum on demand instance count. Instance type will taken from [InstanceType](#instancetype) param. This parameter can be used only when [SpotFleet](#spotfleetmaxprice) is enabled.

| Default value (if HighAvailability disabled) | `1` |

| Default value (if HighAvailability enabled) | `2` |

### SpotFleetTargetType

The unit type used for the Spot Fleet target capacity. This parameter can be used only when [SpotFleet](#spotfleetmaxprice) is enabled.

| Allowed values | `memory-mib`, `units`, `vcpu` |

| Default value | `units` |

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

### SyslogDestination

Syslog address destination, you need to pass the protocol to be used, e.g. `tcp+tls://logsX.syslog.com:1234`.

| Default value | *<blank>* |

### SyslogFormat

Syslog format (low case) to sent to SyslogDestination. See [Docker Syslog](https://docs.docker.com/config/containers/logging/syslog/) and [RFC5424](https://www.rfc-editor.org/rfc/rfc5424#section-6).

| Default value | `rfc5424` |

### Tags

Custom tags to add with AWS resource

| Default value  | *<blank>* |
| Format | `<key>=<val>,<key>=<val>`. example: `key1=val1,key2=val2` |

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

### WhiteList

Comma delimited list of CIDRs, e.g. `10.0.0.0/24,172.10.0.1/32`, to allow access to the rack api.

**Attention!!** Please be careful to consider all required connections to the rack API before enabling Whitelist. You can block your access and ability to edit this parameter from CLI again if misconfigured.

| Default value | *<blank>* |
