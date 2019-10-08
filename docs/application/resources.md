---
title: Resources
---

A resource is a network-attached dependency of your application.

## Definition

Here we define a resource called `primary_database` that is a Postgres database with 100GB of storage, and we then link it to our `web` service, which will be given an [env var](#accessing-resources-through-environment-variables) to connect to the database:

```yaml
resources:
  primary_database:
    type: postgres
    options:
      storage: 100
  services:
    web:
      resources:
        - primary_database
```

The resource name only affects the [environment variable name](#accessing-resources-through-environment-variables) that is passed to your services.  You are free to name it what you wish with no regard to the type of resource.

#### Dynamic Configuration

If you'd like to make your resource configuration variable (e.g. to produce different results between environments like staging vs production) you can use environment variable interpolation:

```yaml
resources:
  primary_database:
    type: postgres
    options:
      storage: ${POSTGRES_STORAGE_SIZE}
```

```
$ convox env set POSTGRES_STORAGE_SIZE=50 --rack=acme/staging
$ convox env set POSTGRES_STORAGE_SIZE=200 --rack=acme/production
```

## Accessing Resources through Environment variables

You can access defined resources from services with environment variables.
In the above example, the `primary_database` resource provides a `PRIMARY_DATABASE_URL` variable that is accessible from the `web` service.
The environment variable name is the resource name converted to all-caps, with a `_URL` suffix.

This would contain the entire connection string you would need, ie:

```
PRIMARY_DATABASE_URL=postgres://username:password@host.com:5432/databaseName
```


## Available Resources

#### memcached

| Option    | Default          | Description       |
|-----------|------------------|-------------------|
| `class`   | `cache.t2.micro` | Instance class    |
| `nodes`   | `1`              | Number of nodes   |
| `version` | `1.4.34`         | Memcached version |

#### mysql

| Option      | Default          | Description                             |
|-------------|------------------|-----------------------------------------|
| `class`     | `db.t2.micro`    | Instance class                          |
| `encrypted` | `false`          | Encrypt data at rest                    |
| `durable`   | `false`          | Automatic failover                      |
| `iops`      |                  | Provisioned IOPS for database disks     |
| `storage`   | `20`             | GB of storage to provision              |
| `version`   | `5.7.22`         | MySQL version                           |

#### postgres

| Option      | Default          | Description                             |
|-------------|------------------|-----------------------------------------|
| `class`     | `db.t2.micro`    | Instance class                          |
| `durable`   | `false`          | Automatic failover                      |
| `encrypted` | `false`          | Encrypt data at rest                    |
| `iops`      |                  | Provisioned IOPS for database disks     |
| `storage`   | `20`             | GB of storage to provision              |
| `version`   | `9.6`            | PostgreSQL version                      |

#### redis

| Option      | Default          | Description          |
|-------------|------------------|----------------------|
| `class`     | `cache.t2.micro` | Instance class       |
| `durable`   | `false`          | Automatic failover   |
| `encrypted` | `false`          | Encrypt data at rest |
| `nodes`     | `1`              | Number of nodes      |
| `version`   | `2.8.24`         | Redis version        |
