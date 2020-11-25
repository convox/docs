---
title: "Private Registries"
---

Convox apps are composed of one or more processes that run inside Docker containers.

In most cases, the Docker images that make up your app are either public images pulled from [Docker Hub](https://hub.docker.com/) or custom images that are built from your codebase. In some cases, however, you might want to pull an image from a private registry.

For example, you might have a private fork of a popular image – like [postgres](https://hub.docker.com/_/postgres/) – in your Docker Hub account. You can specify this image in `convox.yml` so that your app will use it:

    database:
      image: yourname/postgres

But when you try to deploy, Convox will return an error:

    $ convox deploy
    Deploying yourapp
    Creating tarball... OK
    Uploading... OK
    RUNNING: docker pull yourname/postgres
    Pulling repository docker.io/yourname/postgres
    time="2016-01-29T21:22:15Z" level=fatal msg="Error: image yourname/postgres:latest not found"
    ERROR: exit status 1

## Adding a registry

In order to deploy from a private registry you will need to add credentials via the `convox registries add` command:

Continuing with our Docker Hub example, the command would be:

    $ convox registries add index.docker.io/v1/ username password
    Adding registry... OK

Once the registry has been added, you can pull private images:

    $ convox deploy
    Deploying test
    Creating tarball... OK
    Uploading... OK
    RUNNING: docker pull yourname/postgres
    latest: Pulling from yourname/postgres

<div class="block-callout block-show-callout type-info" markdown="1">
Note that you do not need to include the `https://` protocol as part of the registry address.  Convox will add this for you automatically.
</div>

## Removing a registry

To remove private registry info, use the `convox registries remove` command. To remove Docker Hub in our example the command would be:

    $ convox registries remove index.docker.io/v1/
    Removing registry... OK

## Adding an Amazon EC2 Container Registry (ECR) from a different account

Convox is already configured to use ECR in its own AWS account. However, you may also want to pull and build from images stored in the ECR of a different AWS account:

    database:
      image: 901416387788.dkr.ecr.us-east-1.amazonaws.com/postgres

Since [ECR authorization tokens expire every 12 hours](http://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth), you must give Convox IAM access keys that have permission to generate ECR tokens and pull images:

    $ aws iam create-user --user-name ECRReadOnly
    {
        "User": {
            "UserName": "ECRReadOnly",
            "Path": "/",
            "CreateDate": "2016-02-23T00:52:05.930Z",
            "UserId": "AIDAJ6JPEYYKRY5PEVSU6",
            "Arn": "arn:aws:iam::901416387788:user/ECRReadOnly"
        }
    }

    $ aws iam attach-user-policy --user-name ECRReadOnly --policy-arn arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly

    $ aws iam create-access-key --user-name ECRReadOnly
    {
        "AccessKey": {
            "UserName": "ECRReadOnly",
            "Status": "Active",
            "CreateDate": "2016-02-23T00:54:32.475Z",
            "SecretAccessKey": "2yf2HqhykiGHNKlwbvuS66WOBgSTefWXClOQIy0f",
            "AccessKeyId": "AKIAJ7GE3UMOANV37YNQ"
        }
    }

Now pass the access key info to `convox registries add`:

    $ convox registries add 901416387788.dkr.ecr.us-east-1.amazonaws.com AKIAJ7GE3UMOANV37YNQ 2yf2HqhykiGHNKlwbvuS66WOBgSTefWXClOQIy0f
    Adding registry... OK

You can revoke Convox access by deleting the IAM user and removing the registry:

    $ aws iam delete-access-key --user-name ECRReadOnly --access-key-id AKIAJ7GE3UMOANV37YNQ
    $ aws iam detach-user-policy --user-name ECRReadOnly --policy-arn arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly
    $ aws iam delete-user --user-name ECRReadOnly
    $ convox registries remove 901416387788.dkr.ecr.us-east-1.amazonaws.com
    Done.

<div class="block-callout block-show-callout type-info" markdown="1">
Local racks do not authenticate against Amazon EC2 Container Registries (ECR) but do respect `docker login` sessions.  You can authenticate against your chosen ECR's with `docker login`/`aws ecr get-login` (see [here](https://aws.amazon.com/blogs/compute/authenticating-amazon-ecr-repositories-for-docker-cli-with-credential-helper/)) before runnning a local build and your remote images will be pulled succesfully.
</div>
