---
title: "Running Locally"
---

Convox can boot your application locally using Docker in an environment identical to production.

## Installing Docker

#### OS X

Install [Docker for Mac](https://docs.docker.com/engine/installation/mac/#/docker-for-mac) to get a local Docker environment.

#### Linux

Install Docker from your application's default package manager.

## Starting the Application

### convox start

Use `convox start` to build and run your application locally.

    $ cd ~/myapp
    $ convox start
    RUNNING: docker build -t knexsfvjdc ~/myapp
    Sending build context to Docker daemon 8.192 kB
    Step 0 : FROM ruby:2.2.2
     ---> 9664620d4c2a
    Step 1 : EXPOSE 3000
     ---> Running in d2894bf8d64b
    ...
    web    | docker run -i --name myapp-web -p 5000:3000 procfile/web sh -c ruby web.rb
    web    | [2015-09-18 06:16:53] INFO  WEBrick 1.3.1
    web    | [2015-09-18 06:16:53] INFO  ruby 2.2.2 (2015-04-13) [x86_64-linux]
    web    | == Sinatra (v1.4.6) has taken the stage on 3000 for development with backup from WEBrick
    web    | [2015-09-18 06:16:53] INFO  WEBrick::HTTPServer#start: pid=7 port=3000

This will read your `docker-compose.yml` and use the information found there to boot all of your app's processes and apply configured [links](/docs/gen1/linking). Local code changes will be [synced](/docs/code-sync) with your running processes in real time.

To exit `convox start` gracefully, press `Ctrl+C`. To force-quit, press `Ctrl+C` again.

### File syncing

Local code changes will be [synced](/docs/code-sync) with your running processes in real time. To disable this, pass the `--no-sync` flag to `convox start`.

### Caching

Your app is rebuilt each time you run `convox start`, which takes advantages of Docker's built-in caching mechanism by default. To bypass the cache and rebuild from scratch, you can pass the `--no-cache` flag to `convox start`.

### Running without building

As explained above, your app is rebuilt each time you run `convox start`. If you'd like to skip the build step and start your app with previously built images, you can pass the `--no-build` flag to `convox start`.

### Shifting ports

You can offset the external ports of processes run by `convox start` by a given number, allowing you to easily run multiple applications on one host without port conflicts.

There are two ways to accomplish this: via the `--shift` flag, or via service labels.

#### `convox start --shift`

Passing the `--shift` flag to `convox start` will offset the public ports of all services by the provided number. For example, on an app that normally publishes port 80, running `convox start --shift 1` will publish port 81 instead, and update any corresponding port environment variables (`SERVICENAME_PORT`, etc) and container names.

Additionally, all other port labels (`convox.port.<port>.protocol`, `convox.port.<port>.proxy`, `convox.port.<port>.secure`) will have their port values shifted.

```
$ convox start
ERROR: ports in use: [80 443]
```

```
$ convox start --shift 1
build  │ running: docker build -f /home/aj/git/convox/site/Dockerfile -t site-staging/web /home/aj/git/convox/site
Sending build context to Docker daemon 19.95 MB 557.1 kB
[...]
web    │   Server running... press ctrl-c to stop.
```

#### `convox.start.shift` label

You can make this port shifting more persistent on a per-service basis with the [convox.start.shift](/docs/gen1/docker-compose-labels/#convoxstart) label in `docker-compose.yml`:

```
  labels:
    - convox.start.shift=2
```

When shifting ports, you can view the differences with `docker ps` (ports `81` and `444` as opposed to `80` and `443`):

```
$ docker ps
IMAGE               COMMAND                  PORTS                                              NAMES
convox/proxy        "proxy-link 444 4001 "   0.0.0.0:444->444/tcp                               site-staging-web-proxy-444
convox/proxy        "proxy-link 81 4001 t"   0.0.0.0:81->81/tcp                                 site-staging-web-proxy-81
site-staging/web    "_bin/web"               0.0.0.0:32788->4001/tcp, 0.0.0.0:32789->4001/tcp   site-staging-web
```

Your app will then be available at `http://0.0.0.0:81`.

Note: The `--shift` and `convox.start.shift` label values are cumulative. If you have services `x` and `y`, and you define `convox.start.shift=1000` for the `x` service and then run `convox start --shift 2000`, service `x` will have its ports shifted by `3000` and service `y` will have its ports shifted by `2000`.

### Environment

`convox start` will read variables defined in a file called `.env` in the project root directory. For more information, see the [Environment](/docs/gen1/environment/#local) documentation.

## Data persistence

If your app uses a database, you'll find it useful to run a database container locally for development. To persist data for that container between starts, you can use Docker volumes to mount a host directory to the container's data directory. For example, if you're using the `convox/postgres` image, you can persist your data like this:

```yaml
database:
  image: convox/postgres
  ports:
    - 5432
  volumes:
    - /var/lib/postgresql/data
```

Convox does not recommend running datastores as containers in your Rack. Instead, you should use a hosted service, such as the [Postgres](/docs/postgresql) resource that Convox configures using Amazon RDS, or externally-hosted resources like [Compose.io's MongoDB](https://www.compose.com/mongodb). For more information, see [Resources](/docs/about-resources/).

## Interacting with remote resources during development

If you have set up a database (or other) resource, you may wonder how to interact with it during local development.

You have several options:

### Tunnel to remote databases

You can use `convox resources proxy` to tunnel to the remote production (or staging, or any) database as described [here](/docs/remote-resources). This will allow your app to interact with the remote resource as if it were running on your own machine.

### Use local containers

You can also use local containers as defined via the services in `docker-compose.yml`. The environment variables your app should use to communicate between containers will be automatic for linked services as described [here](/docs/gen1/environment#linking).

You'll want to [scale the remote services to `-1`](/docs/gen1/scaling/#scaling-down-unused-services) to avoid creating unnecessary containers and load balancers in production.

## Local Volumes

If your `docker-compose.yml` specifies volumes, they will be created on your local machine at `~/.convox/volumes`.

For more information, see [Persistence for local containers](/docs/volumes/#persistence-for-local-containers).
