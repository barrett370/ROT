---
title: "Home"
date: 2020-02-18T15:28:25Z
---

# Welcome!

We are developing a Open Source Room Occupancy Tracking System for use primarily in educational institutions, providing students with accurate, real-time information on the current occupancy of study spaces across campuses

## Team Members

- Jon Freer
- Sam Barrett
- Charlie de Freitas
- Anant Kapoor
- George Wellington

{{% children %}}

### Summary

- [ ] Problem outline
  - [ ] Why we picked this
- [ ] Design
  - [ ] Early sketches
  - [ ] Current concept
  - [ ] What is currently working
  - [ ] Critique of current concept
- [ ] Research methods
  - [x] Surveys
  - [x] Individual interviews
  - [x] Brainstorming
  - [ ] Personas
  - [ ] Desk research


We have conducted in depth research through a range of techniques to allow us to make use of our own data along with secondary data.

With this foundation of information we along with starting to explore our chosen design in further detail.

You must present the problem, and your conceptual approach to the solution, along with a description of key parts that you have got working so far.  You are expected to have at least a basic concept designed. You should also present a critique of the solution so far with some ideas as to how to develop it further/change it to improve it.

You will be assessed on:

design thinking
clarity of expression
quality and innovation behind idea
implementation approach
critique

{{<mermaid>}}
gantt
        dateFormat DD-MM
        axisFormat %d/%m
        title                                           Outline of Tasks

        section                                         Research
        Research focus points                           :fsc1, 11-02,           3h
        Surveys                                         :fsc2, after fsc1,      3h
        Interviews                                      :fsc3, after fsc1,      5h
        Personas                                        :fsc4, after fsc2,      2h
        Desk Research                                   :fsc4, after fsc2,      2h

        section                                         Conceptual Design
        What data to collect?                           :aat1, after fsc3,      3h
        What do we want to show?                        :fsc5, after fsc3,      3h

        section                                         Technical Design
        Hardware comparison                             :aat2, after aat1,      3h
        Data collection                                 :aat4, after aat2,      2h
        Data distribution                               :aat5, after aat2,      3h
        Data representation                             :fsc5, after aat2,      3h

        section                                         Prototyping
        Language Choice                                 :aat3, after aat4 aat5, 2h
        First device build                              :prt1, after aat4,      3h
        Metric ingress                                  :prt2, after aat5,      3h
{{</mermaid>}}
