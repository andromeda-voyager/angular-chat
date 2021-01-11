import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Login } from './login';
import { Observable } from 'rxjs';
import { Account, NewAccount } from './account';
import { environment } from 'src/environments/environment';
import { NewServer, Server, Invite, Connection } from '../shared/server';

const jsonOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'accept': 'application/json',
  }), credentials: 'same-origin',
  withCredentials: true
};
const formOptions = {
  headers: new HttpHeaders({
    'accept': 'application/json',
  }), credentials: 'same-origin',
  withCredentials: true
};

const createAccountUrl = environment.BaseApiUrl + "/create-account";
const loginUrl = environment.BaseApiUrl + "/login";
const verificationCodeUrl = environment.BaseApiUrl + "/send-verification-code";
const createServerUrl = environment.BaseApiUrl + "/create-server";
const joinServerUrl = environment.BaseApiUrl + "/join-server";
const getPostsURL = environment.BaseApiUrl + "/posts";

@Injectable({
  providedIn: 'root'
})

export class ChatService {
  account!: Account;
  constructor(private http: HttpClient) { }

  login(login: Login): Observable<Account> {
    let k = { Username: "matt" };
    console.log(k);
    return this.http.post<Account>(loginUrl, login, jsonOptions);
  }

  addUserData(account: Account) {
    this.account = account;
  }

  getUserData(): Account {
    return this.account;
  }

  createAccount(file: File, user: NewAccount): Observable<Account>  {
    const formData = new FormData();
    if (file) {
      formData.append("image", file, file.name);
    }
    formData.append("user", JSON.stringify(user));

    return this.http.post<Account>(createAccountUrl, formData);
  }

  createServer(file: File, server: NewServer): Observable<Connection> {
    const formData = new FormData();
    if (file) {
      formData.append("image", file, file.name);
    }
    formData.append("server", JSON.stringify(server));

    return this.http.post<Connection>(createServerUrl, formData, formOptions);
  }

  joinServer(invite: Invite): Observable<Server> {
    return this.http.post<Server>(joinServerUrl, invite);
  }

  getPosts(serverID: number): Observable<Server> {
    return this.http.get<Server>(getPostsURL + "?serverID="+serverID, formOptions);
  }

  sendVerificationCode(_email: string) {
    let email = { email: _email };
    this.http.post(verificationCodeUrl, email).subscribe(() => {
      console.log("user created");
    })
  }

  // uploadImage(file: File) {
  //   const formData = new FormData();
  //   formData.append("file", file, file.name);
  //   this.http.post(uploadImageUrl, formData).subscribe(() => {
  //     console.log("file uploaded");
  //   })

  // }
}
