---
title: "Dockerfile"
---

The `Dockerfile` describes the steps you need to build and run your application.

* `FROM` defines the base image for your application.
* `COPY` moves files from the local directory into the image.
* `RUN` executes a command.
* `ENTRYPOINT` defines a command prefix that should be prepended to any command run on this image.
* `CMD` defines the default command to start application.
* `ARG` allows you to specify build-time variables.

### ARG

Convox respects the `ARG` Dockerfile directive, allowing you to specify build-time variables to be populated.

This is useful for creating dynamic build environments, allowing you to do things like:
- building differently in production and development environments
- specifying environment variables that should be used during a build, but should not be present in the `Dockerfile` itself or in the resulting image.

<div class="block-callout block-show-callout type-warning" markdown="1">
Warning: It is not recommended to use build-time variables for passing secrets. Build-time variable values are visible to any user of the image with the `docker history` command.
</div>

You can declare a build argument in the `Dockerfile`, with either a default value or an empty one:

```
ARG BUNDLE_WITHOUT="development:test"
ARG RAILS_ENV
```

You can send a value (or set an empty value, as below) to be applied during remote builds on your Rack by setting it with `convox env`:

```
$ convox env set BUNDLE_WITHOUT=none --promote
$ convox deploy
```

## See also

- [Builds](/docs/builds)
- [Docker: ARG](https://docs.docker.com/engine/reference/builder/#arg)
- [Docker: -\-build-arg](https://docs.docker.com/engine/reference/commandline/image_build/#/set-build-time-variables---build-arg)
- [Docker: Build your own image](https://docs.docker.com/engine/getstarted/step_four/)
