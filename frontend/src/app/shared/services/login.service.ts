import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Account, LoginResponse } from '../models/user';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { Credentials } from '../models/credentials';

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
  isUserLoggedIn = false;
  constructor(private http: HttpClient) {
  }

  isLoggedIn(): boolean {
    return this.isUserLoggedIn;
  }

  getLoginResponse(): LoginResponse {
    return this.loginResponse;
  }

  logout() {
    this.isUserLoggedIn = false;
    this.http.post(logoutURL, null, jsonOptions).subscribe();
  }

  login(credentials: Credentials): Observable<LoginResponse> {
    return this.http.post<LoginResponse>(loginUrl, credentials, jsonOptions).pipe(tap(loginResponse => {
      if (loginResponse.user) {
        this.loginResponse = loginResponse
        this.isUserLoggedIn = true;
        //  sessionStorage.setItem("loginData", JSON.stringify(loginResponse))
      } else console.log("not logged in");
    }));
  }

  loginWithCookie(): Observable<LoginResponse> {
    return this.http.get<LoginResponse>(loginWithCookieURL, jsonOptions).pipe(tap(loginResponse => {
      if (loginResponse.user) {
        this.loginResponse = loginResponse;
        this.isUserLoggedIn = true;
      }
    }));
  }
}
