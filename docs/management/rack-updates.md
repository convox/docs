---
title: "Rack Updates"
---

As Convox is in active development, new Rack releases are frequently available. There are two ways to update a Rack: [via Console](#via-console) and [via the CLI](#via-the-cli). Also keep in mind that some Rack releases are [required](#required-updates), i.e. cannot be skipped even if there is a newer release available.

## Updating via Console

After logging into [the Convox Console](https://console.convox.com/), you can see your list of Racks in the Racks section:

![List of Racks](/assets/images/docs/what-is-a-rack/list-of-racks.png)

If any of the Racks in the list are outdated, you can click the `Update` button to update to the next version. If the Rack is very outdated, you may need to update more than once (see [Required Updates](#required-updates) below).

## Updating via the CLI

You can update via the CLI by running `convox rack update [--rack <rack name>] [version]`:

```
$ convox rack update --rack dev
Updating to 20161111013816... UPDATING
```

When you update your rack, you will likely need to redeploy applications so that they can pick up the changes and start logging to CloudWatch directly.

## Required Updates

Occasionally, enhancements to Rack will replace an old component with a new one, but in a two-step process to ensure a seamless and stable migration. The mechanism we use to guarantee that a Rack completes these migrations properly is the required update.

Let's say your Rack is running on release version A, and you want to update it to version D, but version B is a required release. Running `convox rack update` would first update your Rack to B to ensure its new component is introduced and put in use. Afterwards, with the first step of the migration complete, you could update the Rack to version C or D, which could safely remove the old component.

In practice, if you notice that a `convox rack update` doesn't update your Rack to the latest release, you've encountered a required release. Running `convox rack update` again will update your Rack to the latest release, unless another required release is encountered first.
