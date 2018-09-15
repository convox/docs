---
title: "App Parameters"
---

Parameters can be used to configure your Convox apps. Below is a list of the available app parameters.

<ul>
  <li><a href="#ltprocessnamegtformation">&lt;ProcessName&gt;Formation</a></li>
  <li><a href="#internal">Internal</a></li>
  <li><a href="#securitygroup">SecurityGroup</a></li>
  <li><a href="#taskrole">TaskRole</a></li>
</ul>

## Setting Parameters

Parameters can be set using the following command.

    convox apps params set Foo=bar

You can also set multiple parameters at once.

    convox apps params set Foo=bar Baz=qux

## &lt;ProcessName&gt;Formation

For a given app, specify the number of processes to run, CPU units to reserve, and MB of RAM to reserve.

A &lt;ProcessName&gt;Formation parameter is created for each app process you define in your `docker-compose.yml`. For example, your app might have `WebFormation` and `DatabaseFormation` parameters.

See the [cpu section](http://docs.aws.amazon.com/AmazonECS/latest/APIReference//API_ContainerDefinition.html#ECS-Type-ContainerDefinition-cpu) of the AWS ContainerDefinition doc for more information about reserving CPU units.

Keep in mind that [ECS will terminate your app](http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html#cfn-ecs-taskdefinition-containerdefinition-memory) if it attempts to use more than the amount of memory you have reserved for it. You may need to [scale up your app's memory](/docs/scaling#memory) if you encounter this kind of termination unexpectedly.

| Default value  | "1,0,256"        |

## Internal

Have the app use Internal ELBs for all processes, i.e. make it unreachable from the Internet. See our [Internal Apps doc](/docs/gen1/internal-apps) for more information.

| Allowed values | "Yes", "No" |
| Default value  | "No"        |

## SecurityGroup

The `SecurityGroup` app parameter can be set to the ID of a custom AWS security group. When this param is set on an app, the security group rules will be applied to requests coming in to any of the app's load balancers.

For details, see [Load balancers: limited application access](/docs/gen1/load-balancers#limited-application-access).

## TaskRole

The `TaskRole` app parameter can be set to the ARN or short name of an IAM Role you wish to apply to the ECS Tasks of this app.
