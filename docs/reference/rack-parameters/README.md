---
title: "Rack Parameters"
description: "Reference for all Rack-level parameters that configure Convox infrastructure, networking, scaling, and security on AWS."
---

# Rack Parameters

## Setting Parameters

Parameters can be set using the following command:

```bash
$ convox rack params set Foo=bar
```

You can also set multiple parameters at once:

```bash
$ convox rack params set Foo=bar Baz=qux
```

## Parameters

| Parameter | Default | Description |
|:----------|:--------|:------------|
| [Ami](/reference/rack-parameters/Ami) | "" | Custom Amazon Machine Image for Rack instances |
| [ApiCount](/reference/rack-parameters/ApiCount) | `2` | Number of Rack API containers to run |
| [ApiCpu](/reference/rack-parameters/ApiCpu) | `128` | CPU units reserved by the API web process |
| [ApiMonitorMemory](/reference/rack-parameters/ApiMonitorMemory) | `128` | Memory (MB) reserved by the API monitor process |
| [ApiRouter](/reference/rack-parameters/ApiRouter) | `ELB` | Legacy load balancer type label for the Rack API |
| [ApiWebMemory](/reference/rack-parameters/ApiWebMemory) | `256` | Memory (MB) reserved by the API web process |
| [Autoscale](/reference/rack-parameters/Autoscale) | `Yes` | Enable or disable Rack instance autoscaling |
| [AutoscaleExtra](/reference/rack-parameters/AutoscaleExtra) | `1` | Extra capacity instances maintained by autoscale |
| [AvailabilityZones](/reference/rack-parameters/AvailabilityZones) | "" | Override the default Availability Zones |
| [BuildCpu](/reference/rack-parameters/BuildCpu) | `256` | CPU units allocated to builds |
| [BuildImage](/reference/rack-parameters/BuildImage) | "" | Override the default builder image |
| [BuildInstance](/reference/rack-parameters/BuildInstance) | `t3.small` | EC2 instance type for the dedicated build instance |
| [BuildInstancePolicy](/reference/rack-parameters/BuildInstancePolicy) | "" | Additional IAM policy ARN for build instances |
| [BuildInstanceSecurityGroup](/reference/rack-parameters/BuildInstanceSecurityGroup) | "" | Security group for build instances |
| [BuildMemory](/reference/rack-parameters/BuildMemory) | `1000` | Memory (MB) allocated to builds |
| [BuildMethod](/reference/rack-parameters/BuildMethod) | `ec2` | Build process type (EC2 or Fargate) |
| [BuildVolumeSize](/reference/rack-parameters/BuildVolumeSize) | `100` | Build disk size in GB |
| [CpuCredits](/reference/rack-parameters/CpuCredits) | "" | CPU credit option for T instances |
| [DefaultAmi](/reference/rack-parameters/DefaultAmi) | `.../amazon-linux-2/recommended/image_id` | Default AMI for x86_64 instances |
| [DefaultAmiArm](/reference/rack-parameters/DefaultAmiArm) | `.../amazon-linux-2/arm64/recommended/image_id` | Default AMI for ARM64 instances |
| [DisableALBPort80](/reference/rack-parameters/DisableALBPort80) | `No` | Disable port 80 on the ALB |
| [DynamoDbTableDeletionProtectionEnabled](/reference/rack-parameters/DynamoDbTableDeletionProtectionEnabled) | `false` | Enable DynamoDB table deletion protection |
| [DynamoDbTablePointInTimeRecoveryEnabled](/reference/rack-parameters/DynamoDbTablePointInTimeRecoveryEnabled) | `false` | Enable DynamoDB point-in-time recovery |
| [EcsContainerStopTimeout](/reference/rack-parameters/EcsContainerStopTimeout) | "" | Custom ECS container stop timeout in seconds |
| [EcsPollInterval](/reference/rack-parameters/EcsPollInterval) | `1` | ECS service update poll interval in seconds |
| [EnableContainerReadonlyRootFilesystem](/reference/rack-parameters/EnableContainerReadonlyRootFilesystem) | `No` | Enable read-only root filesystem for containers |
| [EnableS3Versioning](/reference/rack-parameters/EnableS3Versioning) | `Suspended` | Enable S3 bucket versioning for Rack buckets |
| [EnableSharedEFSVolumeEncryption](/reference/rack-parameters/EnableSharedEFSVolumeEncryption) | `false` | Enable KMS encryption on the shared EFS volume |
| [EncryptEbs](/reference/rack-parameters/EncryptEbs) | `No` | Encrypt EBS volumes at rest |
| [Encryption](/reference/rack-parameters/Encryption) | `Yes` | Encrypt secrets with KMS |
| [ExistingVpc](/reference/rack-parameters/ExistingVpc) | "" | Use an existing VPC by ID |
| [FargateBuildCpu](/reference/rack-parameters/FargateBuildCpu) | "" | CPU for Fargate builds |
| [FargateBuildMemory](/reference/rack-parameters/FargateBuildMemory) | "" | Memory for Fargate builds |
| [HighAvailability](/reference/rack-parameters/HighAvailability) | `true` | Enable High Availability mode (immutable after creation) |
| [HttpProxy](/reference/rack-parameters/HttpProxy) | "" | HTTP proxy for outbound connections |
| [ImagePullBehavior](/reference/rack-parameters/ImagePullBehavior) | `default` | Docker image pull behavior |
| [IMDSHttpPutResponseHopLimit](/reference/rack-parameters/IMDSHttpPutResponseHopLimit) | `1` | IMDS HTTP PUT response hop limit |
| [IMDSHttpTokens](/reference/rack-parameters/IMDSHttpTokens) | `optional` | Require IMDSv2 tokens |
| [InstanceBootCommand](/reference/rack-parameters/InstanceBootCommand) | "" | Cloud-init boot command for instances |
| [InstanceCount](/reference/rack-parameters/InstanceCount) | `3` | Number of EC2 instances (HA mode) |
| [InstancePolicy](/reference/rack-parameters/InstancePolicy) | "" | Additional IAM policy ARN for cluster instances |
| [InstanceRunCommand](/reference/rack-parameters/InstanceRunCommand) | "" | Cloud-init run command for instances |
| [InstanceSecurityGroup](/reference/rack-parameters/InstanceSecurityGroup) | "" | Security group for ECS instances |
| [InstancesIpToIncludInWhiteListing](/reference/rack-parameters/InstancesIpToIncludInWhiteListing) | `Both` | Auto-whitelist instance IPs |
| [InstanceType](/reference/rack-parameters/InstanceType) | `t3.small` | EC2 instance type for the Rack cluster |
| [InstanceUpdateBatchSize](/reference/rack-parameters/InstanceUpdateBatchSize) | `1` | Instances to update per batch during rolling updates |
| [Internal](/reference/rack-parameters/Internal) | `No` | Enable the internal load balancer |
| [InternalOnly](/reference/rack-parameters/InternalOnly) | `No` | Only support internal applications |
| [InternalRouterSuffix](/reference/rack-parameters/InternalRouterSuffix) | `-rti` | Suffix for internal router domain |
| [InternetGateway](/reference/rack-parameters/InternetGateway) | "" | Internet Gateway ID for an existing VPC |
| [Key](/reference/rack-parameters/Key) | "" | SSH key name for cluster access |
| [LoadBalancerIdleTimeout](/reference/rack-parameters/LoadBalancerIdleTimeout) | `3600` | ALB idle timeout in seconds |
| [LogBucket](/reference/rack-parameters/LogBucket) | "" | S3 bucket for load balancer and Rack logs |
| [LogDriver](/reference/rack-parameters/LogDriver) | `CloudWatch` | Log driver for Rack and services |
| [LogRetention](/reference/rack-parameters/LogRetention) | `7` | Days to retain logs |
| [MaintainTimerState](/reference/rack-parameters/MaintainTimerState) | `No` | Maintain timer enable/disable state across deploys |
| [MaxAvailabilityZones](/reference/rack-parameters/MaxAvailabilityZones) | `3` | Maximum Availability Zones to use |
| [NoHAAutoscaleExtra](/reference/rack-parameters/NoHAAutoscaleExtra) | `0` | Extra autoscale capacity for non-HA Racks |
| [NoHaInstanceCount](/reference/rack-parameters/NoHaInstanceCount) | `1` | Number of instances for non-HA Racks |
| [OnDemandMinCount](/reference/rack-parameters/OnDemandMinCount) | `3` | Minimum on-demand instances when using spot |
| [Password](/reference/rack-parameters/Password) | *(required)* | API HTTP password |
| [PlaceLambdaInVpc](/reference/rack-parameters/PlaceLambdaInVpc) | `No` | Place Convox Lambda functions inside the VPC |
| [Private](/reference/rack-parameters/Private) | `No` | Create resources in private subnets |
| [PrivateApi](/reference/rack-parameters/PrivateApi) | `No` | Place the Rack API load balancer in a private network |
| [PrivateApiSecurityGroup](/reference/rack-parameters/PrivateApiSecurityGroup) | "" | Security group for private API access |
| [PrivateBuild](/reference/rack-parameters/PrivateBuild) | `No` | Place build instances in a private network |
| [PruneOlderImagesCronRunFreq](/reference/rack-parameters/PruneOlderImagesCronRunFreq) | `daily` | Cron frequency for Docker image pruning |
| [PruneOlderImagesInHour](/reference/rack-parameters/PruneOlderImagesInHour) | `96` | Prune Docker images older than this many hours |
| [RouterInternalSecurityGroup](/reference/rack-parameters/RouterInternalSecurityGroup) | "" | Security groups for the internal router |
| [RouterMitigationMode](/reference/rack-parameters/RouterMitigationMode) | `defensive` | Load balancer desync mitigation mode |
| [RouterSecurityGroup](/reference/rack-parameters/RouterSecurityGroup) | "" | Custom security group for the Rack router |
| [ScheduleRackScaleDown](/reference/rack-parameters/ScheduleRackScaleDown) | "" | Cron schedule to scale down the Rack |
| [ScheduleRackScaleUp](/reference/rack-parameters/ScheduleRackScaleUp) | "" | Cron schedule to scale up the Rack |
| [SpotFleetAllocationStrategy](/reference/rack-parameters/SpotFleetAllocationStrategy) | `lowestPrice` | Spot Fleet allocation strategy |
| [SpotFleetAllowedInstanceTypes](/reference/rack-parameters/SpotFleetAllowedInstanceTypes) | "" | Allowed instance types for Spot Fleet |
| [SpotFleetExcludedInstanceTypes](/reference/rack-parameters/SpotFleetExcludedInstanceTypes) | "" | Excluded instance types for Spot Fleet |
| [SpotFleetMaxPrice](/reference/rack-parameters/SpotFleetMaxPrice) | "" | Maximum price per hour for Spot Fleet |
| [SpotFleetMinMemoryMiB](/reference/rack-parameters/SpotFleetMinMemoryMiB) | `1000` | Minimum memory (MiB) for Spot Fleet instances |
| [SpotFleetMinOnDemandCount](/reference/rack-parameters/SpotFleetMinOnDemandCount) | `1` (`2` if HA) | Minimum on-demand instances in Spot Fleet |
| [SpotFleetMinVcpuCount](/reference/rack-parameters/SpotFleetMinVcpuCount) | `0` | Minimum vCPU count for Spot Fleet instances |
| [SpotFleetTargetType](/reference/rack-parameters/SpotFleetTargetType) | `units` | Unit type for Spot Fleet target capacity |
| [SpotInstanceBid](/reference/rack-parameters/SpotInstanceBid) | "" | Bid price in dollars for spot instances |
| [SslPolicy](/reference/rack-parameters/SslPolicy) | "" | SSL policy for the Rack load balancer |
| [Subnet0CIDR](/reference/rack-parameters/Subnet0CIDR) | `10.0.1.0/24` | Public Subnet 0 CIDR block |
| [Subnet1CIDR](/reference/rack-parameters/Subnet1CIDR) | `10.0.2.0/24` | Public Subnet 1 CIDR block |
| [Subnet2CIDR](/reference/rack-parameters/Subnet2CIDR) | `10.0.3.0/24` | Public Subnet 2 CIDR block |
| [SubnetPrivate0CIDR](/reference/rack-parameters/SubnetPrivate0CIDR) | `10.0.4.0/24` | Private Subnet 0 CIDR block |
| [SubnetPrivate1CIDR](/reference/rack-parameters/SubnetPrivate1CIDR) | `10.0.5.0/24` | Private Subnet 1 CIDR block |
| [SubnetPrivate2CIDR](/reference/rack-parameters/SubnetPrivate2CIDR) | `10.0.6.0/24` | Private Subnet 2 CIDR block |
| [SwapSize](/reference/rack-parameters/SwapSize) | `5` | Default swap volume size in GB |
| [SyslogDestination](/reference/rack-parameters/SyslogDestination) | "" | Syslog endpoint address |
| [SyslogFormat](/reference/rack-parameters/SyslogFormat) | `rfc5424` | Syslog message format |
| [Tags](/reference/rack-parameters/Tags) | "" | Custom tags for AWS resources |
| [Tenancy](/reference/rack-parameters/Tenancy) | `default` | EC2 instance tenancy (default or dedicated) |
| [Version](/reference/rack-parameters/Version) | *(required)* | Convox release version |
| [VolumeSize](/reference/rack-parameters/VolumeSize) | `50` | EBS volume size in GB per instance |
| [VPCCIDR](/reference/rack-parameters/VPCCIDR) | `10.0.0.0/16` | VPC CIDR block |
| [WhiteList](/reference/rack-parameters/WhiteList) | "" | CIDR allowlist for Rack API access |

## See Also

- [App Parameters](/reference/app-parameters)
- [Rack Statuses](/reference/rack-statuses)
- [Rack Updates](/management/rack-updates)
- [Scaling](/scaling/scaling)
