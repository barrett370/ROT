import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute,  ParamMap } from '@angular/router';
import { switchMap } from 'rxjs/operators';
import { Identifiers } from '@angular/compiler';

@Component({
  selector: 'app-location',
  templateUrl: './location.component.html',
  styleUrls: ['./location.component.css']
})
export class LocationComponent implements OnInit {

  public id: string;
  constructor(route: ActivatedRoute) {
    route.params.subscribe(val => {
      // put the code from `ngOnInit` here
      this.id = val.id;
      console.log(val);
    });
  }
  public topLoc  = [{name: 'Library', percent: 15}, {name: 'CS', percent: 20}, {name: 'Guild of Students', percent: 25}];

  ngOnInit() {

  }

}
