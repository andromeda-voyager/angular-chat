import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Login } from './login';
import { Observable } from 'rxjs';
import { User } from './user';
import { environment } from 'src/environments/environment';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'accept': 'application/json',
  }), credentials: 'same-origin',
  withCredentials: true
};
const httpOptions2 = {
  headers: new HttpHeaders({
    'Content-Type': 'multipart/form-data',
    'accept': 'application/json',
  }), credentials: 'same-origin',
  withCredentials: true
};

const createAccountUrl = environment.BaseApiUrl + "/create-account";
const loginUrl = environment.BaseApiUrl + "/login";
const uploadImageUrl = environment.BaseApiUrl + "/upload-image";

@Injectable({
  providedIn: 'root'
})

export class ChatService {

  constructor(private http: HttpClient) { }

  login(login: Login): Observable<User> {
    let k = { Username: "matt" };
    console.log(k);
    return this.http.post<User>(loginUrl, login, httpOptions);
  }

  createAccount(file: File, user: User) {
    const formData = new FormData();
    formData.append("image", file, file.name);
    formData.append("user", JSON.stringify(user));

    this.http.post(createAccountUrl, formData).subscribe(() => {
      console.log("user created");
    })
  }

  uploadImage(file: File) {
    const formData = new FormData();
    formData.append("file", file, file.name);
    this.http.post(uploadImageUrl, formData).subscribe(() => {
      console.log("file uploaded");
    })

  }
}
