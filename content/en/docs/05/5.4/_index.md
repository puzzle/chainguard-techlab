---
title: "5.4 Vulnerability Scan and CVE comparison"
weight: 54
sectionnumber: 5.4
---

## Scan Images

We will scan and compare both images.

Check the Grype tutorial for installing Grype or using the Grype container: https://edu.chainguard.dev/chainguard/chainguard-images/staying-secure/working-with-scanners/grype-tutorial/

Scan both images to compare the vulnerabilities:

```bash
grype my-spring-jdk:latest

grype my-spring-cg:latest
```

Alternatively with docker:

```bash
docker run --rm -it --volume /var/run/docker.sock:/var/run/docker.sock anchore/grype my-spring-jdk:latest
docker run --rm -it --volume /var/run/docker.sock:/var/run/docker.sock anchore/grype my-spring-cg:latest
```

What are the differences?
