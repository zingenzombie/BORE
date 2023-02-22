import { Component, OnInit } from '@angular/core';
import { UserService } from './user.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  title = 'BORE';
  subTitle = 'Join a room and share files';

  showRectangle = false;

  username!: string | null;
  ngOnInit() {
    this.username = localStorage.getItem('username');
  }

  constructor(public userService: UserService) { }
}