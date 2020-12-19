import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Login } from './login';
import { Observable } from 'rxjs';
import { User } from './user';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'accept': 'application/json',
    // 'withCredentials': 'true'
  }), credentials: 'same-origin',
  withCredentials: true
};

const loginUrl = "http://localhost:8080/login";
@Injectable({
  providedIn: 'root'
})

export class ChatService {
 
  constructor(private http: HttpClient) { }

  login(login: Login) : Observable<User>{
    let k = {Username:"matt"};
    console.log(k);
    return this.http.post<User>(loginUrl, login, httpOptions);
  }

}
