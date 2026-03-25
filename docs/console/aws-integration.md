---
title: "AWS Integration"
description: "How to connect your AWS account to Convox Console for Rack installation and management."
---

# AWS Integration

Convox uses AWS CloudFormation and the Elastic Container Service (ECS) to deploy, monitor, and scale your Apps. To do this, Convox requires access and permission to manage resources in your AWS account.

## Connecting Your AWS Account

Navigate to the integrations section of the Console.

![](/assets/images/docs/console/integrations.png)

Click the plus sign in the Runtime section and select **Amazon Web Services**. When the prompt appears, click **Launch Stack**.

![](/assets/images/docs/console/launch-stack.png)

If you are not already logged into your AWS account, you will be prompted to do so. Once you complete this step, it takes approximately 10 minutes for Convox to create the necessary infrastructure in your AWS account.

## Removing the Integration

You can remove the AWS integration at any time by clicking the gear icon next to your AWS account on the integrations screen and selecting **Remove**. Delete any Racks you have created before removing the integration.

## See Also

- [Import a Rack](/management/import-rack)
- [Core Concepts: Console](/core-concepts/console)
- [Reference: AWS](/reference/aws)
- [Integrations](/console/integrations)
