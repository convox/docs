---
title: "About resources"
---

## What is a resource?

_Resources_ are dependencies of an app that communicate with it over the network and run software outside the scope of your team's focus.

## Resource Creation

You can create resources using the `convox resources create` command.

This will provision the appropriate AWS service and associate it with your active Convox Rack.

Creation can take up to 15 minutes. To check the status, use `convox resources info`. You can also view the resource in the appropriate section of your AWS console (RDS for postgres and mysql, etc).

For advanced creation options, see the documentation for the specific resource type in the sidebar on the left.


## Resource Status and Information

To see relevant info about a resource, use the `convox resources info` command:

    $ convox resources info memcached-5864
    Name    memcached-5864
    Status  running
    Exports
      URL: dev-ca-m3z4ik3n7bej.77prpt.cfg.use1.cache.amazonaws.com:11211

## Accessing the resource 

### Via AWS console

You can view AWS resources in your AWS console, where they will have their own CloudFormation stack. For example, mysql resources are visible in your [RDS Console](https://console.aws.amazon.com/rds/home).

If you don't see what you expect, check that the correct geographic region is active.

### Via proxy

Resources are generally configured so they are not accessible from the public internet (when possible).  You can get proxy access with the `convox resources proxy` command, e.g.:

```
$ convox resources proxy mysql-4624
```

For details, see [convox resources proxy](/docs/remote-resources).

### Resource credentials

For some types of resources, the username and password are contained in the resource URL, which can be retrieved by running `convox resources info <resource name>`.
For example, below we can see the RDS credentials embedded in a `mysql` resource URL:

```
URL: mysql://app:672ffd60a4eb323gf238ff28508f95@dev-mysql-4624.caujtavicybf.us-east-1.rds.amazonaws.com:3306/app
                 ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^                ^^^^^^^^^^^^
```

## Connecting the resource to your app


### Via environment variables

You can add this URL as an environment variable to any application with `convox env set`:

    $ convox env set MEMCACHED_URL='dev-ca-m3z4ik3n7bej.77prpt.cfg.use1.cache.amazonaws.com:11211' --app example-app

### Via `convox resources link`

To forward logs from an application to a syslog forwarder use `convox resources link`:

    $ convox resources link syslog-3785 --app example-app
    Linked syslog-3786 to example-app

Note: This is currently only supported with `syslog` resources. To link other resource types, use `convox env set` as described above.


## Resource Updates

A few types of resources can be updated. For instance, to modify the URL of a syslog resource that has `example.com` as the URL, run `convox resources update <resource name> Option=newvalue`:

```
$ convox resources update syslog-3165 Url=example.net
Updating syslog-3165 (url="example.com")...UPDATING

$ convox resources info syslog-3165
Name    syslog-3165
Status  updating
Exports
  URL: example.net
```

## Scaling down unused services

If you've replaced a service with a resource, you'll want to be sure to [scale the unused service down to `-1`](/docs/scaling/#scaling-down-unused-services) so it doesn't create unnecessary AWS resources in production.

## Resource Deletion

To delete a resource, use the `convox resources delete` command:

    $ convox resources delete memcached-5864
    Deleting memcached-5864... DELETING

Deleting the resource will take several minutes.

<div class="block-callout block-show-callout type-warning" markdown="1">
This action will cause an unrecoverable loss of data.
</div>

