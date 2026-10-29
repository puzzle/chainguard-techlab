---
title: "6.1 Pull-through cache"
weight: 61
sectionnumber: 6.1
description: >
  Using Chainguard images through your own registry.
---

## Pull-through cache

For this techlab we pull the images directly from `cgr.dev`. If you are using Chainguard for your company, it is best practice to set up a pull-through cache in the registry of your choice (for example Harbor, Artifactory, Nexus or Cloudsmith).

The advantages of a pull-through cache are:

* builds and clusters only need access to your own registry, not to `cgr.dev`
* only one technical account is needed to authenticate against Chainguard
* images are cached, which reduces the dependency on `cgr.dev` and speeds up pulls
* existing policies of your registry (scanning, retention, access control) also apply to Chainguard images

This lab is optional and requires a registry in which you can configure a proxy or remote repository.


### Create a pull token

The registry needs credentials to pull from `cgr.dev`. Create a pull token for our organization:

```bash
chainctl auth pull-token --parent=puzzle-partner.com --ttl=720h
```

The command prints a username and a password. Store them safely, they are needed in the next step.

{{% alert title="Note" color="info" %}}
Choose the TTL of the token according to your requirements and renew the token before it expires.
{{% /alert %}}


### Configure the registry

The exact steps depend on your registry. In general you have to:

1. create a remote/proxy repository with the upstream URL `https://cgr.dev`
2. configure the username and password of the pull token
3. pull the images through your registry

Example with Harbor:

1. In "Registries", create a new endpoint with the provider "Docker Registry", the URL `https://cgr.dev` and the pull token credentials
2. Create a new project, enable "Proxy Cache" and select the endpoint
3. Pull the image through the project:

```bash
docker pull <harbor-host>/<project>/puzzle-partner.com/python:latest
```

Guides for the different registries can be found in the [Chainguard documentation](https://edu.chainguard.dev/chainguard/chainguard-registry/pull-through-guides/).
