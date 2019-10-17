---
title: "Import a Rack"
---

Whilst we recommend creating your Racks through the Convox Console, you can create Racks through the CLI or via our new Terraform creation process.

If you wish to then manage and configure your Rack through the Console, you will have to import the Rack.

## Installing a Rack via the Convox CLI

The two required [parameters](/reference/rack-parameters) when installing a Rack are the `Password` and `Version` fields.  Everything else is optional, but you may wish to set many of these [options](/reference/rack-parameters).

```
$ convox rack install aws Password=externalrack Version=20190822194002
Preparing... OK
Installing...  98.61% 6m14s
Starting... OK, convox-608673667.us-east-1.elb.amazonaws.com
```

Once installed, remember the password you have set and take a note of that hostname reported in the output.  You can go to the Console and start importing.

## Installing a Rack via Terraform

To install a Kubernetes-backed Rack on AWS or GCP you can follow the Terraform installation process from the instructions on the [repo here](https://github.com/convox-examples/terraform).


## Importing an AWS Rack to your Console account through your Runtime Integration

If you have an AWS Runtime integration then it's very simple to import your Rack (if you installed it in AWS 😉):

- From the 'Racks' page within the Console, click on the 'Add Rack' dropdown and select your runtime integration.
- Switch to the 'Import Rack' tab within the pop-up that appears.
- Select the Region your Rack was installed in, `us-east-1` by default unless you specified differently during installation.
- Convox will pull a list of all Racks installed in that region, so now choose your Rack's name from the drop-down.  `convox` by default, but if in doubt, it will be the first part of the hostname reported in the output of installation.
- Now you may optionally change the name by which this Rack is known within Console, before clicking on 'Import Rack' to include it within your Console account.

![](/assets/images/docs/import-rack/import-popup.png)

<div class="block-callout block-show-callout type-info" markdown="1">
Please note that importing the Rack in this manner will change the password as you set above.
</div>

## Importing a GCP/AWS Rack to your Console account without a Runtime Integration

If you have installed your Rack in GCP or do not have an AWS Runtime integration you can also easily import your Rack:

- From the 'Racks' page within the Console, click on the 'Add Rack' dropdown and select the 'Manual' option.
- Fill in your chosen name for the Rack, the hostname as reported by the installation, and the password you set during installation.
- Click on 'Add Rack' and the Rack will be imported to your Console account.

![](/assets/images/docs/import-rack/import-rack-manual.png)
