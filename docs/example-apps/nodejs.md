---
title: "Node.js"
---

Our example node.js app can be found [here](https://github.com/convox-examples/nodejs).  You can clone this locally to run and experiment with.

The following is a step-by-step walkthrough of how the app was configured and why. The sample app is a fresh new node.js app, Dockerised, configured and a `convox.yml` added.  See the `README.md` in the repo for the changes made.

### Running Locally

A few steps to get started:

1. Make sure you have [Docker](https://www.docker.com/products/docker-desktop) installed on your local machine. 
2. [Signup for a Convox account](https://convox.com/signup) It's free!
3. Install the [Convox CLI](https://docs.convox.com/development/installation). 
4. Install the [Convox local Rack](https://docs.convox.com/development/running-locally)

Once you are all setup you can switch to your local rack with ```convox switch local``` and start your local application with ```convox start``` (make sure you are in the root directory).

You should now be able to access your application by going to [https://web.nodejs.convox](https://web.nodejs.convox). If you renamed anything you may need to modify your local URL. The format is https://[service name].[app name].convox

### Custom Application Components

#### Dockerfile

Starting from the [node:10.16.3-alpine](https://hub.docker.com/_/ruby/) image, the [Dockerfile](https://github.com/convox-examples/nodejs/blob/master/Dockerfile) simply sets up a basic nodejs app by copying the app files into the container, exposing port 3000 and specifying the command to be executed.

#### convox.yml

The [convox.yml](https://github.com/convox-examples/rails/blob/master/convox.yml) file explains how to run the application. This file only has one section for our node app.

1. Services: This is where we define our application(s). In this case we have a single application called ```web``` which is built from our dockerfile.  In a production application you may have additional sections for resources like databases etc.

### Deploying to production

In order to deploy to production we have to ensure we have completed the following steps:

1. [Signup for a Convox account](https://convox.com/signup)
2. [Connect an AWS account](https://docs.convox.com/introduction/getting-started#connect-an-aws-account)
3. [Install an AWS Rack](https://docs.convox.com/introduction/getting-started#install-an-aws-rack)
4. Make sure your CLI is [logged in](https://docs.convox.com/reference/login-and-authentication) to your Convox account using ```convox login``` and your [CLI Key](https://console.convox.com/account)

Once you are all set here you can see the name of your production rack

```bash 
convox racks
```

And switch your CLI to your production rack

```bash
convox switch [rack name]
```

Now you can create an empty application in your production rack

```bash
convox apps create --wait
```

And you can deploy your application to production (the first time you do this it may take up to 15 minutes to create the necessary resources)

```bash
convox deploy --wait
```

Finally you can retrieve the URL from your production application with

```bash
convox services
```

