---
title: "Frequently Asked Questions"
order: 500
---

#### Can I use Convox for free?

  Definitely.  Single-user Console accounts are [free](https://console.convox.com/signup).

#### Is Convox open source?

  The Convox CLI, Rack platform, and many of our supporting libraries are indeed open source, and available through [GitHub](https://github.com/convox)!  The Convox Enterprise Console codebase is private.

#### Is support available?

  Yes!  Business Day support is available on [Pro plans](/pricing) and 24/7 support available with an [Enterprise license](/enterprise).  [Community support](https://community.convox.com/) is also available!

#### Can I use my own AWS/GCP account with Convox?

  Definitely, that's how it works.  You maintain control and direct access to all of your infrastructure.  Convox makes managing it easier, with no markup based on how many instances or apps you run!

#### Can I use my AWS/GCP credits with Convox?

  Yep!  Use them up as you wish, no restrictions here.

#### What is a Rack?

  [A Rack](/introduction/rack) is your deployment platform for your applications and services.  It's an abstraction of your ECS/EKS/GKE/local cluster, installed in your own cloud account or local machine.  It provides an API to manage your applications, controlled from the Convox CLI or Console.  

#### Do you support Multi-Cloud?

  You can use Convox to install your Racks in your AWS account.  Google Cloud (GCP) support is currently in Beta.  Keep an eye out for further platforms in the future 😉

#### Does Convox integrate with CircleCi, Jenkins, or my custom CI/CD pipeline?

  We have a [CircleCI Orb](https://circleci.com/orbs/registry/orb/convox/orb) ready for use to make your CI/CD easier.  Jenkins and other solutions are also well supported with the use of the Convox CLI and unique [deploy keys](/console/deploy-keys) to provide access from your automated processes.  

#### What kind of apps can I deploy on Convox?

  Any Dockerised app (or even Docker wrapper) can be deployed.  We have example apps in various languages and frameworks available [here](/example-apps/examples), and we are adding more all the time.

#### What types of Databases, Block Storage, Messsage Queues, etc… does Convox support?

  Convox provides native support for common [application resources](/application/resources) such as MySQL/Postgres/Redis/Memcached.  [Rack level resources](/gen1/resources) also provide support for S3/SQS/SNS/Syslog available to all apps.  We add others on a regular basis.  Anything not directly supported can easily be integrated, by passing the endpoint for your external resource into the app with an [environment variable](/application/environment).  

#### How does Convox handle scaling?

  You have multiple options to control and configure your scaling.  [Application scaling](/deployment/scaling) can be static or automatic, based on CPU, Memory, Traffic or custom metrics.  Rack scaling can also be configured, even utilising Spot instances or Fargate for lower pricing and more flexibility!

#### Is Convox PCI/HIPAA compliant?

  Yes, we are an ideal solution for systems which need the highest level of security and accountability.  More details on making your installation HIPAA compliant can be found [here](reference/hipaa-compliance).

#### Can I use Convox with an external Docker Registry?

  Yes, you can access your [private Docker images](/deployment/private-registries) easily!

#### Does Convox provide a solution for local development?

  We do.  [Local rack installation](/development/running-locally) is quick and easy, whether you develop on Mac, Linux or Windows.  This provides a near identical replica to your cloud staging or production environments.  Alternatively, developers can utilise cloud development environments to relieve stress and pressure on their local machines.

#### Can I setup multiple environments (dev/staging/production) with Convox?

  Naturally.  A common topology our customers use is to create separate racks for each environment they wish.  Each Rack can be installed into any pre-existing VPCs you may have for your different environments to provide appropriate isolation and security between them.

#### What’s the difference between EKS and ECS powered Racks?

  From a developer/operations perspective, they are interchangeable.  No changes are needed to your processes when changing between the two!  However, EKS powered racks provide faster rollouts and deployments, and whilst currently in Beta, will allow for a much higher level of customisation around load balancers and other Kubernetes configuration.

#### How many applications can I deploy in a single Rack?

  There is no artificially small limit.  Some of our customers easily run up to 100 apps in a single rack.  There are many variables to determine, and decide how you configure your deployment environments.

#### Does Convox have a solution for secrets management?

  Yes, all environment variables are stored and hardware encrypted by default within KMS for security and resilience.

#### Does Convox have a solution for CI/CD?

  [Workflows](/console/workflows), available within the Console, allow for easy and integrated CI/CD processes for deployments, testing and app review.
