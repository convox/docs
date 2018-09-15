---
title: "HIPAA Compliance"
---

At Convox we understand that for HIPAA-compliant teams, having full control of your customer data is the most important factor when choosing a new technology. We built Convox as a single-tenant solution and made privacy and security fundamental design goals to guarantee that you know exactly who and what has access to your data.

Convox offers HIPAA customer references on request.

## Steps to ensure HIPAA Compliance with Convox

### Install a private rack on dedicated hardware

Use [EC2 Dedicated Instances](https://aws.amazon.com/ec2/purchasing-options/dedicated-instances/) and isolate them from the Internet.

```
convox rack install aws Private=Yes Tenancy=dedicated
```

See our [Private Networking doc](/docs/private-networking/) for more details.

### Run a private Convox console

Instead of managing your teams and racks through the Convox-hosted console at [console.convox.com](https://console.convox.com), deploy a self-hosted console into a rack in your own AWS account. The self-hosted console is available to all [Enterprise Plan](/pricing) customers.

### Sign a BAA with AWS

If you run a self-hosted console, you leave us with no conceivable way to access your PHI, and therefore place Convox out of scope of the BAA requirement. You will still need to [sign a Business Associate Agreement (BAA) with AWS](https://aws.amazon.com/compliance/hipaa-compliance/).

### Review AWS HIPAA Compliance guidelines

AWS makes a clear distinction between its responsibilities ("security _of_ the cloud") and customer responsibilities ("security _in_ the cloud"). For a more explicit delineation, take a look at their [Shared Responsibility Model](https://aws.amazon.com/compliance/shared-responsibility-model/).

Along similar lines, you should be mindful of which AWS services have been deemed HIPAA-eligible:

* DynamoDB
* Elastic Block Store (EBS)
* Elastic Cloud Compute (EC2)
* Elastic Load Balancing (ELB)
* Elastic MapReduce (EMR)
* Glacier
* Redshift
* Relational Database Service (RDS), MySQL and Oracle engines only
* Simple Storage Service (S3)

While these are the only services which can process, store, and transmit PHI according to HIPAA regulations, other AWS services can still be used so long as they do not come in direct contact with PHI. For example, ECS can be used when all PHI is actually processed on EC2.

#### Further Reading

[AWS HIPAA Compliance Overview](https://aws.amazon.com/compliance/hipaa-compliance/)
[AWS HIPAA Compliance Whitepaper](https://d0.awsstatic.com/whitepapers/compliance/AWS_HIPAA_Compliance_Whitepaper.pdf)
[AWS HIPAA Compliance FAQ, Part One](https://blogs.aws.amazon.com/security/post/Tx3TGE4YTL0XK5Z/Frequently-Asked-Questions-About-HIPAA-Compliance-in-the-AWS-Cloud)
[AWS HIPAA Compliance FAQ, Part Two](https://blogs.aws.amazon.com/security/post/Tx3FDPNNKZ5XFEE/Frequently-Asked-Questions-About-HIPAA-Compliance-in-the-AWS-Cloud-Part-Two)
