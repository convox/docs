---
title: "MySQL"
---

## Resource Creation

You can create MySQL databases using the `convox resources create` command:

    $ convox resources create mysql
    Creating mysql-3785 (mysql)... CREATING

This will provision MySQL database on the Amazon RDS service. Creation can take up to 15 minutes. To check the status use `convox resources info`.

### Additional Options

See `convox resources options mysql` for a list of available options.

## Resource Information

To see relevant info about the database, use the `convox resources info` command:

    $ convox resources info mysql-3785
    Name    mysql-3785
    Status  running
    URL     mysql://app:password@my1.cbm068zjzjcr.us-east-1.rds.amazonaws.com:3306/app

## Resource Linking

You can add this URL to any application with `convox env set`:

    $ convox env set 'DATABASE_URL=mysql://mysql:password@my1.cbm068zjzjcr.us-east-1.rds.amazonaws.com:3306/app' --app example-app

## Resource Update

You can change options with `convox resources update`:

    $ convox resources update mysql-3785 MultiAZ=true
    Updating mysql-3785... UPDATING

## Resource Deletion

To delete the database, use the `convox resources delete` command:

    $ convox resources delete mysql-3785
    Deleting mysql-3785... DELETING

Deleting the database will take several minutes.

<div class="block-callout block-show-callout type-warning" markdown="1">
This action will cause an unrecoverable loss of data.
</div>
