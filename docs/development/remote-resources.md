---
title: Remote Resources
order: 500
---

It's often useful to be able to interact with remote services or resources (e.g. databases) during local development. However, access to those remote services is often rightly restricted to traffic coming from the application itself, rather than from the internet.

In these cases, `convox proxy` and `convox resources proxy` can help by proxying traffic from your local development machine to your app's services via your Rack.

## Access external resources

When possible, [resources](/docs/about-resources/) are generally configured so they are not accessible from the public internet.  You can get proxy access with the `convox resources proxy` command, e.g.:

```
$ convox resources proxy mysql-4624
```

Under the hood, this creates an encrypted and secure websocket tunnel through your Rack process that talks to your database via a socket on your local machine. In other words, `convox resources proxy` allows you to communicate with your Postgres database (for example) as if it were running on your own machine.

### Example

In our [Flask example repo](https://github.com/convox-examples/flask), we've created a `counter` service and replaced the `redis` service with a redis resource (referred to here as `redis-9990`).

We can't connect directly to our redis resource because it's inside our app's VPC and only accepts internal requests:

```
$ redis-cli -h lerabcdefghi.jklm5o.ng.0001.use1.cache.amazonaws.com:6379
Could not connect to Redis at lerabcdefghi.jklm5o.ng.0001.use1.cache.amazonaws.com:6379: Connection timed out
```

The manual way to interact with the redis resource would be to use `convox exec` to get a shell in the `counter` container and install `redis-cli` there, but this is neither efficient nor recommended.

`convox resources proxy` makes it much easier. In a terminal, run:

```
$ convox resources proxy redis-9990
proxying 127.0.0.1:6379 to lerabcdefghi.jklm5o.ng.0001.use1.cache.amazonaws.com:6379
```
Then, in another terminal on your local machine, run:

```
$ redis-cli
127.0.0.1:6379>
```
You'll see `connect: 6379` appear in the terminal where you ran `convox resources proxy` as evidence of your connection.

Now you can indirect with your remote redis resource:

```
127.0.0.1:6379> ping "hello world"
"hello world"
```

<div class="block-callout block-show-callout type-info" markdown="1">
**Tip: Alternate ports** 

To avoid local port conflicts (for example, if you already have Postgres running on 5432), you can specify a different local port with `--listen`:

`$ convox resources proxy postgres-6525 --listen 5433`

When you run the command above, you should see output like:

`proxying 127.0.0.1:5433 to dev-postgres-6525.cbm068zjzjcr.us-east-1.rds.amazonaws.com:5432`

and then you should be able to connect to localhost:5433.
</div>

## Access private services with `convox proxy`

`convox proxy` works much the same way as `convox resources proxy` except for internal services (i.e. exposed behind an internal ELB inside a VPC) rather than [external resources](/docs/about-resources).

### Example

Let's say we want to proxy from port 8000 on our local machine to port 80 of the `counter` service in our [example repo](https://github.com/convox-examples/flask):

First get the service endpoint by running `convox services`:

    $ convox services
    SERVICE  DOMAIN                                                       PORTS
    web      flask-counter-NJS6JDQ-842041791.us-east-1.elb.amazonaws.com  80

Then choose a port on your local machine (`8000` in our case), then run the `convox proxy` command in a terminal:

```
$ convox proxy 8000:flask-counter-NJS6JDQ-842041791.us-east-1.elb.amazonaws.com:80
proxying 127.0.0.1:8000 to flask-counter-NJS6JDQ-842041791.us-east-1.elb.amazonaws.com:80
```

In another terminal, run `curl localhost:8000`:

```
$ curl -s localhost:8000 | grep served
        <h2>This request was served by 3d511a719fc7.</h2>
        <h2>3d511a719fc7 served 7 requests so far.</h2>
        <h2>A grand total of 7 requests were served.</h2>
```

The page is returned as if it were being served from `localhost`.

## See also

- [Debugging](/docs/debugging)
- [Running locally](/docs/running-locally)
