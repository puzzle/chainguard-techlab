---
title: "3.4 Scanning and Verification"
weight: 34
sectionnumber: 3.4
---

## Container images - Scanning and Verification

Chainguard images are built to have as few known vulnerabilities (CVEs) as possible. Every image is signed and comes with a build-time SBOM (Software Bill of Materials). The signatures and SBOMs are created with [Sigstore](https://www.sigstore.dev/) and can be verified with `cosign`.

In this lab we scan the images from [lab 3.3](/docs/03/3.3/), verify the signature of a Chainguard image and scan its SBOM.


### Tools

We use the following tools. All of them are available as official container images, so you can run them with Docker without installing them. For each step, the command is shown for Docker and for a local installation.

| Tool | Purpose | Container image |
| --- | --- | --- |
| [Grype](https://github.com/anchore/grype) | vulnerability scanner | `anchore/grype` |
| [cosign](https://docs.sigstore.dev/) | verify signatures and attestations | `ghcr.io/sigstore/cosign/cosign` |
| [jq](https://jqlang.org/) | format and filter JSON output | `ghcr.io/jqlang/jq` |


#### Install the tools locally (optional)

Skip this section if you use the container images.

Install Grype, see the [Grype tutorial](https://edu.chainguard.dev/chainguard/chainguard-images/staying-secure/working-with-scanners/grype-tutorial/):

```bash
curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | sh -s -- -b ~/.local/bin
```

Install [cosign](https://docs.sigstore.dev/cosign/system_config/installation/), for example with Homebrew:

```bash
brew install cosign
```

Install jq with the package manager of your system, for example:

```bash
brew install jq
```

### Compare Image size

Get the size of the images from [lab 3.3](/docs/03/3.3/).

```bash
docker image inspect --format '{{.Size}}' python-app:upstream | numfmt --to=iec
```

```bash
docker image inspect --format '{{.Size}}' python-app:cg | numfmt --to=iec
```

The upstream image is about 1.1G and the Chainguard image 85M (less than 10% of upstream)

Why is the size that differnet?


### Scan images

Scan the upstream image and the Chainguard image from [lab 3.3](/docs/03/3.3/).

With Docker:

```bash
docker run --rm --volume /var/run/docker.sock:/var/run/docker.sock anchore/grype python-app:upstream
docker run --rm --volume /var/run/docker.sock:/var/run/docker.sock anchore/grype python-app:cg
```

{{% alert title="Note" color="primary" %}}
The Docker socket is mounted into the container, so that Grype can access the images on your host.
{{% /alert %}}

With a local installation:

```bash
grype python-app:upstream
grype python-app:cg
```

Example output of the upstream image:

```bash
$ grype python-app:upstream
 ✔ Loaded image     python-app:upstream
 ✔ Parsed image     sha256:ae000a1f3ae6773750b71b87680c26fd7eabdaa429f5a439ce2be66c80508d9c
 ✔ Cataloged contents     f53f824be8198d0643a0d09b01bf962dacf618d093030cbd9ee7cf2f7c613631
   ├── ✔ Packages                        [486 packages]
   ├── ✔ Executables                     [1,411 executables]
   ├── ✔ File metadata                   [21,581 locations] 
   └── ✔ File digests                    [21,581 files]
 ✔ Scanned for vulnerabilities     [5234 vulnerability matches]
   ├── by severity: 591 critical, 3806 high, 3211 medium, 461 low, 1106 negligible (2365 unknown)
   └── by status:   9284 fixed, 2256 not-fixed, 6306 ignored
```

Example output of the Chainguard image:

```bash
$ grype python-app:cg
 ✔ Loaded image      python-app:cg
 ✔ Parsed image      sha256:4d2e40ae50d69ddc6a21da366cba2720c33495240c75e7c2468e1933fe28b893
 ✔ Cataloged contents       9426a10cccf3995efaca48f57ffd248a53a47d8c34e27185b6807231eb989771
   ├── ✔ Packages                        [44 packages]
   ├── ✔ Executables                     [149 executables]
   ├── ✔ File metadata                   [1,846 locations]
   └── ✔ File digests                    [1,846 files]
 ✔ Scanned for vulnerabilities     [4 vulnerability matches]
   ├── by severity: 0 critical, 2 high, 1 medium, 1 low, 0 negligible
   └── by status:   0 fixed, 4 not-fixed, 0 ignored
```

Compare the results:

* How many vulnerabilities are found in each image?
* Where do the vulnerabilities of the upstream image come from (base image, OS packages, Python packages)?

To only show vulnerabilities for which a fix is available, add `--only-fixed`.

With Docker:

```bash
docker run --rm --volume /var/run/docker.sock:/var/run/docker.sock anchore/grype python-app:upstream --only-fixed
```

With a local installation:

```bash
grype python-app:upstream --only-fixed
```

{{% alert title="Note" color="primary" %}}
Other scanners like [Trivy](https://trivy.dev/) or [Docker Scout](https://docs.docker.com/scout/) also support Chainguard images. Scanners need to know the Wolfi security advisories to correctly report the vulnerabilities of Chainguard images.
{{% /alert %}}


### Verify the image signature

Every Chainguard image is signed during the build. With `cosign` we can verify that an image was built by Chainguard and was not modified afterwards.

The public images in `cgr.dev/chainguard` are signed by the Chainguard release workflow on GitHub.

With Docker:

```bash
docker run --rm ghcr.io/sigstore/cosign/cosign:latest verify \
  --certificate-oidc-issuer=https://token.actions.githubusercontent.com \
  --certificate-identity=https://github.com/chainguard-images/images/.github/workflows/release.yaml@refs/heads/main \
  cgr.dev/chainguard/python:latest | docker run --rm -i ghcr.io/jqlang/jq .
```

With a local installation:

```bash
cosign verify \
  --certificate-oidc-issuer=https://token.actions.githubusercontent.com \
  --certificate-identity=https://github.com/chainguard-images/images/.github/workflows/release.yaml@refs/heads/main \
  cgr.dev/chainguard/python:latest | jq
```

The output shows the verified claims of the signature, for example the digest of the image.

Images of our organization are signed with a different identity, see [Verifying Chainguard Containers](https://edu.chainguard.dev/chainguard/chainguard-images/how-to-use/verifying-chainguard-images-and-metadata-signatures-with-cosign/) for details.


### Scan the SBOM

The SBOM lists every package included in the image, including the version and license. It is attached to the image as a signed attestation.

Download the SBOM of the Chainguard Python image as described in [lab 3.2](/docs/03/3.2/), with `IMAGE=cgr.dev/chainguard/python:latest`, and store it as `python-sbom.json`.

Also list and count the packages.

You can scan the SBOM directly, without pulling the image.

With Docker, the current directory is mounted into the container:

```bash
docker run --rm --volume "$PWD":/work --workdir /work anchore/grype sbom:./python-sbom.json
```

With a local installation:

```bash
grype sbom:./python-sbom.json
```
