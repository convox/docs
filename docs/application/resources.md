---
title: Resources
---

A resource is a network-attached dependency of your application.

## Definition

```yaml
resources:
  database:
    type: postgres
    options:
      storage: 100
  services:
    web:
      resources:
        - database
```

#### Dynamic Configuration

If you'd like to make your resource configuration variable (e.g. staging/production) you can use interpolation:

```yaml
resources:
  database:
    type: postgres
    options:
      storage: ${POSTGRES_STORAGE_SIZE}
```

```
$ convox env set POSTGRES_STORAGE_SIZE=200
```

## Available Resources

#### memcached

| Option    | Default          | Description       |
|-----------|------------------|-------------------|
| `class`   | `cache.t2.micro` | Instance class    |
| `nodes`   | `1`              | Number of nodes   |
| `version` | `1.4.34`         | Memcached version |

#### mysql

| Option    | Default          | Description                             |
|-----------|------------------|-----------------------------------------|
| `class`   | `db.t2.micro` | Instance class                          |
| `durable` | `false`          | Automatic failover                      |
| `iops`    |                  | Provisioned IOPS for database disks     |
| `storage` | `20`             | GB of storage to provision              |
| `version` | `5.7.22`         | MySQL version                           |

#### redis

| Option      | Default          | Description          |
|-------------|------------------|----------------------|
| `class`     | `cache.t2.micro` | Instance class       |
| `durable`   | `false`          | Automatic failover   |
| `encrypted` | `false`          | Encrypt data at rest |
| `nodes`     | `1`              | Number of nodes      |
| `version`   | `2.8.24`         | Redis version        |

#### postgres

| Option    | Default          | Description                             |
|-----------|------------------|-----------------------------------------|
| `class`   | `db.t2.micro`    | Instance class                          |
| `durable` | `false`          | Automatic failover                      |
| `iops`    |                  | Provisioned IOPS for database disks     |
| `storage` | `20`             | GB of storage to provision              |
| `version` | `9.6.6`          | PostgreSQL version                      |

