---
title: "Getting Started"
order: 400
---

Getting started with Convox is easy. The instructions below guide you through:

* [Signing up for a Convox account](https://convox.com/signup)
* Install Convox
* Run your first application locally
* Connect your AWS account
* Deploy your first application

This guide takes around 30 minutes to go from zero to your first production deployment.

## Sign Up

First, if you haven't already, you will need to [sign up for Convox](https://convox.com/signup).

## Install the Convox CLI

The `convox` [command line tool](/reference/cli) offers:

* `convox start` - A single command to start your development environment
* `convox deploy` - A single command to deploy your application to a Rack

along with numerous other utilities that make building, configuring, scaling and securing your apps easy.

* [Install the Convox CLI](/development/installation) for your platform.

* Next, click the **[Setup](https://console.convox.com/grid/user/welcome)** button then **[Connect the Convox CLI](https://console.convox.com/grid/user/api_key)** to get your API key.

* Finally, use the `convox login` command with your [API key](https://console.convox.com/grid/user/api_key):

<pre id="login">
$ convox login
API Key: ********-****-****-****-************
Logged in successfully.
</pre>

## Prepare your application

If you already have a [dockerized](https://docs.docker.com/engine/examples/) application, running on Convox is as easy as adding one small file that helps Convox understand how you want your application to run. If you are not already using Docker, don't worry we have sample applications for all popular frameworks that will make it easy to get started.

* If you have an existing application that you want to run on Convox, follow these easy steps [here](/development/preparing-an-application)
* If you are starting from scratch, clone one of our [demo applications](https://github.com/convox-examples) to get started

## Run your application locally
To ensure your production deployments behave exactly the same as your local development environment Convox installs a local [rack](/introduction/rack) for development that mimics your production racks.
If you already have Docker installed, [Installing your local rack](/development/running-locally) is as simple as:
    
    $ sudo convox rack install local
    $ convox switch local

If your application is ready to go, you run it locally with `convox start` 
Once your application is up and running you just need to point your browser at it.


Here is an example using one of our sample apps:
    
    $ git clone https://github.com/convox-examples/rails
    Cloning into 'rails'...
    $ cd rails
    $ convox start 
    build   | uploading source
    build   | starting build
    build   | Building: ...
    
 Once your build completes, you can open a new terminal and see what services you have running locally with `convox services`
    
    $ convox services
    SERVICE  DOMAIN            PORTS           
    web      web.rails.convox  80:3000 443:3000
  
Now if you point your browser at `https://web.rails.convox` you can see your first app in action!

Once you have had a chance to play around with local development it's time to deploy to production. First you will need to connect your AWS account
    

## Connect an AWS Account

Click the **[Setup](https://console.convox.com/grid/user/welcome)** button then **Connect an AWS account**, and give Convox an AWS [access key](https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys). This grants Convox access and permission to help manage resources in your AWS account.

See [AWS Integration](/console/aws-integration) for more details.

## Launch your private PaaS

Next, click on the <img src="/assets/images/docs/add-rack.png" alt="Add Rack" style="height: 1.5em;"> button and then **Install New**, in the top navigation bar. Enter a descriptive Rack name such as `production` if you plan to deploy production services, or `staging` if this is for testing.

See [Installing a Rack](/deployment/installing-a-rack) for more details.


## Deploy to Convox

#### Go to the local directory containing the app you want to deploy

In this case we will use the rails example app we cloned for local development

    $ git clone https://github.com/convox-examples/rails
    $ cd rails
    
#### Switch to your production Rack
Assuming in the step above you created a Rack called `production` we need to point your Convox CLI at that Rack instead of your local Rack so your commands are executed against the staging Rack.

    $ convox racks
    NAME                  STATUS     
    local/convox          running    
    MyCompany/production  running   
    
    $ convox switch MyCompany/production
    Switched to MyCompany/production
    
Don't forget when you want to go back to local development to switch racks with `convox switch local`

#### Create an app in your Rack

Before deploying, create a new app in your Rack.

    $ convox apps create <name>
    
Wait for the underlying components to be created:

    $ convox apps wait

Deploy the application

    $ convox deploy
  
By default convox assumes your application is the same name as the current directory. If you gave your app a custom name you will need to specify the app name using `-a <app_name>`

For example
    
    my-laptop $ cd rails
    my-laptop:rails $ convox apps create demoapp
    my-laptop:rails $ convox deploy -a demoapp

Watch `convox services` to find the load balancer hostnames for the application.

    $ convox services

<div class="block-callout block-show-callout type-info" markdown="1">
When a load balancer is first created it can take a few minutes for its hostname to become available in DNS.
</div>

## Next Steps

Now that you've deployed your first application you can:

* Create a production database like [Postgres](/resources/postgresql) and link it to your app
* [Prepare and deploy more of your own apps](/development/preparing-an-application)
* Grant your team members [access](/console/access-control) to your organization
* Set up Continuous Delivery [Workflows](/console/workflows)
* Install another Rack for isolated development or staging deployments

Or you can easily [uninstall everything](/reference/uninstalling-convox) you just experimented with.

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
