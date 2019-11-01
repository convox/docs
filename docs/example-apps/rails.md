---
title: "Rails"
---

This guide will show you how to install a Rails app to Convox using an [example application](https://github.com/convox-examples/rails).

## Deploying the Application

### Initial Setup

Ensure that you have a working Convox Rack and that your `convox` CLI is pointing at it.

If you don't yet have a Rack, you can follow the instructions in the [Getting Started Guide](/introduction/getting-started).

### Clone the Example App

* `git clone https://github.com/convox-examples/rails`
* `cd rails`

### Deploying

Create the application and wait for it to finish:

* `convox apps create --wait`

Deploy the application and wait for it to finish:

* `convox deploy --wait`

Determine the hostname of the `web` service of the application and visit it in your browser:

* `convox services`

## Application Layout

Here we will describe the changes we made to a vanilla Rails application to allow it to be installed to Convox.

### [Dockerfile](https://github.com/convox-examples/rails/blob/master/Dockerfile)

Starting from the [ruby:2.6.4](https://hub.docker.com/_/ruby/) image, the `Dockerfile` defines the steps necessary to turn the application code into an image that is ready to run. 

This `Dockerfile` has 3 steps and they are executed in a particular order to take advantage of Docker's build caching behavior.

1. `bundle install` and `yarn install` are run to install dependencies after copying just the files needed to run these commands. This will ensure that the output these commands are cached unless one of these files changes.

2. The application source is copied over. These files will change frequently so this step of the build will very rarely be cached.

3. Finally, after setting the appropriate environment variables the assets are precompiled.

### [convox.yml](https://github.com/convox-examples/rails/blob/master/convox.yml)

The `convox.yml` manifest explains how to run the application. The manifest for this application has two sections:

#### resources

These are network-attached dependencies of the application. In this application we have a single resource, a `postgres` database.

<div class="block-callout block-show-callout type-info" markdown="1">
	The implementation of a resource will vary by Rack type. For example, on a local Rack a `postgres` resource will be satisfied with a containerized database while on AWS that same resource definition will be satisfied with an RDS database.
</div>

#### services

These are the web-facing services of the application. This application has a single service named `web` which is built from the local directory.

Because the resource named `database` appears in the `links:` section of this service it will receive an environment variable named `DATABASE_URL` with connection details.

### [config/database.yml](https://github.com/convox-examples/rails/blob/master/config/database.yml)

This file is configured to read database credentials from the `DATABASE_URL` environment variable.

