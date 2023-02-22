import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CircleroomComponent } from './circleroom/circleroom.component';
import { FileComponent } from './file/file.component';
import { UserService } from './user.service';
import { UsernameComponent } from './username/username.component';

@NgModule({
  declarations: [
    AppComponent,
    CircleroomComponent,
    FileComponent,
    UsernameComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    CommonModule,
    FormsModule
  ],
  providers: [UserService],
  bootstrap: [AppComponent]
})
export class AppModule { }
