---
title: "CLI"
---

## Install the CLI

To install the `convox` CLI tool, see [Installing the Convox CLI](/docs/installation/).

## Convox Update

You can easily update the CLI to get bugfixes and features:

    $ convox update
    OK, 20160617165137

    $ convox --version
    client: 20160617165137
    server: 20160615213630 (console.convox.com)

## Overriding App Defaults

You can set a default app name in an app directory. Instead of inferring the default app name from the current directory name, CLI commands will default to the app name specified in `.convox/app`:

    $ cd myapp
    $ mkdir .convox
    $ echo myapp-staging > .convox/app
    $ convox apps info
    Name       myapp-staging
    ...

## Switching Between Racks

The Convox Console makes it easy to install and manage multiple Racks, like one for development, staging and production. The `convox` CLI offers a few strategies to switch between these environments.

#### Switch Command

You can use `convox racks` and `convox switch` to select a Rack that your computer will operate against by default.

    $ convox racks
    myorg/staging                    running
    myorg/production                 running
    personal/staging                 running

    $ convox switch production
    Switched to myorg/production

    $ convox switch staging
    ERROR: You have access to multiple racks with that name, try one of the following:
    convox/staging
    personal/staging

    $ convox switch personal/staging
    Switched to personal/staging

    $ convox switch
    personal/staging

#### Rack Flag

You can specify a specific Rack on a per-command basis with the `--rack` flag:

    $ convox switch
    myorg/staging

    $ convox apps
    APP           STATUS
    api-staging   running
    docs-staging  running

    $ convox apps --rack personal/staging
    APP    STATUS
    httpd  running
    rails  running

#### CONVOX_RACK Environment Variable

You can specify a specific Rack for a new terminal session with the CONVOX_RACK environment variable:

    $ convox switch
    myorg/staging

    $ export CONVOX_RACK=personal/staging

    $ convox switch
    personal/staging

    $ convox apps
    APP    STATUS
    httpd  running
    rails  running

#### App Repository Setting

You can set a default Rack name in an app's repository. This will take precedence over the `convox switch` setting:

    $ cd rails
    $ mkdir .convox
    $ echo personal/staging > .convox/rack
    $ convox deploy

#### Client Rack Header

All of the above tools control what `Rack` header is sent in API commands. You can also control this explicitly when using the API:

    $ curl -u :$API_KEY -H "Rack: personal/staging" https://console.convox.com/apps
    [
      {
        "name": "httpd",
        "release": "RLBTJENIMQI",
        "status": "running"
      },
      {
        "name": "rails",
        "release": "RGPIVDMTGVH",
        "status": "running"
      }
    ]

#### Precedence

The order of precedence is:

* --rack flag
* CONVOX_RACK environment variable
* ./convox/rack app repository setting
* `convox switch` default setting in ~/.convox/rack

When you want to manage multiple racks in multiple terminals you should use `CONVOX_RACK`.

When you want to pin an app to a specific Rack you should use `./convox/rack` which can only be overridden by an explicit `--rack` flag.

## Shell Autocomplete Support

1. Download the [urfave/cli](https://github.com/urfave/cli) completion script.

       $ curl -o ~/.convox/completion.bash \
         https://raw.githubusercontent.com/urfave/cli/master/autocomplete/bash_autocomplete

1. Source the completion script in your shell initialization file.

    a. For **bash**, add the following line to your `.bash_profile`:

       source ~/.convox/completion.bash

    b. For **zsh**, add the following to your `.zshrc`:

       autoload -U compinit && compinit
       autoload -U bashcompinit && bashcompinit
       source ~/.convox/completion.bash

      Users of zsh may need to change the first line of the completion script to:

       : ${PROG:=$(basename $0)}

## Active Rack Command Prompt Helper

In your `.profile`, define `__convox_switch` as a function:

    __convox_switch() {
      [ -e ~/.convox/rack ] && convox switch || echo unknown;
    }

Then include `$(__convox_switch)` in your prompt (`PS1`):

    export PS1="\h:\W \u (\$(__convox_switch))\$(__git_ps1)$ "

or

    PS1+="[$(__convox_switch)] "
    export PS1

If using `zsh` and PowerLevel9K you can use a similar function to define a custom command as per: https://github.com/bhilburn/powerlevel9k#custom_command

    convox_rack(){
      [ -e ~/.convox/rack ] && convox switch || echo unknown;
    }

    POWERLEVEL9K_CUSTOM_CONVOX_RACK="convox_rack"
    POWERLEVEL9K_LEFT_PROMPT_ELEMENTS=(os_icon time root_indicator context dir custom_convox_rack rbenv vcs)

