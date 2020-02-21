---
title: User Interface
description: A look into the design of user interfaces
---
## Introduction

The user interface is of extreme importance to this project, as it is the portal to which the user will actually experience the product. A user interface is make or break when it comes to a users interactions and impressions of your project. If the data is there but it is tedious to access and presented in ways which are not intuitive and easy to read then people will not want to use your product. An example of this is the universities timetabling system, a system so bad users are happy to hand over their login details to a random party just so they can get out of using it.

## Standing on the shoulders of giants

Some of the best design at the moment is carried out by the tech giants. Luckily for us these companies have guides on their styles and designs.

Google is leading edge in design and has comprehensive guides on its style along with a number of different tools such as free icons and visually impaired friendly colour pickers. Along with a whole range of tools which allow for assessing how disability friendly your design is and helping you improve it. 

These resources can be found at https://design.google/resources/. Many other companies have similar style guides.

## Ease

When it comes to UI design our main focus should be around ease of use. This can be achieved by working out which tools are important to the user and putting them as the easiest to get to.

With the idea of ease of use comes cleanliness. Having a clean and simple design can help with ease of use as it means that the user knows exactly where to look for what they need as they are the only things which they can see. Of course you can always go too far in the other direction and make a design which is too clean and in this case you limit the user as it hides the functionality. 

One of the main design principles at the moment is material design. This design can be found across the google suite of products including android. This design methodology strongly follows the idea of clean and simple design in order to help the user. Everything in your design should be intuitive, removing the need for long explanatory tool tips or large manuals.

### "The blah blah blah blah blah"

For our project we will be following these guides.

## Colour

Colour can often be overlooked as just an aesthetic choice, but colour can play a large role in guiding the user through your interface. An example of this is using a consistent colour scheme throughout your interface. Having titles in one colour, subheadings in another and buttons in another, means that the user knows exactly what they are looking at and if they can interact it just by looking at the colour.

It is extremely important that you think about what colours you want to use, especially as your content is going to be viewed on a number of different devices with different screens and by people with different visual abilities. 

Some colour may look great on your device but what about devices which manipulate the colour such as night light software to change the hue to a much warmer one during the night. 

On top of this in this day and age there is a huge demand for a dark mode. Thoughts need to be given to a colour scheme which is dark and can take advantage of OLED technology to save energy.

The final and most important aspect of colour is visibility. The text should be easily readable. This means picking two highly contrasting colours for text and background. The higher the contrasting the easier it is to read for all abilities including those with partial sight. Google has a great tool for this which automatically picks colours for your text based on your background and will tell you how accessable they are: https://material.io/resources/color.

![](/images/uploads/colour.PNG "Google Colour Picker")

## Presenting Data

Another aspect we have thought about is how to present the data. As seen in other sections we having explored technologies such as grafana to do this but on a much lower level we need to think about what ways we want to present data. 

The most simplest of the methods would be a purely text based method which lists the rooms and then lists the occupancy. This approach is simple and easy to read and can often be found on websites status pages. Combined with colour it would make it the ideal choice if it was not for the fact that it struggles to scale as it can get messy with more and more locations.

![](/images/uploads/reddit.PNG)

This can be expanded if you want to represent the status over time as can be seen below:

![](/images/uploads/reddit2.PNG)

The above approach gives an abstracted view of the "occupancy" however it does not provide greater details. If this is needed then a graph like the one below. This graph gives you exact details and it is easy to spot the trend. However it does take up a lot of screen real estate. 

![](/images/uploads/graphana.PNG)

Moving away from the more traditional types of graphs, there are heat maps. These represent the data in a way which is easy to interpret to users who are not fully aware of room names as they can see it based on the floor plan. The downside of the heatmap is that if the user just wants to find a space on campus that is free they will have to check through each of the heat maps, as opposed to a table where they can just sort it via occupancy.

![](/images/uploads/heatmap.png)

## Conclusion

Each design has its pros and cons are serve different use cases. Management who want to monitor occupancy trends would find the graph most useful. 

Students on the other hand looking for a room would find the just list of rooms or heatmap most useful. A compromise between the two would be the most useful. For example at the top you list the rooms which have the lowest occupancy and below you show the heatmap. Therefor the data which you are presenting to the user first which is the most important to them is a suggestion of a couple of spaces which are free, and then if they are more picky about location or not aware where the rooms are they can consult the heat map.

![](/images/uploads/phone.png)
