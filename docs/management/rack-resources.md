---
title: "Rack and External Resources"
description: "Create and manage Rack-level resources (S3, SNS, SQS, syslog, webhook, databases, caches) and integrate external third-party services."
---

# Rack and External Resources

Rack Resources are infrastructure components managed at the Rack level through the CLI, independent of any single App. They are available only on cloud-backed Racks (not local development Racks).

For resources defined in `convox.yml` and tied to a specific App, see [App Resources](/application/resources).

## Managing Rack Resources

### Creating a Resource

Create a Rack Resource with default options or pass custom values at creation time:

```bash
$ convox rack resources create postgres
Creating resource... OK, postgres-8458

$ convox rack resources create postgres MultiAZ=true
Creating resource... OK, postgres-2871
```

### Listing Resources

```bash
$ convox rack resources
NAME                  TYPE     STATUS
console-175092fe6ab1  webhook  running
syslog-2984           syslog   running
postgres-8458         postgres running
postgres-2871         postgres running
```

### Viewing Resource Info

```bash
$ convox rack resources info postgres-8458
```

### Viewing Available Options

```bash
$ convox rack resources options memcached
NAME           DEFAULT         DESCRIPTION
InstanceType   cache.t2.micro  The type of instance to use
NumCacheNodes  1               The number of cache clusters for this replication group
```

### Listing Available Resource Types

```bash
$ convox rack resources types
```

### Updating a Resource

```bash
$ convox rack resources update postgres-8458 MultiAZ=true
```

### Linking a Resource to an App

Rack Resources can be linked to Apps, which injects a connection URL as an environment variable:

```bash
$ convox rack resources link postgres-8458 --app myapp
```

To remove the link:

```bash
$ convox rack resources unlink postgres-8458 --app myapp
```

### Deleting a Resource

```bash
$ convox rack resources delete postgres-8458
Deleting postgres-8458... OK
```

## Available Rack Resource Types

### memcached

