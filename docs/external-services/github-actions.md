---
title: "Github Actions"
---

You can deploy to Convox using our [Github Actions](https://github.com/marketplace/actions/convox-deploy)

## Available Actions
Convox provides a full set of Actions to enable a wide variety of deployment workflows. The following actions are available:
### [Deploy](https://github.com/convox/action-deploy)
Authenticates, builds, and deploys in a single step
### [Login](https://github.com/convox/action-login)
Authenticates your Convox account. You should run this action as the first step in your workflow
### [Build](https://github.com/convox/action-build)
Builds an app and creates a release which can be promoted later
### [Run](https://github.com/convox/action-run) 
Runs a command (such as a migration) using a previously built release before or after it is promoted
### [Promote](https://github.com/convox/action-promote)
Promotes a release 


## Adding a Github Workflow to Your Repository
You can read the complete details about Github Workflows [here](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/configuring-a-workflow) but in the simplest case all you need to do is add a `.github/workflows` directory to your repository and put a `deployment.yml` file in it. In this file you add the [trigger](https://help.github.com/en/articles/events-that-trigger-workflows#webhook-events) for your workflow as well as the  [steps](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/workflow-syntax-for-github-actions#jobsjob_idsteps) you want the workflow to run. In addition, as is the case with the Convox Actions, you may want to add some [secrets](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/creating-and-using-encrypted-secrets) to your repository to reference during your workflow.

## A Simple Deploy
As a simple example, let's say we have Rails app that we want to deploy to our Convox Staging Rack every time a commit is pushed to the master branch. The steps would be as follows:
### Add a Deploy Key to Github
The first thing we will do is add our [Deploy Key](https://docs.convox.com/console/deploy-keys) as an [Encrypted Github Secret](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/creating-and-using-encrypted-secrets) called `CONVOX_DEPLOY_KEY`. This will allow us to securely authenticate with Convox when we deploy without storing any keys in our code.
### Create a Deployment.yml file
For this simple example our `.github/workflows/deployment.yml` file looks like
```
name: CD
on: [push]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - name: checkout
      uses: actions/checkout@v1
    - name: deploy
      uses: convox/action-deploy@v1
      with:
        rack: staging
        app: myrailsapp
        password: ${{ secrets.CONVOX_DEPLOY_KEY }}
```
The steps are very simple.
* Trigger on push
* Checkout the latest code
* Run the Convox Deploy Action
## More Advanced Workflows
You can create more advanced workflows by combining the [individual Convox Github Actions](#available-actions). For example, if we want to build the same Rails app but run migrations before promoting we would a create a `deployment.yml` that looks something like
```
name: CD
on: [push]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - name: checkout
      id: checkout
      uses: actions/checkout@v1
    - name: login
      id: login
      uses: convox/action-login@v1
      with:
        password: ${{ secrets.CONVOX_DEPLOY_KEY }}
    - name: build
      id: build
      uses: convox/action-build@v1
      with:
        rack: staging
        app: myrailsapp
    - name: migrate
      id:migrate
      uses: convox/action-run@v1
      with:
        rack: staging
        app: myrailsapp
        service: web
        command: rake db:migrate
        release: ${{ steps.build.outputs.release }}
    - name: promote
      id: promote
      uses: convox/action-promote@v1
      with:
        rack: staging
        app: myrailsapp
        release: ${{ steps.build.outputs.release }}
```
In this case the steps are as follows
* Trigger on push
* Checkout the latest code
* Authenticate with the [`login`](#loginhttpsgithubcomconvoxaction-login) action
* Build with the [`build`](#buildhttpsgithubcomconvoxaction-build) action
* Run our migrations with the [`run`](#runhttpsgithubcomconvoxaction-run) action using the `release` output variable from the build step to target that specific release for our run
* Promote our release with the [`promote`](#promotehttpsgithubcomconvoxaction-promote) action also using the `release` output variable from the build step

One thing to note is that by default the `run` and `promote` actions will use the release from a previous build step so you only need to pass the `release` value if you want to target a release other than the one you just built.