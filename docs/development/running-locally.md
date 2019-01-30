---
title: "Running Locally"
order: 200
---

Convox can boot your application locally using Docker in an environment identical to production.

## Installing Docker

#### MacOS

Install [Docker Desktop](https://www.docker.com/products/docker-desktop) for MacOS.

Under Docker's preferences, change the following configuration:

* On the **Advanced** tab, increase CPUs to the halfway mark and Memory as high as you are comfortable with on your environment.
* On the **Kubernetes** tab check **Enable Kubernetes**.

#### Linux

Install the `microk8s` snap:

```
$ sudo snap install microk8s --classic
```

Once it is running, enable the DNS service:

```
$ microk8s.enable dns
```

#### Windows

Install [Docker Desktop](https://www.docker.com/products/docker-desktop) for Windows.

Under Docker's preferences, change the following configuration:

* On the **Advanced** tab, increase CPUs to the halfway mark and Memory as high as you are comfortable with on your environment.
* On the **Kubernetes** tab check **Enable Kubernetes**.

## Local Rack

#### Installation

    $ sudo convox rack install local

#### Running Applications

Use `convox start` to run applications against your local Rack.

You can use `convox switch local` to point your CLI at the local rack and use all `convox` commands normally.
