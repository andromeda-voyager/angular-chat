import { Injectable } from '@angular/core';
import { User, Account } from '../models/user';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

const formOptions = {
  headers: new HttpHeaders({
    'accept': 'application/json',
  }), credentials: 'same-origin',
  withCredentials: true
};

const createAccountUrl = environment.BaseApiUrl + "/create-account";
const verificationCodeUrl = environment.BaseApiUrl + "/send-verification-code";

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  constructor(private http: HttpClient) { }

  createAccount(file: File, account: Account): Observable<User>  {
    const formData = new FormData();
    if (file) {
      formData.append("image", file, file.name);
    }
    formData.append("user", JSON.stringify(account));

    return this.http.post<User>(createAccountUrl, formData);
  }

  sendVerificationCode(_email: string) {
    let email = { email: _email };
    this.http.post(verificationCodeUrl, email).subscribe(() => {
      console.log("user created");
    })
  }

}
