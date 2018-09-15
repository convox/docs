---
title: "Installing the Convox CLI"
order: 50
---

The `convox` command line tool makes building, configuring, scaling, and securing your apps easy.

Here are some of the highlights:

* `convox start` - A single command to start your development environment
* `convox deploy` - A single command to deploy your application to a Rack
* `convox rack update` - A single command to deliver API and infrastructure improvements to a Rack

You can install it via the command line:

## OS X

    $ curl -L https://convox.com/cli/osx/convox -o /tmp/convox
    $ sudo mv /tmp/convox /usr/local/bin/convox
    $ sudo chmod 755 /usr/local/bin/convox

#### Homebrew

Alternatively, on OSX you can also install via Homebrew:

    $ brew install convox

## Linux

    $ curl -L https://convox.com/cli/linux/convox -o /tmp/convox
    $ sudo mv /tmp/convox /usr/local/bin/convox
    $ sudo chmod 755 /usr/local/bin/convox

## Windows

    $ curl -L https://convox.com/cli/windows/convox.exe -O

# Next steps

## Logging in to the CLI

After installing Convox, you'll need to `convox login`:

    $ convox login console.convox.com
    Password: <your Console API key>
    Logged in successfully.

## Updating the CLI

To update the CLI you can run `convox update`:

    $ convox update
    Updating convox to 20211231000000: OK

## Further reading

For details about how to use `convox` on the command line, see the [CLI reference](/docs/cli/).
