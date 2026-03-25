---
title: "Rack Updates"
description: "Update your Convox Rack via the CLI and understand required update releases."
---

# Rack Updates

As Convox is in active development, new Rack releases are frequently available. Rack updates are typically applied automatically but you can manually trigger an update [via the CLI](#updating-via-the-cli). Some Rack releases are [required](#required-updates) and cannot be skipped even if there is a newer release available.

## Updating via the CLI

You can update via the CLI by running `convox rack update [--rack <rack name>] [version]`:

```bash
$ convox rack update --rack dev
Updating to 20161111013816... UPDATING
```

> **Note:** Updating your Rack will not cause any application or resource downtime, or require App redeployment (unless explicitly noted to enable new features). Rack updates generally update the Convox software stack running on the Rack in the background. Sometimes they will update the underlying AMIs the Rack runs on, which will initiate a rolling update, automatically migrating your Apps to the new instances without causing any downtime.

## Required Updates

Occasionally, enhancements to Rack will replace an old component with a new one, but in a two-step process to ensure a seamless and stable migration. The mechanism we use to guarantee that a Rack completes these migrations properly is the required update.

Let's say your Rack is running on release version A, and you want to update it to version D, but version B is a required release. Running `convox rack update` would first update your Rack to B to ensure its new component is introduced and put in use. Afterwards, with the first step of the migration complete, you could update the Rack to version C or D, which could safely remove the old component.

In practice, if you notice that a `convox rack update` doesn't update your Rack to the latest release, you've encountered a required release. Running `convox rack update` again will update your Rack to the latest release, unless another required release is encountered first.

## See Also

- [Core Concepts: Rack](/core-concepts/rack)
- [Import a Rack](/management/import-rack)
- [Reference: Rack Parameters](/reference/rack-parameters)
- [Reference: Rack Statuses](/reference/rack-statuses)
