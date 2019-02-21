---
title: Heroku Migration
layout: guide
tags: convox, heroku, ecs, aws, migration
description: Migrate your application from Heroku to AWS with Convox
---

Migrating from Heroku to Convox is pretty straightforward and if you follow these instructions you should be able to easily move your application over to Convox with minimal downtime.

The basic steps are:

1. [Prerequisites](#prerequisites)
2. [Containerize your application](#containerize-your-application)
3. [Write your Convox manifest](#write-your-convox-manifest)
4. [Run your application locally](#run-your-application-locally)
5. [Deploy to Convox](#deploy-to-convox)
6. [Migrate your data](#migrate-your-data)
7. [Make the switch](#make-the-switch)

### Prerequisites

(you can check out our [getting started guide](https://docs.convox.com/introduction/getting-started) if you need more details on any of these)

- If you haven’t already you will need to [sign up](https://console.convox.com/signup)
- Install the [Convox CLI](https://docs.convox.com/introduction/installation) on your local computer
- Install a [local rack](https://docs.convox.com/development/running-locally)
- [Connect your AWS account to Convox](https://docs.convox.com/console/aws-integration)

### Containerize Your Application

If you are already deploying docker images to Heroku using their container registry you can skip this step. If you currently deploy using Heroku Buildpacks then you will need to create a dockerfile for your application. We have [sample applications](https://github.com/convox-examples/) for many popular frameworks that you can copy as a starting point. For a simple rails application a docker file might look like:

```
FROM ruby:2.5
RUN apt-get update -qq && apt-get install -y nodejs postgresql-client
RUN mkdir /myapp
WORKDIR /myapp
COPY Gemfile /myapp/Gemfile
COPY Gemfile.lock /myapp/Gemfile.lock
RUN bundle install
COPY . /myapp`
```

### Write Your Convox Manifest

All of our [sample applications](https://github.com/convox-examples/) have sample [convox.yml](https://docs.convox.com/application/convox-yml) files that you can copy. For a typical application that is running on heroku you are going to want to define a single database [resource](https://docs.convox.com/application/resources) such as postgres. You are also going to want to define a primary [service](https://docs.convox.com/application/services) for you application. Typically your build parameter will be set to “.” which instructs Convox to build the current directory. Your command parameter will typically be whatever you have defined in your current procfile. Again for a simple rails app your `convox.yml` might look like:

```
resources:
  database:
    type: postgres
services:
  web:
    build: .
    command: bundle exec rails server -b 0.0.0.0 -p 3000
    port: 3000
    resources:
      - database
```

### Run Your Application Locally

One of the great advantages of Convox is the built in [local development environment](https://docs.convox.com/development/running-locally) that mimics your production environment. This means no more dealing with setting up your dev environment every time you get a new laptop or add a new team member. It also means no more “it works on my machine” problems. To run locally simply type  `convox start` in a terminal window from your the base directory for your application. Open another terminal and run `convox switch local` followed by ‘convox services` to find the URL of your locally running app.

### Deploy To Convox

Once you have your app running locally and everything looks good it’s time to deploy. First make sure you have created a [production rack](https://docs.convox.com/introduction/getting-started#install-an-aws-rack). You can see a list of all your racks with `convox racks` and then switch to your production rack with `convox switch [rack name]`. Create an application with `convox apps create --wait` then deploy with `convox deploy --wait`. Once complete, you can grab the load balancer hostname for your app with `convox services`

### (Optional) Migrate Heroku Scheduler Jobs To Timers

On Heroku it’s popular to use the scheduler to run periodic tasks. Convox supports the same thing with [timers](https://docs.convox.com/application/timers). which are defined in the `convox.yml`. For example if we want to trigger a task to send emails every five minutes we might modify our `convox.yml` from above to look like

```
resources:
  database:
    ....
services:
  web:
    ....
timers:
  sendemails:
    schedule: "*/5 * * * ?"
    command: bin/rake send_emails
    service: worker
```

### Migrate Your Data

For the purposes of this example we will assume you are using Postgres as your data store but the same concepts apply for MySQL.

- Take a snapshot of your Heroku database with `heroku pg:backups:capture`
- Download the snapshot with `heroku pg:backups:download` taking note of the dumpfile name (typically “latest.dump”)
- Grab the connection string for your Convox production database with `convox resources` The url will be in the format `postgres://[username]:[password]@[hostname]:[port]/[database]`
- Open a [local proxy](https://docs.convox.com/management/resources) to your Convox database using `convox resources proxy [name]`. If you receive an error that the port is in use make sure you don’t have a local postgres database running already. You can also specify a different port with --port to avoid this conflict.
- Restore the backup onto your Convox production database (please note this will overwrite any data you have on your Convox production database) by opening a new terminal window and running `pg_restore --verbose --clean --no-acl --no-owner -h localhost -U [username] -d [database] [dumpfile name]` If you receive an error like `pg_restore: [archiver] unsupported version` you may need to upgrade your local postgres to the latest version
- Test your Convox production application and make sure everything works and your latest data is present

### Make The Switch

Now that you are sure everything is running correctly on Convox it’s time to make the production switch. There will be some small amount of downtime associated with the move so you will likely want to plan your move during off hours and make whatever preparations are appropriate for your application to be down for a short period of time. Once you are all set here are the steps:

- (1 day ahead) Lower the TTL for the DNS host record that you are going to migrate. For example if your app’s URL is www.mycompany.com then you are going to want to go to your DNS provider and lower the TTL for the A record or Cname for www.mycompany.com to something short (typically 60 seconds is reasonable) You can see instructions for how to do this with AWS route 53 [here](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-values-basic.html#rrsets-values-basic-ttl). This will ensure that your DNS change propagates as quickly as possible
- (1 day ahead) Create a certificate
  Use `convox certs generate <domain>` to have Convox generate a certificate for your domain. Have your AWS account administrator be on the lookout for a certificate approval email from AWS and make sure they approve it.
- (1 day ahead) Update the [domain](https://docs.convox.com/deployment/custom-domains) section of your `convox.yml` to the hostname you are migrating (ex: www.mycompany.com) and redeploy your Convox application.
- (right before you are ready to make the switch)
  Disable any [heroku scheduler jobs](https://devcenter.heroku.com/articles/scheduler) you may have. This will ensure that nothing writes to your Heroku database once you have put your app in maintenance mode. You can always recreate these jobs using [timers](#optional-migrate-heroku-scheduler-jobs-to-timers) in Convox
- Put your heroku app in [maintenance mode](https://devcenter.heroku.com/articles/maintenance-mode). At this point your Heroku application will be offline.
- If you have any [worker dynos](https://devcenter.heroku.com/articles/background-jobs-queueing) you will want to give them a few seconds to complete any open tasks and then scale them down to zero.
- Follow the database migration steps outlined [above](#migrate-your-data)
- Grab the router value for your convox rack using `convox rack` and update the CNAME record for your desired url (ex: www.mycompany.com) to [point to that location](https://docs.convox.com/deployment/custom-domains#configuring-dns)
- Wait a few minutes for your DNS to update (you may need to flush your DNS cache on your computer) and verify that your site is live.

Congratulations you are migrated to Convox! Enjoy those much smaller hosting bills.

#### Further Reading:

- [Workflows](https://docs.convox.com/console/workflows)
- [Autoscaling](https://docs.convox.com/deployment/scaling)
- [Rolling Deployments](https://docs.convox.com/deployment/rolling-updates)
