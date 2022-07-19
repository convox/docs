---
title: "CLI Installation"
order: 200
---

The `convox` command line tool makes building, configuring, scaling, and securing your apps easy.

Here are some of the highlights:

* `convox start` - A single command to start your development environment
* `convox deploy` - A single command to deploy your application to a Rack
* `convox rack update` - A single command to deliver API and infrastructure improvements to a Rack

You can install it via the command line:

## OS X

### x86_64 / amd64

    $ curl -L https://convox.s3.amazonaws.com/cli/darwin/convox -o /tmp/convox
    $ sudo mv /tmp/convox /usr/local/bin/convox
    $ sudo chmod 755 /usr/local/bin/convox

### arm64

```html
    $ curl -L https://convox.s3.amazonaws.com/cli/darwin/convox-arm64 -o /tmp/convox
    $ sudo mv /tmp/convox /usr/local/bin/convox
    $ sudo chmod 755 /usr/local/bin/convox
```

## Linux


### x86_64 / amd64

```html
    $ curl -L https://convox.s3.amazonaws.com/cli/linux/convox -o /tmp/convox
    $ sudo mv /tmp/convox /usr/local/bin/convox
    $ sudo chmod 755 /usr/local/bin/convox
```

### arm64

```html
    $ curl -L https://convox.s3.amazonaws.com/cli/linux/convox-arm64 -o /tmp/convox
    $ sudo mv /tmp/convox /usr/local/bin/convox
    $ sudo chmod 755 /usr/local/bin/convox
```

## Windows

    $ curl -L https://convox.s3.amazonaws.com/cli/windows/convox.exe -O

# Next steps

## Logging in to Console

After installing Convox, you'll need to `convox login`:

    $ convox login console.convox.com
    Password: <your Console API key>
    Logged in successfully.

## Updating the CLI

To update the CLI you can run `convox update`:

    $ convox update
    Updating convox to 20211231000000: OK
