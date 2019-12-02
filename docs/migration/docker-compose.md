---
title: docker-compose.yml
layout: guide
description: Migrate your application from docker-compose.yml to convox.yml
---

So, you have your app running with a `docker-compose.yml` and you want to move to use Convox?  Here are some steps to help you migrate your configuration 😊  Given an entirely fictitious compose file, let's see how we can translate that to a shiny new `convox.yml`....

```yml
version: '2'
services:
  web:
    image: httpd
    ports:
      - 80:80
    volumes:
      - /tmp/something
    enviroment:
      - MY_ENVIRONMENT: development
    links:
      - supportservice
      - redis
  web_secondary:
    image: myusername/privateimage:latest
    command: override_command.sh
    ports:
      - 3001
    links:
      - redis
  redis:
    image: "redis:alpine"
  supportservice:
    build:
      context: .
      dockerfile: Dockerfile.support
      links:
        - postgres
  postgres:
    image: postgres:12
```

So here we have 5 different services defined, one a httpd service, one a private image with an overriden command, one a plain Redis service, and one we're building locally from a custom Dockerfile, with a Postgres database image thrown in for good measure.

```yml
resources:
  redis:
    type: redis
  postgres:
    type: postgres
    options:
      version: 12
services:
  web:
    image: httpd
    port: 80
    volumes:
      - /tmp/something
    enviroment:
      - MY_ENVIRONMENT=development
    links:
      - supportservice
    resources:
      - redis
  web_secondary:
    image: myusername/privateimage:latest
    command: override_command.sh
    port: 3001
    resources:
      - redis
  supportservice:
    build:
      path: .
      manifest: Dockerfile.support
    resources:
      - postgres
```

Here you can see how we've turned the Redis and Postgres services into resources, which can be shared between the other services, and backed by the most appropriate infrastructure in your cloud provider.  We've also tweaked some of the other syntax to match, but overall you can see how easy it was to move over.  The most commonly used, and useful, Docker compose configuration options are supported by Convox for deployment on the cloud provider of your choice.  Moving your supplementary services to Resources allows you to utilise the power and efficiency of RDS and Elasticache for instance.
There is also full support for building/pulling your images from a variety of sources.



