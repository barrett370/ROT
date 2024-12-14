---
title: Occupancy Data Interpretation
description: A look into representing occupancy
---
## Percentage %

This metric is similar to colour in the sense that it is abstract. It does not tell the user exactly how many seats are left but does give them an impression of room fullness. This is useful for users who only care about the feel of the room, ie how "packed it is".  


## Seats Left

Seats left provides the most exact details. This metric is perfect for users who don't really care about how full the room is just want to find a seat or a specific number of seats. This method falls short when it comes to rooms with drastically different sizes. For example if a room has 300 seats available and another has 6 then saying 3 seats left, doesn't give a good impression of occupancy as one room is 99% full and the other is 50% full. As can be seen those rooms are not the same in fullness yet the seats left metric suggests they are equals which gives the wrong impression.

## Colour

Colour representations allow you to simplify information to a level where you can just glance at it to a basic insight.

### Gradient

A gradient of colour could be considered a seamless representation of percentage occupancy and could be applied to still provide an interpretation that can be processed from just entering the peripheral view.
![Grafana Gradient](https://grafana.com/static/assets/img/blog/bargauge/gradient.jpg)

### Traffic Light

Using a "traffic light" code to indicate occupancy presents information in an already widely understood standard.
![Grafana Traffic Light](https://grafana.com/static/img/docs/v45/singlestat-color-options.png)

### Accessibility of colours

Depending on colour does again introduce the accessibility issue and solving this with a different colour pallet has the potential to introduce a new issue whereby users may not understand the relation between colour and occupancy when looking directly at it, let alone in passing.

Simplicity
Colour issues
Different types of colour

Seen at a glance, not very detailed
