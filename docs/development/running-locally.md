---
title: "Running Locally"
order: 200
---

Convox can boot your application locally using Docker in an environment identical to production.

## Installing Docker

#### OS X

Install [Docker for Mac](https://docs.docker.com/engine/installation/mac/#/docker-for-mac) to get a local Docker environment.

#### Linux

Install Docker from your application's default package manager.

## Installing the Local Rack

    $ sudo convox rack install local

## Running applications locally

Use `convox start` to run applications against your local Rack.

You can use `convox switch local` to point your CLI at the local rack and use all `convox` commands normally.
