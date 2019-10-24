---
title: "Running Locally"
order: 200
---

Convox can boot your application locally using Docker in an environment identical to production.

## Installing Docker

#### MacOS

Install [Docker Desktop](https://www.docker.com/products/docker-desktop) for MacOS.

Under Docker's preferences, change the following configuration:

* Make sure you are on the **edge** version of Docker Desktop.
* On the **Advanced** tab, increase CPUs to the halfway mark and Memory as high as you are comfortable with on your environment.
* On the **Kubernetes** tab check **Enable Kubernetes**.

#### Linux

Install the `microk8s` snap:

```
$ sudo snap install microk8s --classic --channel=1.13/stable
```

Once it is running, enable a few additional services:

```
$ microk8s.enable dns storage
```

<div class="block-callout block-show-callout type-warning" markdown="1">
  Running this command quickly after installing the `microk8s` snap has been known to fail. If you encounter errors please wait a few minutes and try running it again.
</div>

Create an alias for the `kubectl` binary:

```
$ sudo snap alias microk8s.kubectl kubectl
```

#### Windows

Install [Docker Desktop](https://www.docker.com/products/docker-desktop) for Windows.

Under Docker's preferences, change the following configuration:

* Make sure you are on the **edge** version of Docker Desktop.
* On the **Advanced** tab, increase CPUs to the halfway mark and Memory as high as you are comfortable with on your environment.
* On the **Kubernetes** tab check **Enable Kubernetes**.

## Local Rack

#### Installation

    $ sudo convox rack install local

#### Running Applications

Use `convox start` to run applications against your local Rack.

You can use `convox switch local` to point your CLI at the local rack and use all `convox` commands normally.


<div class="block-callout block-show-callout type-warning" markdown="1">
  If your local Kubernetes setup does not point to a valid cluster, that can slow down your Convox CLI operations as it tries to interrogate the invalid endpoint.  In this case, you can set a local env var `$ export CONVOX_LOCAL=disable` to stop the CLI from doing this and speed up your commands.
</div>