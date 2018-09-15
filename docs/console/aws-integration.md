---
title: "AWS Integration"
---

Convox is designed to give your team agility and consistency on top of AWS. In minutes you can install Convox and stand up new apps that follow AWS and DevOps best practices. Every Convox app leverages AWS CloudFormation and the EC2 Container Service for consistent configuration management and automation to reliably deploy, monitor and scale apps.

To this end, Convox requires access and permission to help manage resources in your AWS account.

The best practice to grant Convox this permission into your account is with <a href="http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html">a cross-account IAM role</a>. If you create a role in your account and trust the Convox account to assume it, AWS operations can be performed without sharing long term credentials.

There are a few ways to create this role and policy.

## AWS Integration Quickstart

1. Log into [Console](https://console.convox.com/)
2. Click the **Integrations** menu button
3. Click the **Enable** button for the AWS provider
4. Drag a `credentials.csv` file or enter your Access Key ID and Secret Access Key into the access key form and click **Integrate AWS**

That's it! Your AWS account now has a `convox/ConvoxRole-$ID` IAM role (where `$ID` is your Console Organization ID) and a `convox/ConvoxPolicy` IAM policy. Convox has permission to use it to [install Racks](/docs/installing-a-rack/).

<div class="block-callout block-show-callout type-info" markdown="1">
### Note on `credentials.csv` format

The CSV file should be in a format similar to either of the following examples:

```
User name,Password,Access key ID,Secret access key,Console login link
aj,,AKIAICABC123A,1ababABCD+TP1Abcabca+TTabcABC,https://1234567.signin.aws.amazon.com/console
```

or:

```
User Name,Access Key Id,Secret Access Key
"aj",AKIAICABC123A,1ababABCD+TP1Abcabca+TTabcABC

```

</div>

## Generating an Administrator Access Key

If you don't have an access key, you can create one with the AWS Management Console.

1. From the AWS [IAM Users dashboard](https://console.aws.amazon.com/iam/home?#/users), select **Add user**.
2. Enter a username (e.g. `convox-setup`). Select the "Programmatic access" checkbox and click **Next: Permissions**.
3. Select the "Attach existing policies directly" button at the top. Search for the `AdministratorAccess` policy, select it, then click **Next: Review**.
4. On the next screen, click **Create user**
5. On the final screen, click the **Download .csv** button, and save the resulting file. Then click **Close**.

Then follow the **AWS Integration Quickstart**.

After the "AWS Integration Created" message, you can delete the `convox-setup` user, which disables the sensitive administrator Access Key.

## AWS CLI Instructions

If you don't want to provide an Access Key to Convox, you can set up the role with the [AWS CLI](https://aws.amazon.com/cli/) that is configured with your administrator Access Key.

1. Log into [Console](https://console.convox.com/)
2. Click the **Integrations** menu button
3. Click the **Enable** button for the AWS provider
4. Copy the commands from the **AWS CLI Authorization** section into your terminal
5. Paste the resulting ARN into the form and click **Integrate AWS**

The CLI commands take the general form of:

```bash
aws iam create-policy --policy-name 'ConvoxPolicy' --path '/convox/' --description 'Policy that Convox can assume' --policy-document '...'
aws iam create-role --role-name 'ConvoxRole-$ID' --path '/convox/' --assume-role-policy-document '{"Version": "2012-10-17","Statement": [{"Effect": "Allow","Principal": {"AWS": "665986001363"},"Condition": {"StringEquals": {"sts:ExternalId": "$ID"}},"Action": ["sts:AssumeRole"]}]}'
aws iam attach-role-policy --role-name 'ConvoxRole-$ID' --policy-arn $(aws iam list-policies --path-prefix '/convox/' --query 'Policies[?PolicyName==`ConvoxPolicy`].Arn' --output text)
aws iam get-role --role-name 'ConvoxRole-$ID' --query Role.Arn --output text
```

But they are custom for your account, where `$ID` is your Console Organization ID.

## Convox Role Review

You can review the IAM Role and Trust Relationship on your own.

1. From the AWS [IAM Roles dashboard](https://console.aws.amazon.com/iam/home?#/roles), search for **ConvoxRole**.
2. Click the **ConvoxRole-$ID** role name
3. On the next screen, select the **Trust Relationships** tab
4. Verify that the role trusts account **665986001363** and has a unique sts:ExternalId

## Removing the AWS Integration

At any time you can delete the `convox/ConvoxRole-$ID` role and `convox/ConvoxPolicy` policy. If you do this, you will not be able to install new Racks, but it will not affect any running apps.

1. Log into [Console](https://console.convox.com/)
2. Click the **Integrations** menu button
3. Click the **Disable** button for the AWS provider

That's it! Your AWS account no longer has a Convox role and policy.

You can also delete the role and policy through the [IAM dashboard](https://console.aws.amazon.com/iam/home) or with the [AWS CLI](https://aws.amazon.com/cli/).
