---
title: "1.4 Prerequisites"
weight: 14
sectionnumber: 1.4
description: >
  Getting started with Chainguard
---


## Getting started with Chainguard

For this techlab it is necessary to install chainctl, if you are using Chainguard for your company however, it is best practice to set up a pull-through cache with the registry of your choice (see [lab 6.1](/docs/06/6.1/)).


### Install chainctl

Chainguard’s `chainctl` command-line interface provides essential tools for managing your container security infrastructure, including image management, identity and access control, and resource monitoring. This CLI enables automation of Chainguard operations and integration with CI/CD pipelines.

For more information on chainctl, see [CGDocs](https://edu.chainguard.dev/chainguard/chainctl-usage/how-to-install-chainctl/).

To begin with, we are creating a temporary directory and moving into it, by executing

```bash
mkdir ~/tmp && cd $_
```


#### Linux

To install chainctl via curl by executing

```bash
curl -o chainctl "https://dl.enforce.dev/chainctl/latest/chainctl_$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m | sed 's/aarch64/arm64/')"
```

Move chainctl into your /usr/local/bin directory and elevate its permissions so that it can execute as needed.

```bash
sudo install -o $UID -g $(id -g) -m 0755 chainctl /usr/local/bin/
```

And you are finished, jump to verifying your installation :)


#### Homebrew for Mac or Linux

It is required that [Xcode Command Line Tools](https://mac.install.guide/commandlinetools/) is installed to install chainctl via Homebrew.

You can do that by running the following command:

```bash
xcode-select --install
```

Before installing chainctl with Homebrew, use brew tap to bring in Chainguard’s repositories.

```bash
brew tap chainguard-dev/tap
```

Next, install chainctl with Homebrew.

```bash
brew install chainctl
```

And you are finished, jump to verifying your installation :)


#### Windows

To download the executable file, run the following command:

```bash
curl -o chainctl.exe https://dl.enforce.dev/chainctl/latest/chainctl_windows_x86_64.exe
```

And you are finished, please be aware, that Windows PowerShell does not load commands from the working directory by default so you will need to include .\ before any chainctl commands you run, as in this example.

```bash
.\chainctl version
```


### Verifying installation

Check if chainctl is installed correctly by running:

```bash
chainctl version
```


### Authentication

For this techlab, you log in with an identity provided by the lab instructor. You will receive an identity ID and an identity token.

Set them as environment variables:

```bash
export CHAINGUARD_IDENTITY_ID=<identity-id>
export CHAINGUARD_IDENTITY_TOKEN=<identity-token>
```

Log in with chainctl:

```bash
chainctl auth login --identity "$CHAINGUARD_IDENTITY_ID" --identity-token "$CHAINGUARD_IDENTITY_TOKEN"
```

Check that you are logged in:

```bash
chainctl auth status
```

{{% alert title="Note" color="info" %}}
The identity and its token are only valid for the duration of the techlab. If you want to explore Chainguard on your own afterwards, create a personal account as described in [lab 1.2](/docs/01/1.2/) and log in with `chainctl auth login`.
{{% /alert %}}


### Updating chainctl

Did you install chainctl a while ago and need to update it?

Just run

```bash
sudo chainctl update
```

And make sure you are logged in as described in "Authentication".
