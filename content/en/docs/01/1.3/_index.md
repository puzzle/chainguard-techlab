---
title: "1.3 Prerequesites"
weight: 13
sectionnumber: 1.3
description: >
  Getting started with Chainguard
---


## Getting started with Chainguard


### Log into Puzzle Partner Account

Enter Chainguard partner console, by going to the [Chainguard Image Directory](https://images.chainguard.dev/), clicking the "sign in" button in the upper right corner and sign-in with the puzzle email address (Google). 

You now should be able to see the [Puzzle Partner space](https://console.chainguard.dev/org/puzzle-partner.com/overview).

If you don't have a Puzzle mail address, or for more information see [PDoc](https://docs.puzzle.ch/user-guides/chainguard-user-guide/index.html#_getting_started).

### Install chainctl

Chainguard’s ´chainctl´ command-line interface provides essential tools for managing your container security infrastructure, including image management, identity and access control, and resource monitoring. This CLI enables automation of Chainguard operations and integration with CI/CD pipelines.

For more information on chainctl, see [CGDocs](https://edu.chainguard.dev/chainguard/chainctl-usage/how-to-install-chainctl/).

#### Linux

#### Windows

#### Mac


### Verifying installation and first steps




Set Puzzle org as default:

```bash
chainctl config set default.group puzzle-partner.com
```

#### Auth

Check that you have the `owner` role in https://console.chainguard.dev/.
This is needed to be able to pull images and libraries.

Have latest [chainctl](https://edu.chainguard.dev/chainguard/chainctl/chainctl-docs/chainctl/) installed, run:

```bash
chainctl update`
rm -f /home/craaflaub/.cache/chainctl/chainctl.bak
```

Chainctl login to Chainguard console:

```bash
chainctl auth login
```

Browser should open and do auth.