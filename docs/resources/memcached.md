---
title: "Memcached"
---

## Resource Creation

You can create Memcached instances using the `convox resources create` command:

    $ convox resources create memcached
    Creating memcached-5864 (memcached)... CREATING

This will provision Memcached on Amazon ElastiCache. Creation can take a few minutes. To check the status use `convox resources info <resource-name>`.

### Additional Options

See `convox resources options memcached` for a list of available options.

## Resource Information

To see relevant info about the Memcached instance, use the `convox resources info` command:

    $ convox resources info memcached-5864
    Name    memcached-5864
    Status  running
    URL     memcached://:password@dev-ca-m3z4ik3n7bej.77prpt.cfg.use1.cache.amazonaws.com:11211

## Resource Linking

You can add this URL as an environment variable to any application with `convox env set`:

    $ convox env set MEMCACHED_URL='memcached://:password@dev-ca-m3z4ik3n7bej.77prpt.cfg.use1.cache.amazonaws.com:11211' --app example-app

## Resource Update

You can change options with `convox resources update`:

    $ convox resources update memcached-5864 InstanceType=cache.t2.small
    Updating memcached-5864... UPDATING

## Resource Deletion

To delete the resource, use the `convox resources delete` command:

    $ convox resources delete memcached-5864
    Deleting memcached-5864... DELETING

Deleting the resource will take several minutes.

<div class="block-callout block-show-callout type-warning" markdown="1">
This action will cause an unrecoverable loss of data.
</div>
