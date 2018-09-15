---
title: Definitions
layout: docs
permalink: /docs/definitions/
---

## Process

See [Service](#service).

## Resource

<div class="alert alert-warning">
Previously called a service.
</div>

Resources are like services, but are external to your application and which your application communicates with over a network.

Examples of typical resources used with Convox are data stores like RDS or redis, mailservers, and so on.

Resources are typically stateful services.

Resources correspond closely to the twelve-factor app notion of [_backing services_](https://12factor.net/backing-services).

## Service

<div class="alert alert-warning">
Previously called a process.
</div>

A _service_ is a component of your app, defined as a command run against an image. An app is composed of one or more services, e.g. `web` and `worker`.
