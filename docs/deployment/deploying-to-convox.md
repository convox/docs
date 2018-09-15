---
title: "Deploying to Convox"
order: 200
---

You have a few options for building and deploying your applications to Convox.

## Automatic Builds

You can easily configure [Console](https://console.convox.com) to build and deploy your application when changes are pushed to your main repository. See [Workflows](/docs/workflows) for more information.

## CLI

* Go to the directory that contains your application.
* Type `convox switch <org>/<rack>` to select the Rack that contains the application to deploy.
* Type `convox deploy --app <appname>`

## Canceling a deployment

To cancel a bad or stuck deployment, run `convox apps cancel -a <app-name>`. Behind the scenes, this cancels the CloudFormation stack update that corresponds to the deployment in progress.

## Troubleshooting

For troubleshooting tips, see [this page](/docs/troubleshooting/#i-get-an-error-when-i-deploy-my-app-to-convox).
