---
title: "Linking"
---

Convox allows you to link containers by declaring associations in the `docker-compose.yml` manifest. Links are created by injecting environment variables that point at the linked container.

This avoids the need for your application to interface with configuration or service discovery mechanisms to find the other services it needs to talk to.

To link containers add a `links` section to the relevant container in your `docker-compose.yml`:

## Defining Links

```bash
web:
  build: .
  ports:
    - 80:80
  links:
    - database
database:
  image: convox/postgres
  ports:
    - 5432
```

Configuring the `links` section in this way will cause the following environment variables to be set for the `web` service:

- `DATABASE_URL`
- `DATABASE_SCHEME`
- `DATABASE_HOST`
- `DATABASE_PORT`
- `DATABASE_PATH`
- `DATABASE_USERNAME`
- `DATABASE_PASSWORD`

Here is an example of what those actually look like with the `convox/postgres` image:

#### Local environment

	DATABASE_HOST=172.17.0.1
	DATABASE_PASSWORD=password
	DATABASE_PATH=/app
	DATABASE_PORT=5432
	DATABASE_SCHEME=postgres
	DATABASE_URL=postgres://postgres:password@172.17.0.1:5432/app
	DATABASE_USERNAME=postgres

#### Production

	DATABASE_HOST=postgres-i-191910196.us-east-1.elb.amazonaws.com
	DATABASE_PASSWORD=password
	DATABASE_PATH=/app
	DATABASE_PORT=5432
	DATABASE_SCHEME=postgres
	DATABASE_URL=postgres://postgres:password@postgres-i-191910196.us-east-1.elb.amazonaws.com:5432/app
	DATABASE_USERNAME=postgres

## Linkable Images

Most images will work out of the box and will use the `tcp` scheme by default. If you'd like to override the scheme or supply default authentication credentials, you can set the following `ENV` variables in your image's `Dockerfile`:

* `LINK_SCHEME`
* `LINK_USERNAME`
* `LINK_PASSWORD`
* `LINK_PORT`
* `LINK_PATH`

See the [convox/redis](https://github.com/convox/redis/blob/9b56f5553ce6dd0a2f72d76b752f1dded287f109/Dockerfile#L10-L13) Dockerfile for an example.

#### Application-level overrides

You can also override these `LINK_` variables in your application's `docker-compose.yml`:

```yaml
database:
  image: example/postgres
  environment:
    - LINK_SCHEME=postgres
    - LINK_PASSWORD=password2
    - LINK_USERNAME=postgres
    - LINK_PORT=5432
  ports:
    - 5432
```

## Ports

When a service declares a link, the linked container (`database` in our example) needs to expose at least one port so convox can create a load balancer and construct the service URL.

When multiple ports are specified, the first one in the list of ports is used to construct the link URL. This can be overridden with the `LINK_PORT` environment variable described above.

## See also

- [Environment](/docs/gen1/environment/)
