---
title: "Releases"
order: 300
---

<div class="block-callout block-show-callout type-info" markdown="1">
Every time you deploy an application or change its environment a new Release is created. You can promote any Release for an application to deploy it. Promoting a Release will cause a [rolling update](/docs/rolling-updates) that gracefully deploys the desired Release.
</div>

Releases will automatically be created:

- Once a build artifact is created through any of the methods mentioned [above](#builds).
- When you update your environment variables in the Console.
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

## Promoting a Release

#### Console

* Click the **Promote** button next to the Release you would like to promote.

#### CLI

* Run `convox releases promote <id>`

<div class="block-callout block-show-callout type-info" markdown="1">
When a Release is promoted, new processes are gracefully [rolled out](/docs/rolling-updates) into production.
If there are errors starting new processes, new processes are not [verified as healthy](/application/health-checks), or the rollout doesn't complete in 10 minutes, an automatic [rollback](/deployment/rolling-back) is performed.
</div>

If running a `convox deploy` then the Release will automatically be promoted with a rolling deployment across your Rack.  Environment variable changes performed from the Console will cause a new Release to be created and then promoted once you click on 'Apply Changes'.

Releases created from a straight `convox build` or an [environment update from the CLI](/application/environment#setting-and-editing-your-environment-variables) will need to be promoted.  This can be done from the Console by clicking on the 'Promote' button when looking at the Release list for an app.  Alternatively you can `convox releases promote <release id>` from the CLI.

## Release Descriptions

When creating a build using the CLI, you can use the `--description`/`-d` flag to set your own description against the release.  If you don't, Convox will use your git commit message (assuming a clean working tree) to populate the description:

    build 7329df90c1 this is my commit message

If you create a new release by updating environment variables, the release description will give a summary of the changes:

    env add:FOO change:BAR remove:BAZ
  