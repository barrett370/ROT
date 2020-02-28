import { Component, OnInit } from "@angular/core";
import locations from "../locations.json";
import { WebServerInterface, IDTypes } from "../shared/shared_functions.js";
import { HttpClient } from "@angular/common/http";

interface Occupancy {
  name: string;
  percent: number;
}
@Component({
  selector: "app-top-rooms",
  templateUrl: "./top-rooms.component.html",
  styleUrls: ["./top-rooms.component.css"]
})
export class TopRoomsComponent implements OnInit {
  webServerInterface = new WebServerInterface(this.http);
  constructor(private http: HttpClient) {}

  get_all_rooms(n: number): Occupancy[] {
    var ret: Occupancy[] = [];
    locations.slice(0,1).forEach(building => {
      building.floors.slice(0,1).forEach(floor => {
        floor.rooms.slice(0, n).forEach(async room => {
          console.log(room.id)
          const occ = await this.webServerInterface.getOccupancy(
            room.id,
            IDTypes.RID
          );
          ret.push({ name: room.name, percent: occ / room.max_occupancy * 100 });
        });
      });
    });
    return ret;
  }

  get_top_rooms() {
    let ret = this.get_all_rooms(3);
    ret.sort((o1, o2) => {
      if (o1.percent < o2.percent) {
        return -1;
      } else if (o1.percent > o2.percent) {
        return 1;
      }
      return 0;
    });
    console.log(ret);
    return ret;
    // console.log(ret.length)
    // console.log(ret);
    // // let sorted_ret: Occupancy[] = ret.sort((o1, o2) => {
    // //   if (o1.percent > o2.percent) {
    // //     return 1;
    // //   }
    // //   if (o1.percent < o2.percent) {
    // //     return -1;
    // //   }
    // //   return 0;
    // // });
    // // let sorted_ret: Occupancy[] = ret.sort((o1, o2) => o1.percent - o2.percent);
    // // console.log(sorted_ret);
    // ret = ret.slice(0,3)

    // // let slice = Array.prototype.slice.call(ret,0,3)
    // console.log(ret);
    // return ret;
  }

  // public topLoc = [
  //   { name: "Room 101", percent: 15 },
  //   { name: "Room 202", percent: 20 },
  //   { name: "Room 303", percent: 25 }
  // ];
  public topLoc = this.get_top_rooms();
  ngOnInit(): void {}
}
