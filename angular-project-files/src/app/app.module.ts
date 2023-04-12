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
import { DescriptionBoxComponent } from './description-box/description-box.component';
import { SampleRoomPageComponent } from './sample-room-page/sample-room-page.component';
import { RouterModule } from '@angular/router';
import { FlexboxcircleroomComponent } from './flexboxcircleroom/flexboxcircleroom.component';
import { FlexboxtriangleroomComponent } from './flexboxtriangleroom/flexboxtriangleroom.component';
import { FlexboxsquareroomComponent } from './flexboxsquareroom/flexboxsquareroom.component';
import { ToolbarComponent } from './toolbar/toolbar.component';
import { SquareroomComponent } from './squareroom/squareroom.component';

@NgModule({
  declarations: [
    AppComponent,
    CircleroomComponent,
    FileComponent,
    FlexboxcircleroomComponent,
    FlexboxtriangleroomComponent,
    FlexboxsquareroomComponent,
    ToolbarComponent,
    SquareroomComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    CommonModule,
    FormsModule,
    RouterModule.forRoot([
    ]),
  ],
  providers: [UserService],
  bootstrap: [AppComponent]
})
export class AppModule { }
