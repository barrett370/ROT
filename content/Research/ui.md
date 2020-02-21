---
title: UI Layout
description: An in-depth look into layout
weight: 7
pre: "<b>1.7 </b>"
---
### Introduction

The layout of the UI is something which should be thought about for multiple reasons. The layout will affect what the user interact with and how easily they interact with it. A poorly laid user interface may mean that your best features never get used. In this section we are going to explore the interface for a typical student user.

#### What's Important?

It's important that we identify what is important to our user so we can make it available to them. According to our research, most of the users don't care about most of the features, they just want to find a space. That's it, any space and they will be satisfied. 

Some students identified that they wanted to find spaces in specific areas of campus or buildings.

#### First things first

As finding any space is the highest important to students, in our design we have decided to put this at the top of the page, right where the user will see when they open the app / website. This section will use an algorithm to determine which spaces the student should go to based on current/expected occupancy and even possibly taking into account the location of the student. It will then display them clearly in a listed button format, dictating the most important information: space name, building, and occupancy. If the user wants more details they can click on the button and it will take them to a view specifically for that item.

#### Less Specific Locations

Often students didn't a specific room, they wanted a much wider impression of occupancy across campus. An example of this is students wanted to know what was more busy, the library or the sports center. They often didn't much care about specific rooms but just wanted know which buildings are low occupancy. 

To help communicate this, information will also be displayed on the main view under the specific rooms. It will show a list of each of the top occupancy for whole locations. This will allow users to know which building is best in general, instead of a specific room. 

A pro of doing building based ranking is that the occupancy will be much less volatile. The amount of people that it will take to fill up a building as opposed to a single room is much higher which means that it is less likely to change faster. This matters because a user doesn't want to walk halfway across campus just for there space to become occupied by the time they get there.

An example of this can been seen below.

![](/images/uploads/home.PNG)

#### Specific Locations

There will also be a list of the locations. Here the user can find the exact room they want by expanding through a series of collapsible buttons of increasing specificity. It is important that there is a balance between the number of clicks it takes to get to your item and also the number of items displayed when an item is expanded. For example you could find what you are looking for if each section is vague and therefore contains a large number of items. But this means it will take quite a bit of searching to find it in these items. You could be more specific and have each collapse only reveal a few children, however this will take a much larger number of clicks to achieve your goal. 

One possible solution would be `Campus Region -> Building -> Floor -> Rooms`, another approach would be to go `Group of Subjects -> Subject -> Building -> Floor -> Rooms`. Each of these methods will have different outcomes on the end user based on how they want to use the app. For example if they know the geographical layout of the campus and only care about buildings in one specific region then the first would be faster but if they are just trying to find a room in chemistry then it would be simpler for them to use the second approach. 

![](/images/uploads/locations.PNG)

# Testing

Whilst our research may show what the users think they want, you cant be sure that it is definitely what they want. To handle this we hope to do a variety of A/B testing of a variety of different layouts. Both in the initial sketches, put also the interactive prototypes.

In this testing we hope to observe a number of usability metrics to assess the designs. Some examples of these is speed and number of clicks to reach a goal. We will also interview the users to try to figure out their experience with each design. Hopefully using this information we will be able to design a product which puts the users first.
