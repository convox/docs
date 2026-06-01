---
title: Heroku
description: Migrate your application from Heroku to AWS with Convox
---

# Heroku

Migrating from Heroku to Convox is straightforward and if you follow these instructions you should be able to move your application over to Convox with minimal downtime.

The basic steps are:

1. [Prerequisites](#prerequisites)
2. [Containerize your application](#containerize-your-application)
3. [Write your Convox manifest](#write-your-convox-manifest)
4. [Run your application locally](#run-your-application-locally)
5. [Deploy to Convox](#deploy-to-convox)
6. [Migrate your data](#migrate-your-data)
7. [Make the switch](#make-the-switch)

## Prerequisites

(you can check out our [getting started guide](/introduction/getting-started) if you need more details on any of these)

- If you haven't already you will need to [sign up](https://console.convox.com/signup)
- Install the [Convox CLI](/introduction/installation) on your local computer
- Install a [local rack](/development/running-locally)
- [Connect your AWS account to Convox](/console/aws-integration)

## Containerize Your Application

If you are already deploying docker images to Heroku using their container registry you can skip this step. If you currently deploy using Heroku Buildpacks then you will need to create a dockerfile for your application. We have [sample applications](https://github.com/convox-examples/) for many popular frameworks that you can copy as a starting point. For a simple rails application a docker file might look like:

```dockerfile
FROM ruby:3.3
RUN apt-get update -qq && apt-get install -y nodejs postgresql-client
RUN mkdir /myapp
WORKDIR /myapp
COPY Gemfile /myapp/Gemfile
COPY Gemfile.lock /myapp/Gemfile.lock
RUN bundle install
COPY . /myapp
```

## Write Your Convox Manifest

All of our [sample applications](https://github.com/convox-examples/) have sample [convox.yml](/application/convox-yml) files that you can copy. For a typical application that is running on heroku you will define a single database [resource](/application/resources) such as postgres and a primary [service](/application/services) for your application. The key fields to set are:

| Field | Value | Notes |
|-------|-------|-------|
| `resources` | a database such as `postgres` | Maps to your Heroku add-on database |
| `build` | `.` | Instructs Convox to build the current directory |
| `command` | the command from your Procfile | Whatever you currently run as your web process |
| `port` | the port your app listens on | The service exposes this port |

Again for a simple rails app your `convox.yml` might look like:

```yaml
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

## Run Your Application Locally

One of the great advantages of Convox is the built in [local development environment](/development/running-locally) that mimics your production environment. This means no more dealing with setting up your dev environment every time you get a new laptop or add a new team member. It also means no more "it works on my machine" problems. To run locally, type `convox start` in a terminal window from your the base directory for your application. Open another terminal and run `convox switch local` followed by `convox services` to find the URL of your locally running app.

## Deploy To Convox

Once you have your app running locally and everything looks good it's time to deploy. First make sure you have created a [production rack](/introduction/getting-started#rack-installation), then:

1. List your racks with `convox racks`.
2. Switch to your production rack with `convox switch [rack name]`.
3. Create an application with `convox apps create --wait`.
4. Deploy with `convox deploy --wait`.
5. Once complete, grab the load balancer hostname for your app with `convox services`.

## (Optional) Migrate Heroku Scheduler Jobs To Timers

On Heroku it's popular to use the scheduler to run periodic tasks. Convox supports the same thing with [timers](/application/timers) which are defined in the `convox.yml`. For example if we want to trigger a task to send emails every five minutes we might modify our `convox.yml` from above to look like

```yaml
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

## Migrate Your Data

For the purposes of this example we will assume you are using Postgres as your data store but the same concepts apply for MySQL.

1. Take a snapshot of your Heroku database with `heroku pg:backups:capture`.
2. Download the snapshot with `heroku pg:backups:download`, taking note of the dumpfile name (typically "latest.dump").
3. Grab the connection string for your Convox production database with `convox resources`. The url will be in the format `postgres://[username]:[password]@[hostname]:[port]/[database]`.
4. Open a [local proxy](/management/resources) to your Convox database using `convox resources proxy [name]`. If you receive an error that the port is in use make sure you don't have a local postgres database running already. You can also specify a different port with --port to avoid this conflict.
5. Restore the backup onto your Convox production database (this will overwrite any data you have on your Convox production database) by opening a new terminal window and running `pg_restore --verbose --clean --no-acl --no-owner -h localhost -U [username] -d [database] [dumpfile name]`. If you receive an error like `pg_restore: [archiver] unsupported version` you may need to upgrade your local postgres to the latest version.
6. Test your Convox production application and make sure everything works and your latest data is present.

## Make The Switch

Now that you are sure everything is running correctly on Convox it's time to make the production switch. There will be some small amount of downtime associated with the move so you will likely want to plan your move during off hours and make whatever preparations are appropriate for your application to be down for a short period of time. Once you are all set here are the steps:

1. (1 day ahead) Lower the TTL for the DNS host record that you are going to migrate. For example if your app's URL is www.mycompany.com then you are going to want to go to your DNS provider and lower the TTL for the A record or Cname for www.mycompany.com to something short (typically 60 seconds is reasonable). You can see instructions for how to do this with AWS route 53 [here](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-values-basic.html#rrsets-values-basic-ttl). This will ensure that your DNS change propagates as quickly as possible.
2. (1 day ahead) Create a certificate. Use `convox certs generate <domain>` to have Convox generate a certificate for your domain. Have your AWS account administrator be on the lookout for a certificate approval email from AWS and make sure they approve it.
3. (1 day ahead) Update the [domain](/deployment/custom-domains) section of your `convox.yml` to the hostname you are migrating (ex: www.mycompany.com) and redeploy your Convox application.
4. (right before you are ready to make the switch) Disable any [heroku scheduler jobs](https://devcenter.heroku.com/articles/scheduler) you may have. This will ensure that nothing writes to your Heroku database once you have put your app in maintenance mode. You can always recreate these jobs using [timers](#optional-migrate-heroku-scheduler-jobs-to-timers) in Convox.
5. Put your heroku app in [maintenance mode](https://devcenter.heroku.com/articles/maintenance-mode). At this point your Heroku application will be offline.
6. If you have any [worker dynos](https://devcenter.heroku.com/articles/background-jobs-queueing) you will want to give them a few seconds to complete any open tasks and then scale them down to zero.
7. Follow the database migration steps outlined [above](#migrate-your-data).
8. Grab the router value for your convox rack using `convox rack` and update the CNAME record for your desired url (ex: www.mycompany.com) to [point to that location](/deployment/custom-domains#configuring-dns).
9. Wait a few minutes for your DNS to update (you may need to flush your DNS cache on your computer) and verify that your site is live.

Congratulations you are migrated to Convox! Enjoy those much smaller hosting bills.

### Further Reading

- [Workflows](/console/workflows)
- [Autoscaling](/scaling/scaling)
- [Rolling Deployments](/deployment/rolling-updates)

## See Also

- [Getting Started](/introduction/getting-started)
- [convox.yml](/application/convox-yml)
- [Custom Domains](/deployment/custom-domains)
- [Timers](/application/timers)
