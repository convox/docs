---
title: "Papertrail"
---

You can easily configure your application to to forward its logs to Papertrail.

## Set up a Syslog Forwarder

* Go to [Log Destinations](https://papertrailapp.com/account/destinations) and click **Create Log Destination**

* Make sure that *TLS encryption* under *TCP* is checked

* Note the hostname and port of your syslog destination

## Create a Syslog Resource

Use the syslog endpoint from the previous step to create a `syslog` resource:

    $ convox rack resources create syslog Url=tcp+tls://logs1.papertrailapp.com:12345
    Creating syslog-3785 (syslog)... CREATING
    
## Link the Resource

Link the resource to any application whose logs you would like to forward to Papertrail.

    $ convox rack resources link syslog-3785 --app example
    Linked syslog-3786 to example
