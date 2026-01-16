---
title: "4. Chainguard Libraries"
weight: 2
sectionnumber: 4.
---

## Chainguard Libraries

To show which libraries you have access to, paste this in your CLI:

```bash
chainctl libraries entitlements list --parent=puzzle-partner.com
```

Should show something like this:

```bash
Ecosystem Library Entitlements for puzzle-partner.com (337c8d93f34ada9c0f59aa998aa2234aca10f7d7)

                            ID                             | ECOSYSTEM  
-----------------------------------------------------------|------------
 337c8d93f34ada9c0f59aa998aa2234aca10f7d7/2b16fc4ebdf97061 | JAVASCRIPT 
 337c8d93f34ada9c0f59aa998aa2234aca10f7d7/37e80f06263e0662 | JAVA
 337c8d93f34ada9c0f59aa998aa2234aca10f7d7/4a3dbce6fc6d2a1a | PYTHON
```
 
This means, that we have access to [Chainguard libraries](https://www.chainguard.dev/libraries) for Java, JavaScript and python.




