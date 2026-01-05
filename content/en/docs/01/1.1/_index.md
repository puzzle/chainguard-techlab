---
title: "1.1 Big Picture"
weight: 11
sectionnumber: 1.1
description: >
  Big picture and motivation.
---

## Security Risks in your Cloud-Native Stack

Every year OWASP publishes their Top 10 Most Critical Web Application Security Risks, in 2025 Software Supply Chain Failures were moved up to the third pace (from 6th), and now share the top three among Broken Acess Control and Security Misconfigurations ([OWASP 2025 Report](https://owasp.org/Top10/2025/0x00_2025-Introduction/)).

Software Supply Chain describes the entirety of physical and software components involved in creating a product.

The modern cloud infrastructure and applications are build on open source software, with 96% of codebases containing free and open source software (Linux Foundation, CITE TODO). 
The perpetual reuse of open source software has many advantages, such as transparency, cost-efficiency, and flexibility just to name a few.

Open source software unfortunately also comes with downsides, many of the used base images come with known vulnerabilities (especially if not regularly updated), and they may include hidden dependencies (typically dependencies of dependencies...).

![Hidden supply chain security risks](infrastructure.png)

Many of the open source projects are underfunded, aren't regularly updated and 32% of them are in a neglected state ([Orca 2025 State of Coud Security Report](https://orca.security/lp/2025-state-of-cloud-security-report/)).
Therefore, many of them accumulate CVEs (common vulnerabiities and exposures).

## How does Chainguard help?

Chainguard has a shift left approach, providing daily updated images with ~0 CVEs.


