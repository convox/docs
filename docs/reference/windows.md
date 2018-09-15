---
title: Windows
---

## Install Docker

To install Docker, follow the [Get started with Docker for Windows](https://docs.docker.com/docker-for-windows/) guide.

<div class="alert alert-warning">
You do not need to enable <a href="https://docs.docker.com/docker-for-windows/#/shared-drives">Shared Drives</a>. Convox does not use shared drives due to compatibility issues between the host NTFS and container filesystems.
</div>

## Install Git

To get a `git` and `bash` environment, install <a href="https://git-for-windows.github.io/">Git for Windows</a>.

<div class="alert alert-warning">
The container filesystem expects Unix-style line endings, not Windows. Select the <b>Checkout as-is, commit Unix-style line endings</b> installer option so `git` doesn't modify line-endings.
<br /><br />
You can also run `git config --global core.autocrlf input`.
</div>

## Install Convox

To get the `convox` command, download the [Convox Windows installer](https://dl.equinox.io/convox/convox/stable).

This is a per-user installer that does not require Administrator priveledges. It is installed into `C:\Users\name\AppData\Local\Programs\convox`.

<div class="alert alert-warning">
Windows 10 warns that the Convox installer is not signed. Click "more info" then the "Run anyway" option.
</div>

## Terminal Support

The `convox` CLI works best with Git BASH. Powershell and Command Prompt also work but may display graphical errors.
