---
title: "App Statuses"
---

As you interact with your Convox apps, you will notice a variety of statuses displayed in the Console web interface and in the output of CLI commands like `convox apps info`. Here's what they mean.

## creating

Expect to see this status after running `convox apps create <app-name>`, before the app boots and transitions to the _running_ status.

## deleting

When you delete an app, you'll see this status until the underlying AWS resources have been deleted.

## rollback

If a deploy fails to complete, or is canceled, the app will display a status of _rollback_ while it reverts to the previous release.

## running

Most of the time, your apps should display the _running_ status and be fully operational.

## updating

When you deploy an app, you'll see a status of _updating_ while the new release is being promoted. Once promotion completes, the status will transition to _running_.
