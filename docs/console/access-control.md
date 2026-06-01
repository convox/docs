---
title: "Access Control"
description: "Manage organization settings, user roles, deploy keys, and role-based permissions in the Convox Console."
---

# Access Control

Once you've created an organization, you'll find extra settings on the "Organization" pane in the Convox Console, allowing you to define access controls for your Convox resources. There, you'll find three sections: [Settings](#settings), [Users](#users), and [Deploy Keys](#deploy-keys).

## Settings

Here you can modify the organization name or delete the organization entirely.

## Users

On this page you can manage which users have access to your team and Racks.

**On the Pro plan and above**, you can assign one of the following roles to each of the users you've added: _Developer_, _Operator_, or _Administrator_, listed below from least to most access. For details about the permissions of each role, see the [Permissions](#permissions) section below.

When users join your organization, each will be assigned a unique API Key that grants access to the organization's Racks.

### Access Model and RBAC

The roles described on this page are the original access model. The newer [Role-Based Access Control (RBAC)](/console/rbac) model lets you define granular, custom roles instead of choosing from these fixed roles.

Custom RBAC roles require the Pro plan or higher. On the Pro plan and above, only the _Administrator_ role grants full organization access, and you can scope every other user with a fixed or custom role. On plans below Pro, custom roles are not available and every user in the organization is effectively an _Administrator_ with full access.

Users who have not been assigned a custom role keep their original permissions. See [RBAC](/console/rbac) for how to create and assign granular roles.

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

## See Also

- [RBAC](/console/rbac)
- [Deploy Keys](/console/deploy-keys)
- [Workflows](/console/workflows)
- [Audit Logs](/console/audit)
- [Enterprise: LDAP Authentication](/enterprise/ldap-authentication)
- [Enterprise: SAML Authentication](/enterprise/saml-authentication)
