---
title: "Monitoring an App"
---

Once you've got an app running on Convox, setting up monitoring can help you keep it running smoothly. As a starting point, we'd recommend the following baseline monitoring coverage when using Convox:

1. [Set up availability monitoring](#set-up-availability-monitoring) to be notified if your app stops responding.
1. [Create a CloudWatch Dashboard](#create-a-cloudwatch-dashboard) to help track multiple metrics of your app.
1. [Create CloudWatch Alarms](#create-cloudwatch-alarms) to be notified of critical changes to your app's metrics.
1. [Track CloudWatch Metrics](#track-cloudwatch-metrics) to keep tabs on your database, etc.

## Set up availability monitoring

Let's say your app runs at example.com. By having a third-party service like [Pingdom](https://www.pingdom.com) send an HTTP request to example.com at frequent intervals, you can ensure that you'll be notified as soon as part of your system fails. Whether you are having problems with your infrastructure, database, web application, or DNS, simple availability monitoring can be used to help you catch many major problems quickly.

We recommend using an availability monitoring service not run on AWS so that an AWS outage doesn't leave you blind to the problem.

### Further Resources

[Datadog: Introducing Availability Monitoring](https://www.datadoghq.com/blog/introducing-availability-monitoring/)
[NewRelic: Downtime alert settings](https://docs.newrelic.com/docs/alerts/alert-policies/downtime-alerts/downtime-alert-settings)
[Pingdom: How to set up an uptime (HTTP) check](https://help.pingdom.com/hc/en-us/articles/203679631-How-to-set-up-an-uptime-HTTP-check)

## Create a CloudWatch Dashboard

Sometimes it's hard to diagnose a problem without stepping back and looking at the bigger picture. By creating a CloudWatch Dashboard that tracks multiple metrics of your application over time, you can more easily spot abnormalities, warning signs, bottlenecks, and more.

To set up a CloudWatch Dashboard with a single widget that tracks CPU and memory utilization of a given app, follow these instructions:

1. Visit your CloudWatch console and navigate to Dashboards.
1. Click the "Create Dashboard" button.
1. Enter a name for the new dashboard. Consider `<rack-name>-<app-name>`, e.g. production-site.
1. Select the "Metric graph" widget type.
1. From the ECS Metrics category, select "ClusterName,ServiceName."
1. Select the CPUUtilization and MemoryUtilization graphs for the service you'd like to monitor. For example, if you wanted to monitor the `web` process of a `foo-bar` app that is running in your `production` Rack, the service name would look like `production-foo-bar-ServiceWeb-B3YF4EMQCN89`.
1. Click the "Create widget" button.
1. Consider resizing the widget (by dragging the bottom right corner) for better readability.

## Create CloudWatch Alarms

If, in addition to being able to monitor your new CloudWatch Dashboard, you'd like to be notified any time a certain metric enters a specific range of values, you can set up a CloudWatch Alarm. For example, if you want to be emailed when your app exceeds an average CPU utilization of 95% for a period of 30 minutes, you could set up that CloudWatch Alarm by taking the following steps:

1. Visit your CloudWatch console and navigate to Alarms.
1. Click the "Create Alarm" button.
1. From the ECS Metrics category, select "ClusterName,ServiceName."
1. Select CPUUtilization for the service you'd like to monitor. For example, if you wanted to monitor the `web` process of a `foo-bar` app that is running in your `production` Rack, the service name would look like `production-foo-bar-ServiceWeb-B3YF4EMQCN89`.
1. Click the "Next" button to move to the "Define Alarm" step.
1. In the Alarm Threshold section, specify a name like "High CPU utilization," and then fill out the inputs to complete the phrase, "Whenever CPUUtilization is > 90 for 2 consecutive periods."
1. In the Actions section, next to "Send notification to:," click "New list," then enter a new notification list topic name and the email addresses that should be notified. Note that this will result in the creation of a new SNS Topic.
1. In the Alarm Preview section, confirm that Period is set to "15 Minutes" and the Statistic is "Average."
1. Confirm that you see your new Alarm listed with State "OK."

<div class="block-callout block-show-callout type-info" markdown="1">
Keep in mind that a notification email sent by a CloudWatch Alarm can be used as a trigger for other services like PagerDuty. See their [Email Integration Guide](https://www.pagerduty.com/docs/guides/email-integration-guide/) for more information.
</div>

## Track CloudWatch Metrics

Once you've [created a CloudWatch Dashboard](#create-a-cloudwatch-dashboard), you can configure additional CloudWatch Metrics and add them to a dashboard as widgets. Let's say you want to keep an eye on the swap usage of your PostgreSQL RDS instance, so you know when to upgrade to a larger instance type. You could add a widget displaying that usage to a dashboard by following these steps:

1. Visit your CloudWatch console and navigate to Metrics.
1. From the RDS Metrics category, select "Per-Database Metrics."
1. Select the SwapUsage Metric Name for the desired RDS instance.
1. Click the "Add to Dashboard" button.
1. In the "Add to Dashboard" dialog, select the desired dashboard and click the "Add to Dashboard" button.
1. Consider resizing the widget (by dragging the bottom right corner) for better readability.
