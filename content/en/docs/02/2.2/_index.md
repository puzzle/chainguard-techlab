---
title: "2.2 More challenging swaps"
weight: 22
sectionnumber: 2.2
---

## Container images - More challenging swaps

<!--

TODO

Adapt Dockerfile

* switch to apk
* user permission ID no longer root, uid=65532
* images set entrypoint explicitly (in python image the entrypoint is python)
  * ´docker run-it python´
  * ´docker run -it python echo "in a shell"´
* Busy Box as default
* groupadd/useradd -> addgroup, adduser
* entrypoints may be different

Migration Checlist: https://edu.chainguard.dev/chainguard/migration/migration-checklist/
Converter: https://edu.chainguard.dev/chainguard/migration/dockerfile-conversion/


-->

Not every image is a drop-in replacement. As soon as you build your own image on top of a base image, a few differences between Chainguard images and the common upstream images become visible:

* Chainguard images are based on Wolfi and use `apk` instead of `apt`/`dnf` as package manager
* the images run as non-root user with the UID `65532` (`nonroot`) instead of `root`
* the runtime images are distroless, they contain no shell and no package manager
* the images set the entrypoint explicitly, which may differ from the upstream image
* the `-dev` images use BusyBox, so some tools have different names or options (for example `adduser`/`addgroup` instead of `useradd`/`groupadd`)

In this lab we migrate a small Python application step by step.

For a complete list of things to check, see the [Migration Checklist](https://edu.chainguard.dev/chainguard/migration/migration-checklist/).


### Prepare the example application

Create a new directory and move into it:

```bash
mkdir -p ~/tmp/python-app && cd $_
```

Create the file `app.py`:

```python
from flask import Flask

app = Flask(__name__)


@app.route("/")
def hello():
    return "Hello from Python!\n"


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
```

Create the file `requirements.txt`:

```text
flask
```

Create the file `Dockerfile` as you would typically write it with the upstream Python image:

```Dockerfile
FROM python:3.12

RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*
RUN useradd -m appuser

WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY app.py .

USER appuser
EXPOSE 8080
CMD ["python", "app.py"]
```

Build and run the image to make sure it works:

```bash
docker build -t python-app:upstream .
docker run --rm -p 8080:8080 python-app:upstream
```

In a second terminal check the application:

```bash
curl http://localhost:8080/
```

Stop the container with `ctrl+c`.


### Entrypoints

Let's first have a look at how the Chainguard Python image behaves. The entrypoint of the image is `python` itself:

```bash
docker run --rm -it cgr.dev/chainguard/python:latest
```

This opens the Python interpreter (exit with `exit()`). Every argument you pass is handed over to `python`:

```bash
docker run --rm cgr.dev/chainguard/python:latest --version
```

Try to run a shell command:

```bash
docker run --rm cgr.dev/chainguard/python:latest echo "in a shell"
```

This fails, because Python tries to open a file called `echo`. The image has no shell, so overwriting the entrypoint does not help either:

```bash
docker run --rm --entrypoint /bin/sh cgr.dev/chainguard/python:latest
```

The `-dev` variant of the image includes a shell and `apk`:

```bash
docker run --rm -it --entrypoint /bin/sh cgr.dev/chainguard/python:latest-dev
```

Inside the container, check the user and the installed tools:

```bash
whoami
id
apk --version
exit
```

{{% alert title="Note" color="info" %}}
Our `Dockerfile` uses `CMD ["python", "app.py"]`. With the Chainguard image this results in `python python app.py`, because `python` is already the entrypoint. Use `ENTRYPOINT` instead of `CMD`, or only pass the arguments, for example `CMD ["app.py"]`.
{{% /alert %}}


### Convert the Dockerfile with dfc

Chainguard provides the Dockerfile Converter [dfc](https://edu.chainguard.dev/chainguard/migration/dockerfile-conversion/). It replaces the base images with Chainguard images, converts package manager commands to `apk` and user commands to the BusyBox syntax.

Install dfc with Homebrew:

```bash
brew install chainguard-dev/tap/dfc
```

Or with Go:

```bash
go install github.com/chainguard-dev/dfc@latest
```

Convert the `Dockerfile` and print the result. With `--org` we tell dfc to use the images of the lab organization `chainguard.dev-partners`:

```bash
dfc --org chainguard.dev-partners ./Dockerfile
```

Compare the output with the original `Dockerfile`:

* Which base image is used now?
* What happened to the `apt-get` command?
* What happened to `useradd`?
* Why did dfc add a `USER root`?

dfc does a mechanical conversion. The result works in most cases, but it still uses a `-dev` image for runtime. To get the full benefit of Chainguard images, we optimize the result manually with a multistage build.


### Multistage build

The recommended pattern is to use a `-dev` image to install dependencies or compile the application and to copy only the required files into a distroless runtime image.

Create the file `Dockerfile.cg`:

```Dockerfile
FROM cgr.dev/chainguard/python:latest-dev AS builder

ENV LANG=C.UTF-8
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV PATH="/app/venv/bin:$PATH"

WORKDIR /app
RUN python -m venv /app/venv
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt


FROM cgr.dev/chainguard/python:latest

ENV PYTHONUNBUFFERED=1
ENV PATH="/venv/bin:$PATH"

WORKDIR /app
COPY app.py .
COPY --from=builder /app/venv /venv

EXPOSE 8080
ENTRYPOINT ["python", "/app/app.py"]
```

Things to note:

* the dependencies are installed into a virtual environment in the `builder` stage and copied into the runtime stage
* `curl` is no longer installed, the runtime image only contains what is needed to run the application
* no user is created, the image already runs as `nonroot` (UID `65532`)
* the application listens on port `8080`, a non-root user cannot bind to ports below `1024`

Build and run the image:

```bash
docker build -t python-app:cg -f Dockerfile.cg .
docker run --rm -p 8080:8080 python-app:cg
```

In a second terminal check the application again:

```bash
curl http://localhost:8080/
```

Stop the container with `ctrl+c`.


### Compare the images

Compare the size of the images:

```bash
docker images python-app
```

Check which user the containers are running as:

```bash
docker inspect --format '{{.Config.User}}' python-app:upstream
docker inspect --format '{{.Config.User}}' python-app:cg
```

We will compare the vulnerabilities of both images in the [next lab](/docs/02/2.3/).
