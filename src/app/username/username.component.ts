import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { User } from './user';
import { UserService } from '../services/user.service';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-username',
  templateUrl: './username.component.html',
  styleUrls: ['./username.component.css']
})
export class UsernameComponent implements OnInit{
  name!: string;
  submitted = false;

  userList: string[] = [];
  
  constructor(public userService: UserService) {}
  //constructor(private http: HttpClient, private usersService: UsersService) {}

  ngOnInit(): void {
    
  }

  onSubmit(name: string) {
    console.log(name);
    this.userService.setUsername(this.name);
    this.submitted = true;
    this.addUser(this.name);
    this.userService.postUsername(name);
    
  }

  addUser(name: string) {
    this.userList = [...this.userList, name];
  }
}

