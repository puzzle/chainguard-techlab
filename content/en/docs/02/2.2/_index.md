---
title: "2.2 Container images - More challenging swaps"
weight: 22
sectionnumber: 2.2
---

## Container images - More challenging swaps

Adapt Dockerfile

- user permission ID no longer root, uid=65532
- images set entrypoint explicitly (in python image the entrypoint is python)
    ´docker run-it python´
    ´docker run -it python echo "in a shell"´