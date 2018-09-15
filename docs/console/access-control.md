---
title: "Access Control"
---

Once you've created an organization, you'll find extra settings on the "Organization" pane in the Convox Console, allowing you to define access controls for your Convox resources. There, you'll find three sections: [Settings](#settings), [Users](#users), and [Deploy Keys](#deploy-keys).

## Settings

Here you can modify the organization name or delete the organization entirely.

## Users

On this page you can manage which users have access to your team and Racks.

**On the Pro plan and above**, you can assign one of the following roles to each of the users you've added: _administrator_, _operator_, or _developer_. For details about the permissions of each role, see the [Permissions section](#permissions) below.

When users join your organization, each will be assigned a unique API Key that grants access the org's Racks.

![Organization Members](/assets/images/docs/rbac/rbac.png){: .screenshot }

## Deploy Keys

Deploy keys are special, limited-scope API keys that allow you to run the `build` and `deploy` commands from remote servers for the purposes of CI.

You should give a Deploy Key to a CI service like CircleCI or Travis CI so it can deploy code but not access or modify any other Rack resources.

For more details, see [Deploy Keys](/docs/deploy-keys).

![Deploy Keys](/assets/images/docs/rbac/deploy-keys.png){: .screenshot }

## Permissions

| Action                      | Developer    | Operator     | Administrator    |
| --------------------------- | ------------ | ------------ | ---------------- |
| Billing:Read                |              |              | X                |
| DeployKey:Read              |              |              | X                |
| Environment:Read            |              | X            | X                |
| Environment:Write           |              | X            | X                |
| Integration:Create          |              | X            | X                |
| Organization:Delete         |              | X            | X                |
| Organization:Read           | X            | X            | X                |
| OrganizationInvite:Create   | X            | X            | X                |
| OrganizationUser:Create     | X            | X            | X                |
| OrganizationUser:Kick       |              |              | X                |
| Rack:Execute                |              | X            | X                |
| Rack:Read                   | X            | X            | X                |
| Rack:Write                  |              | X            | X                |
| RackEvents:Read             | X            | X            | X                |
| User:Read                   | X            | X            | X                |
| User:UpdateRole             |              |              | X                |
| Webhook:Create              |              | X            | X                |
| Workflow:Create             |              | X            | X                |
| Workflow:Delete             |              | X            | X                |
| Workflow:Edit               |              | X            | X                |
| Workflow:Update             |              | X            | X                |
