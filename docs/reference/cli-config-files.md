---
title: "CLI Configuration Files"
---

This page describes a number of files and directories which can exist in `~/.convox/` and/or `./.convox/`, and which are taken into account by the Convox CLI.

Configuration is evaluated in the following order:

1. Environment variables (e.g. `$CONVOX_RACK`)
2. Local configuration directory (`./.convox/`)
3. User configuration directory (`~/.convox/`)

## Global configuration files

### `~/.convox/auth`

This file is written to every time you run `convox login`, whether you're logging into the Convox console or directly into a Rack.

This file contains a JSON struct in the following format:

    $ cat ~/.convox/auth 
    {
      "console.convox.com": "your-convox-console-api-key",
      "your-rack-host.us-east-1.elb.amazonaws.com": "rack-api-key",
    }

Note: When you install a Rack via `convox rack install`, you are logged in automatically to the newly created Rack.

### `~/.convox/host`

This file contains the hostname of the Rack you're currently logged into.

If you're logged into `console.convox.com`, you can `convox switch` between all the Racks you've added or installed via Console.

### `~/.convox/rack`

This file contains the organization and name of the active Rack. Its contents are used as if passed to the `--rack` flag for Convox commands.

It can be overridden by the [`RACK` environment variable](/reference/cli-environment-variables#rack) or with the `--rack` flag.

Under the hood, this file is how `convox rack` determines which Rack you want to know about:


```
$ cat ~/.convox/rack 
personal/dev

$ convox rack
Name     dev
Status   running
Version  20161123204337
Region   us-east-1
Count    3
Type     t2.small
```

If you're logged into console.convox.com, you can change the active Rack by running `convox switch` or by simply overwriting the contents of this file:

```
$ echo 'personal/legit' > ~/.convox/rack 

$ convox rack
Name     legit
Status   running
Version  20161123204337
Region   us-east-1
Count    3
Type     t2.small
```

Running `convox switch` will overwrite this file with the new active Rack's organization/name.

```
$ cat ~/.convox/rack 
personal/legit

$ convox switch personal/dev
Switched to personal/dev

$ cat ~/.convox/rack 
personal/dev
```


## Local configuration files

### `.convox/app`

You can pin a local directory to a specific app by placing a file called `app` containing the app name in the `.convox` directory in the project root.

```
$ echo "myapp" > .convox/app
```

### `.convox/rack`

You can pin a local directory to a specific Rack by placing a file called `rack` containing the Rack name in the `.convox` directory in the project root.

```
$ echo "myorg/staging" > .convox/rack
```

## See also

- [CLI environment variables](/reference/cli-environment-variables)
- [Login and authentication](/reference/login-and-authentication)
