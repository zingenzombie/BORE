import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CircleroomComponent } from './circleroom/circleroom.component';
import { FileComponent } from './file/file.component';
import { UserService } from './services/user.service';
import { UsernameComponent } from './username/username.component';
import { DescriptionBoxComponent } from './description-box/description-box.component';
import { HomePageComponent } from './home-page/home-page.component';
import { SampleRoomPageComponent } from './sample-room-page/sample-room-page.component';
import { RouterModule } from '@angular/router';
import { Component } from '@angular/core';

@NgModule({
  declarations: [
    AppComponent,
    CircleroomComponent,
    FileComponent,
    DescriptionBoxComponent,
    HomePageComponent,
    UsernameComponent,
    SampleRoomPageComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    CommonModule,
    FormsModule,
    RouterModule.forRoot([
      {path: 'room-page', component: SampleRoomPageComponent},
      {path: 'home', component: HomePageComponent}
    ]),
  ],
  exports: [
    DescriptionBoxComponent
  ],
  providers: [UserService],
  bootstrap: [AppComponent, DescriptionBoxComponent]
})
export class AppModule { }
