---
title: "3. Chainguard Helm Charts"
weight: 2
sectionnumber: 3
description: >
  Test Migrations of Helm Charts.
---

## Helm Charts

https://edu.chainguard.dev/chainguard/chainguard-images/how-to-use/use-chainguard-helm-charts/


```bash
chainctl auth login
chainctl auth configure-docker --pull-token --save --ttl=24h


chainctl auth configure-docker --pull-token --save --ttl=24h

    With which location is the pull token associated?      
                                                           
    [puzzle-partner.com] puzzle-partner.com images catalog.
  > └ [iamguarded-charts]                                  

 export HELMUSER=
 export HELMPASS=


helm registry login cgr.dev \
  --username=$HELMUSER \
  --password=$HELMPASS
  
helm template oci://cgr.dev/puzzle-partner.com/iamguarded-charts/kafka
```
