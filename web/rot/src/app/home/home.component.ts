import { Component, OnInit } from '@angular/core';
declare var JustGage: any;
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {

    let g = new JustGage({
      id: 'gauge',
      value: 67,
      min: 0,
      max: 100,
      title: 'Visitors'
    });

  }

}
