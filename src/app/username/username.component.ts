import { Component } from '@angular/core';
import { User } from './user';
import { UserService } from '../user.service';

@Component({
  selector: 'app-username',
  templateUrl: './username.component.html',
  styleUrls: ['./username.component.css']
})
export class UsernameComponent {
  username!: string;
  submitted = false;
  
  constructor(private userService: UserService) {}

  onSubmit() {
    this.userService.setUsername(this.username);
    this.submitted = true;
  }
}
