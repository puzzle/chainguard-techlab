---
title: "5.1 Introduction to Chainguard Libraries"
weight: 51
sectionnumber: 5.1
description: >
  Introduction to Chainguard Libraries.
---

<!-- TODO: placeholder, to be extended -->

## Chainguard Libraries

Most applications depend on many open source libraries, which are usually downloaded as pre-built packages from public registries like Maven Central, PyPI or npm. It is often not traceable how these packages were built, and attackers increasingly target these registries, for example by publishing malicious packages or by compromising the build of popular projects.

[Chainguard Libraries](https://www.chainguard.dev/libraries) address this risk: Chainguard builds the libraries from their source code in its own hardened build infrastructure and provides them in a separate repository. The libraries can be used as a replacement for the packages from the public registries, without changes to the application code.

Chainguard Libraries are available for the following ecosystems:

* Java (Maven)
* Python (PyPI)
* JavaScript (npm)

The libraries are provided through `https://libraries.cgr.dev/` and can be used directly by your build tool or through a repository manager like Artifactory or Nexus.

In the following labs we check which libraries we have access to ([lab 5.2](/docs/05/5.2/)) and use the Java libraries to build an application ([lab 6.4](/docs/06/6.4/)).

For more information, see the [Chainguard Libraries documentation](https://edu.chainguard.dev/chainguard/libraries/overview/).
