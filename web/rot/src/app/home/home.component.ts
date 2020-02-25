import { Component, OnInit, HostListener  } from '@angular/core';
declare var JustGage: any;
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})



export class HomeComponent implements OnInit {
  public mobile = false;
  public innerWidth: any;
  constructor() {}


  @HostListener('window:resize', ['$event']) 
    onresize(event) {
      this.innerWidth = window.innerWidth;
      if (this.innerWidth < 500) {
        this.mobile = true;
      } else {
        this.mobile = false;
      }
    }

  

  ngOnInit() {
      this.innerWidth = window.innerWidth;
      if (this.innerWidth < 500) {
        this.mobile = true;
      }
      let g = new JustGage({
            id: 'gauge',
            value: 67,
            min: 0,
            max: 100,
            title: 'Visitors'
          });
  }


}

