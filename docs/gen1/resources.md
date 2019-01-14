---
title: Resources
---

The following resources are available at the Rack level:

* `memcached`
* `mysql`
* `postgres`
* `redis`
* `s3`
* `sns`
* `sqs`
* `syslog`
* `webhooks`

## Creating a Resource

```
$ convox rack resources create postgres
Creating resource... OK, postgres-8458
```

#### Options

```
$ convox rack resources create postgres Encrypted=true MultiAZ=true
```

You can see the available options for an instance type:

```
$ convox resources options postgres
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

## Updating a Resource

```
$ convox rack resources update postgres-8458 MultiAZ=true
```

## Deleting a Resource

```
$ convox rack resources delete postgres-8458
Deleting postgres-8458... OK
```
