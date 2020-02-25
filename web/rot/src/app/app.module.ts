import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { TopRoomsComponent } from './top-rooms/top-rooms.component';
import { TopLocationsComponent } from './top-locations/top-locations.component';
import { LocationComponent } from './location/location.component';
import { HomeComponent } from './home/home.component';
import { HeaderComponent } from './header/header.component';

@NgModule({
  declarations: [
    AppComponent,
    TopRoomsComponent,
    TopLocationsComponent,
    LocationComponent,
    HomeComponent,
    HeaderComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
