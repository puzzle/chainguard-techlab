---
title: "2.3 Scanning and Verification"
weight: 23
sectionnumber: 2.3
---

## Container images - Scanning and Verification

Chainguard images are built to have as few known vulnerabilities (CVEs) as possible. Every image is signed and comes with a build-time SBOM (Software Bill of Materials). The signatures and SBOMs are created with [Sigstore](https://www.sigstore.dev/) and can be verified with `cosign`.

In this lab we scan the images from [lab 2.2](/docs/02/2.2/), verify the signature of a Chainguard image and download its SBOM.


### Install the tools

Install [Grype](https://github.com/anchore/grype) as vulnerability scanner, see the [Grype tutorial](https://edu.chainguard.dev/chainguard/chainguard-images/staying-secure/working-with-scanners/grype-tutorial/):

```bash
curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | sh -s -- -b ~/.local/bin
```

Install [cosign](https://docs.sigstore.dev/cosign/system_config/installation/), for example with Homebrew:

```bash
brew install cosign
```

We also use [jq](https://jqlang.org/) to format JSON output.


### Scan images

Scan the upstream image and the Chainguard image from [lab 2.2](/docs/02/2.2/):

```bash
grype python-app:upstream
grype python-app:cg
```

If you don't want to install Grype, you can use the Grype container:

```bash
docker run --rm -it --volume /var/run/docker.sock:/var/run/docker.sock anchore/grype python-app:upstream
docker run --rm -it --volume /var/run/docker.sock:/var/run/docker.sock anchore/grype python-app:cg
```

Compare the results:

* How many vulnerabilities are found in each image?
* Where do the vulnerabilities of the upstream image come from (base image, OS packages, Python packages)?

To only show vulnerabilities for which a fix is available, use:

```bash
grype python-app:upstream --only-fixed
```

{{% alert title="Note" color="info" %}}
Other scanners like [Trivy](https://trivy.dev/) or [Docker Scout](https://docs.docker.com/scout/) also support Chainguard images. Scanners need to know the Wolfi security advisories to correctly report the vulnerabilities of Chainguard images.
{{% /alert %}}


### Verify the image signature

Every Chainguard image is signed during the build. With `cosign` we can verify that an image was built by Chainguard and was not modified afterwards.

The public images in `cgr.dev/chainguard` are signed by the Chainguard release workflow on GitHub:

```bash
cosign verify \
  --certificate-oidc-issuer=https://token.actions.githubusercontent.com \
  --certificate-identity=https://github.com/chainguard-images/images/.github/workflows/release.yaml@refs/heads/main \
  cgr.dev/chainguard/python:latest | jq
```

The output shows the verified claims of the signature, for example the digest of the image.

Images of our organization are signed with a different identity, see [Verifying Chainguard Containers](https://edu.chainguard.dev/chainguard/chainguard-images/how-to-use/verifying-chainguard-images-and-metadata-signatures-with-cosign/) for details.


### Retrieve the SBOM

The SBOM lists every package included in the image, including the version and license. It is attached to the image as a signed attestation.

Download the SBOM of the Chainguard Python image:

```bash
cosign download attestation \
  --platform=linux/amd64 \
  --predicate-type=https://spdx.dev/Document \
  cgr.dev/chainguard/python:latest | jq -r .payload | base64 -d | jq .predicate > python-sbom.json
```

List the packages included in the image:

```bash
jq -r '.packages[] | "\(.name) \(.versionInfo)"' python-sbom.json
```

You can also scan the SBOM directly, without pulling the image:

```bash
grype sbom:./python-sbom.json
```

For more information, see [Retrieve image SBOMs](https://edu.chainguard.dev/chainguard/chainguard-images/how-to-use/retrieve-image-sboms/).
