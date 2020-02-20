---
title: Software tooling
draft: false
---

## Data Representation & Storage

### InfluxDB

InfluxDB, 2.0 specifically, is a prime candidate for the storage solution to a project such as ours. Specialising in the storage of time series data, the use of one of their databases would allow us to easily store sensor data in time order and parse to provide accurate, real-time statistics to our end users. SQL based approaches such as PostgreSQL are not built for storing real-time information and as such they introduce significant overhead causing great difficulty in creating real-time systems such as ours.

Featuring heavy Grafana integration in the newest 2.0 version, the representation of data to the developer is very clear and easy to work with. Clearly the data stored in the database would require post-processing in order to derive the occupancy of a given room as data such as ambient temperature or CO2 readings are not of interest to the user. 


#### Interfacing 

Querying an InfluxDB Database requires the use of their own query language *(InfluxQL)*. They also provide client libraries for C#, Golang, Java, Javascript/Node.js, Python and Ruby. 

For the case of pushing data to our database from an embedded system such as a raspberry Pi or Arduino, the use of a lighter-weight, compiled language such as C# or Golang might be better. Further, server-side processing of the data need-not use the same language and such a task may lend itself more to a language such as Python. 

### Grafana

Grafana is an open source solution to easily creating read-outs and dashboards for presenting data to either system administrators or end-users. Its integration with InfluxDB would allow for a canonical solution to data representation throughout our product.

Grafana allows for many presentations of data, and is extensible via a (Java/Type)script API which would allow us to create custom graphs such as heatmap overlays for presenting occupancy visually in relation to location. 

The ability to present the data with a wide variety of visuals will allow for less technical users to easily grasp the concepts without having to necessarily read off charts or graphs, expanding our audience and stopping people from being put off using our service.

