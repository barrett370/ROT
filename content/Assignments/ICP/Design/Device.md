---
title: Device Design Ideas
draft: false
---

There are 3 major sensor-based approaches we have thought of pursuing. 

1. Entryway traffic based tracking
2. Per-table based tracking
3. Ambient environment tracking

Notice that all of these approaches are relatively non-invasive with none requiring the use of cameras. 

## Entryway traffic tracking

This proposed system would work via a doorway mounted laser depth sensor. By mounting the sensor at a (*known*) angle, we can calculate the direction of travel of an individual. This simple sensor and post-processing algorithm would allow us to maintain a record of entrances and exits from a given doorway and with a backing system aware of the relative positions of the devices, we could extrapolate numbers of people in each building or room. 
This system requires no camera or personal data to be collected and only one sensor per device, relatively low in cost.


![ETT](/img/SDSDesignApproach1.jpg)


## Per-table tracking

This system would be quite similar to the one currently implemented, with under-table mounted laser sensors allowing the system to accurately determine how many people are surrounding the given area. A downside to this is it's expense with each table in a given space requiring a sensor array and computer to process and upload the data. This system also does not work for spaces that have non-standard seating arrangements such as single seat booths found in the library and Aston Webb building. 

![PTT](/img/SDSDesignApproach2.jpg)

## Ambient environment Tracking

This system would rely on the gathering of metrics such as ambient temperature, humidity or CO2 levels in order to extrapolate ,via heuristics,the number of people in a given space. 
This system would be more expensive per unit due to the requirement of several delicate sensors. However, it would allow for a more subtle placement not requiring a location above a doorway or under a table. It's accuracy would however, be dependent on experimentation on a per-space basis as occupancies can lead to wildly different reading based on multitude of factors such as: size of room, are there any windows open or whether the heating is on etc. 

![AET](/img/SDSDesignApproach3.jpg)