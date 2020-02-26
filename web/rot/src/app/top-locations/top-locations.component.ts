import { Component, OnInit, AfterViewInit } from "@angular/core";
import locations from "../locations.json";
import { HttpClient, HttpHeaders } from "@angular/common/http";

export interface OccupancyResult {
  occupancy: number;
}

@Component({
  selector: "app-top-locations",
  templateUrl: "./top-locations.component.html",
  styleUrls: ["./top-locations.component.css"]
})
export class TopLocationsComponent implements OnInit {
  corsHeaders = new HttpHeaders({
    "Content-Type": "application/json",
    Accept: "application/json",
    "Access-Control-Allow-Origin": "http://localhost:6969/"
  });
  res: OccupancyResult;
  async getOcc(id: number) {
    let baseURL = "/api/occupancy/?buildingID=" + id;
    const result = await this.http.get<any>(baseURL).toPromise();
    return result.occupancy;
  }

  construct_list() {
    let ret = [];
    locations.forEach(async building => {
      const occ = await this.getOcc(building.id);
      ret.push({ name: building.name, percent: occ / building.max_occupancy });
    });
    return ret;
  }

  constructor(private http: HttpClient) {}

  public topLoc = this.construct_list();

  ngOnInit(): void {}
}
