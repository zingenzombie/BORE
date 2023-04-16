import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { Component, OnInit } from '@angular/core';
import { map } from 'rxjs/operators';

interface Name {
  Name: string;
}

@Injectable({
  providedIn: 'root'
})
export class UserService {
  name!: string;
  debugString!: Name;
  tempString!: string;
  tempString1!: string;
  username!: string;

  tempName!: string;
  ip = location.host;
  private apiUrl = 'http://localhost:3621';
  //private user: UserObject = {name: this.name};

  constructor(private http: HttpClient) {}

  setUsername(name: string) {
    this.name = name;
  }

  ngOnInit() {
    
  }

  updateUsername(): string {
    return this.returnUsername();
  }

  returnUsername(): string {
    const str = this.getUsername();
    const word = this.extractWord(str);
    return word;
  }

  extractWord(str: string): string {
    const match = str.match(/"([^"]*)"/g);
    if (match && match.length > 1) {
      return match[1].replace(/"/g, '');
    }
    return '';
  }
  // getDebug() {
  //    this.tempName = this.apiUrl + '/debug';
  //    return this.http.get<Name>(this.tempName);
  // }

  // getUsername(): Observable<string> {
  //   this.tempName = this.apiUrl + '/debug';
  //   return this.http.get<Name>(this.tempName).pipe(
  //     map((data: Name) => {
  //       this.debugString = data;
  //       console.log(this.debugString.Name); // should output "debug"
  //       return this.debugString.Name;
  //     })
  //   );
  // }
  getUsername(): string {
    this.tempName = this.apiUrl + '/debug';
    this.http.get(this.tempName).pipe(
      map(data => JSON.stringify(data))
    )
    .subscribe(result => {
      this.tempString = result;
      console.log(this.tempString); // Prints the converted string to the console
    });
    return this.tempString;
  }
  
  getUsername1(): string {
    this.tempName = this.apiUrl + '/debug';
    this.http.get<Name>(this.tempName).pipe(
      map(data => (data))
    )
    .subscribe(result => {
      this.debugString = result;
      this.tempString1 = JSON.stringify(result.Name);
      console.log(this.tempString1); // Prints the converted string to the console
    });

    return this.tempString1;
  }
  
  postUsername(name: string) {
    this.tempName = this.apiUrl + '/setName';
    this.http.post(this.tempName, name, {responseType: 'json'})
    .subscribe((res) => {
      console.log(res);
    });
    
  }

  


  retrieveUsername() {
    //return this.http.get<Confix>
  }
}
