import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Account, LoginResponse } from './user';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Credentials } from './credentials';
import { Observable } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

const jsonOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'accept': 'application/json',
  }), credentials: 'same-origin',
  withCredentials: true
};

const loginWithCookieURL = environment.BaseApiUrl + "/login-with-cookie";
const loginUrl = environment.BaseApiUrl + "/login";
const logoutURL = environment.BaseApiUrl + "/logout";

@Injectable({
  providedIn: 'root'
})

export class LoginService {
  loginResponse!: LoginResponse;
  private isLoggedIn: boolean = false;
  constructor(private http: HttpClient) {
  }

  isUserLoggedIn() {
    return this.isLoggedIn;
  }

  getLoginResponse(): LoginResponse {
    return this.loginResponse;
  }

  logout() {
    this.isLoggedIn = false;
    return this.http.post(logoutURL, null, jsonOptions).subscribe();
  }

  login(credentials: Credentials): Observable<LoginResponse> {
    return this.http.post<LoginResponse>(loginUrl, credentials, jsonOptions).pipe(tap(loginResponse => {
      this.loginResponse = loginResponse;
      if (loginResponse.user) {
        this.isLoggedIn = true; 
        console.log("logged in");
      } else console.log("not logged in");
    }));
  }

  loginWithCookie(): Observable<LoginResponse> {
    return this.http.get<LoginResponse>(loginWithCookieURL, jsonOptions).pipe(tap(loginResponse => {
      this.loginResponse = loginResponse;
      if (loginResponse.user) { this.isLoggedIn = true; }
    }));
  }
}
