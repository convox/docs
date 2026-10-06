---
title: "Resources"
description: "Define and configure network-attached dependencies such as databases, caches, and shared filesystems for Convox applications."
---

# Resources

A resource is a network-attached dependency of your application.

Resources are declared in `convox.yml`, the gen2 manifest format that is the current, actively-developed generation; the legacy gen1 format is End-of-Life and is not covered here.

## Definition

Here we define a resource called `mydb` that is a Postgres database with 100GB of storage, and we then link it to our `web` service, which will be given an [env var](#accessing-resources) to connect to the database:

```yaml
resources:
  mydb:
    type: postgres
    options:
      storage: 100
    tags:
      Name: example-database
      Environment: production
services:
  web:
    resources:
      - mydb
```

The resource name only affects the [environment variable name](#accessing-resources) that is passed to your Services.  You are free to name it what you wish with no regard to the type of resource.

You can define multiple resources within one `convox.yml`:

```yaml
resources:
  maindb:
    type: mysql
    tags:
      Name: main-database
      Environment: staging
  gisdb:
    type: postgres
    options:
      version: 17
    tags:
      Name: gis-database
      Project: mapping
  queue:
    type: redis
  cache:
    type: redis
services:
  web:
    resources:
      - maindb
      - gisdb
      - queue
      - cache
```

## Accessing Resources

You can access defined resources from Services with environment variables.
In the above example, the `mydb` resource provides a `MYDB_URL` variable that is accessible from the `web` service.

The environment variable name is the resource name converted to all-caps, with a `_URL` suffix.

This would contain the entire connection string you would need, ie:

```text
MYDB_URL=postgres://username:password@host.com:5432/databaseName
```

### Additional credentials
*Available in rack version **20221013170042 or later***

You can also use the additional credentials to connect to the resource, the credentials will be provided in the environment variables with the resource name prefix and the following suffix: `_USER`, `_PASS`, `_HOST`, `_PORT`, `_NAME`.

Using the example above, the resource name `mydb` will provide the following environment variable:

```text
MYDB_USER=username
MYDB_PASS=password
MYDB_HOST=host.com
MYDB_PORT=5432
MYDB_NAME=databaseName
```

### Dynamic Configuration

If you'd like to make your resource configuration variable (e.g. to produce different results between environments like staging vs production) you can use environment variable interpolation:

```yaml
resources:
  mydb:
    type: postgres
    options:
      storage: ${POSTGRES_STORAGE_SIZE}
    tags:
      Environment: ${ENVIRONMENT}
```

```bash
$ convox env set POSTGRES_STORAGE_SIZE=50 ENVIRONMENT=staging --rack=acme/staging
$ convox env set POSTGRES_STORAGE_SIZE=200 ENVIRONMENT=production --rack=acme/production
```

## Read Replica Support
*Available in rack version **20240513194424 or later***

Read replicas allow you to scale out read-heavy workloads by creating read-only copies of an existing database. This helps distribute database traffic, improve performance, and enhance application scalability without affecting the primary database.

### Defining a Read Replica

To create a read replica using Convox resources, reference an existing database as the `readSourceDB` in your `convox.yml` file. The source database must exist before deploying a read replica.

#### How to Set `readSourceDB`

The value for `readSourceDB` follows this format:

```text
#convox.resources.<source-database-name>
```

The `#convox.resources.` prefix is constant, and the `<source-database-name>` must match the name of the database resource you want to replicate. For example, if your primary database resource is named `maindb`, then the correct `readSourceDB` reference would be:

```yaml
readSourceDB: "#convox.resources.maindb"
```

### Configurable Read Replica Options

While read replicas inherit most configurations from their primary database, certain parameters can be adjusted. The following constraints apply:

| Area | Constraint |
|------|------------|
| Encryption | If the primary database is unencrypted, the read replica must also be unencrypted. |
| Encryption | If the primary database is encrypted, the read replica can be either encrypted or unencrypted. |
| Encryption | You cannot create an encrypted read replica from an unencrypted primary database. |
| Storage capacity | A read replica can have more storage than the primary database, but not less. |
| Storage capacity | Storage capacity cannot be reduced once allocated. |
| Multi-AZ failover (`durable`) | Multi-AZ failover is optional for read replicas. |
| Multi-AZ failover (`durable`) | A read replica is not required to have `durable: true`, even if the primary database does. |

### Important: Ensure the Database Version Matches

When creating a read replica, it is mandatory to explicitly set the same database version as the primary database. A mismatch can lead to deployment failures or unexpected behavior.

Always ensure the `version` field in the read replica matches the source database exactly.

### Example Configuration

```yaml
resources:
  primary-db:
    type: mysql
    options:
      version: "8.4"  # Ensure the version matches the primary database
      class: db.t3.medium
      storage: 50
      encrypted: true
      durable: true
    tags:
      Name: primary-database
      Environment: production

  read-replica:
    type: mysql
    options:
      readSourceDB: "#convox.resources.primary-db"
      version: "8.4"  # The read replica must use the same version as the primary database
      class: db.t3.medium
      storage: 50
      encrypted: true
    tags:
      Name: read-replica-database
      Environment: production
      Purpose: reporting
```

### Converting a Read Replica to an Active Database

If you want to convert a read replica into an independent database, remove the `readSourceDB` option from `convox.yml` and redeploy the application. This process does not affect the original primary database, and the read replica retains the same name.

## EFS Resource
*Available in rack version **20221214201933 or later***


The EFS resource lets you share volumes between Services in different AZs.

EFS resources have additional configurations. The definition is different from the database resources. See Available Resources > [EFS](#efs).
After declaring the resource and the options, the link between the resource and Service you want to expose is required. Example:

```yaml
resources:
  sharedvolume:
    type: efs
    options:
      path: "/root-directory"
    tags:
      Name: shared-filesystem
      BackupSchedule: daily
services:
  web:
    resources:
      - sharedvolume
```

Once the resource is linked to the Service, set the volume and the path you want to expose to the application, e.g. `{efs-resource-name}:/path/to/mount/on/application`.
The path you want to link the application is not bound to the same you declared in the definition, it's the directory that will be mount in the application container. A full example of EFS resource using the `resources` and `volumes`:

```yaml
resources:
  sharedvolume:
    type: efs
    options:
      path: "/root-directory"
    tags:
      Name: shared-filesystem
      Environment: production
      BackupSchedule: daily
services:
  web:
    volumes:
      - sharedvolume:/app/dir
    resources:
      - sharedvolume
```

## Available Resources

### memcached

| Option    | Default          | Description       |
|-----------|------------------|-------------------|
| `class`   | `cache.t3.micro` | Instance class    |
| `nodes`   | `1`              | Number of nodes   |
| `version` | `1.4`            | Memcached version |

### mariadb

| Option      | Default          | Description                             |
|-------------|------------------|-----------------------------------------|
| `class`     | `db.t3.micro`    | Instance class                          |
| `encrypted` |                  | Encrypt data at rest                    |
| `deletionProtection` | `false` | Enable deletion protection              |
| `durable`   | `false`          | Multi-AZ automatic failover             |
| `iops`      |                  | Provisioned IOPS for database disks     |
| `parameterGroupName` |        | Custom DB parameter group name. When blank, uses the default parameter group for the engine version |
| `snapshot`  |                  | ARN of a DB snapshot to restore from    |
| `storage`   | `20`             | GB of storage to provision              |
| `version`   | `11.4`           | MariaDB version. See [Default Versions and Classes](#default-versions-and-classes) |
| `preferredBackupWindow` |  | The daily time range during which automated backups are created if automated backups are enabled, using the `backupRetentionPeriod` option. Must be in the format hh24:mi-hh24:mi. Must be in Universal Coordinated Time (UTC). Must not conflict with the preferred maintenance window. Must be at least 30 minutes.              |
| `backupRetentionPeriod`   | `1`           | The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups. |




### mysql

| Option      | Default          | Description                             |
|-------------|------------------|-----------------------------------------|
| `class`     | `db.t3.micro`    | Instance class                          |
| `encrypted` |                  | Encrypt data at rest                    |
| `deletionProtection` | `false` | Enable deletion protection              |
| `durable`   | `false`          | Multi-AZ automatic failover             |
| `iops`      |                  | Provisioned IOPS for database disks     |
| `parameterGroupName` |        | Custom DB parameter group name. When blank, uses the default parameter group for the engine version |
| `snapshot`  |                  | ARN of a DB snapshot to restore from    |
| `storage`   | `20`             | GB of storage to provision              |
| `version`   | `8.4`            | MySQL version. See [Default Versions and Classes](#default-versions-and-classes) |
| `preferredBackupWindow` |  | The daily time range during which automated backups are created if automated backups are enabled, using the `backupRetentionPeriod` option. Must be in the format hh24:mi-hh24:mi. Must be in Universal Coordinated Time (UTC). Must not conflict with the preferred maintenance window. Must be at least 30 minutes.              |
| `backupRetentionPeriod`   | `1`           | The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups. |

### postgres

| Option      | Default          | Description                             |
|-------------|------------------|-----------------------------------------|
| `class`     | `db.t3.micro`    | Instance class                          |
| `deletionProtection` | `false` | Enable deletion protection              |
| `durable`   | `false`          | Multi-AZ automatic failover             |
| `encrypted` |                  | Encrypt data at rest                    |
| `iops`      |                  | Provisioned IOPS for database disks     |
| `parameterGroupName` |        | Custom DB parameter group name. When blank, uses the default parameter group for the engine version |
| `snapshot`  |                  | ARN of a DB snapshot to restore from    |
| `storage`   | `20`             | GB of storage to provision              |
| `version`   | `17`             | PostgreSQL version. See [Default Versions and Classes](#default-versions-and-classes) |
| `preferredBackupWindow` |  | The daily time range during which automated backups are created if automated backups are enabled, using the `backupRetentionPeriod` option. Must be in the format hh24:mi-hh24:mi. Must be in Universal Coordinated Time (UTC). Must not conflict with the preferred maintenance window. Must be at least 30 minutes.              |
| `backupRetentionPeriod`   | `1`           | The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups. |

### redis

| Option      | Default          | Description                 |
|-------------|------------------|-----------------------------|
| `class`     | `cache.t3.micro` | Instance class              |
| `durable`   | `false`          | Multi-AZ automatic failover. When it is set to `true`, the option `nodes` has to be greater or equal to `2`, otherwise it will fail |
| `encrypted` | `false`          | Encrypt data at rest and in transit. When `true`, the URL uses `rediss://` and includes an auth token |
| `engine`    | `redis`          | Cache engine to use (`redis` or `valkey`) |
| `nodes`     | `1`              | Number of nodes             |
| `version`   | `7.0`            | Redis version               |

### valkey

| Option      | Default          | Description                 |
|-------------|------------------|-----------------------------|
| `class`     | `cache.t3.micro` | Instance class              |
| `durable`   | `false`          | Multi-AZ automatic failover. When it is set to `true`, the option `nodes` has to be greater or equal to `2`, otherwise it will fail |
| `encrypted` | `false`          | Encrypt data at rest and in transit. When `true`, the URL uses `rediss://` and includes an auth token |
| `nodes`     | `1`              | Number of nodes             |
| `version`   | `8.1`            | Valkey version              |

### efs

Use to share volumes between the tasks in different AZs and instances.

| Option        | Default  | Description                                                            |
|---------------|----------|------------------------------------------------------------------------|
| `encrypted`   | `false`  | Encrypt data at rest                                                   |
| `owner-gid`   | `1000`   | POSIX group ID to apply to the `path` directory                        |
| `owner-uid`   | `1000`   | POSIX user ID to apply to the `path` directory                         |
| `path`        | `/`      | The path on the file system used as the root directory by the Services |
| `permissions` | `0777`   | POSIX permissions to apply to the `path` directory                     |

## Default Versions and Classes

| Resource | Option | Default | Before rack version 20261005214736 |
|----------|--------|---------|------------------------------------|
| `postgres` | `version` | `17` | `12` |
| `mysql` | `version` | `8.4` | `5.7` |
| `mariadb` | `version` | `11.4` | `10.4` |
| `memcached`, `redis`, `valkey` | `class` | `cache.t3.micro` | `cache.t2.micro` |

A new resource without `version` or `class` gets the default of the Rack version it is created on. An existing resource keeps its engine version, parameter group and class: removing `version` or `class` from `convox.yml` does not change it, and neither does a Rack update that changes a default.

Changing `version` on an existing database upgrades it in place, including to a new major version. RDS supports only certain upgrade paths; MySQL 5.7, for example, upgrades to 8.0 and not directly to 8.4. RDS bills Extended Support for a database on a major version past its end of standard support until you upgrade it.

We recommend setting `version` on every database and `class` on every cache. Without them, Apps created later from the same `convox.yml`, such as [review workflow](/console/workflows#review-workflows) Apps or a rebuilt staging App, can start on a different version than your existing Apps.

When restoring from `snapshot`, set `version` to the snapshot's engine version. Without it, the restore requests the default version, and the deploy fails when the snapshot is on a different major version.

### TLS Connections

New databases on these versions with the default parameter group need clients that connect with TLS:

| Engine | Default on a new database | Client requirement |
|--------|---------------------------|--------------------|
| PostgreSQL 15 and later | `rds.force_ssl` is `1` | Connect with TLS. A connection without TLS is refused with `no pg_hba.conf entry for host ..., no encryption` |
| MySQL 8.4 | users are created with `caching_sha2_password` | The client must support `caching_sha2_password`. A login without TLS can be refused, so connect with TLS |

The `_URL` environment variable has no TLS parameters. If your database client does not use TLS by default, enable it in the client's connection settings.

## AutoMinorVersionUpgrade

In case you specify the minor version on your resource definition you have to turn off the [AutoMinorVersionUpgrade](/reference/app-parameters/AutoMinorVersionUpgrade) on your app parameter. It's enabled by default and it will update the DB instance during the maintenance window.

## See Also

- [Services](/application/services)
- [Environment](/application/environment)
- [Volumes](/application/volumes)
- [Management: Resources](/management/resources)
