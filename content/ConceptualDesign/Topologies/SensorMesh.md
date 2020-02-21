---
Mermaid: true
title: Sensor Mesh
description: 'Inter connected sensors relaying data across the mesh to a processing node'
---

Where possible all sensors are connected to a mesh network. At the edge of each mesh is an **ingress** node that processes the raw sensor data. This is then streamed into the database which is made available to the user.

{{<mermaid>}}
graph LR;
        subgraph mesh
        A[Sensor #1]
        B[Sensor #2]
        C[Sensor #3]
        A --- B
        A --- C
        C --- B
        end
        A ==> D
        B ==> D
        C ==> D
        subgraph cloud
        D[Ingress]
        E[Database]
        D --> E
        end
        E --> F[User Interface]
{{< /mermaid >}}