---
title: "1.3 Prerequesites"
weight: 13
sectionnumber: 1.3
description: >
  Getting started with Chainguard
---


## Getting started with Chainguard

For this techlab it is necessary to install chainctl, if you are using Chainguard for your company however, it is best practise to set up a pull-through-cache with the registry of your choice.

### Log into Puzzle Partner Account

Enter Chainguard partner console, by going to the [Chainguard Image Directory](https://images.chainguard.dev/), clicking the "sign in" button in the upper right corner and sign-in with the puzzle email address (Google). 

You now should be able to see the [Puzzle Partner space](https://console.chainguard.dev/org/puzzle-partner.com/overview).

If you don't have a Puzzle mail address, or for more information see [PDoc](https://docs.puzzle.ch/user-guides/chainguard-user-guide/index.html#_getting_started).

### Install chainctl

Chainguard’s ´chainctl´ command-line interface provides essential tools for managing your container security infrastructure, including image management, identity and access control, and resource monitoring. This CLI enables automation of Chainguard operations and integration with CI/CD pipelines.

For more information on chainctl, see [CGDocs](https://edu.chainguard.dev/chainguard/chainctl-usage/how-to-install-chainctl/).

To begin with, we are creating a temporary directory and moving into it, by executing

´´´
mkdir ~/tmp && cd $_
´´´

#### Linux

To install chainctl via curl by executing

´´´
curl -o chainctl "https://dl.enforce.dev/chainctl/latest/chainctl_$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m | sed 's/aarch64/arm64/')"
´´´

Move chainctl into your /usr/local/bin directory and elevate its permissions so that it can execute as needed.

´´´
sudo install -o $UID -g $(id -g) -m 0755 chainctl /usr/local/bin/ 
´´´

And you are finshed, jump to verifying your installation :)

#### Homebrew for Mac or Linux

It is required that [Xcode Command Line Tools](https://mac.install.guide/commandlinetools/) is installed to install chainctl via Homebrew.

You can do that by running the following command:

´´´
xcode-select --install
´´´
Before installing chainctl with Homebrew, use brew tap to bring in Chainguard’s repositories.

´´´
brew tap chainguard-dev/tap
´´´
Next, install chainctl with Homebrew.

´´´
brew install chainctl
´´´

And you are finshed, jump to verifying your installation :)

#### Windows

To download the executional file, run the following command:

´´´
curl -o chainctl.exe https://dl.enforce.dev/chainctl/latest/chainctl_windows_x86_64.exe
´´´

And you are finshed, please be aware, that that Windows PowerShell does not load commands from the working directory by default so you will need to include .\ before any chainctl commands you run, as in this example.

´´´
.\chainctl auth login
´´´

### Verifying installation and first steps

Check if chainctl is installed correctly by running:

´´´
chainctl version
´´´

Set Puzzle org as default:

```bash
chainctl config set default.group puzzle-partner.com
```

#### Authentication

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

### Updating chainctl

Did you install chainctl a while ago and need to update it?

Just run

```bash
sudo chainctl update
```

And make sure you are signed in as described in "Verifying installation and first steps".