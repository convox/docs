---
title: "SNS"
---

## Resource Creation

You can create SNS topics using the `convox resources create` command:

    $ convox resources create sns
    Creating sns-3331 (sns)... CREATING

This will provision an SNS topic. Creation will take a few moments. To check the status use `convox resources info`.

### Additional Options

See `convox resources options sns` for a list of available options.

## Resource Information

To see relevant info about the SNS topic, use the `convox resources info` command:

    $ convox resources info sns-3331
    Name    sns-3331
    Status  running
    URL     sns://ACCESS:SECRET@arn:aws:sns:us-east-1:235997312769:development-sns-3331

## Resource Linking

You can add this URL as an environment variable to any application with `convox env set`:

    $ convox env set SNS_URL='sns://ACCESS:SECRET@arn:aws:sns:us-east-1:235997312769:development-sns-3331' --app example-app

## Resource Update

You can change options with `convox resources update`:

    $ convox resources update sns-3331 Queue=arn:...
    Updating sns-3331... UPDATING

## Resource Deletion

To delete the SNS topic, use the `convox resources delete` command:

    $ convox resources delete sns-3331
    Deleting sns-3331... DELETING
