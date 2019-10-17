---
title: "Build & Release"
---

Convox follows some accepted best practices when it comes to a reliable, dependable and repeatable build process.

All of these steps can be automated and run from CI/CD processes through the Convox CLI!

### Builds

<div class="block-callout block-show-callout type-info" markdown="1">
When you [deploy or build](/deployment/builds) your application a "build" artifact is created. This consists of the Docker images that make up your application plus metadata stored in your Rack's database. Each build has a unique ID and is associated with one or more [releases](/docs/releases).
</div>

Builds can either be created:

- Manually from the CLI.  `convox build` will create your build artifact, whereas `convox deploy` will create the artifact and also then promote the subsequent release.
- As part of a Convox [workflow](/console/workflows).  Workflows can connect to your code repository and initiate builds and subsequent activities from code pushes or pull requests.
- As part of an external CI process.  Convox has a [CircleCI Orb](/external-services/circleci) for easy integration.  Other CI services can also easily integrate with Convox by [installing the Convox CLI](/introduction/installation) in your CI environment, utilising [deploy keys](/console/deploy-keys) to keep your credentials secure, and running the `build` or `deploy` commands as appropriate in your CI workflow.

You can see a list of all the builds for your app by running `convox builds` from the CLI.

```
$ convox builds
ID           STATUS    RELEASE      STARTED      ELAPSED  DESCRIPTION
BBLTITAIGHI  complete  ROCHMCOUESG  1 week ago   2m46s    build 7329df90c1 this is my commit message
BNNJBWFAWYJ  complete  RVUGSLANDXU  2 weeks ago  1m19s
BPNRIHAOQGM  failed                 2 weeks ago  11s
BRMUBCYGSWH  complete  RIHVPQXQWSR  3 weeks ago  1m8s
```

### Releases

<div class="block-callout block-show-callout type-info" markdown="1">
Every time you deploy an application or change its environment a new [Release](/docs/releases) is created. You can promote any Release for an application to deploy it. Promoting a Release will cause a [rolling update](/docs/rolling-updates) that gracefully deploys the desired Release.
</div>

Releases will automatically be created:

- Once a build artifact is created through any of the methods mentioned [above](#builds).
- When you update your environment variables the Console.
- When you update your environment variables from the CLI.

You can see the list of all releases, and which one is currently active for your app in the Console under the 'Releases' tab, or via the CLI by running `convox releases`.

```
convox releases
ID           STATUS  BUILD        CREATED      DESCRIPTION
ROCHMCOUESG  active  BBLTITAIGHI  1 week ago   build 7329df90c1 this is my commit message
RCTSJKKLHXJ          BNNJBWFAWYJ  2 weeks ago  env add:FOO change:BAR remove:BAZ
RVUGSLANDXU          BNNJBWFAWYJ  2 weeks ago
RIHVPQXQWSR          BRMUBCYGSWH  3 weeks ago
```

### Release Promotion

<div class="block-callout block-show-callout type-info" markdown="1">
When a Release is promoted, new processes are gracefully [rolled out](/docs/rolling-updates) into production.
If there are errors starting new processes, new processes are not [verified as healthy](/application/health-checks), or the rollout doesn't complete in 10 minutes, an automatic [rollback](/deployment/rolling-back) is performed.
</div>

If running a `convox deploy` then the Release will automatically be promoted with a rolling deployment across your Rack.  Environment variable changes performed from the Console will cause a new Release to be created and then promoted once you click on 'Apply Changes'.

Releases created from a straight `convox build` or an [environment update from the CLI](/application/environment#setting-and-editing-your-environment-variables) will need to be promoted.  This can be done from the Console by clicking on the 'Promote' button when looking at the Release list for an app.  Alternatively you can `convox releases promote <release id>` from the CLI.

### Rollbacks

<div class="block-callout block-show-callout type-info" markdown="1">
Convox keeps history for your Apps and for Rack and CLI releases. It's possible to [roll back](/deployment/rolling-back) to previous versions of all of these.
</div>

If you wish to revert to a previous release for whatever reason, Convox makes that simple.  As Releases also contain state about environment variables, you can even revert to previous versions of these with ease!

In the Console, when looking at the list of releases for an app, ones older than the currently active release can be simply rolled back to by clicking the 'Rollback' button.  Convox will then create a copy of the older Release, place it at the top of your Release list, and promote it.  This causes a rolling update to this copied release across your Rack.

From the CLI, you can `convox releases rollback <release id>` to kick off the rollback process.

Previous releases are not lost or forgotten in the rollback process, they will stay in your Release list allowing you to `rollback` to any other release at any time you wish.
