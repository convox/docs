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
NAME                     TYPE      STATUS
console-v1-175092fe6ab1  webhook   running
syslog-2984              syslog    running
postgres-8458            postgres  running
postgres-2871            postgres  running
```

### Viewing Resource Info

```bash
$ convox rack resources info postgres-8458
```

### Viewing Available Options

```bash
$ convox rack resources options memcached
NAME           DEFAULT         DESCRIPTION
InstanceType   cache.t3.micro  The type of instance to use
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

Only `syslog` Rack Resources can be linked to an App. Linking sends the App's logs to the syslog destination:

```bash
$ convox rack resources link syslog-2984 --app myapp
Linking to myapp... OK
```

To remove the link:

```bash
$ convox rack resources unlink syslog-2984 --app myapp
Unlinking from myapp... OK
```

To use a Rack database or cache from an App, get its URL with `convox rack resources url` and set it as an App environment variable, as described in [External Resources](#external-resources).

### Deleting a Resource

```bash
$ convox rack resources delete postgres-8458
Deleting resource... OK
```

## Available Rack Resource Types

### memcached

On AWS, creates an [ElastiCache](https://docs.aws.amazon.com/elasticache/) Memcached cluster.

```bash
$ convox rack resources options memcached
NAME           DEFAULT         DESCRIPTION
InstanceType   cache.t3.micro  The type of instance to use
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
Engine                    redis           The cache engine to use (redis or valkey)
EngineVersion             7.0             The version of the cache engine
InstanceType              cache.t3.micro  The type of instance to use
NumCacheClusters          1               The number of cache clusters for this replication group
```

### valkey

On AWS, creates an [ElastiCache](https://docs.aws.amazon.com/elasticache/) Valkey cluster.

```bash
$ convox rack resources options valkey
NAME                      DEFAULT         DESCRIPTION
AutomaticFailoverEnabled  false           Indicates whether Multi-AZ is enabled. Must be accompanied with NumCacheClusters=2 or higher.
Database                  0               Default database index
Encrypted                 false           Encrypt at rest and in transit
EngineVersion             8.1             The version of the cache engine
InstanceType              cache.t3.micro  The type of instance to use
NumCacheClusters          1               The number of cache clusters for this replication group
```

> While memcached, redis and valkey are available as Rack Resources, there are advantages to using them as [App Resources](/application/resources) instead: configuration is version controlled in `convox.yml`, and App Resources are automatically available during [review workflows](/console/workflows#review-workflows).

### mysql

On AWS, creates an [RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Welcome.html) MySQL instance.

```bash
$ convox rack resources options mysql
NAME                        DEFAULT      DESCRIPTION
AllocatedStorage            20           Allocated storage size (GB)
AllowMajorVersionUpgrade    false
AutoMinorVersionUpgrade     true
Database                    app          Default database name
DatabaseSnapshotIdentifier               ARN of database snapshot to restore
Encrypted                   false        Encrypt database with KMS
EngineVersion               8.4          Version of MySQL
InstanceType                db.t3.micro  Instance class for database nodes
MultiAZ                     false        Multiple availability zone
Password                    (generated)  Server password
Username                    app          Server username
```

### postgres

On AWS, creates an [RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Welcome.html) PostgreSQL instance.

```bash
$ convox rack resources options postgres
NAME                        DEFAULT      DESCRIPTION
AllocatedStorage            20           Allocated storage size (GB)
AllowMajorVersionUpgrade    false
AutoMinorVersionUpgrade     true
BackupRetentionPeriod       1            The automatic RDS backup retention period, (default 1 day)
Database                    app          Default database name
DatabaseSnapshotIdentifier               ARN of database snapshot to restore
Encrypted                   false        Encrypt database with KMS
EngineVersion               17           Version of Postgres
Family                      postgres17   Postgres version family
InstanceType                db.t3.micro  Instance class for database nodes
MaxConnections                           ParameterGroup max_connections value, i.e. '{DBInstanceClassMemory/15000000}'
MultiAZ                     false        Multiple availability zone
Password                    (generated)  Server password
Username                    postgres     Server username
```

> While mysql and postgres are available as Rack Resources, there are advantages to using them as [App Resources](/application/resources) instead: configuration is version controlled in `convox.yml`, and App Resources are automatically available during [review workflows](/console/workflows#review-workflows).

### Database Versions

`convox rack resources options` shows the default `EngineVersion` for mysql and postgres on your Rack. Existing mysql and postgres Resources keep their `EngineVersion` and `Family` through Rack updates and `convox rack resources update`.

For postgres, an `EngineVersion` without `Family` sets the matching `Family` on create. Requires rack version 20261005214736 or newer; on earlier Racks, pass `Family` with `EngineVersion`.

```bash
$ convox rack resources create postgres EngineVersion=16 --name reports-db --wait
Creating resource... OK, reports-db
```

To restore from `DatabaseSnapshotIdentifier`, also pass the snapshot's `EngineVersion`. Without it, the restore requests the default version, and a postgres restore from a snapshot on a different major version fails.

To upgrade an existing database to a new major version, pass `AllowMajorVersionUpgrade=true` with the new `EngineVersion`, and for postgres the matching `Family`. Requires rack version 20260212195551 or newer.

```bash
$ convox rack resources update reports-db AllowMajorVersionUpgrade=true EngineVersion=17 Family=postgres17 --wait
Updating resource... OK
```

New PostgreSQL 15 and later and MySQL 8.4 databases need clients that connect with TLS. See [TLS Connections](/application/resources#tls-connections).

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
URL      s3://AKIAIOSFODNN7EXAMPLE:EXAMPLEKEYEXAMPLEKEYEXAMPLEKEYEXAMPLEKEY@test-s3-2988
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
URL      sns://AKIAIOSFODNN7EXAMPLE:EXAMPLEKEYEXAMPLEKEYEXAMPLEKEYEXAMPLEKEY@arn:aws:sns:us-east-1:123456789012:test-sns-8309
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
URL      sqs://AKIAIOSFODNN7EXAMPLE:EXAMPLEKEYEXAMPLEKEYEXAMPLEKEYEXAMPLEKEY@sqs.us-east-1.amazonaws.com/123456789012/test-sqs-5495
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

Sends [notifications](/console/notifications) about events within your Apps and Rack to an HTTP endpoint. Console creates one on each Rack it manages, named `console-v1-<id>`, to feed the Console Events tab and notification integrations such as Slack.

```bash
$ convox rack resources options webhook
NAME  DEFAULT  DESCRIPTION
Url            Webhook URL
```

`Url` is required and takes an `http://` or `https://` URL. Use `https://` where the receiver supports it, since `http://` deliveries are unencrypted. Delivery to `http://` URLs requires rack version 20261005214736 or newer.

```bash
$ convox rack resources create webhook Url=https://hooks.example.com/convox --name my-hook --wait
Creating resource... OK, my-hook
```

Each event is a `POST` to the URL with the event JSON as the body:

```json
{"action":"release:create","data":{"app":"myapp","id":"RABCDEFGHIJ","rack":"production"},"status":"success","timestamp":"2026-10-02T12:00:00.123456789Z"}
```

A `release:promote` event has `"status":"start"` when the promotion begins. An event for a failed operation has `"status":"error"` and the error text in `data.message`.

Redirects are not followed: when a receiver redirects `http://` to `https://`, the event never reaches the `https://` URL, so give the webhook the `https://` URL. Events are sent from AWS Lambda, not from your Rack's instances or NAT gateway, so the receiver must be reachable from the internet and cannot allowlist your Rack's egress IPs. An address reachable only inside the Rack's VPC, such as an [internal Service](/networking/internal-services) or the router of an [InternalOnly](/reference/rack-parameters/InternalOnly) Rack, receives nothing.

#### Changing a Webhook URL

Requires rack version 20261005214736 or newer. On older racks the webhook keeps its URL; delete it and create it again with the new URL.

```bash
$ convox rack resources update my-hook Url=https://hooks.example.com/new --wait
Updating resource... OK

$ convox rack resources url my-hook
https://hooks.example.com/new
```

The option name is case-sensitive on update: a lowercase `url=` is ignored and the webhook keeps its URL.

| Command | Result |
|:--|:--|
| `convox rack resources update <name> Url=<http or https URL>` | Events go to the new URL |
| `convox rack resources update <name>` | The webhook keeps its URL |
| `convox rack resources update <name> Url=` | `ERROR: must specify a URL`, webhook unchanged |
| `convox rack resources update <name> Url=ftp://...` | `ERROR: invalid URL scheme: ftp. Allowed schemes are: http, https`, webhook unchanged |
| `convox rack resources update console-v1-<id> Url=...` | `ERROR: webhook console-v1-<id> is managed by Console and its Url cannot be changed`, webhook unchanged |

Any webhook whose name starts with `console-v1-` is treated as Console's and refuses a `Url` update, so give your own webhooks a different name.

#### Applying Webhook Fixes

`convox rack update` does not update existing webhooks. After a Rack update completes (`convox rack update --wait`), run `convox rack resources update <name>` with no options to apply webhook fixes, such as `http://` delivery, to an existing webhook. When the webhook is already current, the command prints `ERROR: ValidationError: No updates are to be performed.` and nothing changes.

Console's `console-v1-<id>` webhook takes the same update. If it is deleted, Console creates it again within about 15 minutes, and the Rack's events do not reach Console until then.

## External Resources

For resource types that Convox does not natively support, you can integrate them by injecting environment variables with connection details.

For example, to connect an App to an externally managed MariaDB instance:

```bash
$ convox env set MARIADB_URL=jdbc:mariadb://mariadb-instance1.123456789012.us-east-1.rds.amazonaws.com:3306/DB?user=myUsername&password=myPassword -a myapplication -r acme/production
```

Add the variable to your Service's `environment` configuration in `convox.yml`:

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

- [App Resources](/application/resources): resources defined in `convox.yml` and linked to Services
- [Accessing Resources](/management/resources): proxying to App Resources for local management
- [Syslogs](/deployment/syslogs): detailed syslog integration configuration
- [Notifications](/console/notifications): webhook notification events
