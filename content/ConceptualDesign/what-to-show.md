---
title: What data do we want to show?
weight: 3
pre: "<b>2.3 </b>"
description: >-
    Given our research, what information do our potential clients and end-users want to see?
---

From [our research](/research/), we now have a good idea of the short comings of the current system, what the information the users want to see and how they want to see it.

Our users have expressed interest in a simple design showing localised occupancy and opening times. Opening times is simple to show with a table next to the occupancy data for a given space.

Localised occupancy can be shown in a number of ways or perhaps as a combination of multiple formats.

The main goal of our system is to be simple, informative and easy to use. We do not want to present the user with information in formats such as this:

![Not this](/images/uploads/notthis.png)

## Heatmap layout


![Heatmap display](/images/uploads/heatmap.png)

This view is very *Tangible* with it being immediately obvious the areas of high density without the need for explicit axis or legends, although one may be included for completeness. One must also consider the case of a colour blind person who may struggle to interpret maps such as the one shown above, instead we may opt for a reduced colour pallette or commit to using multiple different information read-outs including more graphical means.

This sort of representation also takes up quite a lot of space in order for it to be clear. As well as it being expensive in terms of space graphics like these can be computationally expensive to generate and update/ redraw.


## Gauge Representation 

![Gauge Layout](/images/uploads/gauge.png)

Gauges are another simple way to show relative occupancy. They clearly show the current status compared to the extremes of 0 or 100%. An array of these gauges would quickly and simply show to users the spaces with the lowest relative occupancy.

This format is also quite space efficient although may take longer to interpret than the heatmap above due to it being more of an abstract representation.

## Other data to present.

We may also want to present historical data to our users either in the form of previous trends or predictions we have made using our database. 
This data could range from warnings in weeks that have historically had spikes in occupancy such as exam weeks to full occupancy number predictions based on some form of machine learning system.