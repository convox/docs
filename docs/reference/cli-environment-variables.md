---
title: "CLI Environment Variables"
---

<div class="block-callout block-show-callout type-info" markdown="1">
This page is about CLI-specific environment variables. For information about Docker- and app-related environments, see the [Environment documentation](/docs/environment).
</div>

Convox pays attention to the following environment variables.

## `AWS_REGION`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN`

AWS credentials used for `convox rack install aws`.

## `CONVOX_CONFIG`

Location of Convox config directory.

## `CONVOX_RACK`

Defines which Rack `convox` commands should be applied to. Overrides `./.convox/rack` and `~/.convox/rack`. Can be overridden by `--rack` flag.

## `CONVOX_HOST`

Convox Rack or Console endpoint. If not present, `console.convox.com` is assumed. Overrides [`~/.convox/host`](/docs/cli-config-files/#convoxhost).

## `CONVOX_PASSWORD`

Your Convox account password*; used when running [`convox login`](/docs/login-and-authentication/).

_* or Rack API key, if you're logging directly into a Rack._

## `RACK`

Active Rack name. Takes precedence over the [`.convox/rack`](/docs/cli-config-files/#convoxrack-1) repository setting, but is overridden by the `--rack` flag. For details, see [Switching between Racks](/docs/cli#switching-between-racks).

## `VERSION`

Custom Rack version. Used when [installing a new Rack](/docs/installing-a-rack). Defaults to `latest`.
