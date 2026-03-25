---
title: "Rolling Back"
description: "How to roll back Apps, Racks, and the CLI to previous versions."
---

# Rolling Back

Convox keeps history for your Apps and for Rack and CLI releases. You can roll back to previous versions of all of these.

Releases contain both Build artifacts and environment variable state, so rolling back a Release also reverts environment changes.

In the Console, when viewing the list of Releases for an App, click the **Rollback** button on any Release older than the currently active one. Convox creates a copy of that Release, places it at the top of your Release list, and promotes it. This triggers a [rolling update](/deployment/rolling-updates) to the copied Release.

From the CLI, run `convox releases rollback <release id>` to start the rollback process.

Previous Releases are not lost during a rollback. They remain in your Release list, allowing you to roll back to any other Release at any time.

## Apps

To roll back an App, promote a previous Release.

```bash
$ convox releases
ID           CREATED      BUILD        STATUS
RPQXVNXFGXU  1 week ago   BWIWLFJJVIU  active
RHKOIXXVDMJ  2 weeks ago  BAQEPGSROOD
RAWLSJIXNYB  2 weeks ago  BKNVBIUYHIS
RREZLTEMWCP  2 weeks ago  BBDAMGDWUOM
RDSKWHNEQMK  1 month ago  BIVOMOGBRLD

$ convox releases rollback RHKOIXXVDMJ
Rolling back to RHKOIXXVDMJ... OK, RAXHPJHSWZO
Promoting RAXHPJHSWZO...
```

## Rack

To roll back your Rack, update to a specific previous version.

```bash
$ convox rack releases
VERSION                  UPDATED       STATUS
20250301151815           6 days ago    active
20250227211125           6 days ago
20250220212601           1 week ago
20250213200033           2 weeks ago
20250206151159           3 weeks ago

$ convox rack update 20250213200033
Name     dev
Status   updating
Version  20250301151815

Updating to version: 20250213200033
```

## CLI

You can download previous CLI releases from the Convox download server. To revert to a previous CLI version, download the binary for your platform and copy it into your executable path (typically `/usr/local/bin`). See [CLI Installation](/introduction/installation) for details.

## See Also

- [Releases](/deployment/releases)
- [Rolling Updates](/deployment/rolling-updates)
- [Health Checks](/application/health-checks)
- [Rack Updates](/management/rack-updates)
