---
title: "Rails"
---

Our example Rails app can be found [here](https://github.com/convox-examples/rails).  You can clone this locally to run and experiment with.

The following is a step-by-step walkthrough of how the app was configured and why. The sample app is a fresh new Rails app, Dockerised, configured and a `convox.yml` added.  See the `README.md` in the repo for the changes made.

### Running Locally

A few steps to get started:

1. Make sure you have [Docker](https://www.docker.com/products/docker-desktop) installed on your local machine. 
2. [Signup for a Convox account](https://convox.com/signup) It's free!
3. Install the [Convox CLI](https://docs.convox.com/development/installation). 
4. Install the [Convox local Rack](https://docs.convox.com/development/running-locally)

Once you are all setup you can switch to your local rack with ```convox switch local``` and start your local application with ```convox start``` (make sure you are in the root directory).

You should now be able to access your application by going to [https://web.rails.convox](https://web.rails.convox). If you renamed anything you may need to modify your local URL. The format is https://[service name].[app name].convox

### Custom Application Components

#### Dockerfile

Starting from the [ruby:2.6.4](https://hub.docker.com/_/ruby/) image, the [Dockerfile](https://github.com/convox-examples/rails/blob/master/Dockerfile) executes the remaining build steps that your Rails app needs. There are basically 3 steps in this process, and they are executed in a particular order to take advantage of Docker's build caching behavior.

1. `bundle install` and `yarn install` are run to bring in any dependencies that are required by the app.

2. The application source is copied over. These files will change frequently, so this step of the build will very rarely be cached.

3. Finally, after setting the appropriate environment variables, the assets are precompiled.

#### Convox.yml

The [convox.yml](https://github.com/convox-examples/rails/blob/master/convox.yml) file explains how to run the application. This file has two sections.

1. Resources: These are network-attached dependencies of your application. In this case we have a single resource which is a postgres database. When [running locally](https://docs.convox.com/development/running-locally) Convox will automatically startup up a container running Postgres and will inject a ```DATABASE_URL``` environment variable into your application container that points to the Postgres database. When your application is [deployed](https://docs.convox.com/deployment/deploying-to-convox) to production Convox will startup an RDS postgres database for your application to use. 

2. Services: This is where we define our application(s). In this case we have a single application called ```web``` which is built from our dockerfile, and uses the postgres resource for a database.  In a production application you may have additional services defined for things like Sidekiq task workers.

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

