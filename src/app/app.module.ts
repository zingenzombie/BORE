import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DescriptionBoxComponent } from './description-box/description-box.component';
import { CircleRoomComponent } from './circle-room/circle-room.component';
import { RectangleComponent } from './circle-room/rectangle/rectangle.component';
@NgModule({
  declarations: [
    AppComponent,
    DescriptionBoxComponent,
    CircleRoomComponent,
    RectangleComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
