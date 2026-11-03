---
title: "6.3 Vulnerability Scan and CVE comparison"
weight: 63
sectionnumber: 6.3
---

## Scan Images

We will scan and compare both images.

Use your scanner of choice or follow the instructions to use Grype.


### Scan with Grype

Use Grype inside a container or install it.


#### Grype local

Check the Grype tutorial for installing Grype or using the Grype container: https://edu.chainguard.dev/chainguard/chainguard-images/staying-secure/working-with-scanners/grype-tutorial/

Scan both images to compare the vulnerabilities.

Eclipse JDK:

```bash
grype my-spring-jdk:latest
```

Chainguard JDK:

```bash
grype my-spring-cg:latest
```


#### Grype with docker

Scan both images to compare the vulnerabilities.

Eclipse JDK:

```bash
docker run --rm -it --volume /var/run/docker.sock:/var/run/docker.sock anchore/grype my-spring-jdk:latest
```

Chainguard JDK:

```bash
docker run --rm -it --volume /var/run/docker.sock:/var/run/docker.sock anchore/grype my-spring-cg:latest
```


## Comparison

What are the differences?

Compare:

* Image size
* Vulnerabilities
* ...


### Image size

Get the size of the image in a human readable format:

```bash
docker image inspect --format '{{.Size}}' my-spring-jdk:latest | numfmt --to=iec
docker image inspect --format '{{.Size}}' my-spring-cg:latest | numfmt --to=iec
```

To see which layers contribute to the size, use:

```bash
docker history my-spring-jdk:latest
docker history my-spring-cg:latest
```


### Vulnerabilities

Example scan of Eclipse JDK:

```bash
$ grype my-spring-jdk:latest
 ✔ Vulnerability DB                [updated]
 ✔ Loaded image                    my-spring-jdk:latest
 ✔ Parsed image   sha256:f3950876afe7997ebec34f4daa29b48eb013c3df09263862e1ce392b70d3703d
 ✔ Cataloged contents   c480f57a57b6e88521fab7bb417fa8af99512641b1e164953b7683492ac47b04
   ├── ✔ Packages                        [186 packages]
   ├── ✔ Executables                     [610 executables]
   ├── ✔ File metadata                   [3,050 locations]
   └── ✔ File digests                    [3,050 files]
 ✔ Scanned for vulnerabilities     [413 vulnerability matches]
   ├── by severity: 0 critical, 21 high, 197 medium, 190 low, 5 negligible
   └── by status:   3 fixed, 410 not-fixed, 0 ignored 
```


Example scan of Chainguard JDK:

```bash
$ grype my-spring-cg:latest
 ✔ Vulnerability DB                [no update available]
 ✔ Loaded image                    my-spring-cg:latest
 ✔ Parsed image   sha256:f31fff25566a95d8e9e5c6821f68518788fa1a5511eeb2d1383a9728bc38669f
 ✔ Cataloged contents  5f539e4a6f293ce3d746da0a016b565092fd772d82e70c210960c4e3f55262d3
   ├── ✔ Packages                        [86 packages]
   ├── ✔ Executables                     [127 executables]
   ├── ✔ File digests                    [1,015 files]
   └── ✔ File metadata                   [1,015 locations]
 ✔ Scanned for vulnerabilities     [6 vulnerability matches]
   ├── by severity: 0 critical, 3 high, 3 medium, 0 low, 0 negligible
   └── by status:   1 fixed, 5 not-fixed, 0 ignored
```
