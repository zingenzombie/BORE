import { HttpClient, HttpHeaders } from '@angular/common/http';
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
  private apiUrl = 'http://10.136.157.172:3621/setName';
  //private user: UserObject = {name: this.name};

  constructor(private http: HttpClient) {}

  setUsername(name: string) {
    this.name = name;
  }

  getUsername(): any {
    // this.apiUrl = 'http://10.136.157.172:3621/getName';
    // return JSON.stringify(this.http.get<any>('http://10.136.165.182:3621/getName'), null);
    //return JSON.stringify(this.http.get<any>('http://10.136.157.172:3621/getName'), null) + 'i';
    //return this.http.get('http://10.136.157.172:3621/getName').pipe(Map((res) => {return res;}));
    return this.name;
  }
  
  postUsername(username: any) {
    console.log("This is a print statement" , username.username);
    const headers = new HttpHeaders({'Content-Type': 'application/json'});
    this.http.post('http://10.136.157.172:3621/setName', JSON.parse('{"name": "' + username.username + '"}'))
    .subscribe((res) => {
      console.log(res);
    });
    
  }


  // postUsername(User: {name: string}) {
    
  //   this.http.post('http://10.136.157.172:3621/setName', User, {responseType: 'json'})
  //   .subscribe((res) => {
  //     console.log(res);
  //   });
    
  // }

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
