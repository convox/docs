---
title: "Import a Rack"
description: "Import an externally installed Convox Rack into the Console for centralized management."
---

# Import a Rack

While creating Racks through the Convox Console is recommended, you can also create Racks through the CLI or via the [Convox Installer](https://github.com/convox/installer) process.

If you want to manage and configure your Rack through the Console, you will need to import the Rack.

## Importing a GCP/AWS/Digital Ocean Rack to your Convox Console

If you have installed your Rack in GCP, Digital Ocean, or do not have an AWS Runtime integration you can import your Rack:

- From the 'Racks' page within the Console, click on the '+ Import' to open the import context.
- Fill in the exact name of the Rack, the hostname as reported by the installation, and the password you set during installation.
  - To find hostname and password in the AWS Console, go to the region you intalled the Rack: 
    - hostname: go to Cloudformation -> Stacks -> Filter for stack {rackname} -> Outputs (Tab) -> Key: Dashboard's Value is the hostname
    - password: go to AWS Systems Manager -> Parameter Store -> Filter for {rackname}RackApiSecret ->  Value is the password
- Click on 'Add Rack' and the Rack will be imported to your Organization's Convox Console .

![](/assets/images/docs/import-rack/import-rack-manual.png)

## See Also

- [AWS Integration](/console/aws-integration)
- [Rack Updates](/management/rack-updates)
- [Core Concepts: Rack](/core-concepts/rack)
- [Reference: Rack Parameters](/reference/rack-parameters)
