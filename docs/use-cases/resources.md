---
title: "Resources"
---

Resources are those supplementary infrastructure components to your running services, such as Databases, Key-Value stores, Queues, Filestores, Logging endpoints and Webhooks.

Resources can be at the app-level, controlled and configured within your `convox.yml`; at the rack level, created and configured through the CLI; or entirely external to the Convox platform and accessible through imported environment variables.

Through this, you can have the majority of your standard app infrastructure controlled through code, whilst still maintaining flexibility to add extra components in as and when you need them.

# App-level Resources

The following resources are available at the [App level](/application/resources):

* `memcached`
* `mysql`
* `postgres`
* `redis`

When resources are defined in your `convox.yml`, they can be automatically linked to your services, requiring no code changes on your part, beyond reading the resource's URL from an injected environment variable.  On your local rack, appropriate containers will be spun up alongside your service to provide the functionality required.  On a cloud-based rack, Convox will automatically utilise the appropriate AWS/GCP resources to provide this functionality and link to your apps for you.

Further types of App-level resource are on the Convox roadmap to implement.  If you have specific demands, please contact us through the Support page in your console to request them.

For further information see the dedicated page on [App Resources](/application/resources).


# Rack-level Resources

Rack-level resources are not supported on your local rack, only in a cloud backed Rack.  They are created and configured from the Convox CLI. The following resources are available at the Rack level:

* `s3`
* `sns`
* `sqs`
* `syslog`
* `webhook`
* `memcached` *
* `mysql` *
* `postgres` *
* `redis` *

<div class="block-callout block-show-callout type-info" markdown="1">
 * Unless required at the Rack level, we would recommend using these types at the App level.
</div>


### Creating a Rack Resource

You can create a Rack resource very simply with the default options or you can pass in your optional values at creation time.

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
$ convox rack resources options postgres
NAME                        DEFAULT      DESCRIPTION
AllocatedStorage            10           Allocated storage size (GB)
BackupRetentionPeriod       1            The automatic RDS backup retention period, (default 1 day)
Database                    app          Default database name
Encrypted                   false        Encrypt database with KMS
EngineVersion               9.6          Version of Postgres
Family                      postgres9.6  Postgres version family
InstanceType                db.t2.micro  Instance class for database nodes
MultiAZ                     false        Multiple availability zone
Password                    (generated)  Server password
Username                    postgres     Server username
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

[Syslog](/deployment/syslogs) resources are widely used to connect your application logging to external 3rd party services.
Internally, Convox uses the `webhook` resource to enable Slack integration.  You can do similar to receive [notifications](/console/notifications) about important events happening within your apps and Rack.  




# External Resources

For any resource types that Convox does not natively support, it is still very easy to integrate them into your infrastructure and environments.  
Best practice is to inject an environment variable for each external resource that you wish to integrate.  

For example if you needed a MariaDB instance, you would go create that in RDS, ensure that the VPC and Security Group allows access to it from your Rack, get the appropriate endpoint, then in the application that needs access, set the env var for your connection string:

```
$ convox env set MARIADB_URL=jdbc:mariadb://mariadb-instance1.123456789012.us-east-1.rds.amazonaws.com:3306/DB?user=myUsername&password=myPassword -a myapplication -r acme/production
```

<div class="block-callout block-show-callout type-info" markdown="1">
As environment variables are encrypted and stored securely in KMS, your credentials are secure within your environment.  You may also route using the internal IP address if you wish, for locked down environments.
</div>

Through this, you can securely control and configure access to different resources based on app and environment.  It is even possible to integrate resources and infrastructure elements that are external to your cloud environment, as long as they are accessible.