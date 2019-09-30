---
title: "Releases"
order: 300
---

Every time you deploy an application or change its environment a new Release is created. You can promote any Release for an application to deploy it. Promoting a Release will cause a [rolling update](/docs/rolling-updates) that gracefully deploys the desired Release.

## Release Descriptions

When creating a build using the CLI, you can use the `-d` flag to set your own description against the release.  If you don't, Convox will use your git commit message (assuming a clean working tree) to populate the description:

    build 7329df90c1 this is my commit message

If you create a new release by updating environment variables, the release description will give a summary of the changes:

    env add:FOO change:BAR remove:BAZ
  

## Listing Releases

#### Console

* Click on the **Racks** tab and again on the Rack that contains the application.
* Click on the application name.
* Click on the **Releases** tab.

#### CLI

* Run `convox releases`

## Promoting a Release

#### Console

* Click the **Promote** button next to the Release you would like to promote.

#### CLI

* Run `convox releases promote <id>`