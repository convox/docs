---
title: "Debugging"
---

Convox makes it easy to debug your running applications.

#### convox logs

You can view all of the output of your application's processes using `convox logs`:

```
$ convox logs
2016-04-12 19:45:00 i-0234d285 service/web:RSPZQWVWGOP/5e3c8576b942 : 10.0.1.242 - - [12/Apr/2016:19:45:00 +0000] "GET / HTTP/1.1" 200 70 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36"
2016-04-12 19:45:00 i-0234d285 service/web:RSPZQWVWGOP/5e3c8576b942 : 10.0.1.242 - - [12/Apr/2016:19:45:00 +0000] "GET / HTTP/1.0" 200 70 0.0019
```

#### convox exec

You can get an active shell into a running process using `convox exec`. First, find the ID of the process:

```
$ convox ps
ID            NAME  RELEASE      SIZE  STARTED     COMMAND
310481bf223f  web   RSPZQWVWGOP  256   5 days ago  bin/web
5e3c8576b942  web   RSPZQWVWGOP  256   4 days ago  bin/web
```

```
$ convox exec 5e3c8576b942 bash
/app #
```
    
You can use this shell to examine your running application in detail.

#### convox instances ssh

You can also gain access to the individual instances that make up your Rack. First, find the ID of the instance:

```
$ convox instances
ID          AGENT  STATUS  STARTED      PS  CPU    MEM
i-0234d285  on     active  6 days ago   5   0.00%  12.36%
i-14cd0a89  on     active  7 hours ago  2   0.00%  5.18%
i-d8740c43  on     active  5 days ago   4   0.00%  21.04%
```

```    
$ convox instances ssh i-0234d285
[ec2-user@ip-10-0-3-209 ~]$ 
````

<div class="block-callout block-show-callout type-info" markdown="1">
Before you can use `convox instances ssh` the first time, you will need to run `convox instances keyroll` to load an SSH key onto your Rack instances. See [Keyroll](/docs/ssh-keyroll/) for details.
</div>
