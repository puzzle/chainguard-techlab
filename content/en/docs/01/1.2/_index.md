---
title: "1.2 Chainguard Console (optional)"
weight: 12
sectionnumber: 1.2
description: >
  Introduction to the Chainguard Console.
---

## Chainguard Console

Chainguard provides two web interfaces to explore the available images:

* the public [Chainguard Images Directory](https://images.chainguard.dev/), which can be used without an account
* the [Chainguard Console](https://console.chainguard.dev/), which shows the images, Helm charts and packages of your own organization and requires an account

This lab is optional. It shows how to explore the freely available information and how to create a personal account.

For more information, see [Using the Chainguard Console](https://www.chainguard.dev/docs/platform/console/images-directory) and [Using the Chainguard Directory](https://www.chainguard.dev/docs/chainguard/containers/registry/chainguard-directory).


### Images Directory without an account

Open the [Chainguard Images Directory](https://images.chainguard.dev/).

The landing page shows a card for each image with its name, the latest tag and when it last changed. Cards may be marked as **Free** tier image, FIPS or STIG-hardened image.

Use the menu on the left to filter by category, for example **Free**, **Base**, **Application**, **FIPS** or **Helm charts**. Use the search box to find a specific image.

Search for `python` and open the image. The free tier images can be pulled without authentication from `cgr.dev/chainguard`:

```bash
docker pull cgr.dev/chainguard/python:latest
```


### Create a personal account

To use the Chainguard Console you need an account. With the free **Catalog Starter** plan you can create your own organization and try up to five Chainguard images for free.

1. Go to [console.chainguard.dev/signup](https://console.chainguard.dev/signup)
2. Sign up with a business email address, or click **Continue with Google** with a company account. Personal addresses such as Gmail are not accepted
3. Enter your company name and click **Set up an organization**
4. Complete the form and click **Submit**
5. Click **Select images** and choose up to five images, for example `python`

{{% alert title="Note" color="info" %}}
The image selection of the Catalog Starter plan is fixed and cannot be changed later. FIPS images are not included. See [Catalog Starter](https://www.chainguard.dev/docs/chainguard/containers/reference/catalog-starter) for all limits.
{{% /alert %}}

Alternatively you can sign up on the command line, for example with your GitHub account. This requires `chainctl`, see [lab 1.4](/docs/01/1.4/):

```bash
chainctl starter init
```

The images of your organization can then be pulled from `cgr.dev/<your-organization>`, after logging in with `chainctl auth login` and `chainctl auth configure-docker`:

```bash
docker pull cgr.dev/<your-organization>/python:latest
```


### Overview

After signing in, you land on the **Overview** page of your organization.

The navigation on the left contains, among others:

* **Images**: container images
* **Helm charts**: Helm charts, including the `iamguarded` charts as replacement for Bitnami charts (see [lab 3](/docs/03/))
* **Packages**: APK packages in the private repository of the organization


### Images

Click **Images**. The page has two tabs:

* **Organization**: the images your organization can use, with **Name**, **Latest tag**, **Status** and **Pull URL**. The status **Active** means the image can be pulled, **Expired** means the access has lapsed.
* **Chainguard catalog**: all images Chainguard provides, with a description and when they were last updated.

Use the search box to find an image by name or version and the **Filter by** menu to filter by category. Click a column name to sort the list.


### Image details

Search for `python` and open the image. The details page is split into several tabs:

| Tab | Content |
| --- | --- |
| **Tags** | all available tags with **Pull URL**, **Compressed size** and **Last changed**. The **Digest** is shown for images of your organization. The **Variant** drop-down filters for `-dev` and non-dev variants |
| **Overview** | README of the image with getting started and compatibility notes |
| **Comparison** | number of CVEs compared to a non-Chainguard alternative |
| **Provenance** | how to verify the signature and attestations with `cosign` |
| **Specifications** | whether the image includes a shell or `apk`, the default user ID, environment variables and the entrypoint |
| **SBOM** | all packages included in the image, with download as SPDX or CycloneDX |
| **Vulnerabilities** | known CVEs of a version, with severity and fixed version |
| **Advisories** | timeline of the security advisories for the image |

Explore the tabs and answer the following questions:

* Which tags are available, and what is the difference between `latest` and `latest-dev`?
* Which user ID does the image run as, and what is the entrypoint? Compare the **Specifications** of the `latest` and the `latest-dev` variant.
* How many CVEs does the **Comparison** tab show for the upstream image?
* Which packages does the **SBOM** list? Is there a shell included?
* Select an older version in the **Vulnerabilities** tab. Are there any findings?

{{% alert title="Note" color="info" %}}
Most `latest` versions show no vulnerabilities, so you may need to select an older version to see findings.
{{% /alert %}}

We will use the same information with the command line tools in the following labs: digests in [lab 2.4](/docs/02/2.4/), SBOM and signatures in [lab 2.3](/docs/02/2.3/).
