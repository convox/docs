---
title: "Logs"
---

You can view the live logs for a Convox application using `convox logs`:

    $ convox logs
    2016-04-12 19:45:00 i-0234d285 service/web:RSPZQWVWGOP/5e3c8576b942 : 10.0.1.242 - - [12/Apr/2016:19:45:00 +0000] "GET / HTTP/1.1" 200 70 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36"
    2016-04-12 19:45:00 i-0234d285 service/web:RSPZQWVWGOP/5e3c8576b942 : 10.0.1.242 - - [12 Apr/2016:19:45:00 +0000] "GET / HTTP/1.0" 200 70 0.0019

### Additional Options

<table>
  <tr><th>Option</th><th>Description</th></tr>
  <tr><td><code>--filter=<b><i>POST</i></b></code></td><td>Return only the logs that match all the filters. Filters are case sensitive and non-alphanumeric terms must be inside double quotes.</td></tr>
  <tr><td><code>--since=<b><i>20m</i></b></code></td><td>Return logs starting this duration ago. Values are a duration like <code>10m</code> or <code>48h</code>.</td></tr>
</table>

You can tie all these together to generate a report from the logs from a single container over the last 2 days:

    $ convox ps
    ID            NAME  RELEASE      SIZE  STARTED     COMMAND
    310481bf223f  web   RSPZQWVWGOP  256   2 days ago  bin/web
    5e3c8576b942  web   RSPZQWVWGOP  256   2 days ago  bin/web

    $ convox logs --filter=GET --since=48h
    2016-04-25T21:12:02Z service/web:RSPZQWVWGOP/5e3c8576b942 10.0.3.243 - - [25/Apr/2016:21:12:02 +0000] "GET / HTTP/1.1" 200 226
    2016-04-25T21:12:03Z service/web:RSPZQWVWGOP/5e3c8576b942 10.0.3.243 - - [25/Apr/2016:21:12:02 +0000] "GET / HTTP/1.1" 200 226
    2016-04-25T21:12:04Z service/web:RSPZQWVWGOP/5e3c8576b942 10.0.3.243 - - [25/Apr/2016:21:12:02 +0000] "GET / HTTP/1.1" 200 226
    ...

### AWS Logs

We include events from AWS services in the output of `convox logs` to help you understand what's going on behind the scenes.

    2017-03-24T21:34:08Z system/aws/ecs service="legit-cv-soulshake-net-ServiceWopr-1TNE86TXRESN5" task="46a3ea28-4868-4620-b43c-5e4af6732219" status="STOPPED" container="wopr"
    2017-03-24T21:34:09Z system/aws/ecs service="legit-cv-soulshake-net-ServiceWopr-1TNE86TXRESN5" task="46a3ea28-4868-4620-b43c-5e4af6732219" status="RUNNING" container="wopr"

    2017-03-24T21:59:56Z system/aws/cfm resource="LaunchConfiguration" status="UPDATE_IN_PROGRESS" reason="Resource creation Initiated"
    2017-03-24T21:59:56Z system/aws/cfm resource="LaunchConfiguration" status="UPDATE_COMPLETE" reason=""
    2017-03-24T22:00:04Z system/aws/cfm resource="Instances" status="UPDATE_IN_PROGRESS" reason=""

These events can be useful for identifying issues with a deployment or an app. For example, when your Rack is in a "converging" state, i.e. some instances or ECS Tasks haven't stabilized yet, there are often AWS events in the app/Rack logs that will show a service crashing, a health check failing, or a placement error due to insufficient resources.

### Rack Logs

You can view the logs for a Convox Rack itself using the `convox rack logs` command:

    $ convox rack logs
    2017-03-24T21:59:57Z service/web:20170322201601/0b92eed79c1d ns=provider.aws at=fetchLogs start=1490392796065 events=0 state=success end=1490392796066 elapsed=225.020
    2017-03-24T22:16:15Z service/web:20170322201601/e378ddb167fd who="EC2/ASG" what="Terminating EC2 instance: i-02ce4f07da10a5333" why="At 2017-03-24T22:14:38Z a user request update of AutoScalingGroup constraints to min: 3, max: 1000, desired: 3 changing the desired capacity from 4 to 3.  At 2017-03-24T22:15:02Z an instance was taken out of service in response to a difference between desired and actual capacity, shrinking the capacity from 4 to 3.  At 2017-03-24T22:15:02Z instance i-02ce4f07da10a5333 was selected for termination."
