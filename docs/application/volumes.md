---
title: "Volumes"
---

You can use Docker volumes to make data persist between runs of a given service's containers across restarts and instance replacements. This is useful for applications like WordPress or Jenkins that need to store data on the filesystem.

## Definition

```yaml
services:
  web:
    volumes:
      - /my/shared/data
```

In this example Convox would mount a persistent data volume for processes in the `web` service at `/my/shared/data`.

Data written to this volume will be persisted and shared between all processes of the `web` service.

## Example

WordPress is a popular PHP blogging platform. It expects a persistent filesystem for storing themes, plugins, and media uploads. You can persist this data by specifying a shared volume at `/var/www/html`:

```yaml
services:
  web:
    image: wordpress:4.5.2-apache
    port: 80
    volumes:
      - /var/www/html
```

## Host Volumes

Certain applications and containers may wish to access host volumes rather than mount a shared volume across the Rack.  You can do this by specifying the host path you wish to mount, colon separated with the local path in the container, as per usual Docker syntax:

```yaml
    volumes:
      - /sys/fs/cgroup/:/host/sys/fs/cgroup/
      - /proc/:/host/proc/
      - /var/run/docker.sock:/var/run/docker.sock
```

<div class="block-callout block-show-callout type-info" markdown="1">
  Only certain specific host paths are supported for security reasons. (`/cgroup/`, `/dev/log`, `/etc/passwd`, `/proc/`, `/sys/fs/cgroup/`, `/var/log/audit/`, `/var/run/`, `/var/run/docker.sock`)<br/>
</div>
