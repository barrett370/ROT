import { Component, OnInit, AfterViewInit } from "@angular/core";
import locations from "../locations.json";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { WebServerInterface, IDTypes } from "../shared/shared_functions.js";

export interface OccupancyResult {
  occupancy: number;
}

@Component({
  selector: "app-top-locations",
  templateUrl: "./top-locations.component.html",
  styleUrls: ["./top-locations.component.css"]
})
export class TopLocationsComponent implements OnInit {
  construct_list() {
    let ret = [];
    locations.forEach(async building => {
      const occ = await this.webServerInterface.getOccupancy(
        building.id,
        IDTypes.BID
      );
      ret.push({ name: building.name, percent: occ / building.max_occupancy *100});
    });
    return ret;
  }
  webServerInterface = new WebServerInterface(this.http);
  constructor(private http: HttpClient) {}

  public topLoc = this.construct_list();

  ngOnInit(): void {}
}
