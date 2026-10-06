---
title: "Notifications"
description: "Receive Slack notifications for Rack events including app deployments, resource changes, and infrastructure updates."
---

# Notifications

Console users can get notification for common Rack events. Below you can find a list of the types of notifications you will receive.

## Notification Types

### [*rack*] Created app *example*

A convox app has been created, as with `convox apps create` and is ready to accept deployments.

### [*rack*] Build `BNJFGEQEXOK` failed for app *example*

An app build failed. Run `convox builds logs <id>` to view build logs.

### [*rack*] Created release `RMDKLNZIACD` for app *example*

A new release has been created and is ready to be deployed. Releases are created when new builds complete or when the App’s environment variables are changed. You can promote a new release with the `convox releases promote` command.

### [*rack*] Started promoting release `RMDKLNZIACD` for app *example*

A release promotion has started, as with `convox releases promote` or `convox deploy`.

### [*rack*] Promoted release `RMDKLNZIACD` for app *example*

The App's CloudFormation update for the release has completed and the release is live. Sent for Generation 2 Apps only.

### [*rack*] Promoting release `RMDKLNZIACD` for app *example* failed

The App's CloudFormation update for the release failed and was rolled back, or the rollback failed. Sent for Generation 2 Apps only.

### [*rack*] Created postgres resource *pg1*

A Rack Resource (such as postgres, mysql, redis, etc) is being created, as with `convox rack resources create`. The notification is sent when the create starts, and is not sent for Resources defined in `convox.yml`. It tells you the resource type and resource name, respectively.

### [*rack*] Deleted postgres resource *pg1*

A resource has been deleted by Convox, as with `convox rack resources delete`.

### [*rack*] Updating rack to: version *20160405223647*

A Rack update has been initiated. Rack updates can take from a few seconds to several minutes to complete, depending on whether they require the backing EC2 instances to be restarted. Most updates do not require instance restarts.

### [*rack*] Updating rack to: count *3*

A request has been received to alter the number of instances in your Rack’s cluster. In some cases this can require processes to be re-launched and can take a few minutes to complete.

### [*rack*] Updating rack to: instance type *t3.medium*

A request has been received to alter the type of instances in your Rack’s cluster. This will require processes to be re-launched and can take a few minutes to complete.

## See Also

- [Integrations](/console/integrations)
- [Workflows](/console/workflows)
- [Rack Updates](/management/rack-updates)
- [Builds](/deployment/builds)
