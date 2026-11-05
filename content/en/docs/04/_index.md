---
title: "4. Chainguard Helm Charts"
weight: 4
sectionnumber: 4
description: >
  Test Migrations of Helm Charts.
---

## Helm Charts

Besides container images, Chainguard also provides Helm charts. There are two families:

* **Upstream charts** under `cgr.dev/<org>/charts/<chart>`: the community charts of a project, but configured to use Chainguard images.
* **iamguarded charts** under `cgr.dev/<org>/iamguarded-charts/<chart>`: forks of the Bitnami charts, which use the `-iamguarded` images (see [lab 3.1](/docs/03/3.1/)). They are meant as a drop-in replacement when migrating away from Bitnami.

In this lab we use the iamguarded Kafka chart of the `chainguard.dev-partners` organization.

For more information see [CGDocs](https://edu.chainguard.dev/chainguard/chainguard-images/how-to-use/use-chainguard-helm-charts/).


### Prerequisites

You need the `helm` CLI (version 3.8 or newer for OCI support). Check it with:

```bash
helm version
```

If it is missing, install it as described in the [Helm documentation](https://helm.sh/docs/intro/install/).


### Find the Helm charts

Open the [Chainguard Console](https://console.chainguard.dev/) and select **Helm charts** in the menu on the left. You see all charts that are available in your organization. Select the `kafka` chart to see its versions, its default values and the images it uses.

{{% alert title="Note" color="primary" %}}
Charts have to be added to an organization before they can be pulled. In the `chainguard.dev-partners` organization this is already done for you.
{{% /alert %}}


### Authenticate Helm

The Helm client does not use the docker login from [lab 2](/docs/02/). Charts are pulled with a pull token, which belongs to the charts location of the organization.

Make sure you are logged in with chainctl:

```bash
chainctl auth login
```

Create a pull token, valid for 24 hours:

```bash
chainctl auth configure-docker --pull-token --ttl=24h
```

You are asked for the location of the pull token. Select `iamguarded-charts`:

```
    With which location is the pull token associated?

    [chainguard.dev-partners] chainguard.dev-partners images catalog.
  > └ [iamguarded-charts]
```

The output shows a `docker login` command with a username (the identity ID of the pull token) and a password (the pull token). Store them in environment variables:

```bash
export HELMUSER=<username>
export HELMPASS=<password>
```

{{% alert title="Note" color="primary" %}}
We do not use `--save` here. It would write the pull token to your docker configuration and replace the login from lab 2.
{{% /alert %}}

Log in to the registry with Helm:

```bash
helm registry login cgr.dev \
  --username=$HELMUSER \
  --password=$HELMPASS
```

You should get a `Login Succeeded` message.


### Get the Helm chart

Show the chart metadata:

```bash
helm show chart oci://cgr.dev/chainguard.dev-partners/iamguarded-charts/kafka
```

Download and unpack the chart to look at it:

```bash
helm pull oci://cgr.dev/chainguard.dev-partners/iamguarded-charts/kafka --untar
ls kafka
```

Show the default values:

```bash
helm show values oci://cgr.dev/chainguard.dev-partners/iamguarded-charts/kafka | less
```

Search for `registry:`, `repository:` and `digest:`. Note that the images are pinned by digest, not only by tag.


### Which Chainguard images are used?

A chart deploys more than one image. Besides the application itself, there are helper images, for example for init containers or metrics.

The chart lists its images as an annotation in `Chart.yaml`:

```bash
yq '.annotations.images' kafka/Chart.yaml
```

Alternatively, render the chart and list the images of all resources:

```bash
helm template kafka oci://cgr.dev/chainguard.dev-partners/iamguarded-charts/kafka \
  --set "global.org=chainguard.dev-partners" \
  | grep 'image:' | sort -u
```

{{% alert title="Note" color="primary" %}}
Many helper images (for example for metrics or volume permissions) are only rendered if the corresponding feature is enabled in the values. The annotation in `Chart.yaml` lists all of them.
{{% /alert %}}

chainctl can list the pinned image references as well:

```bash
chainctl images helm refs cgr.dev/chainguard.dev-partners/iamguarded-charts/kafka
```

Which images are used? Are they all `-iamguarded` images? Compare them with the images you find in the Console under **Organization images**.

{{% details title="Hint" %}}
The Kafka chart uses `kafka-iamguarded` for the brokers and helper images such as `os-shell-iamguarded`, `kubectl-iamguarded` or `jmx-exporter-iamguarded`. The exact list depends on the chart version.
{{% /details %}}


### Use the Helm chart

The iamguarded charts need the organization set with `global.org`, so the images are pulled from your organization:

```bash
helm template kafka oci://cgr.dev/chainguard.dev-partners/iamguarded-charts/kafka \
  --set "global.org=chainguard.dev-partners" \
  > kafka.yaml
```

Check the rendered manifests in `kafka.yaml`. Every image should point to `cgr.dev/chainguard.dev-partners/...`.

To install the chart into a Kubernetes cluster, the cluster also needs credentials to pull the images. Create a pull secret with the same pull token and reference it in the values:

```bash
kubectl create secret docker-registry chainguard-pull-secret \
  --docker-server=cgr.dev \
  --docker-username=$HELMUSER \
  --docker-password=$HELMPASS

helm install kafka oci://cgr.dev/chainguard.dev-partners/iamguarded-charts/kafka \
  --set "global.org=chainguard.dev-partners" \
  --set "global.imagePullSecrets[0].name=chainguard-pull-secret"
```

{{% alert title="Note" color="primary" %}}
A cluster is not part of this techlab. If you want to try the installation, a local [kind](https://kind.sigs.k8s.io/) cluster is enough: `kind create cluster`.
{{% /alert %}}


### Migration from Bitnami

The iamguarded charts keep the structure and values of the Bitnami charts. For a migration, mostly the chart reference changes:

```bash
# before
helm upgrade kafka oci://registry-1.docker.io/bitnamicharts/kafka -f my-values.yaml

# after
helm upgrade kafka oci://cgr.dev/chainguard.dev-partners/iamguarded-charts/kafka -f my-values.yaml \
  --set "global.org=chainguard.dev-partners"
```

Things to watch out for:

* If your values override `image.registry`, `image.repository` or `image.tag`, remove them or point them to the Chainguard images.
* The images are pinned by digest. If you override a tag, also override or clear the digest, otherwise the digest wins: `--set image.tag=<tag> --set image.digest=""`.
* If you mirror the images to your own registry, chainctl generates the matching overrides:

```bash
chainctl images helm values cgr.dev/chainguard.dev-partners/iamguarded-charts/kafka \
  --registry myregistry.example.com
```
