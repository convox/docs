---
title: "LogDNA"
---

You can easily configure your application to to forward its logs to LogDNA.

## Set up a Syslog Forwarder

* Go to [https://app.logdna.com/pages/add-source](https://app.logdna.com/pages/add-source) and select Syslog as your provider

* Use the provided url as the syslog resource Url (with `tcp+tls://`)

* The suggested Format is: `Format="<22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} {MESSAGE}"`

## Create a Syslog Resource

Use the syslog endpoint and Format from the previous step to create a `syslog` resource:

    $ convox rack resources create syslog Url=tcp+tls://syslog-a.logdna.com:6514 Format="<22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} {MESSAGE}
    Creating syslog-3785 (syslog)... CREATING
    
## Link the Resource

Link the resource to any application whose logs you would like to forward to Papertrail.

    $ convox rack resources link syslog-3785 --app example
    Linked syslog-3786 to example
