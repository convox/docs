---
title: "Travis CI"
order: 400
---

You can streamline your workflow by integrating Convox and Travis CI. At a high level, you'll be using familiar CLI commands like `convox build` and `convox deploy`, only from your Travis CI build servers.

## Modifying .travis.yml

First you need to tell Travis CI to install the Convox CLI by adding this to your `.travis.yml`:

```
before_install: |
  curl -O https://bin.equinox.io/c/jewmwFCp7w9/convox-stable-linux-amd64.tgz &&\
  tar zxvf convox-stable-linux-amd64.tgz -C /tmp
```


The [after_success section](https://docs.travis-ci.com/user/deployment/custom/) of `.travis.yml` lets you specify commands to run after a successful build. In the example below, a successful build would trigger a deployment of `example-app` to the `org/staging` Rack.

    after_success:
      - convox deploy --app <app name> --rack <org name>/<rack name>

### Example .travis.yml

Here's an example `.travis.yml` that installs the Convox CLI, runs `convox doctor`, and deploys the app:

<pre class="file yaml" title=".travis.yml">
sudo: required

services:
  - docker

before_install: |
  curl -O https://bin.equinox.io/c/jewmwFCp7w9/convox-stable-linux-amd64.tgz &&\
  sudo tar zxvf convox-stable-linux-amd64.tgz -C /usr/local/bin

script:
  convox doctor

after_success:
  - convox deploy --app cv-soulshake-net --rack personal/legit
</pre>

## Authentication

You'll also need to enable the Travis CI build server to authenticate with your Rack before it can run commands like `convox build` or `convox deploy`. When using the CLI from your development machine, you'd typically `convox login` to do so, but when using Travis CI, you'll want to set `CONVOX_HOST` and `CONVOX_PASSWORD` instead. These Convox credentials are sensitive, so you should set them in your [Travis CI Repository Settings](https://docs.travis-ci.com/user/environment-variables/#Defining-Variables-in-Repository-Settings), if not with the even more secure [Travis CI Encrypted Variables](https://docs.travis-ci.com/user/environment-variables/#Encrypted-Variables), rather than directly in `.travis.yml`.

### Authenticating with Console deploy keys

If you use [Console](https://console.convox.com/) to manage access to your Racks, you'll need to set the following environment variables in Travis CI:

    CONVOX_HOST=console.convox.com
    CONVOX_PASSWORD=<deploy key>

For more information, see [deploy keys](/docs/deploy-keys).
