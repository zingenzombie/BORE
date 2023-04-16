import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { UserService } from './services/user.service';
import { DescriptionBoxComponent } from './description-box/description-box.component';
import { ViewChild } from '@angular/core';
import { UsernameComponent } from './username/username.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  userComponent!: UsernameComponent;
  
  title = 'BORE';
  subTitle = 'Join a room and share files';
  user = this.userService.getUsername();
  showRectangle = false;

  
  
  username!: string | null;

  

  ngOnInit() {
    this.username = localStorage.getItem('username');
  }

  constructor(public userService: UserService) { }
}