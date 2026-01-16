---
title: "1.2 Different Chainguard images"
weight: 12
sectionnumber: 1.2
description: >
  Introduction to Chainguard.
---

## Different Chainguard images

Chainguard offers container images, virtual machine (VM) images, helm charts and libraries.

### Container images

Chainguard offers a wide variety of container images. Let's check out the different types:

##### FIPS images

![Chainguard library FIPS container image](iamguarded.png)

In the image directory of Chainguard some images are tagged with ´FIPS validated´. These images comply strict cryptographic requirements to be used for the U.S. federal agencies, defense contractors, and regulated industries.

##### iamguarded images

![Chainguard library iamguarded container image](iamguarded.png)

If the image has the "-iamguarded" in its name, it is composed to be a base image for a Helm Chart. 


#### Tags and Digests

Image Tags stay the same as official version, but "digest" might update to account for security patches.
when the image tag version is bumped, the previous security patches are either amended or rebased on the new commit.

