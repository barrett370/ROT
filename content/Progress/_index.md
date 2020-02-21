---
title: Progress
hidden: true
---

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

{{<mermaid>}}
gantt
        dateFormat DD-MM
        axisFormat %d/%m
        title                                           Initial concept prototype and improvements

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
