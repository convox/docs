---
title: "Builds"
description: "Create, inspect, export, and import application build artifacts on Convox."
---

# Builds

When you deploy or build your application a "build" artifact is created. This consists of the Docker images that make up your application plus metadata stored in your Rack's database. Each build has a unique ID and is associated with one or more [Releases](/deployment/releases).

Builds can either be created:

- Manually from the CLI.  `convox build` will create your build artifact, whereas `convox deploy` will create the artifact and also then promote the subsequent release.
- As part of a Convox [workflow](/console/workflows).  Workflows can connect to your code repository and initiate builds and subsequent activities from code pushes or pull requests.
- As part of an external CI process.  Convox has a [CircleCI Orb](/integrations/circleci) as well as a set of [Github Actions](/integrations/github-actions) for integration.  Other CI services can also integrate with Convox by [installing the Convox CLI](/introduction/installation) in your CI environment, utilising [deploy keys](/console/deploy-keys) to keep your credentials secure, and running the `build` or `deploy` commands as appropriate in your CI workflow.

You can see a list of all the builds for your app by running `convox builds` from the CLI.

```bash
$ convox builds
ID           STATUS    RELEASE      STARTED      ELAPSED  DESCRIPTION
BBLTITAIGHI  complete  ROCHMCOUESG  1 week ago   2m46s    build 7329df90c1 this is my commit message
BNNJBWFAWYJ  complete  RVUGSLANDXU  2 weeks ago  1m19s
BPNRIHAOQGM  failed                 2 weeks ago  11s
BRMUBCYGSWH  complete  RIHVPQXQWSR  3 weeks ago  1m8s
```

## Creating a Build

Build your application by running `convox build` from the top-level directory:

```bash
$ convox build
```

When you run this command the following steps are executed:

- The Convox CLI builds a tarball from all the files in your project (except the ones specified in `.dockerignore`)
- The CLI uploads the tarball to your Rack
- The Rack extracts the tarball and reads `convox.yml`
- Docker images are built or pulled as specified by `convox.yml`
- The images are tagged and pushed into your Rack's private Amazon ECR registry
- Build metadata is saved to the Rack
- A new [Release](/deployment/releases) is created from the build and its metadata is saved too

The newly created release will not be promoted (made active) until you run `convox releases promote <release ID>`.

If you'd like to build your app and promote the release in a single step, you can run `convox deploy` rather than `convox build`.

## Inspecting Builds

Run `convox builds info <build ID>` to view metadata for a particular build.

Run `convox builds logs <build ID>` to view the logs for a particular build.

## Moving Builds

It's possible to export a build from one app and import it to another app, even if the apps are on different Racks.

To move a build, first export it:

```bash
$ convox builds export <build ID> -a <appname> > build.tgz
```

You can then import the build into another app, even on a different Rack:

```bash
$ convox builds import -a <appname> < build.tgz
```

You can even pipe these commands together directly:

```bash
$ convox builds export <build ID> -a <app1> -r <rack1> | convox builds import -a <app2> -r <rack2>
```

## Build Options

The `convox build` and `convox deploy` commands accept several flags to customize the build process.

| Flag | Description |
|------|-------------|
| `--build-args KEY=VALUE` | Pass Docker build arguments. Can be specified multiple times. |
| `--description`, `-d` | Set a description for the Build and its Release. Defaults to the current git commit message. |
| `--development` | Create a development Build. When set, the build automatically targets a Dockerfile stage named `development` (via `--target development`) if one exists. |
| `--manifest`, `-m` | Specify an alternate manifest file (default: `convox.yml`). |
| `--no-cache` | Build without using the Docker layer cache. |
| `--wildcard-domain` | Enable wildcard domain support for the Release created from this Build. Adds `*.<app>-<service>.<rack-host>` listener rules and provisions a wildcard ACM certificate. |

### Build arguments

Convox also respects the `ARG` Dockerfile directive. For more information, see [Dockerfile: ARG](/application/dockerfile#arg).

## See Also

- [Releases](/deployment/releases)
- [Dockerfile](/application/dockerfile)
- [Rolling Updates](/deployment/rolling-updates)
- [CI Integrations](/integrations/integrations)
- [CI/CD Integrations](/integrations/cicd)
