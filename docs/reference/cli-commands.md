---
title: "CLI Commands"
---

## Commands available

Most commands can also be viewed with this command:

    convox help

Or further assistance with the help flag:

    convox examplecommand -h

Commands valid as of version `20190801142802`

### `api get`

query the rack api

### `apps`

list apps

### ``apps cancel``

cancel an app update

### `apps create`

create an app

### `apps delete`

delete an app

### `apps export`

export an app

### `apps import`

import an app

### `apps info`

get information about an app

### `apps lock`

enable termination protection

### `apps params`

display app parameters

### `apps params set`

set app parameters

### `apps unlock`

disable termination protection

### `apps wait`

wait for an app to finish updating

### `build`

create a build

### `builds`

list builds

### `builds export`

export a build

### `builds import`

import a build

### `builds info`

get information about a build

### `builds logs`

get logs for a build

### `certs`

list certificates

### `certs delete`

delete a certificate

### `certs generate`

generate a certificate

### `certs import`

import a certificate

### `cp`

copy files

### `deploy`

create and promote a build

### `env`

list env vars

### `env edit`

edit env interactively

### `env get`

get an env var

### `env set`

set env var(s)

### `env unset`

unset env var(s)

### `exec`

execute a command in a running process

### `instances`

list instances

### `instances keyroll`

roll ssh key on instances

### `instances ssh`

run a shell on an instance

### `instances terminate`

terminate an instance

### `logs`

get logs for an app

### `proxy`

proxy a connection inside the rack

### `ps`

list app processes

### `ps info`

get information about a process

### `ps stop`

stop a process

### `rack`

get information about the rack

### `rack install`

install a rack (you can configure the install passing [Rack Parameters](https://docsv2.convox.com/reference/rack-parameters) i.e. `convox rack install aws -n sandbox Private=Yes`)

### `rack logs`

get logs for the rack

### `rack params`

display rack parameters

### `rack params set`

set rack parameters

### `rack ps`

list rack processes
### `rack releases`

list rack version history

### `rack runtimes`

List of attachable runtime integrations

```html
convox rack runtimes
```

### `rack runtime attach`

Attach runtime integration to the rack

```html
convox rack runtime attach <runtime_id>
```

### `rack scale`

scale the rack
### `rack uninstall`

uninstall a rack
### `rack update`

update the rack
### `rack wait`

wait for rack to finish updating

### `racks`

list available racks

### `registries`

list private registries

### `registries add`

add a private registry

### `registries remove`

remove private registry

### `releases`

list releases for an app

### `releases info`

get information about a release

### `releases manifest`

get manifest for a release

### `releases promote`

promote a release

### `releases rollback`

copy an old release forward and promote it

### `resources`

list resources

### `resources info`

get information about a resource

### `resources proxy`

proxy a local port to a resource

### `resources url`

get url for a resource

### `restart`

restart all processes for an app on the rack

### `rack resources`

list resources

### `rack resources create`

create a resource

### `rack resources delete`

delete a resource

### `rack resources info`

get information about a resource

### `rack resources link`

link a resource to an app

### `rack resources options`

list options for a resource type

### `rack resources proxy`

proxy a local port to a rack resource

### `rack resources types`

list resource types

### `rack resources update`

update resource options

### `rack resources unlink`

unlink a resource from an app

### `rack resources url`

get url for a resource

### `rack access credential`

Generates rack access credential

> supported version: >= 20230704162933

* Usage

```html
convox rack access --role [role] --duration-in-hour [duration]
```

flags:
  - `role`: Access role for the credential. Allowed roles are: `read` or `write`

  - `duration-in-hour`: TTL for the credential.

* Examples

```html
$ convox rack access --role read --duration-in-hour 1

RACK_URL=https://...

$ export RACK_URL=https://...

$ convox rack

Name      v3-rack
Provider  aws
Router    router.convox
Status    running
Version   ...
```

### `rack access key rotation`

Rotates the rack access key that is used for rack access credential. It will invalidate previous all the credential generated from ` convox rack access --role [role] --duration-in-hour [duration]`.

> supported version: >= 20230704162933

* Usage

```html
convox rack access key rotate
```

* Examples

```html
$ convox rack access key rotate

OK
```

### `run`

execute a command in a new process

### `scale`

scale a service

### `services`

list services for an app

### `services restart`

restart a particular service within an app

### `ssl`

list certificate associates for an app

### `ssl update`

update certificate for an app

### `start`

start an application for local development

### `test`

run tests

### `version`

display version information
