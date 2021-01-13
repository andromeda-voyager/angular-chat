import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Account } from './account';
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
  account!: Account;
  private isLoggedIn: boolean = false;
  constructor(private http: HttpClient) {
  }

  isUserLoggedIn() {
    return this.isLoggedIn;
  }

  getUserData(): Account {
    return this.account;
  }

  logout() {
    this.isLoggedIn = false;
    return this.http.post(logoutURL, null, jsonOptions).subscribe();
  }

  login(credentials: Credentials): Observable<Account> {
    return this.http.post<Account>(loginUrl, credentials, jsonOptions).pipe(tap(account => {
      this.account = account;
      if(account.email)
      {this.isLoggedIn = true;}
    }));
  }

  loginWithCookie(): Observable<Account> {
    return this.http.get<Account>(loginWithCookieURL, jsonOptions).pipe(tap(account => {
      this.account = account;
      if(account.email)
      {this.isLoggedIn = true;}
    }));
  }
}
