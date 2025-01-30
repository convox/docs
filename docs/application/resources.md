---
title: Resources
---

A resource is a network-attached dependency of your application.

## Definition

Here we define a resource called `mydb` that is a Postgres database with 100GB of storage, and we then link it to our `web` service, which will be given an [env var](#accessing-resources-through-environment-variables) to connect to the database:

```yaml
resources:
  mydb:
    type: postgres
    options:
      storage: 100
services:
  web:
    resources:
      - mydb
```

The resource name only affects the [environment variable name](#accessing-resources-through-environment-variables) that is passed to your services.  You are free to name it what you wish with no regard to the type of resource.

You can easily define multiple resources within one `convox.yml`:

```yaml
resources:
  maindb:
    type: mysql
  gisdb:
    type: postgres
    options:
      version: 12
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



## Read Replica Support

Read replicas allow you to scale out read-heavy workloads by creating **read-only copies** of an existing database. This helps distribute database traffic, improve performance, and enhance application scalability without affecting the primary database.

### Defining a Read Replica

To create a read replica using Convox resources, reference an existing database as the `readSourceDB` in your `convox.yml` file. The source database **must exist** before deploying a read replica.

Example:
```yaml
resources:
  mysql-main:
    type: mysql
    options:
      class: db.t3.micro
      encrypted: false
  mysql-replica:
    type: mysql
    options:
      readSourceDB: "#convox.resources.mysql-main"
      class: db.t3.micro
      encrypted: true
```

### Converting a Read Replica to an Active Database

If you want to convert a read replica into an independent database, remove the `readSourceDB` option from `convox.yml` and redeploy the application. This process **does not affect** the original primary database, and the read replica retains the same name.

### EFS Resource (version 20221214201933+)

The EFS resource lets you share volumes between services in different AZs.

EFS resources have additional configurations. The definition is different from the database resources. See Available Resources > [EFS](#efs).
After declaring the resource and the options, the link between the resource and service you want to expose is required. Example:

```
resources:
  sharedvolume:
    type: efs
    options:
      path: "/root-directory"
services:
  web:
    resources:
      - sharedvolume
```

Once the resource is linked to the service, set the volume and the path you want to expose to the application, e.g. `{efs-resource-name}:/path/to/mount/on/application`.
The path you want to link the application is not bound to the same you declared in the definition, it's the directory that will be mount in the application container. A full example of EFS resource using the `resources` and `volumes`:

```
resources:
  sharedvolume:
    type: efs
    options:
      path: "/root-directory"
services:
  web:
    volumes:
      - sharedvolume:/app/dir
    resources:
      - sharedvolume
```

#### Dynamic Configuration

If you'd like to make your resource configuration variable (e.g. to produce different results between environments like staging vs production) you can use environment variable interpolation:

```yaml
resources:
  mydb:
    type: postgres
    options:
      storage: ${POSTGRES_STORAGE_SIZE}
```

```
$ convox env set POSTGRES_STORAGE_SIZE=50 --rack=acme/staging
$ convox env set POSTGRES_STORAGE_SIZE=200 --rack=acme/production
```

## Accessing Resources

You can access defined resources from services with environment variables.
In the above example, the `mydb` resource provides a `MYDB_URL` variable that is accessible from the `web` service.

The environment variable name is the resource name converted to all-caps, with a `_URL` suffix.

This would contain the entire connection string you would need, ie:

```
MYDB_URL=postgres://username:password@host.com:5432/databaseName
```

#### Additional credentials (20221013170042 or newer)
You can also use the additional credentials to connect to the resource, the credentials will be provided in the environment variables with the resource name prefix and the following suffix: `_USER`, , `_PASS`, `_HOST`, `_PORT`, `_NAME`.

Using the example above, the resource name `mydb` will provide the following environment variable:

```
MYDB_USER=username
MYDB_PASS=password
MYDB_HOST=host.com
MYDB_PORT=5432
MYDB_NAME=databaseName
```

## Available Resources

#### memcached

| Option    | Default          | Description       |
|-----------|------------------|-------------------|
| `class`   | `cache.t2.micro` | Instance class    |
| `nodes`   | `1`              | Number of nodes   |
| `version` | `1.4.34`         | Memcached version |

#### mariadb

| Option      | Default          | Description                             |
|-------------|------------------|-----------------------------------------|
| `class`     | `db.t2.micro`    | Instance class                          |
| `encrypted` | `false`          | Encrypt data at rest                    |
| `deletionProtection` | `false` | Enable deletion protection              |
| `durable`   | `false`          | Multi-AZ automatic failover             |
| `iops`      |                  | Provisioned IOPS for database disks     |
| `storage`   | `20`             | GB of storage to provision              |
| `version`   | `10.4`           | MariaDB version                         |
| `preferredBackupWindow` |  | The daily time range during which automated backups are created if automated backups are enabled, using the `backupRetentionPeriod`` option. Must be in the format hh24:mi-hh24:mi.Must be in Universal Coordinated Time (UTC). Must not conflict with the preferred maintenance window. Must be at least 30 minutes.              |
| `backupRetentionPeriod`   | `1`           | The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups. |




#### mysql

| Option      | Default          | Description                             |
|-------------|------------------|-----------------------------------------|
| `class`     | `db.t2.micro`    | Instance class                          |
| `encrypted` | `false`          | Encrypt data at rest                    |
| `deletionProtection` | `false` | Enable deletion protection              |
| `durable`   | `false`          | Multi-AZ automatic failover             |
| `iops`      |                  | Provisioned IOPS for database disks     |
| `storage`   | `20`             | GB of storage to provision              |
| `version`   | `5.7.22`         | MySQL version                           |
| `preferredBackupWindow` |  | The daily time range during which automated backups are created if automated backups are enabled, using the `backupRetentionPeriod`` option. Must be in the format hh24:mi-hh24:mi.Must be in Universal Coordinated Time (UTC). Must not conflict with the preferred maintenance window. Must be at least 30 minutes.              |
| `backupRetentionPeriod`   | `1`           | The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups. |

#### postgres

| Option      | Default          | Description                             |
|-------------|------------------|-----------------------------------------|
| `class`     | `db.t2.micro`    | Instance class                          |
| `deletionProtection` | `false` | Enable deletion protection              |
| `durable`   | `false`          | Multi-AZ automatic failover             |
| `encrypted` | `false`          | Encrypt data at rest                    |
| `iops`      |                  | Provisioned IOPS for database disks     |
| `storage`   | `20`             | GB of storage to provision              |
| `version`   | `12`             | PostgreSQL version                      |
| `preferredBackupWindow` |  | The daily time range during which automated backups are created if automated backups are enabled, using the `backupRetentionPeriod`` option. Must be in the format hh24:mi-hh24:mi.Must be in Universal Coordinated Time (UTC). Must not conflict with the preferred maintenance window. Must be at least 30 minutes.              |
| `backupRetentionPeriod`   | `1`           | The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups. |

#### redis

| Option      | Default          | Description                 |
|-------------|------------------|-----------------------------|
| `class`     | `cache.t2.micro` | Instance class              |
| `durable`   | `false`          | Multi-AZ automatic failover. When it is set to `true`, the option `nodes` has to be greater or equal to `2`, otherwise it will fail |
| `encrypted` | `false`          | Encrypt data at rest        |
| `nodes`     | `1`              | Number of nodes             |
| `version`   | `2.8.24`         | Redis version               |

#### efs

Use to share volumes between the tasks in different AZs and instances.

| Option        | Default  | Description                                                            |
|---------------|----------|------------------------------------------------------------------------|
| `encrypted`   | `false`  | Encrypt data at rest                                                   |
| `owner-gid`   | `1000`   | POSIX group ID to apply to the `path` directory                        |
| `owner-uid`   | `1000`   | POSIX user ID to apply to the `path` directory                         |
| `path`        | `/`      | The path on the file system used as the root directory by the services |
| `permissions` | `0777`   | POSIX permissions to apply to the `path` directory                     |

## AutoMinorVersionUpgrade

In case you specify the minor version on your resource definition you have to turn off the [AutoMinorVersionUpgrade](../reference/app-parameters#autominorversionupgrade) on your app parameter. It's enabled by default and it will update the DB instance during the maintenance window.
