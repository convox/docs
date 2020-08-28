---
title: "Resources"
---

Resources are those supplementary infrastructure components to your running services, such as Databases, Key-Value stores, Queues, Filestores, Logging endpoints and Webhooks.

Resources can be at the app-level, controlled and configured within your `convox.yml`; at the rack level, created and configured through the CLI; or entirely external to the Convox platform and accessible through imported environment variables.

- App Resources are closely-tied to your services, such as a datastore that your app has an intimate dependency on for succesful startup and function.

- Rack Resources allow for more loosely-coupled dependencies on infrastructure services, or datastores that multiple independent apps may wish to interrogate.

- External Resources allow you to incorporate any external, 3rd party or highly customized components into your infrastructure environment.

Through this, you can have the majority of your standard app infrastructure controlled through code, whilst still maintaining flexibility to add extra components in as and when you need them.

# App Resources

The following resources are available at the [App level](/application/resources):

* `mariadb`
* `memcached`
* `mysql`
* `postgres`
* `redis`

When resources are defined in your `convox.yml`, they can be automatically linked to your services, requiring no code changes on your part, beyond reading the resource's URL from an injected environment variable.  App Resources are also automatically available during [review workflows](/console/workflows#review-workflows).  On your local rack, appropriate containers will be spun up alongside your service to provide the functionality required.  On a cloud-based rack, Convox will automatically utilise the appropriate AWS/GCP resources to provide this functionality and link to your apps for you.

Further types of App Resource are on the Convox roadmap to implement.  If you have specific requests, please contact us through the Support page in your Console to request them.

For further information see the dedicated page on [App Resources](/application/resources).


# Rack Resources

Rack Resources are provider specific, only available from a Rack installed in your cloud account.  They will not be available from a Rack installed on your local development machine.  They are created and configured from the Convox CLI. The following resources are available from an AWS-backed Rack:

* `memcached`
* `mysql`
* `postgres`
* `redis`
* `s3`
* `sns`
* `sqs`
* `syslog`
* `webhook`

### Creating a Rack Resource

You can create a Rack Resource very simply with the default options or you can pass in your optional values at creation time.

```
$ convox rack resources create postgres
Creating resource... OK, postgres-8458

$ convox rack resources create postgres MultiAZ=true
Creating resource... OK, postgres-2871
```

### Listing Current Rack Resources

```
$ convox rack resources
NAME                  TYPE     STATUS
console-175092fe6ab1  webhook  running
syslog-2984           syslog   running
postgres-8458         postgres running
postgres-2871         postgres running
```

#### Options

You can see the available options for a resource type using the CLI:

```
$ convox rack resources options memcached
NAME           DEFAULT         DESCRIPTION
InstanceType   cache.t2.micro  The type of instance to use
NumCacheNodes  1               The number of cache clusters for this replication group
```

### Updating a Rack Resource

```
$ convox rack resources update postgres-8458 MultiAZ=true
```

### Deleting a Rack Resource

```
$ convox rack resources delete postgres-8458
Deleting postgres-8458... OK
```

## Use Cases

### Memcached / Redis

```
$ convox rack resources options memcached
NAME           DEFAULT         DESCRIPTION
InstanceType   cache.t2.micro  The type of instance to use
NumCacheNodes  1               The number of cache clusters for this replication group

$ convox rack resources options redis
NAME                      DEFAULT         DESCRIPTION
AutomaticFailoverEnabled  false           Indicates whether Multi-AZ is enabled. Must be accompanied with NumCacheClusters=2 or higher.
Database                  0               Default database index
Encrypted                 false           Encrypt at rest and in transit
EngineVersion             3.2.6
InstanceType              cache.t2.micro  The type of instance to use
NumCacheClusters          1               The number of cache clusters for this replication group
```

