---
title: "3.2 Drop-in replacements"
weight: 32
sectionnumber: 3.2
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

One can also pull specific image digests, the advantage of this being reproducibility, as it will ensure that you are using the same image each time (versus the tag that may receive updates, see [section 3.1](/docs/03/3.1/)).
A specific digest can be pulled via this docker command:

```bash
docker pull cgr.dev/chainguard/git@sha256:f6658e10edde332c6f1dc804f0f664676dc40db78ba4009071fea6b9d97d592f
```


### Retrieve SBOM

Every Chainguard image comes with an SBOM (Software Bill of Materials). The SBOM lists all packages included in the image, with their version and license. It is created at build time and attached to the image as a signed attestation.


#### SBOM in the Chainguard Console

Open the [Chainguard Console](https://console.chainguard.dev/), click **Images** and search for the `git` image.

Open the image and select the **SBOM** tab:

* select the version (tag) and the architecture (`x86_64` or `arm64`)
* use the search box to find a package, for example `openssl`
* click **Download** to export the SBOM in SPDX or CycloneDX format

Questions:

* How many packages are included in the image?
* Which licenses are used?
* Compare the SBOM of the `latest` and the `latest-dev` variant. Which additional packages are included in the `-dev` variant?


#### SBOM with cosign

The SBOM can also be downloaded on the command line with [cosign](https://docs.sigstore.dev/cosign/system_config/installation/). This is useful to automate SBOM processing, for example in a CI/CD pipeline.

Set the image we want to inspect:

```bash
export IMAGE=cgr.dev/chainguard/git:latest
```

You can either use cosign in a container or install it locally.

**Option 1: cosign container**

Chainguard provides cosign as a container image. Download the SBOM attestation and decode it with `jq` on your host:

```bash
docker run --rm cgr.dev/chainguard/cosign:latest download attestation \
  --platform=linux/amd64 \
  --predicate-type=https://spdx.dev/Document \
  $IMAGE | jq -r .payload | base64 -d | jq .predicate > git-sbom.json
```

**Option 2: local cosign binary**

Install cosign, for example with Homebrew:

```bash
brew install cosign
```

Or download the binary for Linux:

```bash
curl -sSfL -o cosign https://github.com/sigstore/cosign/releases/latest/download/cosign-linux-amd64
sudo install -m 0755 cosign /usr/local/bin/cosign
```

Download the SBOM attestation and decode it:

```bash
cosign download attestation \
  --platform=linux/amd64 \
  --predicate-type=https://spdx.dev/Document \
  $IMAGE | jq -r .payload | base64 -d | jq .predicate > git-sbom.json
```

**Inspect the SBOM**

List the packages with their version:

```bash
jq -r '.packages[] | "\(.name) \(.versionInfo)"' git-sbom.json
```

Count the packages and compare the result with the Console:

```bash
jq '.packages | length' git-sbom.json
```

{{% alert title="Note" color="primary" %}}
The attestation is a signed in-toto statement. The SBOM itself is the `predicate`, which is base64 encoded in the `payload`. How to verify the signature is shown in [lab 3.4](/docs/03/3.4/).
{{% /alert %}}

For more information, see [Retrieve image SBOMs](https://edu.chainguard.dev/chainguard/chainguard-images/how-to-use/retrieve-image-sboms/).


### Test image

Now let's test the `git` image by using it instead of a locally installed git.

The entrypoint of the image is `git`, so every argument you pass to `docker run` is handed over to `git`:

```bash
docker run --rm cgr.dev/chainguard/git:latest --version
```

Create a directory for the test and move into it:

```bash
mkdir -p ~/tmp/git-test && cd $_
```

Clone a repository into the current directory. The directory is mounted into the container as `/work`:

```bash
docker run --rm \
  --user "$(id -u):$(id -g)" \
  --volume "$PWD":/work \
  --workdir /work \
  cgr.dev/chainguard/git:latest clone https://github.com/puzzle/backstage-techlab.git
```

{{% alert title="Note" color="primary" %}}
The Chainguard images run as non-root user with the UID `65532`. With `--user "$(id -u):$(id -g)"` the container runs with your user ID, so it is allowed to write into the mounted directory and the cloned files belong to you.
{{% /alert %}}

Check that the repository was cloned to your host:

```bash
ls backstage-techlab
```

Show the last commits of the repository, again with the container:

```bash
docker run --rm \
  --user "$(id -u):$(id -g)" \
  --volume "$PWD/backstage-techlab":/work \
  --workdir /work \
  cgr.dev/chainguard/git:latest log --oneline -5
```

To compare, run the same commands with a commonly used upstream git image, for example `alpine/git`:

```bash
docker run --rm alpine/git:latest --version
```

Both images use `git` as entrypoint, so the Chainguard image can be used as drop-in replacement.

Clean up the test directory:

```bash
cd ~ && rm -rf ~/tmp/git-test
```
