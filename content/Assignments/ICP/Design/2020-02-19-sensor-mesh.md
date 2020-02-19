---
Mermaid: 'false'
title: Sensor Mesh
date: 2020-02-19T15:56:45.245Z
---
{{<mermaid>}} graph LR;
        subgraph mesh
        A\[Sensor #1]         B\[Sensor #2]         C\[Sensor #3]         A --- B
        A --- C
        C --- B
        end
        A ==> D
        B ==> D
        C ==> D
        subgraph cloud
        D\[Ingress]         E\[Database]         D --> E
        end
        E --> F\[User Interface] {{< /mermaid >}}
