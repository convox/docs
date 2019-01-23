---
title: "Workflows"
---

***Please note this document refers to the new workflows available as of the December 2018 update of the Convox Console. Legacy workflows created prior to this update will still work, however they can no longer be created or edited. The new workflows provide much improved functionality over the legacy workflows and we highly encourage everyone to migrate to the new workflows.***
***


Convox allows you to manage continuous integration and continuous delivery using workflows.  You can use workflows to build automation such as deploying when you push to GitHub. You can create workflows by clicking on the workflow link in the navigation bar

![](/assets/images/docs/workflows/workflows-start.png)

Convox has two types of workflows:

* [Review Workflows](#review-workflows)
* [Deployment Workflows](#deployment-workflows)

## Review Workflows

A review workflow allows you to review a new version of your application based on a pull request. When you define a review workflow for a connected Github or Gitlab repository, Convox will build your application whenever a pull request is created for that repository. If selected, Convox will also run your tests against that build and can even create a running review application so you can test your latest changes before you merge the pull request.

#### Creating a Review Workflow

* Repository - This is the linked source control repository from either Github or Gitlab from which creating pull requests will trigger your review workflow
* Manifest - The Convox manifest file to use for this workflow. This defaults to [convox.yml](/application/convox-yml) however, if you have custom needs for your review workflows you can specify a custom manifest
* Rack - This specifies what Rack the demo application will be deployed in 
* Run tests - Checking this box will run whatever command is in the test directive in your [convox.yml](/application/convox-yml). One thing to note is when run tests is enabled, two releases are created for every build and both releases will appear in the release list for your application.
* Deploy Demo - Checking this box will instruct Convox to create a demo application as part of the review workflow. The demo application will be deployed in the specified rack and will be deleted automatically once the pull request is merged
* Before Promote - Here you can specify a service and a command to be run before your application is promoted. This is useful for things like database migrations
* After Promote - Here you can specify a service and a command to be run after your application is promoted. This can be useful for things like notifications or cleanup scripts
* Demo Environment - Here you can specify any environment variables you want set or overridden for your demo application

## Deployment Workflows

A deployment workflow is how you can manage the regular deployment of your applications to staging and production. A deployment workflow is triggered whenever code is merged into the specified repository/branch on either Github or Gitlab

#### Creating a Deployment Workflow

* Repository - This is the linked source control repository from either Github or Gitlab 
* Workflow Name - A name for your workflow ex: staging deploy
* Branch - The branch which merging code into will cause this workflow to trigger
* Manifest - The Convox manifest file to use for this workflow. This defaults to [convox.yml](/application/convox-yml) however if you have custom needs for a specific deployment workflow you can specify a custom manifest
* App - Deployment workflows allow you to specify multiple applications to deploy to and you can specify whether or not to automatically promote each application. For example on a merge to your master branch you might choose to deploy to a staging application with automatic promotion and simultaneously deploy to a production application with manual promotion.
* Run tests - Checking this box will run whatever command is in the test directive in your [convox.yml](/application/convox-yml). One thing to note is when run tests is enabled, two releases are created for every build and both releases will appear in the release list for your application.
* After Promote - Here you can specify a service and a command to be run after your application is promoted. This can be useful for things like notifications or cleanup scripts.
* Demo Environment - Here you can specify any environment variables you want set or overridden for your demo application.

While deployment workflows are triggered by merges to the specified repository/branch you can also run a deployment workflow manually by clicking <img src="/assets/images/docs/workflows/workflow-play.png"  style="height: 1.5em;">

## Workflow Jobs

Whenever a workflow is triggered it creates a job. You can view the jobs history by clicking on the jobs link in the navigation bar

![](/assets/images/docs/workflows/workflow-jobs.png)

Here you can see a complete history of all your review and deployment workflow jobs. You can also re-run any job. You can click on any job in the history list to review all steps and output from that job run.

![](/assets/images/docs/workflows/job-detail.png)

## Example Workflows

The flexibility of Convox workflows should meet the needs of almost any team but here are a few popular workflow examples:

### Review Workflow for Testing 

1. Create a review workflow for your repository
2. Assign the workflow to your staging rack
3. Select "run tests" and "deploy demo"
4. Specify any commands such as running migrations that you might need for your application to work
5. Specify any specific environment variables you might need set or overridden for demo applications

Now whenever a developer on your team opens a pull request a demo application will be created and deployed in your staging rack. 

The URL for the demo application will be available in the job detail screen once the review workflow is complete.

### A Gitflow Deployment Strategy

If your team practices [Gitflow](https://nvie.com/posts/a-successful-git-branching-model/) where you have a staging environment which is automatically built whenever a commit is made to your dev branch and you do production deploys by either merging your dev branch to your master branch, or by committing a hotfix directly to master, then you will want to do something like:

#### Staging Deployment
1. Create a deployment workflow for your repository specifying your dev branch
2. Add your staging application and choose automatic promotion
3. Choose whether or not you want to run tests 

#### Production Deployment
1. Create a deployment workflow for your repository specifying your master branch
2. Add your production application and decide whether you want to promote automatically or manually
3. Choose whether or not you want to run tests 

Now whenever you merge to dev your staging application will be built and whenever you merge to master your production application will be built.   

### A Double Build Deployment Strategy

We are particularly fond of this strategy at Convox. With a double build, whenever code is merged into your master branch your staging application is automatically built and deployed and your production application is automatically built and ready to be deployed. With this strategy, once you are satisfied with your testing on staging all you need to do is click promote to move your code into production.

1. Create a deployment workflow for your repository specifying your master branch
2. Add your staging application and choose automatic promotion
3. Choose whether or not you want to run tests on staging
4. Add your production application and choose manual promotion

Now whenever your merge to master your staging application will be automatically built and your production application will be built and ready to promote. Once you are satisfied with your testing on staging you can click on your production rack, select your application, find the latest release and click promote.

![](/assets/images/docs/workflows/promote-release.png)
