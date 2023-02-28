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


@NgModule({
  declarations: [
    AppComponent,
    CircleroomComponent,
    FileComponent,
    UsernameComponent,
    DescriptionBoxComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    CommonModule,
    FormsModule
  ],
  providers: [UserService],
  bootstrap: [AppComponent]
})
export class AppModule { }
