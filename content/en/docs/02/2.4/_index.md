---
title: "2.4 Keeping images up to date"
weight: 24
sectionnumber: 2.4
---

## Container images - Keeping images up to date

Chainguard images are rebuilt as often as needed, depending on upstream releases and CVE fixes. Even if the version tag stays the same, the underlying digest changes with every new build.


### Tags and digests

A tag like `python:3.12` is a moving pointer. Chainguard uses different kinds of tags:

| Tag | Example | Meaning |
| --- | --- | --- |
| Major/minor version | `python:3.12` | always points to the latest build of the latest 3.12 release |
| Full version | `python:3.12.11` | points to the latest build of this release |
| Epoch | `python:3.12.11-r1` | the build number of this release, increased when Chainguard rebuilds the package, for example to fix a vulnerability |
| Digest | `python@sha256:...` | the immutable identifier of exactly one image |

Only the latest release of a version stream receives updates. As soon as `3.12.12` is available, `3.12.11` is no longer rebuilt.

Therefore the recommendation is:

* reference the version stream you want to use, for example `3.12`
* pin the image with its digest to get reproducible builds
* update the digest automatically and regularly


### Get the digest of an image

Pull the image and show its digest:

```bash
docker pull cgr.dev/chainguard/python:latest
docker inspect --format '{{index .RepoDigests 0}}' cgr.dev/chainguard/python:latest
```

You can also look up the digest without pulling the image:

```bash
docker buildx imagetools inspect cgr.dev/chainguard/python:latest
```

The digests and tags of the images of our organization are also shown in the [Chainguard console](https://console.chainguard.dev/).

Pin the base images in `Dockerfile.cg` from [lab 2.2](/docs/02/2.2/) with the digest:

```Dockerfile
FROM cgr.dev/chainguard/python:latest-dev@sha256:<digest> AS builder
...
FROM cgr.dev/chainguard/python:latest@sha256:<digest>
```

The tag is kept for readability, but the digest decides which image is used.


### Automate digest updates

Pinned digests need to be updated, otherwise you miss the security fixes. Chainguard provides [digestabot](https://github.com/chainguard-dev/digestabot), a GitHub Action that looks for image digests in your repository, checks if a newer digest is available for the tag and creates a pull request with the updates.

Example workflow `.github/workflows/digestabot.yaml`:

```yaml
name: Image digest update

on:
  workflow_dispatch:
  schedule:
    - cron: "0 1 * * *"

permissions:
  contents: write
  pull-requests: write
  id-token: write

jobs:
  image-update:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: chainguard-dev/digestabot@v1
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
```

If you are not using GitHub Actions, [Renovate](https://docs.renovatebot.com/docker/#digest-pinning) can pin and update digests as well.

For more information, see [Using digestabot](https://edu.chainguard.dev/chainguard/chainguard-images/staying-secure/updating-images/digestabot/).
