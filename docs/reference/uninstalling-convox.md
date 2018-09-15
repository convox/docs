---
title: "Uninstalling Convox"
---

<div class="alert alert-warning">
This action will cause an unrecoverable loss of Convox-created data and resources.
</div>

At any time you can easily remove all the AWS resources Convox uses for your Services, Apps and Racks. This makes experimenting with Convox very easy.

Uninstall will take approximately 15 minutes to complete.

You can uninstall a Rack by running `convox rack uninstall <provider> <stack-name>`.

<div class="block-callout block-show-callout type-info" markdown="1">
`provider` will be one of `local` or `aws`.
</div>

<div class="block-callout block-show-callout type-info" markdown="1">
`stack-name` will correspond to the name of the Rack as shown in `convox rack`:

```
$ convox rack --rack personal/example
Name     example
Status   running
Version  20161102160040
Region   us-east-1
Count    3
Type     t2.small

```
</div>

<div class="block-callout block-show-callout type-danger" markdown="1">
## Removing a Rack

If you simply want to unlink a Rack from Convox without deleting the associated resources, you can do so via the [web console](https://console.convox.com/). Click on the Rack name, then navigate to the `Settings` tab and click `Remove Rack`.
</div>
