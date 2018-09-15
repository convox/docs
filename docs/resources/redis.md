---
title: "Redis"
---

## Resource Creation

You can create redis databases using the `convox resources create` command:

    $ convox resources create redis
    Creating redis-3785 (redis)... CREATING

This will provision Redis on the Amazon ElastiCache. Creation can take a few minutes. To check the status use `convox resources info`.

### Additional Options

See `convox resources options redis` for a list of available options.

## Resource Information

To see relevant info about the database, use the `convox resources info` command:

    $ convox resources info redis-3785
    Name    redis-3785
    Status  running
    URL     redis://:password@atb1alu32d6lfy19.c63i2h.ng.0001.use1.cache.amazonaws.com:6379/0

## Resource Linking

You can add this URL as an environment variable to any application with `convox env set`:

    $ convox env set REDIS_URL='redis://:password@atb1alu32d6lfy19.c63i2h.ng.0001.use1.cache.amazonaws.com:6379/0' --app example-app

## Resource Update

You can change options with `convox resources update`:

    $ convox resources update redis-3785 InstanceType=cache.t2.small
    Updating redis-3785... UPDATING

## Resource Deletion

To delete the database, use the `convox resources delete` command:

    $ convox resources delete redis-3785
    Deleting redis-3785... DELETING

Deleting the database will take several minutes.

<div class="block-callout block-show-callout type-warning" markdown="1">
This action will cause an unrecoverable loss of data.
</div>
