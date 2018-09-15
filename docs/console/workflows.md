---
title: "Workflows"
---

[Convox Console](https://console.convox.com) supports the creation of Workflows. A Workflow is a list of tasks executed in response to a trigger. You can use Workflows to build automation--like deploying when you push to GitHub--into an organization.

1. [Creating a Workflow](#creating-a-workflow)
  - [Defining the Trigger](#defining-the-trigger)
  - [Defining Tasks](#defining-tasks)
1. [Managing Workflows](#managing-workflows)
1. [Workflow History](#workflow-history)
1. [Example Workflows](#example-workflows)

## Creating a Workflow

In [Console](https://console.convox.com/), click the Workflows tab. Click the "Create Workflow" button to get started.

![](/assets/images/docs/workflows/tab.png)

### Defining the Trigger

First, choose a trigger for the Workflow. This is the action that will kick off execution of the Workflow tasks.

![](/assets/images/docs/workflows/trigger.png)

Choose either GitHub or GitLab as your integration, then choose a repository and a branch.

When changes are pushed to the specified repository branch, Workflow tasks will be triggered.

### Defining Tasks

Next, define the tasks to be executed once the trigger fires. Tasks are actions that can be executed for any app in your organization. You can currently choose from the following task types:

- **Build**: create a Build from the specified branch.
- **Promote**: take a Build and promote its corresponding Release.
- **Run**: execute a command against a particular Service of a Build.
- **Copy**: export a Build from one App and import it to another, even across Racks.

Add additional tasks to your Workflow by clicking the green "Add Task" button. Once you've defined all of your tasks, click "Create Workflow".

### Naming the Workflow

Finally, give the Workflow a unique, recognizable name.

![Assign workflow name](/assets/images/docs/workflows/name.png) *This is the name that will be used to identify the overall purpose of the workflow.*

#### An Example

Here is an example Build. It promotes the Release corresponding to the resulting Build to complete the automated deployment.

![](/assets/images/docs/workflows/task.png)

## Managing Workflows

Existing Workflows can be managed via the main Workflows tab. Use the buttons to the right if you wish to edit or remove a Workflow.

![](/assets/images/docs/workflows/manage.png)

## Workflow History

The execution of a Workflow is referred to as a "job." From the main Workflows page, click the Jobs button in the right-side column to view a list of previously executed Workflows. A red or green indicator will tell you whether the job completed successfully.

![](/assets/images/docs/workflows/jobs.png)

Click the timestamp link associated with a job to see more detailed information including logs and the Build and Release IDs involved in each task.

## Example Workflows

As mentioned above, Workflows can be used along with other tools that set commit statuses to build Continuous Integration pipelines. Here are a few Workflows you might try with your team.

### Double Build

1. Create a Workflow that triggers on your repo's master branch
2. Define a task to build and promote your staging app
3. Define a second task to build (but not promote) your production app

With this setup, a merge into your master branch will automatically execute everything needed right up to going live in production. If you are satisfied with the state of your staging app when the Workflow completes, you can simply click "Promote" on your production app in Console to complete the process.

For a fully automated production pipeline, enable promotion in step 3. Since a failure of any task in a Workflow halts execution for remaining tasks, you can be sure that your production app will only be promoted if the staging app was successfully promoted (i.e. all containers were rolled out and passed [health checks](/docs/health-checks)).

### CI for Pull Requests

1. Create a Workflow that triggers on a Pull Request event
2. Define a task to create a build based on the PR
3. Define a second task to run your test suite against that Build

With this setup, each PR in your repo would be tested automatically upon submission.

### Single Build with CI

1. Create a Workflow that triggers when the master branch has been pushed
2. Define a task to create a Build in your staging Rack
3. Define a second task to run your test suite against that Build (Workflow aborts if tests fail)
4. Define a third task to promote that Build's release in the staging Rack
5. Define a fourth task to copy that passing Build from the staging to the production Rack
6. After completing QA on the staging app, promote the release in production, or...
   Define a final task to promote the Build automatically in production

With this setup, merging changes to your master branch and pushing them to your GitHub or GitLab remote would trigger a new Build in your staging Rack. If that Build were to pass your test suite, its changes would go live (i.e. its Release would be promoted) for that app in your staging Rack. That same Build would be copied to your production Rack in anticipation of wanting to promote its Release as well. You could then manually promote the Release corresponding to that build at any time to complete the deployment to production, allowing your team to complete QA on the staging app first.
