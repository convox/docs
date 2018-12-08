---
title: "AWS Integration"
---

Convox is designed to give your team agility and consistency on top of AWS. In minutes you can install Convox and stand up new apps that follow AWS and DevOps best practices. Every Convox app leverages AWS CloudFormation and the EC2 Container Service for consistent configuration management and automation to reliably deploy, monitor and scale apps.

To this end, Convox requires access and permission to help manage resources in your AWS account.

Granting Convox access to your AWS account is extremely easy. Simply navigate to the integrations section of the console.

![](/assets/images/docs/console/integrations.png)

Then click on the plus sign in the Runtime section and select Amazon Web Services. When the prompt comes up click ***Launch Stack***


![](/assets/images/docs/console/launch-stack.png)

If you are not already logged into your AWS account you will be prompted to do so. Once you complete this step it will take approximately 10 minutes for Convox to create the necessary infrastructure in your AWS account. 

You can remove the AWS integration at any time by clicking the gear icon next to your AWS account on the integrations screen and selecting ***Remove*** just make sure you delete any Racks you have created before removing the integration.


