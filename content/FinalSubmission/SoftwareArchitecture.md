---
title: Underlying Software Architecture
weight: 1
pre: "<b>5.1 </b>"
---

Our software stack is comprised of physical hardware sensors right up to real-time cloud databases and front end web apps. Each component in this stack has been developed using software and tooling specifically chosen or developed, bespoke, for this application.


## Backend Stack

### Raspberry PI 

Our backend stack, conceptually starts from the Raspberry PI connected to the embedded arduino sensor board. The Raspberry PI is connected via serial to the arduino and receives data from the attached sensor in the form of a serial data stream. The software running on the Pi was written by us in Golang. 
Golang is a language developed by Google specifically to be clean, concise and efficient with an emphasis on concurrency and low overhead computation. As it is a compiled language with static library linking we saw it as the perfect tool for cross-architecture scripting (ARM/ x86) Using a feature of Golang known as *'compile time variable injection'* we were able to effectively *bake in* access tokens to our binaries, obscuring them from anyone who may gain access to the embedded system. With the reduced performance of a Raspberry Pi when compared with a conventional computer the efficiency of Golang also made it an appealing choice.
Going forward we could easily set up CI/CD pipelines using CircleCI or Github Actions to build our Go binaries for each system we are deploying them to and furnish each device with a simple script to fetch any new build artifacts allowing for seamless updating of our deployed devices, requiring no on-site maintenance. 
The concurrency features of Golang could facilitate a Pi being connected to multiple input streams/ sensors and being able to parse the input from a single script, further reducing overhead and allowing for much of the required logic to be recycled and used for different inputs.

Our current prototype features a script that can handle the input of a single serial data stream. Using this data the script can track the occupancy of a single room, writing any changes back to the InfluxDB via the InfluxDB API. The script using little of the power available to us on the Raspberry PI B+ model we are using, raising the question of whether we could instead use a different model such as the Raspberry PI Zero. Such a decision would reduce power usage, deployment cost and space required for a deployment.


### InfluxDB 

Next in our stack is our real-time database hosted by InfluxDB. We chose to use this system due to the structure of the data we will be collecting. Occupancy metrics are only useful in a real time context, delayed readings are useless when trying to accurately report on room usage to users. InfluxDB uses a bucket based data organisational schema, however, also allows for each reading added to the database to have custom tags attached to it. When querying the database you can filter by these tags. We chose to organise all our data in a single bucket but tag each entry with a unique room, building and floor ID. This way no matter how we want to interpret and display the data we can accurately query the database using Influx's own *Flux* Query language. 

### Web API


In order to easily query our database we abstracted the Flux queries away with our RestAPI featuring endpoints to query by the pre-defined room, floor and building IDs.
Also written in Golang, we are again able to embed access tokens into the binary. With low overhead it can run on the smallest of cloud compute units or even inside a docker container. 

In the future it may be useful to run this inside a docker container as it would allow for a fully Kubernetes based scalable cloud architecture which would allow our system to scale up and down depending on usage. I.e. higher capacity during peak times, lower over holidays or overnight. This would reduce running costs and greatly increase reliability to the user.

An area for improvement to the current API would be the implementation of some form of token based authentication to prevent unauthorised agents from gaining information relating to occupancy. Such information is somewhat sensitive and could be used to coordinate illegal activity such as robberies etc.
