---
title: "Running Locally"
description: "Run your Convox application locally using Docker and a local Rack for development and testing."
---

# Running Locally

Convox can boot your application locally using Docker in an environment identical to production.

## Installing Docker

### MacOS

Install [Docker Desktop](https://www.docker.com/products/docker-desktop) for MacOS.

Under Docker's preferences, change the following configuration:

* Make sure you are on the **edge** version of Docker Desktop.
* On the **Advanced** tab, increase CPUs to the halfway mark and Memory as high as you are comfortable with on your environment.
* On the **Kubernetes** tab check **Enable Kubernetes**.

### Linux

Install the `microk8s` snap:

```bash
$ sudo snap install microk8s --classic --channel=1.13/stable
```

Once it is running, enable a few additional services:

```bash
$ microk8s.enable dns storage
```

> **Warning:** Running this command quickly after installing the `microk8s` snap has been known to fail. If you encounter errors, wait a few minutes and try running it again.

Create an alias for the `kubectl` binary:

```bash
$ sudo snap alias microk8s.kubectl kubectl
```

### Windows

Install [Docker Desktop](https://www.docker.com/products/docker-desktop) for Windows.

Under Docker's preferences, change the following configuration:

* Make sure you are on the **edge** version of Docker Desktop.
* On the **Advanced** tab, increase CPUs to the halfway mark and Memory as high as you are comfortable with on your environment.
* On the **Kubernetes** tab check **Enable Kubernetes**.

## Local Rack

### Installation

```bash
$ sudo convox rack install local
```

### Running Applications

Use `convox start` to run applications against your local Rack.

You can use `convox switch local` to point your CLI at the local rack and use all `convox` commands normally.


> **Warning:** If your local Kubernetes setup does not point to a valid cluster, that can slow down your Convox CLI operations as it tries to interrogate the invalid endpoint. In this case, you can set a local env var `export CONVOX_LOCAL=disable` to stop the CLI from doing this and speed up your commands.

## See Also

- [Code Sync](/development/code-sync)
- [convox.yml Reference](/application/convox-yml)
- [Builds](/deployment/builds)
- [One-off Commands](/management/one-off-commands)