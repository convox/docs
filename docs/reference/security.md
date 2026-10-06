---
title: "Security"
description: "Overview of Convox security features including AWS isolation, IAM permissions, VPC networking, load balancers, and dedicated instances."
---

# Security

Convox provides multiple layers of security for your Apps and infrastructure. This page describes the key security features.

## AWS Isolation

Your Convox Rack is installed in your own AWS account. Unlike a multi-tenant PaaS, you never have to worry about other users gaining access to your infrastructure.

[Convox Console](https://console.convox.com) is a multi-tenant management layer that proxies API requests to your Rack and provides team management and workflow tools. Many security measures are baked in such as rollable API keys, programmatic [deploy keys](/console/deploy-keys), and user roles. However, if you don't want to use the multi-tenant Console you can run your own! [Contact us](mailto:support@convox.com) for information about plans and pricing.

## AWS Permissions

The Rack API, builds and the instance autoscaler run as one IAM role, `ApiRole`, which is also the role CloudFormation uses for Rack, App and resource updates. `ApiRole` carries these managed policies:

| Policy | Grants |
|:-------|:-------|
| `PowerUserAccess` (AWS managed), or the Rack's `ApiPolicyScoped` when [ApiRoleScoped](/reference/rack-parameters/ApiRoleScoped) is `Yes` | Every AWS service except IAM, AWS Organizations and account management. `ApiPolicyScoped` limits this to the services the Rack uses |
| `ApiPolicyV2` | IAM reads and writes on roles, policies, instance profiles and users under the `/convox/` path in the Rack's own account; deleting, and listing the policies of, roles outside `/convox/` whose ARN matches `role/<rack>-*-????????????`, and reading instance profiles whose ARN matches `instance-profile/<rack>-*-????????????`, the names CloudFormation generates for the Rack's stacks; `iam:GetRole` and `iam:PassRole` on every role in the account; `iam:GetPolicy` on the [PermissionsBoundary](/reference/rack-parameters/PermissionsBoundary) policy while it is set; IAM server certificates |
| `CMKPolicy` | The Rack's KMS key |

Its inline policies add `lambda:GetFunction`, read access to the Rack's own API secret in Systems Manager Parameter Store, Secrets Manager access to the Rack's own secrets, ECR image pushes and pulls, and ECR Public pulls.

Every IAM role, policy, instance profile and user the Rack creates is under `/convox/`. To cap what those roles and users can do, including App roles that get grants from [IamPolicy](/reference/app-parameters/IamPolicy) or a Service's `policies`, set [PermissionsBoundary](/reference/rack-parameters/PermissionsBoundary).

Deploys and parameter changes leave in place any policy added outside Convox to a role the Rack created. Deleting an App, a Service, a Timer or a Rack resource removes those policies from its roles: managed policies are detached and kept, and inline policies are deleted with the role. This requires rack version 20261005214736 or newer; on older Racks such a policy stops the delete.

On rack version 20261005214736 or newer, when IAM refuses to create a role during `convox apps create` or `convox rack resources create`, the stack rolls back, the new App or Rack resource shows `failed`, and `convox apps delete` or `convox rack resources delete` removes it. On a Rack whose name is longer than 25 characters or contains `--`, some generated role names do not match `role/<rack>-*-????????????`, and the App or Rack resource can show `unknown` instead; see [Recovering a stuck rollback](/reference/rack-parameters/PermissionsBoundary#recovering-a-stuck-rollback).

## VPC Isolation

All of the infrastructure that Convox creates for you runs inside a Virtual Private Cloud ([VPC](https://aws.amazon.com/vpc/)). This provides additional isolation at the networking layer. By default, all resources such as datastores are created in such a way that they can only be accessed from inside the VPC.

## Load Balancers

Convox uses AWS Load Balancers to route traffic to your application. Load balancers and their Security Groups are set up to only listen on the ports you specify and only route requests to the relevant application services.

## Private Networking

If you'd like to take network isolation one step further you can run your Rack in private networking mode, where the Rack instances run in private subnets that access the Internet through NATs. These instances are not routable via the public Internet. Read more on the private networking [doc](/networking/private-networking).

## Dedicated Instances

If you would like to ensure hardware single tenancy all the way down to the AWS infrastructure level you can do so by passing `Tenancy=dedicated` option to the `convox rack install` command when setting up your Rack.

## See Also

- [HIPAA Compliance](/reference/hipaa-compliance)
- [ApiRoleScoped](/reference/rack-parameters/ApiRoleScoped)
- [PermissionsBoundary](/reference/rack-parameters/PermissionsBoundary)
- [Private Networking](/networking/private-networking)
- [Console Access Control](/console/access-control)
- [AWS Infrastructure Details](/reference/aws)
