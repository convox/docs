---
title: "Getting Started"
order: 100
---

Getting started with Convox is easy. The instructions below guide you through:

* [Signing up](https://convox.com/signup)
* Installing the CLI
* Running your application locally
* Deploying your application to AWS

This guide takes around 30 minutes to go from zero to your first production deployment.

## Sign up

First you will need to [sign up for a Convox account](https://convox.com/signup).

## Install the CLI

The `convox` [command line tool](/reference/cli) offers:

* `convox start` - Start an application in development mode
* `convox deploy` - Deploy an application

along with numerous other commands that make configuring, scaling, and securing your apps simple.

* [Install the Convox CLI](/development/installation) for your platform.

* Next, click the **[Setup](https://console.convox.com/grid/user/welcome)** button then **[Connect the Convox CLI](https://console.convox.com/grid/user/api_key)** to get your API key.

* Finally, use the `convox login` command with your [API key](https://console.convox.com/grid/user/api_key):

<pre id="login">
$ convox login
Password: ********************
Authenticating with console.convox.com... OK
</pre>

## Prepare your application

If you already have a [Dockerized](https://docs.docker.com/engine/examples/) application, running on Convox is as easy as adding one small file that describes your application. If you are not already using Docker, don't worry we have sample applications for all popular frameworks that will make it easy to get started.

* If you have an existing application follow these easy steps [here](/development/preparing-an-application).
* If you are starting from scratch you can clone a [demo application](https://github.com/convox-examples) to get started.

## Run locally for development

### Install a local Rack

To ensure your production deployments behave exactly the same as your local development environment Convox installs a local [Rack](/introduction/rack) for development that mimics your production Racks.

If you already have Docker installed, [Installing your local Rack](/development/running-locally) is as simple as:

    $ sudo convox rack install local
    $ convox switch local

### Start your application in development mode

If your application is ready to go, you run it locally with `convox start`.

Here is an example using one of our sample apps:

    $ git clone https://github.com/convox-examples/rails
    Cloning into 'rails'...
    $ cd rails
    $ convox start
    build   | uploading source
    build   | starting build
    build   | Building: ...

Once your build completes you can open a new terminal and run `convox services` to find the hostname for your application.

    $ convox services
    SERVICE  DOMAIN            PORTS
    web      web.rails.convox  80:3000 443:3000

Now if you point your browser at `https://web.rails.convox` you can see your application in action!

You can make changes to your local directory and refresh to see those changes reflected in your browser.

## Run on AWS for Production

### Connect an AWS Account

Click on the integrations link in the main navigation button then click the plus sign on the runtime integration section. 

![](/assets/images/docs/console/integrations.png)

This grants Convox access and permission to help manage resources in your AWS account.

See [AWS Integration](/console/aws-integration) for more details.

### Install an AWS Rack

Next, click on  ***Racks*** in the main navigation and click the <img src="/assets/images/docs/add-rack-new.png" alt="Add Rack" style="height: 1.5em;"> button and select your AWS account. Enter a descriptive Rack name such as `production` if you plan to deploy production services, or `staging` if this is for testing.

See [Installing a Rack](/deployment/installing-a-rack) for more details.

You can now switch your CLI to point at your new Rack.

    $ convox racks
    NAME             STATUS
    local/convox     running
    acme/production  running

    $ convox switch acme/production
    Switched to acme/production

### Deploy your application

Before deploying, create a new app in your Rack.

    $ convox apps create --wait

Deploy the application

    $ convox deploy --wait

Once complete, run `convox services` to find the load balancer hostnames for the application.

    $ convox services
    SERVICE  DOMAIN                                     PORTS
    web      rails-web-123456789.us-east-1.convox.site  80:3000 443:3000

<div class="block-callout block-show-callout type-info" markdown="1">
When a load balancer is first created it can take a few minutes for its hostname to become available in DNS.
</div>

## Next Steps

Now that you've deployed your first application you can:

* Create a production database like [Postgres](/resources/postgresql) and link it to your app
* [Prepare and deploy more of your own apps](/development/preparing-an-application)
* Grant your team members [access](/console/access-control) to your organization
* Set up Continuous Delivery [Workflows](/console/workflows)
* Install another Rack for a staging environment

<script>
$(document).ready(function() {
  if (navigator.platform.indexOf('Win') > -1) {
    $('#install-windows').removeClass('hidden')
    $('#install-mac').addClass('hidden')
    $('#install-linux').addClass('hidden')
  }

  if (navigator.platform.indexOf('Linux') > -1) {
    $('#install-linux').removeClass('hidden')
    $('#install-mac').addClass('hidden')
    $('#install-windows').addClass('hidden')
  }
});
</script>
