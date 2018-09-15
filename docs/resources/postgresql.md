---
title: "PostgreSQL"
---

## Resource Creation

You can create PostgreSQL databases using the `convox resources create` command:

    $ convox resources create postgres
    Creating postgres-3785 (postgres)... CREATING

This will provision postgres database on the Amazon RDS service. Creation can take up to 15 minutes. To check the status use `convox resources info`.

### Additional Options

See `convox resources options postgresql` for a list of available options.

## Resource Information

To see relevant info about the database, use the `convox resources info` command:

    $ convox resources info postgres-3785
    Name    postgres-3785
    Status  running
    URL     postgres://postgres:i3tNTHpZ8wmCn88nvN2c@dev-postgres-3785.cbm068zjzjcr.us-east-1.rds.amazonaws.com:5432/app

## Resource Linking

You can add this URL to any application with `convox env set`:

    $ convox env set 'DATABASE_URL=postgres://postgres:i3tNTHpZ8wmCn88nvN2c@dev-postgres-3785.cbm068zjzjcr.us-east-1.rds.amazonaws.com:5432/app' --app example-app

## Resource Update

You can change options with `convox resources update`:

    $ convox resources update postgres-3785 MultiAZ=true
    Updating postgres-3785... UPDATING

## Resource Deletion

To delete the database, use the `convox resources delete` command:

    $ convox resources delete postgres-3785
    Deleting postgres-3785... DELETING

Deleting the database will take several minutes.

<div class="block-callout block-show-callout type-warning" markdown="1">
This action will cause an unrecoverable loss of data.
</div>
