---
title: "Frequently Asked Questions"
description: "Common questions about Convox pricing, features, scaling, compliance, and integrations"
---

# Frequently Asked Questions

## Can I use Convox for free?

Single-user Console accounts are [free](https://console.convox.com/signup).

## Is Convox open source?

The Convox CLI, Rack platform, and many of our supporting libraries are open source and available through [GitHub](https://github.com/convox). The Convox Enterprise Console codebase is private.

## Is support available?

Yes. Business Day support is available on Pro plans and 24/7 support is available with an Enterprise license. [Community support](https://community.convox.com/) is also available. See [Support](/help/support) for details.

## Can I use my own AWS account with Convox?

Yes, that's how it works. You maintain control and direct access to all of your infrastructure. Convox makes managing it easier, with no markup based on how many instances or Apps you run.

## Can I use my AWS credits with Convox?

Yes. Use them as you wish, no restrictions.

## Can I use Reserved Instances with Convox?

Yes. Any reserved instance arrangements you have in place (if they match with the instance types you are running in your Racks) will automatically apply and you will benefit from your reduced pricing.

## What is a Rack?

[A Rack](/core-concepts/rack) is your deployment platform for your applications and services. It's an abstraction of your ECS cluster, installed in your own AWS account. It provides an API to manage your applications, controlled from the Convox CLI or Console.

## Do you support Multi-Cloud?

Convox v2 Racks are AWS-only. For multi-cloud support across AWS (EKS), Google Cloud (GKE), Azure (AKS), and DigitalOcean (DOKS), see [Convox v3](https://docs.convox.com/getting-started/introduction). A [migration guide](https://www.convox.com/blog/convox-v2-to-v3-migration-guide) is available.

## Does Convox integrate with CircleCi, Jenkins, or my custom CI/CD pipeline?

We have a [CircleCI Orb](https://circleci.com/orbs/registry/orb/convox/orb) as well as a set of [GitHub Actions](/integrations/github-actions) ready for use to make your CI/CD easier. Jenkins and other solutions are also well supported with the use of the Convox CLI and unique [deploy keys](/console/deploy-keys) to provide access from your automated processes.

## What kind of Apps can I deploy on Convox?

Any Dockerized App (or even Docker wrapper) can be deployed. We have example Apps in various languages and frameworks available [here](/example-apps/examples), and we are adding more all the time.

## What types of Databases, Block Storage, Message Queues, etc... does Convox support?

Convox provides native support for common [application resources](/application/resources) such as MariaDB/MySQL/Postgres/Redis/Memcached. [Rack level resources](/management/rack-resources) also provide support for S3/SQS/SNS/Syslog available to all Apps. We add others on a regular basis. Anything not directly supported can be integrated by passing the endpoint for your external resource into the App with an [environment variable](/application/environment).

## How does Convox handle scaling?

You have multiple options to control and configure your scaling. [Application scaling](/scaling/scaling) can be static or automatic, based on CPU, Memory, Traffic or custom metrics. Rack scaling can also be configured, even using Spot instances or Fargate for lower pricing and more flexibility!

## Is Convox PCI/HIPAA compliant?

It depends on the level of compliance you need and how you run Convox.

SOC2 and similar fintech compliance requirements are achievable with the SaaS Convox Console, since your applications and infrastructure always run in your own AWS account.

HIPAA is different. To run HIPAA workloads you need an Enterprise self-hosted Rack, including a self-hosted Console deployed into your own AWS account, and you sign the Business Associate Agreement (BAA) directly with AWS. Because Convox runs entirely within your AWS account, this is not a capability of the SaaS-hosted Console. More details on making your installation HIPAA compliant can be found [here](/reference/hipaa-compliance).

## Can I use Convox with an external Docker Registry?

Yes, you can access your [private Docker images](/integrations/private-registries).

## Does Convox provide a solution for local development?

Yes. [Local Rack installation](/development/running-locally) works on Mac, Linux, or Windows. This provides a near identical replica to your cloud staging or production environments. Alternatively, developers can use cloud development environments to relieve stress and pressure on their local machines.

## Can I setup multiple environments (dev/staging/production) with Convox?

A common topology our users use is to create separate Racks for each environment they wish. Each Rack can be installed into any pre-existing VPCs you may have for your different environments to provide appropriate isolation and security between them.

## What about EKS-powered Racks?

EKS-powered Racks are available in [Convox v3](https://docs.convox.com/getting-started/introduction), which is cloud-agnostic and runs on Kubernetes (EKS, GKE, AKS, DOKS). V3 Racks provide faster rollouts and deployments, and allow for a much higher level of customization around load balancers and other Kubernetes configuration. See the [v2 to v3 migration guide](https://www.convox.com/blog/convox-v2-to-v3-migration-guide) for details.

## How many applications can I deploy in a single Rack?

There is no artificially small limit. Some of our users run up to 100 Apps in a single Rack. There are many variables to determine, and decide how you configure your deployment environments.

## Does Convox have a solution for secrets management?

Yes, all environment variables are stored and hardware encrypted by default within KMS for security and resilience.

## Does Convox have a solution for CI/CD?

[Workflows](/console/workflows), available within the Console, allow for integrated CI/CD processes for deployments, testing and app review.

## See Also

- [Getting Started](/introduction/getting-started)
- [Core Concepts: Rack](/core-concepts/rack)
- [Core Concepts: Console](/core-concepts/console)
- [CLI Installation](/introduction/installation)
- [Migrating from Heroku](/migration/heroku)
