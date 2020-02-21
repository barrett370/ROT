---
Mermaid: true
title: Topologies
description: Early technical proposals
weight: 2
pre: "<b>2.2 </b>"
---

In the most abstract sense our current design consists of the following:

{{<mermaid align="left">}}
graph LR;
        A[Sensors] --> B[Processing]
        B[Processing] --> C[Insights]
{{< /mermaid >}}

---

As follows are early technical architecture propositions.

{{% children description="true" %}}
