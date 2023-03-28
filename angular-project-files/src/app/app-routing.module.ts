import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { SampleRoomPageComponent } from './sample-room-page/sample-room-page.component';

const routes: Routes = [];

@NgModule({
  //imports: [RouterModule.forRoot(routes)],

  imports: [RouterModule.forRoot([
    {path: 'room-page', component: SampleRoomPageComponent},
    {path: 'home', component: HomePageComponent}
  ]),],


  exports: [RouterModule]
})
export class AppRoutingModule { }
