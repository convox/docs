---
title: "Notifications"
---

Console users can get notification for common Rack events. Below you can find a list of the types of notifications you will receive.

#### [*rack*] Created app *example*

A convox app has been created, as with `convox apps create` and is ready to accept deployments.

#### [*rack*] Deleted app *example*

A convox app has been deleted, as with `convox apps delete`.

#### [*rack*] Build `BNJFGEQEXOK` failed for app *example*

An app build failed. Run `convox builds info <id>` to view build logs.

#### [*rack*] Created release `RMDKLNZIACD` for app *example*

A new release has been created and is ready to be deployed. Releases are created when new builds complete or when the app’s environment variables are changed. You can promote a new release with the `convox releases promote` command.

#### [*rack*] Finished rolling deploy of release `RMDKLNZIACD` for app *example*

A deployment is totally finished. All of the new processes have been booted and all of the old processes have been stopped.

#### [*rack*] Promoted release `RMDKLNZIACD` for app *example*

A release has been promoted to be the live version on the application. This does not mean the deployment is totally complete. See the “Finished rolling deploy” notification above.

#### [*rack*] Scaled release `RMDKLNZIACD` for app *example*

A `convox scale` command has been received, instructing the Rack to run a different number of copies for a specific process.

#### [*rack*] Created postgres resource *pg1*

A resource (such as postgres, mysql, redis, etc) has been created by Convox, as with `convox resources create`. This notification tells you the resource type and resource name, respectively.

#### [*rack*] Deleted postgres resource *pg1*

A resource has been deleted by Convox, as with `convox resources delete`.

#### [*rack*] Updating rack to: version *20160405223647*

A rack update has been initiated. Rack updates can take from a few seconds to several minutes to complete, depending on whether they require the backing EC2 instances to be restarted. Most updates do not require instance restarts.

#### [*rack*] Updating rack to: count *3*

A request has been received to alter the number of instances in your Rack’s cluster. In some cases this can require processes to be re-launched and can take a few minutes to complete.

#### [*rack*] Updating rack to: instance type *db.t2.medium*

A request has been received to alter the type of instances in your Rack’s cluster. This will require processes to be re-launched and can take a few minutes to complete.