Memcached and Redis are both popular types of in-memory key-value datastores, allowing for extremely fast read/writes of simple data.  In practice they are utilized in a number of different ways, but can be employed as counters, queues, or caches.  On a AWS-backed Rack, Convox will create [Elasticache](https://docs.aws.amazon.com/elasticache/?id=docs_gateway) instances for you that can be accessed from your Apps via their endpoint URL:

```
$ convox rack resources info redis-8193 | grep URL
URL      redis://tes19oh8aty8gzb4.iigghx.ng.0001.use1.cache.amazonaws.com:6379/0
```

<div class="block-callout block-show-callout type-info" markdown="1">
While Memcached and Redis are available as Rack Resources if needed, there are advantages to utilising them as App Resources instead:
- Your configuration of the Resource is stored in the code of your `convox.yml`, and can therefore be version controlled.
- App Resources are automatically available during [review workflows](/console/workflows#review-workflows).
</div>

### MySQL / Postgres

```
$ convox rack resources options mysql
NAME                        DEFAULT      DESCRIPTION
AllocatedStorage            10           Allocated storage size (GB)
Database                    app          Default database name
DatabaseSnapshotIdentifier               ARN of database snapshot to restore
Encrypted                   false        Encrypt database with KMS
EngineVersion               5.7.16       Version of MySQL
InstanceType                db.t2.micro  Instance class for database nodes
MultiAZ                     false        Multiple availability zone
Password                    (generated)  Server password
Username                    app          Server username

$ convox rack resources options postgres
NAME                        DEFAULT      DESCRIPTION
AllocatedStorage            10           Allocated storage size (GB)
BackupRetentionPeriod       1            The automatic RDS backup retention period, (default 1 day)
Database                    app          Default database name
DatabaseSnapshotIdentifier               ARN of database snapshot to restore
Encrypted                   false        Encrypt database with KMS
EngineVersion               9.6          Version of Postgres
Family                      postgres9.6  Postgres version family
InstanceType                db.t2.micro  Instance class for database nodes
MaxConnections                           ParameterGroup max_connections value, i.e. '{DBInstanceClassMemory/15000000}'
MultiAZ                     false        Multiple availability zone
Password                    (generated)  Server password
Username                    postgres     Server username
```

MySQL and Postgres are two of the most popular relational database types in use today, allowing for more complex tables of data to be utlised in your application.  Creating a Rack Resource for either of these on an AWS-backed Rack will initiate the creation of an [RDS instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Welcome.html) of the appropriate flavour for your use via their endpoint URL:

```
$ convox rack resources info postgres-3894 | grep URL
URL      postgres://postgres:d45eea0e492cf51b2ee83328e72284@test-postgres-3894.cm9ftl2mwswu.us-east-1.rds.amazonaws.com:5432/app
```

<div class="block-callout block-show-callout type-info" markdown="1">
While MySQL and Postgres are available as Rack Resources if needed, there are advantages to utilising them as App Resources instead:
- Your configuration of the Resource is stored in the code of your `convox.yml`, and can therefore be version controlled.
- App Resources are automatically available during [review workflows](/console/workflows#review-workflows).
</div>

### S3

```
$ convox rack resources options s3
NAME        DEFAULT  DESCRIPTION
Topic                SNS resource name for change notifications
Versioning  false    Enable versioning
```

On an AWS-backed Rack, creating a S3 Resource will automatically create and configure a [S3 bucket](https://docs.aws.amazon.com/AmazonS3/latest/gsg/GetStartedWithS3.html) for you for persistent file storage.  You can use this bucket to put and retrieve data objects that will persist independently of your App.  

```
$ convox rack resources info s3-2988 | grep URL
URL      s3://AKIAUTMLCEATCHNWWFHP:joZJ61gHiak1Teix1ABdeXS2U8N95jnNQxSGqw3m@test-s3-2988
```

### SNS

```
$ convox rack resources options sns
NAME   DEFAULT  DESCRIPTION
Queue           SQS resource name to subscribe to this SNS topic
```

On an AWS-backed Rack, creating a SNS Resource will automatically create and configure a [SNS topic](https://docs.aws.amazon.com/sns/?id=docs_gateway) for you for easy pub/sub messaging.  You can use this topic to send and receive messages from the cloud in a scalable and reliable way.  With SNS, messages that are published to the topic will be pushed out to any subscribers automatically.  Multiple subscribers may receive the message.
Publishers/producers and subscribers/consumers can connect to your topic easily through the endpoint URL exposed by your Resource:

```
$ convox rack resources info sns-8309 | grep URL
URL      sns://AKIAUTLLCEATCAHANQIG:G2JwSb43AndwTDwQde891Hy8JnXvKV0d47b9PLBQ@arn:aws:sns:us-east-1:316441501734:test-sns-8309
```

### SQS

```
$ convox rack resources options sqs
NAME                    DEFAULT  DESCRIPTION
MessageRetentionPeriod  345600   Number of seconds that a message should be retained on the queue
ReceiveMessageWaitTime  0        Number of seconds that ReceiveMessage should wait for new messages before returning
VisibilityTimeout       30       Number of seconds that a message should wait for confirmation before being returned to the queue
```

On an AWS-backed Rack, creating a SQS Resource will automatically create and configure a [SQS Queue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/welcome.html) for you for reliable, secure and durable queue processing.  You can use this queue for any appropriate queue processing needs you may have, decoupling communication between your services for example.  With SQS, messages that are sent to the queue are not pushed out to any receivers, they must poll the queue to receive the messages, and any one receiver can retrieve, process and delete a message.
Your services can connect to your queue easily through the endpoint URL exposed by your Resource:

```
$ convox rack resources info sqs-5495 | grep URL
URL      sqs://AKIAUTLLCEATA27HTUUN:t7mfL6nodQzEcN9LrxypYkNooGUJzsBjze37NDkl@sqs.us-east-1.amazonaws.com/316441501734/test-sqs-5495
```

### Syslog

```
$ convox rack resources options syslog
NAME    DEFAULT                                                   DESCRIPTION
Format  <22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} - - {MESSAGE}  Syslog format string
Url                                                               Syslog URL, e.g. 'tcp+tls://logs1.papertrailapp.com:11235'
```

[Syslog](/deployment/syslogs) resources are widely used to connect your application logging to external 3rd party logging services.
We have a dedicated documentation on this [here](/deployment/syslogs).

### Webhook

```
$ convox rack resources options webhook
NAME  DEFAULT  DESCRIPTION
Url            Webhook URL
```

Internally, Convox uses the `webhook` resource to enable Slack integration.  You can do similar to receive [notifications](/console/notifications) about important events happening within your apps and Rack.  

# External Resources

For any resource types that Convox does not natively support, it is still very easy to integrate them into your infrastructure and environments.  
Best practice is to inject an environment variable for each external resource that you wish to integrate.  

For example if you needed a MariaDB instance, you would go create that in RDS, ensure that the VPC and Security Group allows access to it from your Rack, get the appropriate endpoint, then in the application that needs access, set the env var for your connection string:

```
$ convox env set MARIADB_URL=jdbc:mariadb://mariadb-instance1.123456789012.us-east-1.rds.amazonaws.com:3306/DB?user=myUsername&password=myPassword -a myapplication -r acme/production
```

You will also need to add the `MARIADB_URL` env var to the `environment` configuration of your service in your `convox.yml`:

```
services:
  app:
    ...
    environment:
      - MARIADB_URL
      - ...
    ...  
```

<div class="block-callout block-show-callout type-info" markdown="1">
As environment variables are encrypted and stored securely in KMS, your credentials are secure within your environment.  You may also route using the internal IP address if you wish, for locked down environments.
</div>

- For External Resources that are still within your AWS infrastructure, you should create them to be accessible to your Rack.  Either within the same VPC as your Rack, within a VPC that your Rack VPC [peers with](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-peering.html), or in a way that is network accessible (with a public endpoint).  This will be dependent on your security needs, and the type of Resource you are creating.  Any Security Group [configuration](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html#SecurityGroupRules) for your Resource should also allow access to your Resource's port from the Rack Security Group for your instances.  If you did not specify a custom Security Group for your Rack, you can discover it through your AWS Console, named as `{RACK_NAME}-instances`.  If you have set a custom Security Group for your Rack instances, you can remind yourself of it from your rack params: `$ convox rack params | grep InstanceSecurityGroup`.

- External Resource that are outside of your AWS infrastructure will need to be network-accessible to your Rack instances. For [private Racks](/management/private-networking), traffic will come from the IPs of your NAT gateways (you can find out the addresses of your NAT Gateways through your AWS Console, in the VPC Dashboard for your Rack's region). If your Rack is not set to private the IPs will not be easily predictable and you'll likely need to make the resource available to the public internet.
