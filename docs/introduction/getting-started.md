---
title: "Getting Started"
order: 100
---

Getting started with Convox is easy. The instructions below guide you through:

* [Signing up](https://convox.com/signup)
* Installing the CLI
* Deploying your application to AWS

This guide takes around 30 minutes to go from zero to your first production deployment.

## Sign up

First you will need to [sign up for a Convox account](https://convox.com/signup).

## Install the CLI

* [Install the Convox CLI](/introduction/installation) for your platform.

* Next, click the **[Setup](https://console.convox.com/grid/user/welcome)** button then **[Connect the Convox CLI](https://console.convox.com/grid/user/api_key)** to get your API key.

* Finally, use the `convox login` command with your [API key](https://console.convox.com/grid/user/api_key):

<pre id="login">
$ convox login
Password: ********************
Authenticating with console.convox.com... OK
</pre>

## Prepare your application

If you already have a [Dockerized](https://docs.docker.com/engine/examples/) application, running on Convox is as easy as adding one small file that describes your application. If you are not already using Docker, don't worry we have sample applications for all popular frameworks that will make it easy to get started.

* If you have an existing application see the [convox.yml](/application/convox-yml) section of these docs.
* If you are starting from scratch you can clone a [demo application](https://github.com/convox-examples) to get started.

## Deploy to AWS

### Connect an AWS Account

Click on the integrations link in the main navigation button then click the plus sign on the runtime integration section.

![](/assets/images/docs/console/integrations.png)

This grants Convox access and permission to help manage resources in your AWS account.

See [AWS Integration](/console/aws-integration) for more details.

### Install an AWS Rack

Next, click on  ***Racks*** in the main navigation and click the <img src="/assets/images/docs/add-rack-new.png" alt="Add Rack" style="height: 1.5em;"> button and select your AWS account. Enter a descriptive Rack name such as `production` if you plan to deploy production services, or `staging` if this is for testing.

Wait for the Rack to finish installing.

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

* Add a database like [Postgres](/application/resources)
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
