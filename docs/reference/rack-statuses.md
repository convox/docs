---
title: "Rack Statuses"
---

As you interact with your Rack, you will notice a variety of statuses displayed in the Console web interface and in the output of CLI commands like `convox rack`. Here's what they mean.

## converging

Your Rack is running, but it is rearranging itself. For example, you might see the _converging_ status during a routine instance replacement, or when deploying an app to the Rack. If it seems like your Rack is _converging_ for a long time, you can run `convox apps` to see if any of your applications are still updating.

## deleting

When you [uninstall a Rack](/docs/uninstalling-convox/), you'll see the _deleting_ status as the script deletes all of the underlying AWS resources.

This can correspond to `DELETE_IN_PROGRESS` in the AWS CloudFormation console.

## installing

After [installing a Rack](/docs/installing-a-rack/) from the Console web interface or the CLI via `convox rack install`, the Rack will remain in the _installing_ status while resources are created and configured in your AWS account.

## rollback

If a Rack update fails to complete, or is canceled, the Rack will display a status of _rollback_ while it reverts to the last known good state.

This can correspond to `UPDATE_ROLLBACK_IN_PROGRESS`, `UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS` or `ROLLBACK_IN_PROGRESS` in the AWS CloudFormation console.

## running

Most of the time, your Rack should display the _running_ status and be fully operational.

This can correspond to `CREATE_COMPLETE`, `DELETE_FAILED`, `UPDATE_COMPLETE` or `UPDATE_ROLLBACK_COMPLETE` in the AWS CloudFormation console.

## unreachable

When Console or your CLI are unable to communicate with a Rack, that Rack will have a status of _unreachable_.

## updating

When you [update a Rack](/docs/rack-updates/), the Rack will display a status of _updating_ until the update completes and transitions to a _running_ status or fails and transitions to a _rollback_ status.

This can correspond to `UPDATE_IN_PROGRESS` or `UPDATE_COMPLETE_CLEANUP_IN_PROGRESS` in the AWS CloudFormation console.

## See also

* [Rolling Updates](/docs/rolling-updates/)
