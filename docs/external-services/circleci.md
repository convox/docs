---
title: "CircleCI"
---

You can deploy to Convox as part of a CircleCI workflow using the [Convox Orb](https://circleci.com/orbs/registry/orb/convox/orb)

## Sign Up for CircleCI

If you don’t have an account already, [sign up for CircleCI](https://circleci.com/signup/). 

## Enable 3rd Party Orbs

Because the Convox Orb is new, it has not yet completed the CircleCI certification process. For now, the Admin of your CircleCI org must opt-in to 3rd-party uncertified orb usage on the Settings > Security page for your org

## Configure Your CircleCI Project to Build With Convox

If you don't already have a project setup in CircleCI you will need to [add one](https://circleci.com/docs/2.0/gh-bb-integration/#section=projects)

Once you have your project added, login to the Convox console and click on the settings tab in the left side navigation bar. 

From here generate a [deploy key](/console/deploy-keys) and copy it to your local clipboard.

In CircleCI, add an [environment variable](https://circleci.com/docs/2.0/env-vars/#setting-an-environment-variable-in-a-project) with the name `CONVOX_DEPLOY_KEY` and paste your deploy key as the value

## Build a CircleCI config.yml Using the Convox Orb
***
***Note: You must specify `version: 2.1`  in your config.yml in order to use [Orbs](https://circleci.com/docs/2.0/using-orbs/)***
***
The Convox Orb contains a single `deploy` command and matching job for deploying your app to Convox using CircleCI.

The job is self contained. If you would prefer to use the command you must run the `checkout` command before using `deploy`.


The `deploy` Job/Command accepts the following parameters:

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
    <th>Required</th>
    <th>Default Value</th>
  </tr>
  <tr>
    <td><code>rack</code></td>
    <td>The Rack to deploy to</td>
    <td>Yes</td>
    <td></td>
  </tr>
  <tr>
    <td><code>app</code></td>
    <td>The App to deploy</td>
    <td>No</td>
    <td>Project Repository Name</td>
  </tr>
  <tr>
    <td><code>host</code></td>
    <td>Your console hostname</td>
    <td>No</td>
    <td>console.convox.com</td>
  </tr>
</table>

## Examples

The simplest version of a config.yml would look like:

```
version: 2.1
orbs:
  convox: convox/orb@1.4.1
workflows:
  deploy:
    jobs:
      - convox/deploy:
          rack: production
          app: example
```
This will build and deploy your app in a single step

If you would prefer to use the Convox Orb commands directly you need to run the `checkout` command before deploy:
```
version: 2.1
orbs:
  convox: convox/orb@1.4.1
workflows:
  deploy:
    jobs:
      - deploy
jobs:
  deploy:
    executor: convox/cli
    steps:
      - checkout
      - run: make test
      - convox/deploy:
          rack: production
          app: example
```