---
title: "Volumes"
---

You can use Docker volumes to make data persist between runs of a given service's containers across restarts and instance replacements. This is useful for applications like WordPress or Jenkins that need to store data on the filesystem.

## Definition

```
services:
  web:
    volumes:
      - /my/shared/data
```

In this example Convox would mount a persistent data volume for processes in the `web` service at `/my/shared/data`.

Data written to this volume will be persisted and shared between all processes of the `web` service.

## Example

WordPress is a popular PHP blogging platform. It expects a persistent filesystem for storing themes, plugins, and media uploads. You can persist this data by specifying a shared volume at `/var/www/html`:

```
services:
  web:
    image: wordpress:4.5.2-apache
    port: 80
    volumes:
      - /var/www/html
```
