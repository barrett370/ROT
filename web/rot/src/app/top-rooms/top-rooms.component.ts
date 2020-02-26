import { Component, OnInit } from "@angular/core";
import locations from "../locations.json";
import { WebServerInterface, IDTypes } from "../shared/shared_functions.js";
import { HttpClient } from "@angular/common/http";

@Component({
  selector: "app-top-rooms",
  templateUrl: "./top-rooms.component.html",
  styleUrls: ["./top-rooms.component.css"]
})
export class TopRoomsComponent implements OnInit {
  webServerInterface = new WebServerInterface(this.http);
  constructor(private http: HttpClient) {}

  construct_list() {
    let ret = [];
    locations.forEach(building => {
      building.floors.forEach(floor => {
        floor.rooms.forEach(async room => {
          const occ = await this.webServerInterface.getOccupancy(
            room.id,
            IDTypes.RID
          );
          ret.push({ name: room.name, percent: occ / room.max_occupancy });
        });
      });
    });
    console.log(ret);
    let sorted_ret:any[] = ret.sort((o1, o2) => {
      if (o1.percent > o2.percent) {
        return 1;
      }
      if (o1.percent < o2.percent) {
        return -1;
      }
      return 0;
    });
    console.log(sorted_ret);
    console.log(ret.slice(0, 4));
    return ret
  }

  // public topLoc = [
  //   { name: "Room 101", percent: 15 },
  //   { name: "Room 202", percent: 20 },
  //   { name: "Room 303", percent: 25 }
  // ];
  public topLoc = this.construct_list();
  ngOnInit(): void {}
}
