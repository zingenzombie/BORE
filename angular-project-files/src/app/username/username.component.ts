import { Component, NgModule, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { User } from './user';
import { UserService } from '../services/user.service';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-username',
  templateUrl: './username.component.html',
  styleUrls: ['./username.component.css']
})

export class UsernameComponent implements OnInit{
  name!: string;
  submitted = false;
  username!: string;
  showUsers = false;

  userList: string[] = [];
  
  constructor(public userService: UserService) {}
  //constructor(private http: HttpClient, private usersService: UsersService) {}

  ngOnInit(): void {
    
  }

  toggleUsers() {
    this.showUsers = !this.showUsers;
  }

  onSubmit(name: string) {
    console.log(name);
    this.submitHelper(name);
    this.userService.setUsername(this.name);
    
    this.addUser(this.name);
      
    this.username = this.name;
    //get username
    this.userService.updateUsername();
  }

  submitHelper(name: string) {
    this.userService.postUsername(name);
    this.submitted = true;
  }

  addUser(name: string) {
    this.userList = [...this.userList, name];
  }

  // url/getAllMembers returns a string separated by commas
  // puts users into a list
  getAllUsers(input: string): string[] {
    if (!input) {
      return [];
    }
    const list: string[] = input.split(',');
    return list;
  }

  

}

