import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-top-locations',
  templateUrl: './top-locations.component.html',
  styleUrls: ['./top-locations.component.css']
})
export class TopLocationsComponent implements OnInit {

  constructor() { }
  public topLoc  = [{name: 'Library', percent: 15}, {name: 'CS', percent: 20}, {name: 'Guild of Students', percent: 25}];
  ngOnInit(): void {
  }

}
