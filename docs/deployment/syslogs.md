---
title: "Syslogs"
---

Native convox logging information can be found [here](/management/logs)...

### Routing Logs to a 3rd Party

Convox supports routing your logs to any third party that can accept data from a syslog forwarder.

You can create a syslog forwarder with the following command:

    $ convox rack resources create syslog [destination]

For example:

    $ convox rack resources create syslog Url=tcp+tls://logs1.papertrailapp.com:12345
    Creating syslog-3785 (syslog)... CREATING

Additionally, you can pass a `Format` parameter to amend the default logging format to match the requirements of your receiving service.

For example for Datadog:

    $ convox rack resources update syslog-3785 Format="123457890abcdef1234567890 <22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} - - [metas ddsource=\"{GROUP}\" ddtags=\"container_id:{CONTAINER}\"] {MESSAGE}" Url=tcp+tls://intake.logs.datadoghq.com:10516

The `{DATE} / {GROUP} / {SERVICE} / {CONTAINER} / {MESSAGE}` variables within the format are dynamically replaced with the correct values and injected into the log output in the required format.


You can view the forwarder setup at any time using `convox rack resources info`

    $ convox rack resources info syslog-3785
    Name     syslog-3785
    Status   running
    Options  Format=<22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} - - {MESSAGE}
             Url=tcp+tls://logs1.papertrailapp.com:12345
    URL      tcp+tls://logs1.papertrailapp.com:12345

In order to start sending logs from an app to the forwarder you need to link it to the app with `convox rack resources link`

    $ convox rack resources link syslog-3785 --app example
    Linked syslog-3786 to example

Some common 3rd party logging services include:

* [Papertrail](https://papertrailapp.com)

  * Go to [Log Destinations](https://papertrailapp.com/account/destinations) to get the forwarding destination
  * The Default `Format` is suggested. (`<22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} - - {MESSAGE}`)

* [Datadog](https://www.datadoghq.com)

  * Go to [Syslog-Ng Integration](https://docs.datadoghq.com/integrations/syslog_ng/?tab=datadogussite) to check the forwarding destination.  This currently differs between the US site (`intake.logs.datadoghq.com:10516`) and the EU site (`tcp-intake.logs.datadoghq.eu:443`)
  * Suggested `Format="INSERT-YOUR-API-KEY-HERE <22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} - - [metas ddsource=\"{GROUP}\" ddtags=\"container_id:{CONTAINER}\"] {MESSAGE}"` where you replace `INSERT-YOUR-API-KEY-HERE` with your Datadog API key 😉

* [LogDNA](https://logdna.com/)

  * Select Syslog as your provider and then select `provision a syslog port` to get the forwarding destination
  * Suggested `Format="<22>1 {DATE} {GROUP} {SERVICE} ${CONTAINER} [logdna@48950 key=\"INSERT-YOUR-INGESTION-KEY-HERE\"] {MESSAGE}` and replacing `INSERT-YOUR-INGESTION-KEY-HERE` with your LogDNA Ingestion key 😉

* [Loggly](https://www.loggly.com/)

  * Syslog configuration is documented [here](https://www.loggly.com/docs/syslog-ng-manual-configuration/)
  * Suggested `Format="<22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} [INSERT-YOUR-CUSTOMER-TOKEN-HERE@41058 tag=\"INSERT-YOUR-TAG-HERE\" ] {MESSAGE}"` and replacing `INSERT-YOUR-CUSTOMER-TOKEN-HERE` with your Loggly [customer token](https://www.loggly.com/docs/customer-token-authentication-token/) and `INSERT-YOUR-TAG-HERE` with a [tag](https://www.loggly.com/docs/tags/) that describes your source


