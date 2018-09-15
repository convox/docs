---
title: "SQS"
---

## Resource Creation

You can create SQS queues using the `convox resources create` command:

    $ convox resources create sqs
    Creating sqs-3785 (sqs)... CREATING

This will provision an SQS queue. Creation will take a few moments. To check the status use `convox resources info`.

### Additional Options

See `convox resources options sqs` for a list of available options.

## Resource Information

To see relevant info about the queue, use the `convox resources info` command:

    $ convox resources info sqs-3785
    Name    sqs-3785
    Status  running
    URL     sqs://ACCESS:SECRET@sqs.us-east-1.amazonaws.com/ACCOUNT/QUEUE

## Resource Linking

You can add this URL as an environment variable to any application with `convox env set`:

    $ convox env set SQS_URL='sqs://ACCESS:SECRET@sqs.us-east-1.amazonaws.com/ACCOUNT/QUEUE' --app example-app

## Resource Update

You can change options with `convox resources update`:

    $ convox resources update sqs-3785 VisibilityTimeout=10
    Updating sqs-3785... UPDATING

## Resource Deletion

To delete the queue, use the `convox resources delete` command:

    $ convox resources delete sqs-3785
    Deleting sqs-3785... DELETING
