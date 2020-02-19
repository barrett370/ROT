---
title: Software tooling
draft: true
---

## Data Representation & Storage

### InfluxDB

InfluxDB, 2.0 specifically, is a prime candidate for the storage solution to a project such as ours. Specialising in the storage of time series data, the use of one of their databases would allow us to easily store sensor data in time order and parse to provide accurate, real-time statistics to our end users.

Featuring heavy Grafana integration in the newest 2.0 version, the representation of data to the developer is very clear and easy to work with. Clearly the data stored in the database would require post-processing in order to derive the occupancy of a given room as data such as ambient temperature or CO2 readings are not of interest to the user. 


### Grafana

Grafana is an open source solution to easily creating read-outs and dashboards for presenting data to either system administrators or end-users. Its integration with InfluxDB would allow for a canonical solution to data representation throughout our product.


