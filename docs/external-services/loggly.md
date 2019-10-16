---
title: "Loggly"
---

You can easily configure your application to to forward its logs to Loggly.

## Set up a Syslog Forwarder

* The Syslog configuration is documented [here](https://www.loggly.com/docs/syslog-ng-manual-configuration/)

* The suggested Format is: `Format="<22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} [INSERT-YOUR-CUSTOMER-TOKEN-HERE@41058 tag=\"INSERT-YOUR-TAG-HERE\" ] {MESSAGE}"` and replacing `INSERT-YOUR-CUSTOMER-TOKEN-HERE` with your Loggly [customer token](https://www.loggly.com/docs/customer-token-authentication-token/) and `INSERT-YOUR-TAG-HERE` with a [tag](https://www.loggly.com/docs/tags/) that describes your source

## Create a Syslog Resource

Use the syslog endpoint from the previous step to create a `syslog` resource:

    $ convox rack resources create syslog Url=tcp+tls://logs-01.loggly.com:514 Format="<22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} [69af9e44-e002-4f67-8777-7b3b93ab0297@41058 tag=\"example-app\" ] {MESSAGE}"
    Creating syslog-3785 (syslog)... CREATING
    
## Link the Resource

Link the resource to any application whose logs you would like to forward to Papertrail.

    $ convox rack resources link syslog-3785 --app example
    Linked syslog-3786 to example
