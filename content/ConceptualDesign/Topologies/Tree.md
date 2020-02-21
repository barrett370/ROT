---
Mermaid: true
title: Worker Sensors
description: 'Sensors feed directly into processing system'
---

Each sensor unit will process the raw input and feed this directly into the data store

{{<mermaid>}}
graph LR;
        subgraph sensors
        A[#1]
        B[#2]
        C[#3]
        end
        A --> D
        B --> D
        C --> D
        subgraph cloud
        D[Database]
        end
        D --> E[User Interface]
{{< /mermaid >}}
