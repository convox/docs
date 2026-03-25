---
title: "Mezmo"
description: "Forward application logs to Mezmo (formerly LogDNA) via syslog"
---

# Mezmo

You can configure your application to forward its logs to Mezmo (formerly LogDNA).

## Set up a Syslog Forwarder

* Go to [https://app.mezmo.com/pages/add-source](https://app.mezmo.com/pages/add-source) and select Syslog as your provider

* Use the provided url as the syslog resource Url (with `tcp+tls://`)

* The suggested Format is: `Format="<22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} {MESSAGE}"`

## Create a Syslog Resource

Use the syslog endpoint and Format from the previous step to create a `syslog` resource:

```bash
$ convox rack resources create syslog Url=tcp+tls://syslog-a.mezmo.com:6514 Format="<22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} {MESSAGE}
Creating syslog-3785 (syslog)... CREATING
```

## Link the Resource

Link the resource to any application whose logs you would like to forward to Mezmo.

```bash
$ convox rack resources link syslog-3785 --app example
Linked syslog-3786 to example
```

## See Also

- [Syslogs](/deployment/syslogs)
- [Papertrail](/integrations/papertrail)
- [Loggly](/integrations/loggly)
- [Logs](/management/logs)
