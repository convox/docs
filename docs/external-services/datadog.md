---
title: "Datadog"
---

You can add operational visibility to your Convox environments with [Datadog](https://docs.datadoghq.com/integrations/convox).

## Sign up for Datadog

If you don’t have an account already, [sign up for Datadog](https://app.datadoghq.com/signup). You need an API key that enables you to send data from Convox into Datadog services.

## Deploy the Datadog Agent

You can deploy the Datadog Agent as a Convox app by using a `convox.yml` manifest.

```
environment:
        - DD_API_KEY=<YOUR_API_KEY>
        - DD_APM_ENABLED=true
        - DD_PROCESS_AGENT_ENABLED=true
        - DD_AC_EXCLUDE="name:datadog-agent name:ecs-agent"
services:
  datadog-agent:
    agent:
      enabled: true
      ports:
        - 8125/udp
        - 8126/tcp
      image: datadog/agent:latest
      privileged: true
      scale:
        cpu: 128
        memory: 128
      volumes:
        - /sys/fs/cgroup/:/host/sys/fs/cgroup/
        - /proc/:/host/proc/
        - /var/run/docker.sock:/var/run/docker.sock
```

Run `convox deploy` to deploy Datadog Agent into ECS.

### Application Metrics

To forward application metrics to Datadog, you need the host IP address. 

You can get that by using the following:

    $ ip route list match 0/0 | awk '{print $3}'

## Logging Endpoint

To integrate Datadog as a logging endpoint with our [Syslog](/deployment/syslogs) resource:

  * Go to [Syslog-Ng Integration](https://docs.datadoghq.com/integrations/syslog_ng/?tab=datadogussite) to check the forwarding destination.  This currently differs between the US site (`intake.logs.datadoghq.com:10516`) and the EU site (`tcp-intake.logs.datadoghq.eu:443`).
  * Suggested `Format="INSERT-YOUR-API-KEY-HERE <22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} - [metas ddsource=\"{GROUP}\" ddtags=\"container_id:{CONTAINER}\"] {MESSAGE}"` where you replace `INSERT-YOUR-API-KEY-HERE` with your Datadog API key 😉

For example:

    $ convox rack resources create syslog Format="123457890abcdef1234567890 <22>1 {DATE} {GROUP} {SERVICE} {CONTAINER} - - [metas ddsource=\"{GROUP}\" ddtags=\"container_id:{CONTAINER}\"] {MESSAGE}" Url=tcp+tls://intake.logs.datadoghq.com:10516

Link the created Syslog resource to your app:

    $ convox rack resources link syslog-3785 --app example
