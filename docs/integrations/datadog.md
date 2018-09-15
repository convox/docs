---
title: "Datadog"
order: 300
---


You can add operational visibility to your Convox environments with Datadog.

## Sign up for Datadog

If you don’t have an account already, [sign up for Datadog](https://app.datadoghq.com/signup). You’ll need an API key that lets you send data from Convox to the Datadog service.

## Deploy the Datadog Agent

You can deploy the datadog agent as a Convox app with a very simple `convox.yml` manifest:

```
services:
  datadog:
    agent: true
    environment:
      - API_KEY=${DATADOG_API_KEY}
      - DD_APM_ENABLED=true
    image: datadog/docker-dd-agent
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - /proc/:/host/proc
      - /cgroup:/host/sys/fs/cgroup
```
