import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-top-rooms',
  templateUrl: './top-rooms.component.html',
  styleUrls: ['./top-rooms.component.css']
})
export class TopRoomsComponent implements OnInit {

  constructor() { }
  public topLoc  = [{name: 'Room 101', percent: 15}, {name: 'Room 202', percent: 20}, {name: 'Room 303', percent: 25}];
  ngOnInit(): void {
  }

}
