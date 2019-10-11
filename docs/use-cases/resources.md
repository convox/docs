---
title: "Resources"
---

Resources are those supplementary infrastructure components to your running services, such as Databases, Key-Value stores, Queues, Filestores, Logging endpoints and Webhooks.

Resources can be at the app-level, controlled and configured within your `convox.yml`; at the rack level, created and configured through the CLI; or entirely external to the Convox platform and accessible through imported environment variables.

# App-level Resources

The following resources are available at the [App level](/application/resources):

* `memcached`
* `mysql`
* `postgres`
* `redis`

When resources are defined in your `convox.yml`, they can be automatically linked to your services.  On your local rack, appropriate containers will be spun up alongside your service one/s to provide the functionality required.  On a cloud-based rack, Convox will automatically utilise the appropriate AWS/GCP resources to provide this functionality and link to your apps for you.

Further types of App-level resource are on the Convox roadmap to implement.


# Rack-level Resources

Rack-level resources are not supported on your local rack, only in an AWS/GCP backed Rack.  The following resources are available at the Rack level:

* `memcached`
* `mysql`
* `postgres`
* `redis`
* `s3`
* `sns`
* `sqs`
* `syslog`
* `webhook`

## Creating a Rack Resource

You can create a Rack resource very simply with the default options or you can pass in your optional values at creation time.

```
$ convox rack resources create postgres
Creating resource... OK, postgres-8458

$ convox rack resources create postgres MultiAZ=true
Creating resource... OK, postgres-2871
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

## Updating a Rack Resource

```
$ convox rack resources update postgres-8458 MultiAZ=true
```

## Deleting a Rack Resource

```
$ convox rack resources delete postgres-8458
Deleting postgres-8458... OK
```


# External Resources

For any resource types that Convox does not natively support, it is still very easy to integrate them into your infrastructure and environments.