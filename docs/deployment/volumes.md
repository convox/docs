---
title: "Volumes"
order: 600
---

You can use Docker volumes to make data persist between runs of a given service's containers across restarts and instance replacements. This is useful for applications like WordPress or Jenkins that need to store data on the filesystem.

<div class="block-callout block-show-callout type-warning" markdown="1">
<b>Note:</b> It's not currently possible to share volumes between containers of different services in the same app.
</div>

## Shared Filesystem

Convox uses a network filesystem backed by [Amazon EFS](https://aws.amazon.com/efs/) that is shared among all of the instances in your Rack.

<div class="block-callout block-show-callout type-warning" markdown="1">
### Supported Regions

Shared volumes are only possible in AWS regions that support Elastic File System (EFS). As of the time of this writing, the following regions are supported:

- ap-southeast-2 (Sydney)
- eu-central-1
- eu-west-1 (Ireland)
- us-east-1 (N. Virginia)
- us-east-2 (Ohio)
- us-west-1 (N. California)
- us-west-2 (Oregon)

To check whether EFS is available in a given region, select your region from the dropdown in [the EFS section in the AWS console](https://console.aws.amazon.com/efs/home).

Volumes used in regions without EFS will not persist or be shared across instances.
</div>

## Sharing Data

You can mount a shared volume by specifying a container path in the `volumes:` section of your `convox.yml`: 

```
services:
  web:
    volumes:
      - /my/shared/data
```

If you specify your volume path this way, Convox will persist data on your Rack instances in an application-namespaced path under `/volumes`.

<div class="block-callout block-show-callout type-info" markdown="1">
  By default, volumes are located at `/volumes/<rack>-<app>/<service>/*`.
  You can also specify the volume in the more explicit `host:container` format, e.g. `/host/path:/container/path`. This allows you to set exactly where on the host instance to persist the data.
</div>

## Permissions

It's not currently possible to mount EFS volumes as read-only with Convox. If this option would be useful to you, feel free to open a support ticket so we can track it as a feature request.

## Persistence

Files and directories written to the volume inside the container will be persisted between processes of the same service type at that location.

## Example: WordPress

WordPress is a popular PHP blogging platform. It expects a persistent filesystem for storing themes, plugins, and media uploads. You can persist this data by specifying a shared volume at `/var/www/html`:

```
services:
  web:
    image: wordpress:4.5.2-apache
    ports:
      - 80:80
    volumes:
      - /var/www/html
```

This configuration will work with both `convox start` and `convox deploy`. Files written to `/var/www/html` will be persisted on the Docker host.