On AWS, creates an [ElastiCache](https://docs.aws.amazon.com/elasticache/) Memcached cluster.

```bash
$ convox rack resources options memcached
NAME           DEFAULT         DESCRIPTION
InstanceType   cache.t2.micro  The type of instance to use
NumCacheNodes  1               The number of cache clusters for this replication group
```

### redis

On AWS, creates an [ElastiCache](https://docs.aws.amazon.com/elasticache/) Redis cluster.

```bash
$ convox rack resources options redis
NAME                      DEFAULT         DESCRIPTION
AutomaticFailoverEnabled  false           Indicates whether Multi-AZ is enabled. Must be accompanied with NumCacheClusters=2 or higher.
Database                  0               Default database index
Encrypted                 false           Encrypt at rest and in transit
EngineVersion             3.2.6
InstanceType              cache.t2.micro  The type of instance to use
NumCacheClusters          1               The number of cache clusters for this replication group
```

> While memcached and redis are available as Rack Resources, there are advantages to using them as [App Resources](/application/resources) instead: configuration is version controlled in `convox.yml`, and App Resources are automatically available during [review workflows](/console/workflows#review-workflows).

### mysql

On AWS, creates an [RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Welcome.html) MySQL instance.

```bash
$ convox rack resources options mysql
NAME                        DEFAULT      DESCRIPTION
AllocatedStorage            10           Allocated storage size (GB)
AutoMinorVersionUpgrade     true         Automatically update minor versions
Database                    app          Default database name
DatabaseSnapshotIdentifier               ARN of database snapshot to restore
Encrypted                   false        Encrypt database with KMS
EngineVersion               5.7.16       Version of MySQL
InstanceType                db.t2.micro  Instance class for database nodes
MultiAZ                     false        Multiple availability zone
Password                    (generated)  Server password
Username                    app          Server username
```

### postgres

On AWS, creates an [RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Welcome.html) PostgreSQL instance.

```bash
$ convox rack resources options postgres
NAME                        DEFAULT      DESCRIPTION
AllocatedStorage            10           Allocated storage size (GB)
AutoMinorVersionUpgrade     true         Automatically update minor versions
BackupRetentionPeriod       1            The automatic RDS backup retention period, (default 1 day)
Database                    app          Default database name
DatabaseSnapshotIdentifier               ARN of database snapshot to restore
Encrypted                   false        Encrypt database with KMS
EngineVersion               12           Version of Postgres
Family                      postgres12   Postgres version family
InstanceType                db.t2.micro  Instance class for database nodes
MaxConnections                           ParameterGroup max_connections value, i.e. '{DBInstanceClassMemory/15000000}'
MultiAZ                     false        Multiple availability zone
Password                    (generated)  Server password
Username                    postgres     Server username
```

> While mysql and postgres are available as Rack Resources, there are advantages to using them as [App Resources](/application/resources) instead: configuration is version controlled in `convox.yml`, and App Resources are automatically available during [review workflows](/console/workflows#review-workflows).

### s3

On AWS, creates an [S3 bucket](https://docs.aws.amazon.com/AmazonS3/latest/gsg/GetStartedWithS3.html) for persistent file storage.

```bash
$ convox rack resources options s3
NAME        DEFAULT  DESCRIPTION
Topic                SNS resource name for change notifications
Versioning  false    Enable versioning
```

```bash
$ convox rack resources info s3-2988 | grep URL
URL      s3://AKIAUTMLCEATCHNWWFHP:joZJ61gHiak1Teix1ABdeXS2U8N95jnNQxSGqw3m@test-s3-2988
```

### sns

On AWS, creates an [SNS topic](https://docs.aws.amazon.com/sns/) for pub/sub messaging. Messages published to the topic are pushed to all subscribers automatically.

```bash
$ convox rack resources options sns
NAME   DEFAULT  DESCRIPTION
Queue           SQS resource name to subscribe to this SNS topic
```

```bash
$ convox rack resources info sns-8309 | grep URL
URL      sns://AKIAUTLLCEATCAHANQIG:G2JwSb43AndwTDwQde891Hy8JnXvKV0d47b9PLBQ@arn:aws:sns:us-east-1:316441501734:test-sns-8309
```

### sqs

On AWS, creates an [SQS queue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/welcome.html) for reliable queue processing. Unlike SNS, messages must be polled by receivers and are processed by a single consumer.

```bash
$ convox rack resources options sqs
NAME                    DEFAULT  DESCRIPTION
MessageRetentionPeriod  345600   Number of seconds that a message should be retained on the queue
ReceiveMessageWaitTime  0        Number of seconds that ReceiveMessage should wait for new messages before returning
VisibilityTimeout       30       Number of seconds that a message should wait for confirmation before being returned to the queue
```

```bash
$ convox rack resources info sqs-5495 | grep URL
URL      sqs://AKIAUTLLCEATA27HTUUN:t7mfL6nodQzEcN9LrxypYkNooGUJzsBjze37NDkl@sqs.us-east-1.amazonaws.com/316441501734/test-sqs-5495
```

### syslog

Connects Rack logging to external logging services. See [Syslogs](/deployment/syslogs) for detailed configuration.

```bash
$ convox rack resources options syslog
NAME    DEFAULT                                                   DESCRIPTION
Format  <22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} - - {MESSAGE}  Syslog format string
Url                                                               Syslog URL, e.g. 'tcp+tls://logs1.papertrailapp.com:11235'
```

### webhook

Sends [notifications](/console/notifications) about events within your Apps and Rack to an HTTP endpoint. Convox uses this internally for Slack integration.

```bash
$ convox rack resources options webhook
NAME  DEFAULT  DESCRIPTION
Url            Webhook URL
```

## External Resources

For resource types that Convox does not natively support, you can integrate them by injecting environment variables with connection details.

For example, to connect an App to an externally managed MariaDB instance:

```bash
$ convox env set MARIADB_URL=jdbc:mariadb://mariadb-instance1.123456789012.us-east-1.rds.amazonaws.com:3306/DB?user=myUsername&password=myPassword -a myapplication -r acme/production
```

Add the variable to your service's `environment` configuration in `convox.yml`:

```yaml
services:
  app:
    environment:
      - MARIADB_URL
```

> Environment variables are encrypted and stored securely in KMS. You may also route using internal IP addresses for locked-down environments.

### Networking Requirements

- **Resources within AWS**: Create them in the same VPC as your Rack, in a [peered VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-peering.html), or with a public endpoint. Security Groups for the resource must allow access from the Rack instance Security Group (`{RACK_NAME}-instances`, or check `convox rack params | grep InstanceSecurityGroup` for custom groups).

- **Resources outside AWS**: Must be network-accessible to your Rack instances. For [private Racks](/networking/private-networking), outbound traffic comes from NAT gateway IPs (find these in the AWS VPC Dashboard). Non-private Racks do not have predictable egress IPs, so the resource may need to be available on the public internet.

## See Also

- [App Resources](/application/resources) — resources defined in `convox.yml` and linked to services
- [Accessing Resources](/management/resources) — proxying to App Resources for local management
- [Syslogs](/deployment/syslogs) — detailed syslog integration configuration
- [Notifications](/console/notifications) — webhook notification events
