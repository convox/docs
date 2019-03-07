---
title: "Circleci"
---

You can deploy to Convox as part of a Circleci workflow using the [Convox Orb](https://circleci.com/orbs/registry/orb/convox/orb)

## Sign up for Circleci

If you don’t have an account already, [sign up for Circleci](https://circleci.com/signup/). 

## Enable 3rd party Orbs

Because the Convox Orb is new, it has not yet completed the Circle certification process. For now, the Admin of your Circleci org must opt-in to 3rd-party uncertified orb usage on the Settings > Security page for your org

## Configure your Circleci project to build with Convox

If you don't already have a project setup in Circleci you will need to [add one](https://circleci.com/docs/2.0/gh-bb-integration/#section=projects)

Once you have your project added, login to the Convox console and click on the settings tab in the left side navigation bar. 

From here generate a [deploy key](/console/deploy-keys) and copy it to your local clipboard.

In Circleci, add an [environment variable](https://circleci.com/docs/2.0/env-vars/#setting-an-environment-variable-in-a-project) with the name `CONVOX_DEPLOY_KEY` and paste your deploy key as the value

## Build a Circleci config.yml using the Convox Orb

Note: You must specify `version: 2.1` in your config.yml in order to use [Orbs](https://circleci.com/docs/2.0/using-orbs/)

The Convox Orb contains the following as both [Jobs](https://circleci.com/docs/2.0/configuration-reference/#jobs) and [Commands](https://circleci.com/docs/2.0/configuration-reference/#commands-requires-version-21). 

<table>

  <tr>
    <td><code>build</code></td>
    <td>Builds the current project</td>
  </tr>
  <tr>
    <td><code>promote</code></td>
    <td>Promotes the build you just ran</td>
  </tr>
  <tr>
    <td><code>deploy</code></td>
    <td>Builds and promotes in one step </td>
  </tr>
</table>

The jobs are self contained. If you would prefer to use the commands you must run the `convox/install` and `checkout` commands before using any of the other commands.
In addition the `promote` command requires that the `build` command is run first.

Each Job/Command accepts the following parameters:

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
    <th>Required</th>
    <th>Default Value</th>
  </tr>
  <tr>
    <td><code>rack</code></td>
    <td>The Rack to build/deploy to</td>
    <td>Yes</td>
    <td></td>
  </tr>
  <tr>
    <td><code>app</code></td>
    <td>The App to build/deploy</td>
    <td>No</td>
    <td>Project Repository Name</td>
  </tr>
  <tr>
    <td><code>convox-host</code></td>
    <td>For private console orgs only: your console url</td>
    <td>No</td>
    <td>console.convox.com</td>
  </tr>
</table>

## Examples

The simplest version of a config.yml would look like:

```
version: 2.1
orbs:
  convox: convox/orb@1.2.0

workflows:
  "Convox build":
    jobs:
      - convox/deploy:
          rack: myrack
```
This will build and deploy your app in a single step

If you prefer to build and promote as two steps you need to make sure you run the build step first:

```
version: 2.1
orbs:
  convox: convox/orb@1.2.0

workflows:
  "Convox deploy":
    jobs:
      - convox/build:
          rack: myrack
      - convox/promote:
          rack: myrack
          requires:
            - convox/build
```

If you would prefer to use the Convox Orb commands directly you need to run the `convox/install` and `checkout` commands before any others:
```
version: 2.1
orbs:
  convox: convox/orb@1.2.0

workflows:
  convox_deploy:
    jobs:
      - deploy

jobs:
  deploy:
    executor: convox/default
    steps:
      - checkout
      - convox/install
      - convox/deploy:
          rack: myrack
```