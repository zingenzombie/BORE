import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

interface UserObject {
  username: string;
}

@Injectable({
  providedIn: 'root'
})
export class UserService {
  name!: string;
  private apiUrl = 'http://10.136.165.182:3621/setName';
  //private user: UserObject = {name: this.name};

  constructor(private http: HttpClient) {}

  setUsername(name: string) {
    this.name = name;
  }

  getUsername(): string {
    return this.name;
  }
  
  
  postUsername(name: string) {
    this.http.post(this.apiUrl, name, {responseType: 'json'})
    .subscribe((res) => {
      console.log(res);
    });
    
  }

  // postUsername(user: {username: string}) {
  //   this.http.post(this.apiUrl, user)
  //   .subscribe((res) => {
  //     console.log(res);
  //   });
  //   //return this.http.post(this.apiUrl, this.user);
  // }

  retrieveUsername() {
    //return this.http.get<Confix>
  }
}
