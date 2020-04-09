import { HttpClient, HttpHeaders } from "@angular/common/http";
export enum IDTypes{
    "BID",
    "FID",
    "RID",
}
export class WebServerInterface {
  constructor(private http: HttpClient) {}
    async getOccupancy(id: number, type: IDTypes) {
        var baseURL: string
        switch ( type ){
            case IDTypes.BID:
                baseURL = "/api/occupancy/?buildingID=" + id;
                break;
            case IDTypes.FID:
                baseURL = "/api/occupancy/?floorID=" + id;
                break;
            case IDTypes.RID:
                baseURL = "/api/occupancy/?roomID=" + id;
                break;
            default:
                console.log("ERROR incorrect ID type")
      }
    const result = await this.http.get<any>(baseURL).toPromise();
    return result.occupancy;
  }
}
