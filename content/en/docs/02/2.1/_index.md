---
title: "2.1 Drop-in replacements"
weight: 21
sectionnumber: 2.1
---

## Container images - Drop-in replacements


### Check for Chainguard images

Many container images for infrastructure applications are build to be drop-in replacements to their open-source counter parts.
Let's try this out!

List the repos of the default organization:

```bash
chainctl images repos list
```

To list the repos of a specific organization, use the `--parent` flag:

```bash
chainctl images repos list --parent=chainguard.dev-partners
```

Should show something like this:

```bash
[cgr.dev/chainguard.dev-partners]
├ [adoptium-jdk]
├ [adoptium-jre]
├ [curl]
...
```

List `python` image tags that we can use in the lab organization:

```bash
chainctl images list --repo=python
```


### Replace image

What image should we use? infra structure app but easy to check...


#### Pulling images

Chainguard images can be pulled the the docker pull command. Here we are pulling the Tag `latest-glibc`:

```bash
docker pull cgr.dev/chainguard/git:latest-glibc
```
If you don't specify a tag, docker will pull the latest image. Also please note, that not all images are available for this workshop.

One can also pull specific image digests, the advantage of this being reproducibility, as it will ensure that you are using the same image each time (versus the tag that may receive updates, see [section 1.3](/docs/01/1.3/)).
A specific digest can be pulled via this docker command:

```bash
docker pull cgr.dev/chainguard/git@sha256:f6658e10edde332c6f1dc804f0f664676dc40db78ba4009071fea6b9d97d592f
```


### Retrieve SBOM

TODO

https://edu.chainguard.dev/chainguard/chainguard-images/how-to-use/retrieve-image-sboms/


### Test image

