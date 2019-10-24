---
title: "Rolling Back"
---

<div class="block-callout block-show-callout type-info" markdown="1">
Convox keeps history for your Apps and for Rack and CLI releases. It's possible to roll back to previous versions of all of these.
</div>

If you wish to revert to a previous release for whatever reason, Convox makes that simple.  As Releases also contain state about environment variables, you can even revert to previous versions of these with ease!

In the Console, when looking at the list of releases for an app, ones older than the currently active release can be simply rolled back to by clicking the 'Rollback' button.  Convox will then create a copy of the older Release, place it at the top of your Release list, and promote it.  This causes a rolling update to this copied release across your Rack.

From the CLI, you can `convox releases rollback <release id>` to kick off the rollback process.

Previous releases are not lost or forgotten in the rollback process, they will stay in your Release list allowing you to `rollback` to any other release at any time you wish.

## Apps

To roll back an app, promote a previous release.

```
$ convox releases
ID           CREATED      BUILD        STATUS
RPQXVNXFGXU  1 week ago   BWIWLFJJVIU  active
RHKOIXXVDMJ  2 weeks ago  BAQEPGSROOD
RAWLSJIXNYB  2 weeks ago  BKNVBIUYHIS
RREZLTEMWCP  2 weeks ago  BBDAMGDWUOM
RDSKWHNEQMK  1 month ago  BIVOMOGBRLD
RYXPVTEFWKV  1 month ago  BONLNHSTFYQ
RCLSGJFVJJC  1 month ago  BRSRESRRETT
RAYQZFFJYVZ  1 month ago  BNSPUDINUIE
RTAITJQWMGG  1 month ago  BVZFLIJCRTV
RPNUPSAXEFG  1 month ago  BRTFWFZLPFZ
RACMHKRWPXU  1 month ago  BLFRJGIKRQM

$ convox releases rollback --id RHKOIXXVDMJ
Rolling back to RHKOIXXVDMJ... OK, RAXHPJHSWZO
Promoting RAXHPJHSWZO...
```

## Rack

To roll back your Rack, update to a specific release.

```
$ convox rack releases
VERSION                           UPDATED       STATUS
20160722051815                    6 days ago    active
20160720211125                    6 days ago
20160712212601                    1 week ago
20160706200033                    2 weeks ago
20160701151159                    3 weeks ago
20160627075406-mountable-volumes  1 month ago
20160615213630                    1 month ago
20160613225525                    1 month ago
20160610201355                    1 month ago
20160528002649                    2 months ago
20160526133520                    2 months ago
20160514051157                    2 months ago
20160427051705                    3 months ago
20160420221228                    3 months ago
20160415191544                    3 months ago
20160415002713-proxy-labels       3 months ago
20160413152142                    3 months ago
20160409181028                    3 months ago
20160408144745                    3 months ago
20160408135032                    3 months ago

$ convox rack update 20160706200033
Name     dev
Status   updating
Version  20160722051815
Count    3
Type     t2.medium

Updating to version: 20160706200033
```

## CLI

Convox uses [equinox.io](https://equinox.io/) to distribute and update the CLI. You can find archives of previous CLI releases [here](https://dl.equinox.io/convox/convox/stable/archive). To revert to a previous release, download and unzip the binary for your OS, and copy it into your executable path (typically `/usr/local/bin`).
