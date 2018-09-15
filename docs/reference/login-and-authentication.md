---
title: "Login and Authentication"
---

## Login, authorization, and keyroll FAQ

### What's the difference between `convox login` and `convox switch`?

`convox login` tells your CLI what hostname to send requests to. The target can be Console or a Rack. This is stored in [`~/.convox/host`](/docs/cli-config-files/#configuration-files).

When the CLI is logged into Console, `convox switch` tells Console which Rack to proxy your commands to. `convox switch` is only available with Racks that have been installed via Console, or which have been installed via CLI and manually added to the Console web interface.

### Difference between `convox login console.convox.com` and `convox login <rack hostname>`?

When you're logged into Console, Console acts as a proxy for your active Rack. This is only available for Racks that have been installed via the Console web interface (or installed via CLI and manually added to the Console web interface). If you installed a Rack via `convox rack install` command, and want to log into it directly, you have to log into it by its hostname, e.g. `convox login <hostname> --password <password>` as they appear in `~/.convox/auth`.

You can switch between Racks which have been installed or manually added to Console with `convox switch <org>/<rackname>`.

Since Console is a proxy for a Rack, Console stores the encrypted original Rack password (generated automatically when you install a Rack via Console) and authenticates using its own API key (which you can regenerate on your Account Settings screen).

### How do I log out (or switch accounts) using the CLI?

Being "logged in" simply means that `~/.convox/auth` contains the hostname in `~/.convox/host` along with the corresponding Console or Rack API key (referred to in some places as a "Rack password").

In other words: if `auth` contains the correct hostname + API key pair corresponding to the hostname in `host`, then you're "logged in."

To switch accounts, run `convox login console.convox.com --password=CONVOX_API_KEY`, replacing `CONVOX_API_KEY` with your user API key, which can be regenerated in [your Convox account settings](https://console.convox.com/grid/user/profile).

## See also

- [CLI configuration files](/docs/cli-config-files/)
- [CLI environment variables](/docs/cli-environment-variables/)
