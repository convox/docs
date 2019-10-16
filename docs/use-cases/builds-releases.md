---
title: "Build & Release"
---

Convox follows some accepted best practices when it comes to a reliable, dependable and repeatable build process.

- When you [deploy or build](/deployment/builds) your application a "build" artifact is created. This consists of the Docker images that make up your application plus metadata stored in your Rack's database. Each build has a unique ID and is associated with one or more [releases](/docs/releases).

- Every time you deploy an application or change its environment a new [Release](/docs/releases) is created. You can promote any Release for an application to deploy it. Promoting a Release will cause a [rolling update](/docs/rolling-updates) that gracefully deploys the desired Release.

- When a Release is promoted, new processes are gracefully [rolled out](/docs/rolling-updates) into production.
If there are errors starting new processes, new processes are not [verified as healthy](/application/health-checks), or the rollout doesn't complete in 10 minutes, an automatic [rollback](/deployment/rolling-back) is performed.

- Convox keeps history for your Apps and for Rack and CLI releases. It's possible to [roll back](/deployment/rolling-back) to previous versions of all of these.

All of these steps can be automated and run from CI/CD processes through the Convox CLI!

### Builds

Builds can either be created:

- Manually from the CLI.  `convox build` will create your build artifact, whereas `convox deploy` will create the artifact and also then promote the subsequent release.
- As part of a Convox [workflow](/console/workflows).  Workflows can connect to your code repository and initiate builds and subsequent activities from code pushes or pull requests.
- As part of an external CI process.  Convox has a [CircleCI Orb](/external-services/circleci) for easy integration.  Other CI services can also easily integrate with Convox by [installing the Convox CLI](/introduction/installation) in your CI environment, utilising [deploy keys](/console/deploy-keys) to keep your credentials secure, and running the `build` or `deploy` commands as appropriate in your CI workflow.

You can see a list of all the builds for your app by running `convox builds` from the CLI.

### Releases

Releases will automatically be created once a build artifact is created.  They are also created when a change happens to your environment variables, either through the Console or from the CLI.

You can see the list of all releases, and which one is currently active for your app in the Console under the 'Releases' tab, or via the CLI by running `convox releases`.

If running a `convox deploy` then the Release will automatically be promoted with a rolling deployment across your Rack.  Environment variable changes performed from the Console will also be automatically released once you click on 'Apply Changes'.

Releases created from a straight `convox build` or an [environment update from the CLI](/application/environment#setting-and-editing-your-environment-variables) will need to be promoted.  This can be done from the Console by clicking on 'Promote' button when looking at the Release list for an app.  Alternatively you can `convox releases promote <release id>` from the CLI.

### Rollbacks

If you wish to revert to a previous release for whatever reason, Convox makes that simple.  As Releases also contain state about environment variables, you can even revert to previous versions of these with ease!

In the Console, when looking at the list of releases for an app, ones older than the currently active release can be simply rolled back to by clicking the 'Rollback' button.  Convox then handles the rolling update to this older release across your Rack.

From the CLI, you can `convox releases rollback <releease id>` to kick off the rollback process.

Newer releases are not lost or forgotten in the rollback process so you can just as easily `promote` and roll forward to a more recent release whenever you like as well.
