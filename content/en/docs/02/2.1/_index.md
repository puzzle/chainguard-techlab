---
title: "2.1 Container images - Drop-in replacements"
weight: 21
sectionnumber: 2.1
---

## Container images - Drop-in replacements

### Check for Chainguard images

Many container images for infrastructure applications are build to be drop-in replacements to their open-source counter parts. 
Let's try this out:

List Puzzle repos:

```bash
chainctl images repos list

# use this or check setup to set the puzzle org as default
chainctl images repos list --parent=puzzle-partner.com
```

Should show something like this:

```bash
[cgr.dev/puzzle-partner.com]
├ [adoptium-jdk]
├ [adoptium-jre]
├ [curl]
...
```

List `python` image tags that we can use at Puzzle:

```bash
chainctl images list --repo=python
```

### Replace image

### Test image
