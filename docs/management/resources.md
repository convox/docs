---
title: Accessing Resources
---

Application [Resources](/application/resources) are only accessible inside the Rack.

If you need to access these resources to import/export data or perform other management functions, you can use `convox resources proxy`.

```
$ convox resources proxy mysql-4624
```

## Examples

Assuming the following `convox.yml`:

```yaml
resources:
  cache:
    type: redis
```

Open a terminal and run:

```
$ convox resources proxy cache
Proxying localhost:6379 to redis-0001.internal.cache.amazonaws.com
```

Then, in another terminal on your local machine, run:

```
$ redis-cli -h localhost -p 6379
127.0.0.1:6379>
```

You'll see `connect: 6379` appear in the terminal where you ran `convox resources proxy` as evidence of your connection.

Now you can interact with your remote redis resource:

```
127.0.0.1:6379> ping "hello world"
"hello world"
```

#### Options

To avoid local port conflicts you can specify a different local port with `--port`:

```
$ convox resources proxy cache --port 6380
```
