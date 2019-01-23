---
title: "Access Control"
---

Once you've created an organization, you'll find extra settings on the "Organization" pane in the Convox Console, allowing you to define access controls for your Convox resources. There, you'll find three sections: [Settings](#settings), [Users](#users), and [Deploy Keys](#deploy-keys).

## Settings

Here you can modify the organization name or delete the organization entirely.

## Users

On this page you can manage which users have access to your team and Racks.

**On the Pro plan and above**, you can assign one of the following roles to each of the users you've added: _Administrator_, _Operator_, or _Developer_. For details about the permissions of each role, see the [Permissions](#permissions) section below.

When users join your organization, each will be assigned a unique API Key that grants access the organization's Racks.

## Deploy Keys

Deploy keys are limited-scope API keys that allow you to run the `build` and `deploy` commands from remote servers.

You should give a Deploy Key to an external CI service so it can deploy code but not access or modify any other Rack resources.

For more details, see [Deploy Keys](/console/deploy-keys).

## Permissions

| Action               | Developer | Operator | Administrator |
|----------------------|:---------:|:--------:|:-------------:|
| Manage apps          |     X     |     X    |       X       |
| View workflow jobs   |     X     |     X    |       X       |
| Open support tickets |     X     |     X    |       X       |
| Manage racks         |           |     X    |       X       |
| Manage workflows     |           |     X    |       X       |
| View audit log       |           |     X    |       X       |
| Manage integrations  |           |     X    |       X       |
| Manage users         |           |          |       X       |
| Manage billing       |           |          |       X       |
| Manage organization  |           |          |       X       |
